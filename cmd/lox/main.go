package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dcaiafa/lox/internal/codegen"
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
		flagTerminals = flag.Bool("terminals", false, "list terminals and quit")
	)
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		return fmt.Errorf("<path> required")
	}
	dir := flag.Arg(0)

	grammar, err := codegen.ParseGrammar(dir)
	if err != nil {
		return err
	}

	state := codegen.NewParserGenState(dir, grammar)

	if *flagTerminals {
		for _, terminal := range state.Grammar.Terminals {
			fmt.Println(terminal.Name)
		}
		return nil
	}

	state.ConstructParseTables()
	state.Grammar.Print(os.Stdout)

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
