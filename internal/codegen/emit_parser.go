package codegen

import (
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
)

const parserPlaceholderTemplate = `
type Error struct {
	Token    Token
	Expected []int
}

type lox struct {}

type _Lexer interface {
	ReadToken() (Token, int)
}

func (p *lox) parse(l _Lexer) bool {
	panic("not-implemented")
}

func (p *lox) recoverLookahead(typ int, tok Token) {
	panic("not-implemented")
}
`

const parserTemplate = `
var _rules = []int32 {
	{{ lhs() | array }}
}

var _termCounts = []int32 {
	{{ term_counts() | array }}	
}

var _actions = []int32 {
	{{ actions() | array }}
}

var _goto = []int32 {
	{{ goto() | array }}
}

type _Bounds struct {
	Begin Token
	End   Token
	Empty bool
}

func _cast[T any](v any) T {
	cv, _ := v.(T)
	return cv
}

type Error struct {
	Token    Token
	Expected []int
}

func _Find(table []int32, y, x int32) (int32, bool) {
	i := int(table[int(y)])
	count := int(table[i])
	i++
	end := i + count
	for ; i < end; i+=2 {
		if table[i] == x {
			return table[i+1], true
		}
	}
	return 0, false
}

type _Lexer interface {
	ReadToken() (Token, int)
}

type _item struct {
	State int32
	Sym   any
	{{- if emit_bounds }}
	Bounds _Bounds
	{{- end }}
}

type lox struct {
	_lex   _Lexer
	_stack _Stack[_item]

	_la     int
	_lasym  any

	_qla    int
	_qlasym any
}

func (p *{{parser}}) parse(lex _Lexer) bool {
  const accept = {{ accept }}

	p._lex = lex
	p._qla = -1
	p._stack.Push(_item{})

	p._readToken()

	for {
		topState := p._stack.Peek(0).State
		action, ok := _Find(_actions, topState, int32(p._la))
		if !ok {
			if !p._recover() {
				return false
			}
			continue
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			{{- if emit_bounds }}
			latok := p._lasym.(Token)
			{{- end }}
			p._stack.Push(_item{
				State: action,
				Sym:   p._lasym,
				{{- if emit_bounds }}
				Bounds: _Bounds{
					Begin: latok,
					End:   latok,
				},
				{{- end }}
			})
			p._readToken()
		} else { // reduce
			prod := -action
			termCount := _termCounts[int(prod)]
			rule := _rules[int(prod)]
			res := p._act(prod)
			{{- if emit_bounds }}

			// Compute reduction token bounds.
			// Trim leading and trailing empty bounds.
			boundSlice := p._stack.PeekSlice(int(termCount))
			for len(boundSlice) > 0 && boundSlice[0].Bounds.Empty {
				boundSlice = boundSlice[1:]
			}
			for len(boundSlice) > 0 && boundSlice[len(boundSlice)-1].Bounds.Empty {
				boundSlice = boundSlice[:len(boundSlice)-1]
			}
			var bounds _Bounds
			if len(boundSlice) > 0 {
				bounds.Begin = boundSlice[0].Bounds.Begin
				bounds.End = boundSlice[len(boundSlice)-1].Bounds.End
			} else {
				bounds.Empty = true
			}
			if !bounds.Empty {
				p.{{ on_bounds_method }}(res, bounds.Begin, bounds.End)
			}

			{{- end }}
			p._stack.Pop(int(termCount))
			topState = p._stack.Peek(0).State
			nextState, _ := _Find(_goto, topState, rule)
			p._stack.Push(_item{
				State: nextState,
				Sym:   res,
				{{- if emit_bounds }}
				Bounds: bounds,
				{{- end }}
			})
		}
	}

	return true
}

// recoverLookahead can be called during an error production action (an action
// for a production that has a @error term) to recover the lookahead that was
// possibly lost in the process of reducing the error production.
func (p *{{parser}}) recoverLookahead(typ int, tok Token) {
	if p._qla != -1 {
  	panic("recovered lookahead already pending")
	}

	p._qla = p._la
	p._qlasym = p._lasym
	p._la = typ
	p._lasym = tok
}

func (p *{{parser}}) _readToken() {
	if p._qla != -1 {
		p._la = p._qla
		p._lasym = p._qlasym
		p._qla = -1
		p._qlasym = nil
		return
	}

	p._lasym, p._la = p._lex.ReadToken()
	if p._la == ERROR {
		p._lasym = p._makeError()
	}
}

func (p *{{parser}}) _recover() bool {
	errSym, ok := p._lasym.(Error)
	if !ok {
		errSym = p._makeError()
	}

	for p._la == ERROR {
		p._readToken()
	}

	for {
		save := p._stack

		for len(p._stack) >= 1 {
			state := p._stack.Peek(0).State

			for {
				action, ok := _Find(_actions, state, int32(ERROR))
				if !ok {
					break
				}

				if action < 0 {
					prod := -action
					rule := _rules[int(prod)]
					state, _ = _Find(_goto, state, rule)
					continue
				}

				state = action

				_, ok = _Find(_actions, state, int32(p._la))
				if !ok {
					break
				}

				p._qla = p._la
				p._qlasym = p._lasym
				p._la = ERROR
				p._lasym = errSym
				return true
			}

			p._stack.Pop(1)
		}

		if p._la == EOF {
			return false
		}

		p._stack = save
		p._readToken()
	}
}

func (p *{{parser}}) _makeError() Error {
	e := Error{
		Token: p._lasym.(Token),
	}

	// Compile list of allowed tokens at this state.
	// See _Find for the format of the _actions table.
	s := p._stack.Peek(0).State
	i := int(_actions[int(s)])
	count := int(_actions[i])
	i++
	end := i + count
	for ; i < end; i += 2 {
		e.Expected = append(e.Expected, int(_actions[i]))
	}

	return e
}

func (p *{{parser}}) _act(prod int32) any {
	switch prod {
{{- range prod_index, prod := grammar.Prods }}
	{{- rule := prod.Rule }}
	{{- generated := rule_generated(rule) }}
	{{- if generated == "not_generated" }}
		{{- method := methods[prod] }}
			case {{ prod_index }}:
				return p.{{ method.Name() }}(
				{{- range param_index, param := method.Params }}
				  _cast[{{ go_type(param) }}](p._stack.Peek({{ len(method.Params) - param_index - 1 }}).Sym),
				{{- end }}
		    )
	{{- else if generated == "one_or_more" }}
	case {{ prod_index }}:  // OneOrMore
		{{- if len(prod.Terms) == 1 }}
			{{- term_go_type := go_type(get_term_go_type(prod.Terms[0])) }}
		  return []{{ term_go_type }}{
				_cast[{{ term_go_type }}](p._stack.Peek(0).Sym),
			}
		{{- else }}
			{{- term_go_type := go_type(get_term_go_type(prod.Terms[1])) }}
			return append(
				_cast[[]{{term_go_type}}](p._stack.Peek(1).Sym),
				_cast[{{ term_go_type }}](p._stack.Peek(0).Sym),
			)
		{{- end }}
	{{- else if generated == "list" }}
	case {{ prod_index }}:  // List
		{{- if len(prod.Terms) == 1 }}
			{{- term_go_type := go_type(get_term_go_type(prod.Terms[0])) }}
		  return []{{ term_go_type }}{
				_cast[{{ term_go_type }}](p._stack.Peek(0).Sym),
			}
		{{- else }}
			{{- term_go_type := go_type(get_term_go_type(prod.Terms[2])) }}
			return append(
				_cast[[]{{ term_go_type }}](p._stack.Peek(2).Sym),
				_cast[{{ term_go_type }}](p._stack.Peek(0).Sym),
			)
		{{- end }}
	{{- else if generated == "zero_or_one" }}
  case {{ prod_index }}:  // ZeroOrOne
		{{- term_go_type := go_type(rule_go_types[rule]) }}
		{{- if len(prod.Terms) == 1 }}
			return _cast[{{ term_go_type }}](p._stack.Peek(0).Sym)
		{{- else }}
			{
				var zero {{term_go_type}}
				return zero
			}
		{{- end }}
	{{- else if generated == "zero_or_more" }}
  case {{ prod_index }}:  // ZeroOrMore
		{{- term_go_type := go_type(rule_go_types[rule]) }}
		{{- if len(prod.Terms) == 1 }}
			return _cast[{{ term_go_type }}](p._stack.Peek(0).Sym)
		{{- else }}
			{
				var zero {{term_go_type}}
				return zero
			}
		{{- end }}
	{{- end }}
{{- end }}
	default:
		panic("unreachable")
	}
}
`

type parserTemplateInputs struct {
	Placeholder bool
	Package     string
	PackagePath string
}

func renderParserTemplate(in *parserTemplateInputs) string {
	template := parserTemplate
	if in.Placeholder {
		template = parserPlaceholderTemplate
	}
	vars := make(jet.VarMap)
	return renderTemplate(in.Package, in.PackagePath, template, vars)
}

func (c *context) EmitParser() bool {
	vars := make(jet.VarMap)

	const accept int32 = math.MaxInt32

	vars.Set("accept", accept)
	vars.Set("parser", c.ParserType.Obj().Name())
	vars.Set("grammar", c.ParserGrammar)
	vars.Set("methods", c.ActionMethods)
	vars.Set("rule_generated", RuleGenerated)
	vars.Set("get_term_go_type", c.getTermGoType)
	vars.Set("rule_go_types", c.RuleGoTypes)
	vars.Set("emit_bounds", c.EmitBounds)
	vars.Set("on_bounds_method", OnBoundsMethodName)
	vars.Set("on_error_method", OnErrorMethodName)

	vars.Set("array", func(arr []int32) string {
		var str strings.Builder
		WriteArray(&str, arr)
		return str.String()
	})

	vars.Set("actions", func() []int32 {
		table := newTable[int32]()
		for _, state := range c.ParserTable.States {
			var row []int32
			actions := c.ParserTable.Actions(state)
			for _, terminal := range actions.Terminals() {
				action := actions.Get(terminal).Get(0)
				row = append(row, int32(terminal.Index))
				switch action.Type {
				case lr1.ActionShift:
					row = append(row, int32(action.ShiftState.Index))
				case lr1.ActionReduce:
					row = append(row, int32(action.Prods[0].Index)*-1)
				case lr1.ActionAccept:
					row = append(row, accept)
				default:
					panic("unreachable")
				}
			}
			table.AddRow(state.Index, row)
		}
		return table.Array()
	})

	vars.Set("goto", func() []int32 {
		table := newTable[int32]()
		for _, from := range c.ParserTable.States {
			var row []int32
			transitions := c.ParserTable.Transitions(from)
			for _, input := range transitions.Inputs() {
				rule, ok := input.(*lr1.Rule)
				if !ok {
					continue
				}
				to := transitions.Get(rule)
				row = append(row, int32(rule.Index), int32(to.Index))
			}
			table.AddRow(from.Index, row)
		}
		return table.Array()
	})

	vars.Set("lhs", func() []int32 {
		lhs := make([]int32, len(c.ParserGrammar.Prods))
		for i, prod := range c.ParserGrammar.Prods {
			lhs[i] = int32(prod.Rule.Index)
		}
		return lhs
	})

	vars.Set("term_counts", func() []int32 {
		termCounts := make([]int32, len(c.ParserGrammar.Prods))
		for i, prod := range c.ParserGrammar.Prods {
			termCounts[i] = int32(len(prod.Terms))
		}
		return termCounts
	})

	parserGen := renderTemplate(
		c.GoPackageName, c.GoPackagePath, parserTemplate, vars)

	err := os.WriteFile(
		filepath.Join(c.Dir, parserGenGo), []byte(parserGen), 0666)
	if err != nil {
		c.Errs.GeneralError(err)
		return false
	}

	return true
}
