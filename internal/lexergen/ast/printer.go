package ast

import (
	"fmt"
	"io"
)

type Printer struct {
	out    io.Writer
	indent int
}

func NewPrinter(out io.Writer) *Printer {
	return &Printer{
		out: out,
	}
}

func (p *Printer) WithIndent(n int) *Printer {
	np := new(Printer)
	*np = *p
	np.indent += n
	return np
}

func (p *Printer) Printf(s string, args ...any) {
	for i := 0; i < p.indent; i++ {
		fmt.Fprint(p.out, " ")
	}
	fmt.Fprintf(p.out, s+"\n", args...)
}
