package codegen

import (
	"bytes"
	"fmt"
	"runtime"

	"github.com/CloudyKit/jet/v6"
)

func assert(p bool) {
	if !p {
		msg := "assertion failed"
		// Include information about the assertion location. Due to panic recovery,
		// this location is otherwise buried in the middle of the panicking stack.
		if _, file, line, ok := runtime.Caller(1); ok {
			msg = fmt.Sprintf("%s:%d: %s", file, line, msg)
		}
		panic(msg)
	}
}

func unreachable() {
	panic("unreachable")
}

func renderTemplate(templ string, vars jet.VarMap) string {
	loader := jet.NewInMemLoader()
	loader.Set("lox", templ)

	set := jet.NewSet(loader, jet.WithSafeWriter(nil))
	t, err := set.GetTemplate("lox")
	if err != nil {
		panic(err)
	}

	body := &bytes.Buffer{}
	err = t.Execute(body, vars, nil)
	if err != nil {
		panic(err)
	}

	return body.String()
}
