package ast_test

import (
	"strings"
	"testing"
)

func TestParserTermNormalize(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		spec, ctx := parseAndAnalyze(t, `
@lexer
P = '+' ;
@parser
b = P ;
c = P ;
a = b c+ ;
`)
		var str strings.Builder
		ctx.Print(spec, &str)
		requireEqualStr(t, str.String(), `
LexerTokenRule: Name: P
  LexerExpr:
    LexerFactor:
      LexerTermCard:
        LexerTermLiteral: "+"
ParserRule: Name: b
  Prod:
    Term: Name: P Type: Simple
ParserRule: Name: c
  Prod:
    Term: Name: P Type: Simple
ParserRule: Name: a
  Prod:
    Term: Name: b Type: Simple
    Term: Name: a$4 Type: Simple
ParserRule: Name: a$4 Generated: OneOrMore
  Prod:
    Term: Name: a$4 Type: Simple
    Term: Name: c Type: Simple
  Prod:
    Term: Name: c Type: Simple
		`)
	})

}
