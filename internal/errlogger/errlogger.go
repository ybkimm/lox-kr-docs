package errlogger

import (
	"fmt"

	"github.com/dcaiafa/lox/internal/fileloc"
)

type Error struct {
	Loc     fileloc.FileLoc
	Desc    string
	Details []*Detail
}

func (e *Error) Detailf(msg string, args ...any) *Error {
	det := &Detail{
		Desc: fmt.Sprintf(msg, args...),
	}
	e.Details = append(e.Details, det)
	return e
}

type Detail struct {
	Loc  fileloc.FileLoc
	Desc string
}

type ErrLogger struct {
	Errors []*Error
}

func New() *ErrLogger { return &ErrLogger{} }

func (l *ErrLogger) Errorf(msg string, args ...any) *Error {
	err := &Error{
		Desc: fmt.Sprintf(msg, args...),
	}
	l.Errors = append(l.Errors, err)
	return err
}

func (l *ErrLogger) HasErrors() bool {
	return len(l.Errors) != 0
}
