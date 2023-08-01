package errlogger

import (
	"fmt"
	gotoken "go/token"
	"os"
)

type ErrLogger struct {
	fset      *gotoken.FileSet
	hasErrors bool
}

func New(fset *gotoken.FileSet) *ErrLogger {
	return &ErrLogger{
		fset: fset,
	}
}

func (l *ErrLogger) HasError() bool {
	return l.hasErrors
}

func (l *ErrLogger) Error(pos gotoken.Pos, err error) {
	l.hasErrors = true
	position := l.fset.Position(pos)
	fmt.Fprintf(os.Stderr, "%v: %v\n", position.String(), err.Error())
}

func (l *ErrLogger) Info(pos gotoken.Pos, err error) {
	position := l.fset.Position(pos)
	fmt.Fprintf(os.Stderr, "%v: %v\n", position.String(), err.Error())
}
