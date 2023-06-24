package grammar

type Qualifier int

const (
	NoQualifier Qualifier = iota
	ZeroOrMore            // *
	OneOrMore             // +
	ZeroOrOne             // ?
)

type Syntax struct {
	Productions []*Production
}

type Production struct {
	Name  string
	Terms []*Term
}

type Term struct {
	Factors []*Factor
	Label   *Label
}

type Factor struct {
	Name      string
	Literal   string
	Qualifier Qualifier
}

type Label struct {
	Label string
}
