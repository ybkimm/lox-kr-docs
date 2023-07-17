package codegen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/CloudyKit/jet/v6"
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

func (s *State) Generate2() error {
	s.imports = newImportBuilder()

	vars := jet.VarMap{}
	vars.Set("accept", accept)
	vars.Set("imp", s.templImport)
	vars.Set("p", "_lx")
	vars.Set("actions", s.templActions)
	vars.Set("array", s.templArray)
	vars.Set("goto", s.templGoto)
	vars.Set("lhs", s.templLHS)
	vars.Set("term_counts", s.templTermCounts)

	body := renderTemplate(loxGenJet, vars)

	out := bytes.NewBuffer(make([]byte, 0, len(body)+2048))
	fmt.Fprintf(out, "package %v\n\n", s.PackageName)
	s.imports.WriteTo(out)
	out.WriteString(body)
	err := os.WriteFile(filepath.Join(s.ImplDir, parserGenGoName), out.Bytes(), 0666)
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

func (s *State) templLHS() []int32 {
	prods := make([]int32, len(s.Grammar.Prods))
	for prodIndex, prod := range s.Grammar.Prods {
		rule := s.Grammar.ProdRule(prod)
		prods[prodIndex] = int32(s.Grammar.RuleIndex(rule))
	}
	return prods
}

func (s *State) templArray(arr []int32) string {
	var str strings.Builder
	table.WriteArray(&str, arr)
	return str.String()
}

func (s *State) templTermCounts() []int32 {
	prods := make([]int32, len(s.Grammar.Prods))
	for prodIndex, prod := range s.Grammar.Prods {
		prods[prodIndex] = int32(len(prod.Terms))
	}
	return prods
}

func (s *State) templActions() []int32 {
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
	return actionTable.Array()
}

func (s *State) templGoto() []int32 {
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
	return gotoTable.Array()
}

const loxGenJet = `
var {{p}}LHS = []int32 {
	{{ lhs() | array }}
}

var {{p}}TermCounts = []int32 {
	{{ term_counts() | array }}	
}

var {{p}}Actions = []int32 {
	{{ actions() | array }}
}

var {{p}}Goto = []int32 {
	{{ goto() | array }}
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
  const accept = {{ accept }}

	p.loxParser.state.Push(0)
	lookahead, tok := lex.NextToken()

	for {
		topState := p.loxParser.state.Peek(0)
		action, ok := {{p}}Find({{p}}Actions, topState, int32(lookahead))
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
			termCount := {{p}}TermCounts[int(prod)]
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
	{{ imp("fmt") }}.Println("ERROR:", err)
	{{ imp("os") }}.Exit(1)
}

`
