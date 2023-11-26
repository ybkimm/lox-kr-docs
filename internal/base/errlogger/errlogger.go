package errlogger

import (
	"fmt"
	gotoken "go/token"
	"io"
	"os"
	"path/filepath"
)

type ErrLogger struct {
	hasErrors bool
	fset      *gotoken.FileSet
	out       io.Writer
}

func New(fset *gotoken.FileSet, out io.Writer) *ErrLogger {
	return &ErrLogger{
		fset: fset,
		out:  out,
	}
}

func (l *ErrLogger) Output() io.Writer {
	return l.out
}

func (l *ErrLogger) HasError() bool {
	return l.hasErrors
}

func (l *ErrLogger) GeneralError(err error) {
	l.hasErrors = true
	_, _ = fmt.Fprintf(l.out, "%v\n", err.Error())
}

func (l *ErrLogger) GeneralErrorf(msg string, args ...any) {
	l.hasErrors = true
	msg = fmt.Sprintf(msg, args...)
	_, _ = fmt.Fprintf(l.out, "%v\n", msg)
}

func (l *ErrLogger) Errorf(pos gotoken.Pos, msg string, args ...any) {
	l.Errorpf(l.fset.Position(pos), msg, args...)
}

func (l *ErrLogger) Errorpf(pos gotoken.Position, msg string, args ...any) {
	l.hasErrors = true
	msg = fmt.Sprintf(msg, args...)
	_, _ = fmt.Fprintf(l.out, "%v: %v\n", rel(pos).String(), msg)
}

func (l *ErrLogger) Infof(pos gotoken.Pos, msg string, args ...any) {
	l.Infopf(l.fset.Position(pos), msg, args...)
}

func (l *ErrLogger) Infopf(pos gotoken.Position, msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)
	_, _ = fmt.Fprintf(l.out, "%v: %v\n", rel(pos).String(), msg)
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
