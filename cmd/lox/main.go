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

	cfg := &codegen.Config{
		Errs:    errLogger,
		ImplDir: dir,
		Grammar: grammar,
	}

	if *flagAnalyze {
		cfg.AnalysisWriter = os.Stdout
		cfg.AnalysisOnly = true
	}

	codegen.Generate(cfg)

	if cfg.Errs.HasError() {
		return fmt.Errorf("errors ocurred")
	}

	return nil
}
