package parsergen

import (
	"errors"
	"fmt"
	"strings"
)

type Errors []error

func (e Errors) Error() string {
	str := new(strings.Builder)
	fmt.Fprintf(str, "errors ocurred:")
	for _, err := range e {
		fmt.Fprintf(str, "\n  %v", err.Error())
	}
	return str.String()
}

type RedeclaredError struct {
	Sym   Symbol
	Other Symbol
}

func (e *RedeclaredError) Error() string {
	return fmt.Sprintf("%q redeclared", e.Sym.SymName())
}

type UndefinedError struct {
	Term *Term
	Prod *Prod
	Rule *Rule
}

func (e *UndefinedError) Error() string {
	return fmt.Sprintf("undefined: %s/%s", e.Rule.Name, e.Term.Name)
}

var ErrConflict = errors.New("grammar has conflict(s). Checkout spec for details")
