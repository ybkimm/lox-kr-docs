package parser

import (
	gotoken "go/token"
	"os"
	"testing"

	"github.com/dcaiafa/lox/internal/errlogger"
)

var testInput = `
@lexer
C =  '\\' ;

INFO = 'info' ;
EQ   = '==' ;

CHAR = '\'' (ESCAPE_CHAR | ~['\r\n\\]+) '\'' ;
@macro ESCAPE_CHAR = '\\' [nrt'\\] ;

@macro HEX_DIGIT = [a-fA-F0-9] ;
@macro ID_CHAR = [a-zA-Z_-] ;
@macro ID_CHAR2 = [\u0041-\u005A\u0061-\u007a\u002d\u005f] ;

COMMENT = '#' ~[\n]* @skip ;
WS      = [ \t] ;

OCURLY = '{'   @push_mode(DEFAULT) ;
CCURLY = '}'   @pop_mode ;

EXEC_PREFIX = 'e\''  @push_mode(EXEC) ;

@mode EXEC {
	@frag '"' @push_mode(EXEC_DQUOTE) ;

	EXEC_WS      = [ \t\r\n]+ ;
  EXEC_HOME    = '~' ;
  EXEC_LITERAL = ~[ \t\r\n~{"']+ ;
	EXEC_OCURLY  = '{'  @push_mode(DEFAULT) ;
  EXEC_SUFFIX  = '\'' @pop_mode ;
}

@mode EXEC_DQUOTE {
	@frag ~["\r\n\\] ;
	@frag '\\' ([nrt"\\] | ('x' HEX_DIGIT HEX_DIGIT)) ;
	EXEC_DQUOTE_LITERAL = '"' @pop_mode ;
}
`

func TestParser(t *testing.T) {
	fset := gotoken.NewFileSet()
	data := []byte(testInput)
	file := fset.AddFile("input.lox", -1, len(data))
	errs := errlogger.New(os.Stderr)
	Parse(file, []byte(data), errs)
	if errs.HasError() {
		t.Fatal("Parse failed")
	}
}

func TestUnescape(t *testing.T) {
	input := `abc\n\r\t\'\\\-\U00101234\u12e4\xff\x07012`
	expect := "abc\n\r\t'\\-\U00101234\u12e4\xff\x07012"
	output := unescape([]byte(input))
	if output != expect {
		t.Fatalf("expected: %q (%d); actual: %q (%d)", expect, len(expect), output, len(output))
	}
}
