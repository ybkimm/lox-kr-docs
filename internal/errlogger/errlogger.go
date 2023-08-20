package errlogger

import (
	"fmt"
	gotoken "go/token"
	"os"
)

type ErrLogger struct {
	hasErrors bool
}

func New() *ErrLogger {
	return &ErrLogger{}
}

func (l *ErrLogger) HasError() bool {
	return l.hasErrors
}

func (l *ErrLogger) Errorf(pos gotoken.Position, msg string, args ...any) {
	l.hasErrors = true
	msg = fmt.Sprintf(msg, args...)
	fmt.Fprintf(os.Stderr, "%v: %v\n", pos.String(), msg)
}

func (l *ErrLogger) Infof(pos gotoken.Position, msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)
	fmt.Fprintf(os.Stderr, "%v: %v\n", pos.String(), msg)
}

func (l *ErrLogger) Error(pos gotoken.Pos, err error) {
	l.hasErrors = true
	fmt.Fprintf(os.Stderr, "%v\n", err.Error())
}
