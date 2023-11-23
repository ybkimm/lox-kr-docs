package ast_test

import (
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/testutil"
)

func TestSpec_Calc(t *testing.T) {
	spec, ctx := parseAndAnalyze(t, `
@lexer

NUM = [0-9]+ ;
ADD = '+' ;
SUB = '-' ;
MUL = '*' ;
DIV = '/' ;
REM = '%' ;
POW = '^' ;
O_PAREN = '(' ;
C_PAREN = ')' ;

@parser

@start
S = expr ;

expr = expr '+' expr  @left(1)
     | expr '-' expr  @left(1)
     | expr '*' expr  @left(2)
     | expr '/' expr  @left(2)
     | expr '%' expr  @left(2)
     | expr '^' expr  @right(3)
     | '(' expr ')'
     | num ;

num = NUM
    | '-' NUM ;
`)

	// TODO: turn this into a proper test
	_ = spec
	ctx.Grammar.SetStart(ctx.StartParserRule.Rule)
	var printedGrammar strings.Builder
	ctx.Grammar.Print(&printedGrammar)

	testutil.RequireEqualStr(t, printedGrammar.String(), `
Terminals
=========
EOF
ERROR
NUM
ADD
SUB
MUL
DIV
REM
POW
O_PAREN
C_PAREN

Rules
=====
S' = S
S = expr
expr = expr ADD expr  @left(1)
  | expr SUB expr  @left(1)
  | expr MUL expr  @left(2)
  | expr DIV expr  @left(2)
  | expr REM expr  @left(2)
  | expr POW expr  @right(3)
  | O_PAREN expr C_PAREN
  | num
num = NUM
  | SUB NUM
`)
}
