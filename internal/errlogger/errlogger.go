package errlogger

import (
	"fmt"
	gotoken "go/token"
	"os"
	"path/filepath"
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

func (l *ErrLogger) GeneralError(err error) {
	l.hasErrors = true
	fmt.Fprintf(os.Stderr, "%v\n", err.Error())
}

func (l *ErrLogger) GeneralErrorf(msg string, args ...any) {
	l.hasErrors = true
	msg = fmt.Sprintf(msg, args...)
	fmt.Fprintf(os.Stderr, "%v\n", msg)
}

func (l *ErrLogger) Errorf(pos gotoken.Position, msg string, args ...any) {
	l.hasErrors = true
	msg = fmt.Sprintf(msg, args...)
	fmt.Fprintf(os.Stderr, "%v: %v\n", rel(pos).String(), msg)
}

func (l *ErrLogger) Infof(pos gotoken.Position, msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)
	fmt.Fprintf(os.Stderr, "%v: %v\n", rel(pos).String(), msg)
}

func rel(pos gotoken.Position) gotoken.Position {
	if pos.Filename == "" {
		return pos
	}
	localDir, err := os.Getwd()
	if err != nil {
		return pos
	}
	relFilename, err := filepath.Rel(localDir, pos.Filename)
	if err != nil {
		return pos
	}
	pos.Filename = relFilename
	return pos
}
