package testutil

import (
	"reflect"
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

func RequireEqual[T any](t *testing.T, actual, expected T) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("not equal:\nexpected:\n%+v\nactual:\n%+v", expected, actual)
	}
}

func RequireTrue(t *testing.T, v bool) {
	t.Helper()
	if !v {
		t.Fatalf("expected true")
	}
}

func RequireFalse(t *testing.T, v bool) {
	t.Helper()
	if v {
		t.Fatalf("expected false")
	}
}
