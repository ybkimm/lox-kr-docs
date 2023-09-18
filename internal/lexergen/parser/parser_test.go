package parser

import (
	gotoken "go/token"
	"testing"

	"github.com/dcaiafa/lox/internal/errlogger"
)

var testInput = `

INFO = 'info' ;
EQ   = '==' ;

CHAR = '\'' (ESCAPE_CHAR | ~['\r\n\\]+) '\'';
@macro ESCAPE_CHAR = '\\' [nrt'\\] ;

@macro HEX_DIGIT = [a-fA-F0-9] ;

COMMENT = '#' ~[\n]* -> @skip ;
WS      = [ \t] ;

OCURLY = '{' -> @push_mode(DEFAULT) ;
CCURLY = '}' -> @pop_mode ;

EXEC_PREFIX = 'e\'' -> @push_mode(EXEC) ;

@mode EXEC {
	@frag '"' -> @push_mode(EXEC_DQUOTE) ;

	EXEC_WS      = [ \t\r\n]+ ; 
  EXEC_HOME    = '~' ;
  EXEC_LITERAL = ~[ \t\r\n~{"']+ ;
	EXEC_OCURLY  = '{' -> @push_mode(DEFAULT) ;
  EXEC_SUFFIX  = '\'' -> @pop_mode ;
}

@mode EXEC_DQUOTE {
	@frag ~["\r\n\\] ;
	@frag '\\' ([nrt"\\] | ('x' HEX_DIGIT HEX_DIGIT)) ;
	EXEC_DQUOTE_LITERAL = '"' -> @pop_mode ;
}

`

func TestParser(t *testing.T) {
	fset := gotoken.NewFileSet()
	data := []byte(testInput)
	file := fset.AddFile("input.loxl", -1, len(data))
	errs := errlogger.New()
	Parse(file, []byte(data), errs)
	if errs.HasError() {
		t.Fatal("Parse failed")
	}
}
