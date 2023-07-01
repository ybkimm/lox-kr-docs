package parsergen

import (
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
	Def   Def
	Other Def
}

func (e *RedeclaredError) Error() string {
	return fmt.Sprintf("%q redeclared", e.Def.DefName())
}

type UndefinedError struct {
	Term *Term
	Prod *Prod
	Rule *Rule
}

func (e *UndefinedError) Error() string {
	return fmt.Sprintf("undefined: %s/%s", e.Rule.Name, e.Term.Name)
}
