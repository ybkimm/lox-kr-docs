package codegen2

import (
	"bytes"
	"fmt"
	"sort"
)

type imports struct {
	buf     bytes.Buffer
	imports map[string]string
}

func newImports() *imports {
	return &imports{
		imports: make(map[string]string),
	}
}

func (b *imports) Import(path string) string {
	alias, ok := b.imports[path]
	if ok {
		return alias
	}
	alias = fmt.Sprintf("_i%d", len(b.imports))
	b.imports[path] = alias
	return alias
}

func (b *imports) WriteTo(w *bytes.Buffer) {
	if len(b.imports) == 0 {
		return
	}

	fmt.Fprintf(w, "import (\n")
	paths := make([]string, 0, len(b.imports))
	for path := range b.imports {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		alias := b.imports[path]
		fmt.Fprintf(w, "  %v %q\n", alias, path)
	}
	fmt.Fprintf(w, ")\n")
}
