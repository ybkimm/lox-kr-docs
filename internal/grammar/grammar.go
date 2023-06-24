package grammar

type Qualifier int

const (
	NoQualifier Qualifier = iota
	ZeroOrMore            // *
	OneOrMore             // +
	ZeroOrOne             // ?
)

type Syntax struct {
	Rules []*Rule
}

type Rule struct {
	Name        string
	Productions []*Production
}

type Production struct {
	Terms []*Term
	Label *Label
}

type Term struct {
	Name      string
	Literal   string
	Qualifier Qualifier
}

type Label struct {
	Label string
}
