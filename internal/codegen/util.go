package codegen

import (
	"bytes"
	"fmt"
	goparser "go/parser"
	gotoken "go/token"
	"os"
	"path/filepath"

	"github.com/CloudyKit/jet/v6"
)

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

func computePackageName(dir string) (string, error) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	var oneSourceName string
	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() &&
			filepath.Ext(dirEntry.Name()) == ".go" &&
			dirEntry.Name() != parserGenGoName {
			oneSourceName = filepath.Join(dir, dirEntry.Name())
		}
	}
	if oneSourceName == "" {
		return "", fmt.Errorf("package contains no source files")
	}

	oneSource, err := goparser.ParseFile(gotoken.NewFileSet(), oneSourceName, nil, 0)
	if err != nil {
		return "", fmt.Errorf("%v: %w", oneSourceName, err)
	}

	return oneSource.Name.Name, nil
}
