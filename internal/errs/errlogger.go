package errs

import (
	"fmt"
	"io"

	"github.com/dcaiafa/lox/internal/loc"
)

type Error struct {
	Loc     loc.Loc
	Desc    string
	Details []*Detail
}

func (e *Error) String() string {
	return fmt.Sprintf("%s:%v: %v", e.Loc.Filename, e.Loc.Line, e.Desc)
}

func (e *Error) Detailf(loc loc.Loc, msg string, args ...any) *Error {
	det := &Detail{
		Desc: fmt.Sprintf(msg, args...),
	}
	e.Details = append(e.Details, det)
	return e
}

type Detail struct {
	Loc  loc.Loc
	Desc string
}

func (d *Detail) String() string {
	return fmt.Sprintf("%s:%v: %v", d.Loc.Filename, d.Loc.Line, d.Desc)
}

type Errs struct {
	Errors []*Error
}

func New() *Errs { return &Errs{} }

func (l *Errs) Errorf(loc loc.Loc, msg string, args ...any) *Error {
	err := &Error{
		Loc:  loc,
		Desc: fmt.Sprintf(msg, args...),
	}
	l.Errors = append(l.Errors, err)
	return err
}

func (l *Errs) HasErrors() bool {
	return len(l.Errors) != 0
}

func (l *Errs) Dump(w io.Writer) {
	for _, e := range l.Errors {
		fmt.Fprintln(w, e.String())
		for _, d := range e.Details {
			fmt.Fprintln(w, "    ", d.String())
		}
	}
}
