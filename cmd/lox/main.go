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

	state := codegen.NewState(dir, dir)
	err := state.ParseGrammar()
	if err != nil {
		return err
	}

	if *flagTerminals {
		for _, terminal := range state.Grammar.Terminals {
			fmt.Println(terminal.Name)
		}
		return nil
	}

	state.ConstructParseTables()
	err = state.ParseGo()
	if err != nil {
		return err
	}

	state.Grammar.Print(os.Stdout)
	state.ParserTable.Print(os.Stdout)

	err = state.MapReduceActions()
	if err != nil {
		return err
	}
	err = state.Generate()
	if err != nil {
		return err
	}

	return nil
}
