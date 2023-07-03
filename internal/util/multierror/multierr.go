package multierror

import (
	"fmt"
	"strings"
)

type MultiError []error

func (e *MultiError) Add(err error) {
	*e = append(*e, err)
}

func (e MultiError) Error() string {
	str := new(strings.Builder)
	fmt.Fprintf(str, "errors ocurred:")
	for _, err := range e {
		fmt.Fprintf(str, "\n  %v", err.Error())
	}
	return str.String()
}

func (e MultiError) ToError() error {
	if len(e) == 0 {
		return nil
	}
	return e
}
