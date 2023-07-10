package codegen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
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

func (s *State) Generate() error {
	s.imports = newImportBuilder()
	s.writeProdTables()
	s.writeActionTable()
	s.writeGotoTable()
	s.writeOwl()

	out := bytes.NewBuffer(make([]byte, 0, s.body.Len()+2048))
	fmt.Fprintf(out, "package %v\n", s.PackageName)
	s.body.WriteTo(out)
	err := os.WriteFile(filepath.Join(s.ImplDir, loxGenGoName), out.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}

func (s *State) writeProdTables() {
	prods := make([]int32, len(s.Grammar.Prods))
	for prodIndex, prod := range s.Grammar.Prods {
		rule := s.Grammar.ProdRule(prod)
		prods[prodIndex] = int32(s.Grammar.RuleIndex(rule))
	}
	fmt.Fprintf(&s.body, "var loxProdRule = []int32 {\n")
	writeArray(&s.body, prods)
	fmt.Fprintf(&s.body, "}\n")

	for prodIndex, prod := range s.Grammar.Prods {
		prods[prodIndex] = int32(len(prod.Terms))
	}
	fmt.Fprintf(&s.body, "var loxProdTerms = []int32 {\n")
	writeArray(&s.body, prods)
	fmt.Fprintf(&s.body, "}\n")
}

func (s *State) writeActionTable() {
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

	fmt.Fprintf(&s.body, "var loxAction = []int32 {\n")
	writeArray(&s.body, actionRows.ToArray())
	fmt.Fprintf(&s.body, "}\n")

	actionIndexInt := make([]int32, len(actionIndex))
	for stateIndex, row := range actionIndex {
		actionIndexInt[stateIndex] = int32(row.Index)
	}

	fmt.Fprintf(&s.body, "var loxActionIndex = []int32 {\n")
	writeArray(&s.body, actionIndexInt)
	fmt.Fprintf(&s.body, "}\n")
}

func (s *State) writeGotoTable() {
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

	fmt.Fprintf(&s.body, "var loxGoto = []int32 {\n")
	writeArray(&s.body, gotoRows.ToArray())
	fmt.Fprintf(&s.body, "}\n")

	gotoIndexInt := make([]int32, len(gotoIndex))
	for stateIndex, row := range gotoIndex {
		if row != nil {
			gotoIndexInt[stateIndex] = int32(row.Index)
		} else {
			gotoIndexInt[stateIndex] = -1
		}
	}

	fmt.Fprintf(&s.body, "var loxGotoIndex = []int32 {\n")
	writeArray(&s.body, gotoIndexInt)
	fmt.Fprintf(&s.body, "}\n")
}

func (s *State) writeOwl() {
	data := map[string]any{
		"accept": accept,
	}
	err := loxGenGoTemplate.Execute(&s.body, data)
	if err != nil {
		panic(err)
	}
}

var loxGenGoTemplate = template.Must(template.New("lox").Parse(`

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
	fmt.Println("ERROR:", err)
	os.Exit(1)
}

`))
