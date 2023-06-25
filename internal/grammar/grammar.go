package grammar

type Qualifier int

const (
	NoQualifier Qualifier = iota
	ZeroOrMore            // *
	OneOrMore             // +
	ZeroOrOne             // ?
)

type Decl interface {
	DeclName() string
}

type Syntax struct {
	Decls []Decl
}

type Rule struct {
	Name  string
	Prods []*Prod
}

func (r *Rule) DeclName() string { return r.Name }

type Prod struct {
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

type Token struct {
	Name string
}
