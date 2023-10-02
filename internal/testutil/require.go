package testutil

import (
	"strings"
	"testing"
)

func RequireEqualStr(t *testing.T, actual, expected string) {
	t.Helper()
	actual = strings.TrimSpace(actual)
	expected = strings.TrimSpace(expected)
	if actual != expected {
		t.Fatalf("not equal:\nexpected:\n%v\nactual:\n%v", expected, actual)
	}
}
