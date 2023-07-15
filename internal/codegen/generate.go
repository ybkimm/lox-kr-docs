package codegen

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/dcaiafa/lox/internal/codegen/table"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
)

type importBuilder struct {
	buf bytes.Buffer

	imports map[string]string
}

func newImportBuilder() *importBuilder {
	return &importBuilder{
		imports: make(map[string]string),
	}
}

func (b *importBuilder) Import(path string) string {
	alias, ok := b.imports[path]
	if ok {
		return alias
	}
	alias = fmt.Sprintf("_i%d", len(b.imports))
	b.imports[alias] = path
	return alias
}

func (b *importBuilder) WriteTo(w *bytes.Buffer) {
	fmt.Fprintf(w, "import (\n")
	aliases := make([]string, 0, len(b.imports))
	for alias := range b.imports {
		aliases = append(aliases, alias)
	}
	sort.Strings(aliases)
	for _, alias := range aliases {
		path := b.imports[alias]
		fmt.Fprintf(w, "  %v %q\n", alias, path)
	}
	fmt.Fprintf(w, ")\n")
}

func (s *State) Generate() error {
	s.imports = newImportBuilder()

	templ, err := template.New("lox").
		Funcs(map[string]any{
			"p":         func() string { return "_lx" },
			"lhs":       s.templLHS,
			"reduction": s.templReduction,
			"action":    s.templAction,
			"goto":      s.templGoto,
			"import":    s.templImport,
		}).
		Parse(loxGenTemplate)
	if err != nil {
		panic(err)
	}

	body := &bytes.Buffer{}
	err = templ.Execute(body, map[string]any{
		"accept":    math.MaxInt32,
		"terminals": s.terminals(),
	})
	if err != nil {
		panic(err)
	}

	out := bytes.NewBuffer(make([]byte, 0, body.Len()+2048))
	fmt.Fprintf(out, "package %v\n\n", s.PackageName)
	s.imports.WriteTo(out)
	body.WriteTo(out)
	err = os.WriteFile(filepath.Join(s.ImplDir, loxGenGoName), out.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}

func (s *State) terminals() map[int]string {
	terminals := make(map[int]string)
	for i, terminal := range s.Grammar.Terminals {
		terminals[i] = terminal.Name
	}
	return terminals
}

func (s *State) templImport(path string) string {
	return s.imports.Import(path)
}

func (s *State) templLHS() string {
	str := new(strings.Builder)
	prods := make([]int32, len(s.Grammar.Prods))
	for prodIndex, prod := range s.Grammar.Prods {
		rule := s.Grammar.ProdRule(prod)
		prods[prodIndex] = int32(s.Grammar.RuleIndex(rule))
	}
	table.WriteArray(str, prods)
	return str.String()
}

func (s *State) templReduction() string {
	str := new(strings.Builder)
	prods := make([]int32, len(s.Grammar.Prods))
	for prodIndex, prod := range s.Grammar.Prods {
		prods[prodIndex] = int32(len(prod.Terms))
	}
	table.WriteArray(str, prods)
	return str.String()
}

func (s *State) templAction() string {
	actionTable := table.New()
	for _, state := range s.ParserTable.States.States() {
		var row []int32
		s.ParserTable.Actions.ForEachActionSet(s.Grammar, state,
			func(sym grammar.Symbol, actions []lr1.Action) {
				action := actions[0]
				terminal := sym.(*grammar.Terminal)
				terminalIndex := s.Grammar.TerminalIndex(terminal)
				row = append(row, int32(terminalIndex))

				switch action.Type {
				case lr1.ActionShift:
					row = append(row, int32(action.Shift.Index))
				case lr1.ActionReduce:
					row = append(row, int32(s.Grammar.ProdIndex(action.Prod)*-1))
				case lr1.ActionAccept:
					row = append(row, accept)
				default:
					panic("unreachable")
				}
			})
		actionTable.AddRow(state.Index, row)
	}
	return actionTable.String()
}

func (s *State) templGoto() string {
	gotoTable := table.New()
	for stateIndex, state := range s.ParserTable.States.States() {
		var row []int32
		s.ParserTable.Transitions.ForEach(
			state,
			func(sym grammar.Symbol, to *lr1.ItemSet) {
				rule, ok := sym.(*grammar.Rule)
				if !ok {
					return
				}
				ruleIndex := s.Grammar.RuleIndex(rule)
				row = append(row, int32(ruleIndex), int32(to.Index))
			})
		gotoTable.AddRow(stateIndex, row)
	}
	return gotoTable.String()
}

const loxGenTemplate = `

const (
{{- range $index, $name := .terminals }}
  {{ $name }} = {{ $index }}
{{- end}}
)

var {{p}}LHS = []int32 {
	{{ lhs }} 
}

var {{p}}Reduction = []int32 {
	{{ reduction }}
}

var {{p}}Action = []int32 {
	{{ action }}
}

var {{p}}Goto = []int32 {
	{{ goto }}
}

type {{p}}Stack[T any] []T

func (s *{{p}}Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *{{p}}Stack[T]) Pop(n int) {
	*s = (*s)[:len(*s)-n]
}

func (s {{p}}Stack[T]) Peek(n int) T {
	return s[len(s)-n-1]
}

func {{p}}Find(table []int32, y, x int32) (int32, bool) {
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

type {{p}}Lexer interface {
	NextToken() (int, Token)
}

type loxParser struct {
	state {{p}}Stack[int32]
	sym   {{p}}Stack[any]
}

func (p *Parser) parse(lex {{p}}Lexer) {
  const accept = {{ .accept }}

	p.loxParser.state.Push(0)
	lookahead, tok := lex.NextToken()

	for {
		topState := p.loxParser.state.Peek(0)
		action, ok := {{p}}Find({{p}}Action, topState, int32(lookahead))
		if !ok {
			p.onError(tok, "boom")
			return
		}
		if action == accept {
    	break
		} else if action >= 0 { // shift
			p.loxParser.state.Push(action)
			p.loxParser.sym.Push(tok)
			lookahead, tok = lex.NextToken()
		} else { // reduce
			prod := -action
			termCount := {{p}}Reduction[int(prod)]
			rule := {{p}}LHS[int(prod)]
			p.loxParser.state.Pop(int(termCount))
			p.loxParser.sym.Pop(int(termCount))
			topState = p.loxParser.state.Peek(0)
			nextState, _ := {{p}}Find({{p}}Goto, topState, rule)
			p.loxParser.state.Push(nextState)
			p.loxParser.sym.Push(nil)
		}
	}
}

func (p *Parser) onError(tok Token, err string) {
	{{ import "fmt" }}.Println("ERROR:", err)
	{{ import "os" }}.Exit(1)
}

`
