package parsergen

type Qualifier int

const (
	NoQualifier Qualifier = iota
	ZeroOrMore            // *
	OneOrMore             // +
	ZeroOrOne             // ?
)

type Syntax struct {
	rules     []*Rule
	terminals []*Terminal
	defs      map[string]any
}

type Rule struct {
	Name  string
	Prods []*Prod
}

type Prod struct {
	Terms []*Term
}

type Term struct {
	Name      string
	Qualifier Qualifier
}

type Terminal struct {
	Name string
}
