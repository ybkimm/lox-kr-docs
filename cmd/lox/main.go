package main

import (
	"flag"
	"fmt"
	gotoken "go/token"
	"os"
	"runtime/pprof"

	"github.com/dcaiafa/lox/internal/codegen2"
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
		//flagAnalyze = flag.Bool("analyze", false, "")
		flagProf = flag.String("cpu-prof", "", "")
	)

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		return fmt.Errorf("<path> required")
	}
	dir := flag.Arg(0)

	if *flagProf != "" {
		f, err := os.Create(*flagProf)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	fset := gotoken.NewFileSet()
	errs := errlogger.New()

	ok := codegen2.Generate(&codegen2.Config{
		Fset: fset,
		Errs: errs,
		Dir:  dir,
	})
	if !ok {
		return fmt.Errorf("errors ocurred")
	}

	return nil
}
