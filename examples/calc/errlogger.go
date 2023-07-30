package main

import (
	"errors"
	"fmt"
	gotoken "go/token"
	"strings"
)

type ErrLogger struct {
	Fset *gotoken.FileSet
	str  strings.Builder
}

func (l *ErrLogger) Error(pos gotoken.Pos, err error) {
	if l.str.Len() > 0 {
		l.str.WriteString("\n")
	}
	position := l.Fset.Position(pos)
	if position.IsValid() {
		fmt.Fprintf(&l.str, "%v:%v: ", position.Line, position.Column)
	}
	l.str.WriteString(err.Error())
}

func (l *ErrLogger) Err() error {
	if l.str.Len() == 0 {
		return nil
	}
	return errors.New(l.str.String())
}
