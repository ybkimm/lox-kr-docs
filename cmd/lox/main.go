package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dcaiafa/lox/internal/errlogger"
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
	grammarData, err := ioutil.ReadFile(grammarFile)
	if err != nil {
		return err
	}
	errs := errlogger.New()
	parser.Parse(grammarFile, grammarData, errs)
	if errs.HasErrors() {
		return fmt.Errorf("errors ocurred")
	}
	return nil
}
