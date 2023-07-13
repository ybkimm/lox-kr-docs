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
			"loxProdRule":  s.templLoxProdRule,
			"loxProdTerms": s.templLoxProdTerms,
			"loxAction":    s.templLoxAction,
			"loxGoto":      s.templLoxGoto,
			"import":       s.templImport,
			"terminals":    s.templTerminals,
		}).
		Parse(loxGenTemplate)
	if err != nil {
		panic(err)
	}

	body := &bytes.Buffer{}
	err = templ.Execute(body, map[string]any{
		"accept":    math.MaxInt32,
		"terminals": s.Grammar.Terminals,
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

func (s *State) templTerminals() string {
	str := new(strings.Builder)

	fmt.Fprintf(str, "const (\n")
	for terminalIndex, terminal := range s.Grammar.Terminals {
		fmt.Fprintf(str, "  %v = %v\n", terminal.Name, terminalIndex)
	}
	fmt.Fprintf(str, ")\n")

	return str.String()
}

func (s *State) templImport(path string) string {
	return s.imports.Import(path)
}

func (s *State) templLoxProdRule() string {
	str := new(strings.Builder)
	prods := make([]int32, len(s.Grammar.Prods))
	for prodIndex, prod := range s.Grammar.Prods {
		rule := s.Grammar.ProdRule(prod)
		prods[prodIndex] = int32(s.Grammar.RuleIndex(rule))
	}
	fmt.Fprintf(str, "var loxProdRule = []int32 {\n")
	writeArray(str, prods)
	fmt.Fprintf(str, "}\n")
	return str.String()
}

func (s *State) templLoxProdTerms() string {
	str := new(strings.Builder)
	prods := make([]int32, len(s.Grammar.Prods))
	for prodIndex, prod := range s.Grammar.Prods {
		prods[prodIndex] = int32(len(prod.Terms))
	}
	fmt.Fprintf(str, "var loxProdTerms = []int32 {\n")
	writeArray(str, prods)
	fmt.Fprintf(str, "}\n")
	return str.String()
}

func (s *State) templLoxAction() string {
	str := new(strings.Builder)
	actionRows := newRows()
	actionIndex := make([]*row, len(s.ParserTable.States.States()))
	for i, state := range s.ParserTable.States.States() {
		row := new(row)
		s.ParserTable.Actions.ForEachActionSet(s.Grammar, state,
			func(sym grammar.Symbol, actions []lr1.Action) {
				action := actions[0]
				terminal := sym.(*grammar.Terminal)
				terminalIndex := s.Grammar.TerminalIndex(terminal)
				row.Add(int32(terminalIndex))

				switch action.Type {
				case lr1.ActionShift:
					row.Add(int32(action.Shift.Index))
				case lr1.ActionReduce:
					row.Add(int32(s.Grammar.ProdIndex(action.Prod) * -1))
				case lr1.ActionAccept:
					row.Add(accept)
				default:
					panic("unreachable")
				}
			})
		row = actionRows.Add(row)
		actionIndex[i] = row
	}

	fmt.Fprintf(str, "var loxAction = []int32 {\n")
	writeArray(str, actionRows.ToArray())
	fmt.Fprintf(str, "}\n")

	actionIndexInt := make([]int32, len(actionIndex))
	for stateIndex, row := range actionIndex {
		actionIndexInt[stateIndex] = int32(row.Index)
	}

	fmt.Fprintf(str, "var loxActionIndex = []int32 {\n")
	writeArray(str, actionIndexInt)
	fmt.Fprintf(str, "}\n")

	return str.String()
}

func (s *State) templLoxGoto() string {
	str := new(strings.Builder)

	gotoRows := newRows()
	gotoIndex := make([]*row, len(s.ParserTable.States.States()))
	for stateIndex, state := range s.ParserTable.States.States() {
		row := new(row)
		s.ParserTable.Transitions.ForEach(
			state,
			func(sym grammar.Symbol, to *lr1.ItemSet) {
				rule, ok := sym.(*grammar.Rule)
				if !ok {
					return
				}
				ruleIndex := s.Grammar.RuleIndex(rule)
				row.Add(int32(ruleIndex))
				row.Add(int32(to.Index))
			})
		if len(row.Cols) == 0 {
			continue
		}
		row = gotoRows.Add(row)
		gotoIndex[stateIndex] = row
	}

	fmt.Fprintf(str, "var loxGoto = []int32 {\n")
	writeArray(str, gotoRows.ToArray())
	fmt.Fprintf(str, "}\n")

	gotoIndexInt := make([]int32, len(gotoIndex))
	for stateIndex, row := range gotoIndex {
		if row != nil {
			gotoIndexInt[stateIndex] = int32(row.Index)
		} else {
			gotoIndexInt[stateIndex] = -1
		}
	}

	fmt.Fprintf(str, "var loxGotoIndex = []int32 {\n")
	writeArray(str, gotoIndexInt)
	fmt.Fprintf(str, "}\n")

	return str.String()
}

const loxGenTemplate = `

{{ terminals }}

{{ loxProdRule }}

{{ loxProdTerms }}

{{ loxAction }}

{{ loxGoto }}

const loxAccept = {{.accept}}

type loxStack[T any] []T

func (s *loxStack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *loxStack[T]) Pop(n int) {
	*s = (*s)[:len(*s)-n]
}

func (s loxStack[T]) Peek(n int) T {
	return s[len(s)-n-1]
}

func loxFind(index []int32, data []int32, s, k int32) (int32, bool) {
	i := index[int(s)]
	n := data[int(i)]
	for i = i+1; i < n; i+=2 {
		if data[i] == k {
			return data[i+1], true
		}
	}
	return 0, false
}

type loxLexer interface {
	Token() (int, Token)
}

type loxParser struct {
	state loxStack[int32]
	sym   loxStack[any]
}

func (p *Parser) parse(lex loxLexer) {
	p.loxParser.state.Push(0)
	lookahead, tok := lex.Token()

	for {
		topState := p.loxParser.state.Peek(0)
		action, ok := loxFind(loxActionIndex, loxAction, topState, int32(lookahead))
		if !ok {
    	p.onError(tok, "boom")
			return
		}
		if action == loxAccept {
			break
		} else if action >= 0 { // shift
			p.loxParser.state.Push(action)
			p.loxParser.sym.Push(tok)
		} else if action < 0 { // reduce
			prod := -action
			terms := loxProdTerms[int(prod)]
			rule := loxProdRule[int(prod)]
			p.loxParser.state.Pop(int(terms))
			p.loxParser.sym.Pop(int(terms))
			topState = p.loxParser.state.Peek(0)
			nextState, _ := loxFind(loxGotoIndex, loxGoto, topState, rule)
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
