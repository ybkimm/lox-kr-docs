package codegen2

import (
	"os"
	"path/filepath"
)

func (c *context) EmitBase() bool {
	baseGen := renderBaseTemplate(&baseTemplateInputs{
		Package:     c.GoPackageName,
		PackagePath: c.GoPackagePath,
		Terminals:   c.ParserGrammar.Terminals,
	})

	err := os.WriteFile(filepath.Join(c.Dir, baseGenGo), []byte(baseGen), 0666)
	if err != nil {
		c.Errs.GeneralError(err)
		return false
	}

	return true
}
