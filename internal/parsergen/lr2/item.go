package lr2

import (
	"fmt"
	"strings"
)

type Item struct {
	Prod      int
	Dot       int
	Lookahead int
}

func (i Item) IsKernel() bool {
	return i.Prod == sprimeProd || i.Dot != 0
}

func (i Item) ToString(g *Grammar) string {
	var str strings.Builder

	prod := g.prods[i.Prod]
	fmt.Fprintf(&str, "%v = ", prod.Rule.Name)
	for j, term := range prod.Terms {
		if j != 0 {
			str.WriteString(" ")
		}
		if j == i.Dot {
			str.WriteString(".")
		}
		str.WriteString(g.GetSymbolName(term))
	}
	if i.Dot == len(prod.Terms) {
		str.WriteString(".")
	}
	str.WriteString(", ")
	str.WriteString(g.GetSymbolName(i.Lookahead))
	return str.String()
}
