package lr2

type ParserTable struct {
	Grammar *Grammar
	States  *StateSet
}

func NewParserTable(g *Grammar) *ParserTable {
	return &ParserTable{
		Grammar: g,
		States:  NewStateSet(),
	}
}
