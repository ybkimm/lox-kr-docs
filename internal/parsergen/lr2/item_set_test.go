package lr2

import (
	"fmt"
	"testing"

	"github.com/dcaiafa/lox/internal/testutil"
)

func TestItemSetLR0Key(t *testing.T) {
	var is ItemSet
	is.Add(Item{Prod: 3, Dot: 4, Lookahead: 10})
	is.Add(Item{Prod: 1, Dot: 2, Lookahead: 10})
	is.Add(Item{Prod: 1, Dot: 2, Lookahead: 20})

	testutil.RequireEqualStr(
		t, fmt.Sprintf("%x", is.LR0Key()),
		"00000001000000020000000300000004")
}
