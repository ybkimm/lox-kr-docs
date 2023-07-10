package codegen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

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
	s.writeActionTable()

	out := bytes.NewBuffer(make([]byte, 0, s.body.Len()+2048))
	fmt.Fprintf(out, "package %v\n", s.PackageName)
	s.body.WriteTo(out)
	err := os.WriteFile(filepath.Join(s.ImplDir, loxGenGoName), out.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
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
