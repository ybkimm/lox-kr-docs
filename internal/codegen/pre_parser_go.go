package codegen

import (
	goparser "go/parser"
	gotoken "go/token"
	"os"
	"path/filepath"
)

func (c *context) PreParseGo() bool {
	dirEntries, err := os.ReadDir(c.Dir)
	if err != nil {
		c.Errs.GeneralError(err)
		return false
	}

	var oneSourceName string
	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() &&
			filepath.Ext(dirEntry.Name()) == ".go" &&
			dirEntry.Name() != lexerGenGo &&
			dirEntry.Name() != parserGenGo {
			oneSourceName = filepath.Join(c.Dir, dirEntry.Name())
		}
	}
	if oneSourceName == "" {
		c.Errs.GeneralErrorf("package contains no Go sources")
		return false
	}

	oneSource, err := goparser.ParseFile(
		gotoken.NewFileSet(), oneSourceName, nil, 0)
	if err != nil {
		c.Errs.GeneralError(err)
		return false
	}

	c.GoPackageName = oneSource.Name.Name
	return true
}
