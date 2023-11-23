package codegen

import (
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/dcaiafa/lox/internal/parsergen/lr2"
)

const parserPlaceholderTemplate = `
type lox struct {}

type _Lexer interface {
	ReadToken() (Token, int)
}

func (p *lox) parse(l _Lexer) bool {
	panic("not-implemented")
}

func (p *lox) errorToken() Token {
	panic("not-implemented")
}
`

const parserTemplate = `
var _LHS = []int32 {
	{{ lhs() | array }}
}

var _TermCounts = []int32 {
	{{ term_counts() | array }}	
}

var _Actions = []int32 {
	{{ actions() | array }}
}

var _Goto = []int32 {
	{{ goto() | array }}
}

type _Bounds struct {
	Begin Token
	End   Token
}

func _cast[T any](v any) T {
	cv, _ := v.(T)
	return cv
}

var _errorPlaceholder = {{imp("errors")}}.New("error placeholder")

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

type lox struct {
	_lex   _Lexer
	_state _Stack[int32]
	_sym   _Stack[any]
	{{- if has_on_reduce }}
	_bounds _Stack[_Bounds]
	{{- end }}

	_lookahead     Token
	_lookaheadType int
	_errorToken    Token
}

func (p *{{parser}}) parse(lex _Lexer) bool {
  const accept = {{ accept }}

	p._lex = lex

	p._state.Push(0)
	p._ReadToken()

	for {
		if p._lookaheadType == ERROR {
			_, ok := p._Recover()
			if !ok {
				return false
			}
		}
		topState := p._state.Peek(0)
		action, ok := _Find(
			_Actions, topState, int32(p._lookaheadType))
		if !ok {
			action, ok = p._Recover()
			if !ok {
				return false
			}
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			p._state.Push(action)
			p._sym.Push(p._lookahead)
			{{- if has_on_reduce }}
			p._bounds.Push(
				_Bounds{Begin: p._lookahead,
				End: p._lookahead})
			{{- end }}
			p._ReadToken()
		} else { // reduce
			prod := -action
			termCount := _TermCounts[int(prod)]
			rule := _LHS[int(prod)]
			res := p._Act(prod)
			{{- if has_on_reduce }}
			if termCount > 0 {
				bounds := _Bounds{
					Begin: p._bounds.Peek(int(termCount-1)).Begin,
					End: p._bounds.Peek(0).End,
				}
				p.onReduce(res, bounds.Begin, bounds.End)
				p._bounds.Pop(int(termCount))
				p._bounds.Push(bounds)
			} else {
				bounds := p._bounds.Peek(0)
				bounds.Begin = bounds.End
				p._bounds.Push(bounds)
			}
			{{- end }}
			p._state.Pop(int(termCount))
			p._sym.Pop(int(termCount))
			topState = p._state.Peek(0)
			nextState, _ := _Find(_Goto, topState, rule)
			p._state.Push(nextState)
			p._sym.Push(res)
		}
	}

	return true
}

func (p *{{parser}}) errorToken() Token {
	return p._errorToken
}

func (p *{{parser}}) _ReadToken() {
	p._lookahead, p._lookaheadType = p._lex.ReadToken()
}

func (p *{{parser}}) _Recover() (int32, bool) {
	p._errorToken = p._lookahead

	for {
		for p._lookaheadType == ERROR {
			p._ReadToken()
		}

		saveState := p._state
		saveSym := p._sym
		{{- if has_on_reduce }}
			saveBounds := p._bounds
		{{- end }}

		for len(p._state) > 1 {
			topState := p._state.Peek(0)
			action, ok := _Find(_Actions, topState, int32(ERROR))
			if ok {
				action2, ok := _Find(
					_Actions, action, int32(p._lookaheadType))
				if ok {
					p._state.Push(action)
					p._sym.Push(_errorPlaceholder)
					{{- if has_on_reduce }}
					  p._bounds.Push(_Bounds{})
					{{- end }}
					return action2, true
				}
			}
			p._state.Pop(1)
			p._sym.Pop(1)
			{{- if has_on_reduce }}
				p._bounds.Pop(1)
			{{- end }}
		}

		if p._lookaheadType == EOF {
			p.onError()
			return 0, false
		}

		p._ReadToken()
		p._state = saveState
		p._sym = saveSym
		{{- if has_on_reduce }}
		p._bounds = saveBounds
		{{- end }}
	}
}

func (p *{{parser}}) _Act(prod int32) any {
	switch prod {
{{- range prod_index, prod := grammar.Prods }}
	{{- rule := prod.Rule }}
	{{- generated := rule_generated(rule) }}
	{{- if generated == "not_generated" }}
		{{- method := methods[prod] }}
			case {{ prod_index }}:
				return p.{{ method.Name() }}(
				{{- range param_index, param := method.Params }}
				  _cast[{{ go_type(param) }}](p._sym.Peek({{ len(method.Params) - param_index - 1 }})),
				{{- end }}
		    )
	{{- else if generated == "one_or_more" }}
	case {{ prod_index }}:  // OneOrMore
		{{- if len(prod.Terms) == 1 }}
			{{- term_go_type := go_type(get_term_go_type(prod.Terms[0])) }}
		  return []{{ term_go_type }}{
				_cast[{{ term_go_type }}](p._sym.Peek(0)),
			}
		{{- else }}
			{{- term_go_type := go_type(get_term_go_type(prod.Terms[1])) }}
			return append(
				_cast[[]{{term_go_type}}](p._sym.Peek(1)),
				_cast[{{ term_go_type }}](p._sym.Peek(0)),
			)
		{{- end }}
	{{- else if generated == "list" }}
	case {{ prod_index }}:  // List
		{{- if len(prod.Terms) == 1 }}
			{{- term_go_type := go_type(get_term_go_type(prod.Terms[0])) }}
		  return []{{ term_go_type }}{
				_cast[{{ term_go_type }}](p._sym.Peek(0)),
			}
		{{- else }}
			{{- term_go_type := go_type(get_term_go_type(prod.Terms[2])) }}
			return append(
				_cast[[]{{ term_go_type }}](p._sym.Peek(2)),
				_cast[{{ term_go_type }}](p._sym.Peek(0)),
			)
		{{- end }}
	{{- else if generated == "zero_or_one" }}
  case {{ prod_index }}:  // ZeroOrOne
		{{- term_go_type := go_type(rule_go_types[rule]) }}
		{{- if len(prod.Terms) == 1 }}
			return _cast[{{ term_go_type }}](p._sym.Peek(0))
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
			return _cast[{{ term_go_type }}](p._sym.Peek(0))
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
	vars.Set("has_on_reduce", c.HasOnReduce)

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
				case lr2.ActionShift:
					row = append(row, int32(action.ShiftState.Index))
				case lr2.ActionReduce:
					row = append(row, int32(action.Prods[0].Index)*-1)
				case lr2.ActionAccept:
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
				rule, ok := input.(*lr2.Rule)
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
