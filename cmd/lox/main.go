package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dcaiafa/lox/internal/analysis"
	"github.com/dcaiafa/lox/internal/errs"
	"github.com/dcaiafa/lox/internal/parser"
)

func main() {
	flag.Parse()

	grammarFile := flag.Arg(0)
	err := run(grammarFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed: %v", err)
		os.Exit(1)
	}
}

func run(grammarFile string) error {
	grammarData, err := os.ReadFile(grammarFile)
	if err != nil {
		return err
	}
	errs := errs.New()
	syntax := parser.Parse(grammarFile, grammarData, errs)
	if errs.HasErrors() {
		errs.Dump(os.Stderr)
		return fmt.Errorf("errors ocurred")
	}
	analysis.Analyze(syntax, errs)
	if errs.HasErrors() {
		errs.Dump(os.Stderr)
		return fmt.Errorf("errors ocurred")
	}
	return nil
}
