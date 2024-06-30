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

	"github.com/CloudyKit/jet/v6"
	"github.com/dcaiafa/lox/internal/base/baseline"
	"github.com/dcaiafa/lox/internal/base/errlogger"
	"github.com/dcaiafa/lox/internal/codegen"
	"gopkg.in/yaml.v3"
)

type TestFile struct {
	Name    string    `yaml:"-"`
	Sources []*Source `yaml:"sources"`
	Tests   []*Test   `yaml:"tests"`
}

type Source struct {
	Name string `yaml:"name"`
	Data string `yaml:"data"`
}

type Test struct {
	Name   string         `yaml:"name"`
	Vars   map[string]any `yaml:"vars"`
	Inputs []string       `yaml:"inputs"`
}

func TestLox(t *testing.T) {
	testFilenames, err := filepath.Glob("*.yaml")
	if err != nil {
		t.Fatal(err)
	}

	for _, testFilename := range testFilenames {
		testFileData, err := os.ReadFile(testFilename)
		if err != nil {
			t.Fatal(err)
		}

		var testFile TestFile

		err = yaml.Unmarshal(testFileData, &testFile)
		if err != nil {
			t.Fatal(err)
		}

		testFile.Name = strings.TrimSuffix(
			filepath.Base(testFilename), filepath.Ext(testFilename))

		for _, test := range testFile.Tests {
			runTest(t, &testFile, test)
		}
	}
}

func runTest(t *testing.T, testFile *TestFile, test *Test) {
	testName := testFile.Name + "_" + test.Name
	t.Run(testName, func(t *testing.T) {
		tmpDir, err := os.MkdirTemp(".", "test_")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		for _, source := range testFile.Sources {
			rendered := renderSource(source.Data, test.Vars)
			err = os.WriteFile(
				filepath.Join(tmpDir, source.Name),
				rendered,
				0600)
			if err != nil {
				t.Fatal(err)
			}
		}

		resBuf := new(strings.Builder)

		fset := gotoken.NewFileSet()
		errs := errlogger.New(fset, resBuf)
		ok := codegen.Generate(&codegen.Config{
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

func renderSource(source string, vars map[string]any) []byte {
	loader := jet.NewInMemLoader()
	loader.Set("source", source)

	set := jet.NewSet(loader, jet.WithSafeWriter(nil))
	jetTemplate, err := set.GetTemplate("source")
	if err != nil {
		panic(err)
	}

	jetVars := make(jet.VarMap)
	for argName, argValue := range vars {
		jetVars.Set(argName, argValue)
	}

	rendered := &bytes.Buffer{}
	err = jetTemplate.Execute(rendered, jetVars, nil)
	if err != nil {
		panic(err)
	}

	return rendered.Bytes()
}
