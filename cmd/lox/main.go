package main

import (
	"flag"
	"fmt"
	gotoken "go/token"
	"os"

	"github.com/dcaiafa/lox/internal/codegen"
	"github.com/dcaiafa/lox/internal/errlogger"
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

	errLogger := errlogger.New()

	grammar := codegen.ParseGrammar(fset, dir, errLogger)
	if errLogger.HasError() {
		return fmt.Errorf("failed to parse grammar")
	}

	state := codegen.NewParserGenState(dir, grammar, errLogger)

	state.ConstructParseTables()
	if *flagAnalyze {
		state.ParserTable.Print(os.Stdout)
		return nil
	}

	lexerState := codegen.NewLexerGenState(dir, grammar)
	err := lexerState.Generate()
	if err != nil {
		return err
	}

	state.ParseGo()
	if errLogger.HasError() {
		return fmt.Errorf("failed to parse Go package")
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
