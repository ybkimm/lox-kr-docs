package grammar

import (
	"errors"
	"fmt"
)

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
