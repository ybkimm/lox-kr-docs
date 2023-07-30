package main

import (
	"flag"
	"fmt"
	gotoken "go/token"
	"os"

	"github.com/dcaiafa/lox/internal/codegen"
	"github.com/dcaiafa/lox/internal/parser"
)

func main() {
	err := realMain()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func realMain() error {
	var (
		flagAnalyze = flag.Bool("analyze", false, "")
	)

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		return fmt.Errorf("<path> required")
	}
	dir := flag.Arg(0)

	fset := gotoken.NewFileSet()

	errLogger := &parser.ErrLogger{
		Fset: fset,
	}

	grammar, err := codegen.ParseGrammar(fset, dir, errLogger)
	if err != nil {
		return err
	}

	state := codegen.NewParserGenState(dir, grammar)

	state.ConstructParseTables()
	if *flagAnalyze {
		state.ParserTable.Print(os.Stdout)
		return nil
	}

	lexerState := codegen.NewLexerGenState(dir, grammar)
	err = lexerState.Generate()
	if err != nil {
		return err
	}

	err = state.ParseGo()
	if err != nil {
		return err
	}

	state.ParserTable.Print(os.Stdout)

	err = state.MapReduceActions()
	if err != nil {
		return err
	}
	err = state.Generate2()
	if err != nil {
		return err
	}

	return nil
}
