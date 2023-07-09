package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dcaiafa/lox/internal/codegen"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	dir := flag.Arg(0)

	err := realMain(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func realMain(dir string) error {
	_, err := codegen.Parse(dir)
	return err
}
