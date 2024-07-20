package ast_test

import (
	"strings"
	"testing"
)

func TestParserTermNormalize(t *testing.T) {
	t.Run("ZeroOrMore", func(t *testing.T) {
		spec, ctx := parseAndAnalyze(t, `
@lexer
P = '+'
@parser
b = P
c = P
@start S = a | d
a = b c*
d = c*
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
ParserRule: Name: S
  Prod:
    Term: Name: a Type: Simple
  Prod:
    Term: Name: d Type: Simple
ParserRule: Name: a
  Prod:
    Term: Name: b Type: Simple
    Term: Name: c* Type: Simple
ParserRule: Name: d
  Prod:
    Term: Name: c* Type: Simple
ParserRule: Name: c*
  Prod:
    Term: Name: c+ Type: Simple
  Prod:
ParserRule: Name: c+
  Prod:
    Term: Name: c+ Type: Simple
    Term: Name: c Type: Simple
  Prod:
    Term: Name: c Type: Simple
		`)
	})

	t.Run("OneOrMore", func(t *testing.T) {
		spec, ctx := parseAndAnalyze(t, `
@lexer
P = '+'
@parser
b = P
c = P
@start a = b c+
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
    Term: Name: c+ Type: Simple
ParserRule: Name: c+
  Prod:
    Term: Name: c+ Type: Simple
    Term: Name: c Type: Simple
  Prod:
    Term: Name: c Type: Simple
		`)
	})

	t.Run("ZeroOrOne", func(t *testing.T) {
		spec, ctx := parseAndAnalyze(t, `
@lexer
P = '+'
@parser
b = P
c = P
@start a = b c?
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
    Term: Name: c? Type: Simple
ParserRule: Name: c?
  Prod:
    Term: Name: c Type: Simple
  Prod:
		`)
	})

	t.Run("List", func(t *testing.T) {
		spec, ctx := parseAndAnalyze(t, `
@lexer
P = '+'
@parser
b = P
c = P
@start S = a | d
a = b @list(c, P)
d = @list(c, P)
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
ParserRule: Name: S
  Prod:
    Term: Name: a Type: Simple
  Prod:
    Term: Name: d Type: Simple
ParserRule: Name: a
  Prod:
    Term: Name: b Type: Simple
    Term: Name: @list(c,P) Type: Simple
ParserRule: Name: d
  Prod:
    Term: Name: @list(c,P) Type: Simple
ParserRule: Name: @list(c,P)
  Prod:
    Term: Name: @list(c,P) Type: Simple
    Term: Name: P Type: Simple
    Term: Name: c Type: Simple
  Prod:
    Term: Name: c Type: Simple
		`)
	})

	t.Run("ListOpt", func(t *testing.T) {
		spec, ctx := parseAndAnalyze(t, `
@lexer
P = '+'
@parser
b = P
c = P
@start S = a | d
a = b @list(c, P)?
d = @list(c, P)
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
ParserRule: Name: S
  Prod:
    Term: Name: a Type: Simple
  Prod:
    Term: Name: d Type: Simple
ParserRule: Name: a
  Prod:
    Term: Name: b Type: Simple
    Term: Name: @list(c,P)? Type: Simple
ParserRule: Name: d
  Prod:
    Term: Name: @list(c,P) Type: Simple
ParserRule: Name: @list(c,P)?
  Prod:
    Term: Name: @list(c,P) Type: Simple
  Prod:
ParserRule: Name: @list(c,P)
  Prod:
    Term: Name: @list(c,P) Type: Simple
    Term: Name: P Type: Simple
    Term: Name: c Type: Simple
  Prod:
    Term: Name: c Type: Simple
    `)
	})

	t.Run("ListOnlySupportOpt", func(t *testing.T) {
		failParse(t, `
@lexer
P = '+'
@parser
b = P
c = P
@start S = a
a = b @list(c, P)+
`, "@list term can only use the zero-or-more '?' cardinality")

	})

	t.Run("ListEntryMustBeSimple1", func(t *testing.T) {
		parseButFailAnalyze(t, `
@lexer
C = ','
N = '1'
@parser
@start S = @list(@list(N, C), C)
n = N N
		`, "@list entry param must be a simple token or rule")
	})

	t.Run("ListEntryMustBeSimple2", func(t *testing.T) {
		parseButFailAnalyze(t, `
@lexer
C = ','
N = '1'
@parser
@start S = @list(@list(N, C), C)?
n = N N
		`, "@list entry param must be a simple token or rule")
	})

	t.Run("ListSepMustBeSimple", func(t *testing.T) {
		_, _ = parseAndAnalyze(t, `
@lexer
C = ','
N = '1'
@parser
@start S = @list(N, n)
n = N N
		`)
	})

	t.Run("ListSepMustBeSimple", func(t *testing.T) {
		parseButFailAnalyze(t, `
@lexer
C = ','
N = '1'
@parser
@start S = @list(N, @list(N, C))
n = N N
		`, "@list separator param must be a simple token or rule")
	})
}
