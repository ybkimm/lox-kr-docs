package main

import (
	"bytes"
	"fmt"
	gotoken "go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/codegen2"
	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/base/baseline"
	"gopkg.in/yaml.v3"
)

type Source struct {
	Name string `yaml:"name"`
	Data string `yaml:"data"`
}

type Test struct {
	Sources []*Source `yaml:"sources"`
	Inputs  []string  `yaml:"inputs"`
}

func TestLox(t *testing.T) {
	testFiles, err := filepath.Glob("*.yaml")
	if err != nil {
		t.Fatal(err)
	}

	for _, testFile := range testFiles {
		testFileData, err := os.ReadFile(testFile)
		if err != nil {
			t.Fatal(err)
		}

		var test Test
		err = yaml.Unmarshal(testFileData, &test)
		if err != nil {
			t.Fatal(err)
		}

		testName := strings.TrimSuffix(
			filepath.Base(testFile), filepath.Ext(testFile))
		t.Run(testName, func(t *testing.T) {
			tmpDir, err := os.MkdirTemp(".", "test_")
			if err != nil {
				t.Fatal(err)
			}
			defer os.RemoveAll(tmpDir)

			for _, source := range test.Sources {
				err = os.WriteFile(
					filepath.Join(tmpDir, source.Name),
					[]byte(source.Data),
					0600)
				if err != nil {
					t.Fatal(err)
				}
			}

			resBuf := new(strings.Builder)

			fset := gotoken.NewFileSet()
			errs := errlogger.New(resBuf)
			ok := codegen2.Generate(&codegen2.Config{
				Fset:   fset,
				Errs:   errs,
				Dir:    tmpDir,
				Report: resBuf,
			})

			if !ok {
				fmt.Fprintln(resBuf, "Failed to generate")
				baseline.Assert(t, resBuf.String())
				return
			}

			fmt.Fprintln(resBuf, "Tests")
			fmt.Fprintln(resBuf, "=====")

			for _, input := range test.Inputs {
				fmt.Fprintln(resBuf, "Input:")
				fmt.Fprintln(resBuf, input)
				fmt.Fprintln(resBuf, "")

				goRun := exec.Command("go", "run", tmpDir)
				goRun.Stdin = bytes.NewReader([]byte(input))

				out, err := goRun.CombinedOutput()
				if err != nil {
					fmt.Fprintln(resBuf, "go run failed:", err)
				}

				fmt.Fprintln(resBuf, "Output:")
				resBuf.Write(out)
				fmt.Fprintln(resBuf, "")
			}

			baseline.Assert(t, resBuf.String())
		})
	}
}
