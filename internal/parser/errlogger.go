package parser

import (
	"fmt"
	gotoken "go/token"
)

type ErrLogger struct {
	Fset      *gotoken.FileSet
	HasErrors bool
}

func (l *ErrLogger) Error(pos gotoken.Pos, err error) {
	position := l.Fset.Position(pos)
	fmt.Printf("%v: %v\n", position.String(), err.Error())
}
