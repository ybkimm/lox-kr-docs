package baseline

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var updateBaseline = flag.Bool("update-baseline", false, "update _baseline/*.txt files")

// Assert compares output against the contents of file
// _baseline/<TestName>.txt. If the contents don't match,
// the diff is displayed and the test fails.
//
// To regenerated the baseline, run:
//
//	go test -run <your-test> -args -update-baseline
func Assert(t *testing.T, output string) {
	baselineFilename := filepath.Join("_baseline", t.Name()+".txt")

	t.Log("Output:\n", output)

	if *updateBaseline {
		err := os.MkdirAll(filepath.Dir(baselineFilename), 0755)
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile(baselineFilename, []byte(output), 0644)
		if err != nil {
			t.Fatal(err)
		}
		return
	}

	baselineData, err := os.ReadFile(baselineFilename)
	if err != nil {
		t.Fatal(err)
	}

	baseline := string(baselineData)

	if baseline != output {
		t.Log("Diff:\n" + cmp.Diff(baseline, output))
		t.Fatalf("Output does not match baseline")
	}
}
