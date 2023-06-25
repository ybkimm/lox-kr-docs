package parser

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/errs"
)

func TestLexer(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		const input = `
@lexer

@custom ID '*'

@parser

syntax     = decl* '=' qterm+ label? # foo123 .
`

		errs := errs.New()
		lex := newLex("input.lox", bytes.TrimSpace([]byte(input)), errs)

		res := new(strings.Builder)

		for {
			var sym yySymType
			tok := lex.Lex(&sym)
			pos := sym.tok.Pos
			fmt.Fprintf(res, "%d:%d: %v\n", pos.Line, pos.Column, tokenName(tok))
			if tok == 0 || tok == LEXERR {
				break
			}
		}

		expected := `
1:1: kLEXER
3:1: kCUSTOM
3:9: ID
3:12: LITERAL
5:1: kPARSER
7:1: ID
7:12: '='
7:14: ID
7:18: '*'
7:20: LITERAL
7:24: ID
7:29: '+'
7:31: ID
7:36: '?'
7:38: '#'
7:40: ID
7:47: '.'
0:0: $end
`
		expected = strings.TrimSpace(expected)
		actual := strings.TrimSpace(res.String())

		if expected != actual {
			t.Fatalf("Expected:\n%v\nActual:\n%v", expected, actual)
		}
	})

}

// tokenName returns a human-friendly name for the token integer code returned
// by the lexer. This is based on yylex1().
func tokenName(char int) string {
	var token int
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	return yyTokname(token)
}
