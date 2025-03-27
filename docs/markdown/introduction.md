# Introduction

The easiest way to get a feeling for Lox is to look at an example. The following
implements a simple terminal calculator. You can copy the files to an empty Go
module, run `lox .`, then run `go run .`.


```lox
// === parser.lox ===

@lexer // Token definitions.

OPAREN = '('
CPAREN = ')'
SUB    = '-'
ADD    = '+'
MUL    = '*'
DIV    = '/'

NUMBER = ('0' | [1-9][0-9]*) ('.' [0-9]+)?

@frag [ \r\n\t]+  @discard // Discard whitespaces.

@parser // Parsing rules.

@start goal = expr

expr = expr '+' expr @left(1)
     | expr '-' expr @left(1)
     | expr '*' expr @left(2)
     | expr '/' expr @left(2)
     | number
     | '(' expr ')'

number = NUMBER | '-' NUMBER
```
```go
// === parser.go ===

package main

import (
	"bufio"
	"fmt"
	gotoken "go/token"
	"os"
	"strconv"

	"github.com/dcaiafa/loxlex/simplelexer"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		res, err := eval(scanner.Text())
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println(res)
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
		os.Exit(1)
	}
}

func eval(input string) (float64, error) {
	// While Lox generates the lexer state machine, you are responsible for
	// providing the actual lexer implementation. In most cases, `simplelexer` is
	// sufficient. If you prefer not to have external dependencies, you can
	// implement your own lexer or copy `simplelexer` locally.
	lex := simplelexer.New(simplelexer.Config{
		StateMachine: new(_LexerStateMachine),
		File:         gotoken.NewFileSet().AddFile("input", -1, len(input)),
		Input:        []byte(input),
	})

	var parser myParser

	// parse is a method implemented in the embedded `lox` type.
	ok := parser.parse(lex)
	if !ok {
		// For production use, you will likely want to implement more robust error
		// handling. See "Error Handling and Recovery" for more details.
		return 0, fmt.Errorf("failed to parse")
	}

	return parser.result, nil
}

// You must define a type named `Token` because the `lox` tool expects your
// lexer to produce tokens of this type. Since we are using `simplelexer` as the
// lexer implementation, we need to alias its `Token` type to match the expected
// type.
type Token = simplelexer.Token

// You must define a parser type. The name of this type is not important, but it
// must embed the generated `lox` type. The `lox` tool will match productions to
// methods of your type.
type myParser struct {
	// Identifies this type as THE parser. Also provides parser state and
	// primitives.
	lox

	// You can have anything else you want in your parser. In this case we need a
	// field to store the result of the expression.
	result float64
}

// The "on_" prefix indicates to Lox that this method is the action associated
// with a production rule named after the "on_" prefix.
//
// The method parameters must correspond to the terms of the production. If a
// term's value is not needed, its corresponding parameter can be named "_" to
// indicate it is ignored, but the parameter must still be present in the method
// signature. The type of each parameter must match the term's value type. For
// token terms, the type is always `Token`, while for rule (non-terminal) terms,
// the type is determined by the return type of the referenced rule's action
// method.
//
// Lox expects *all* parser productions to be matched to exactly one "on_"
// method. If a rule has multiple productions, a double-underscore suffix can be
// used to make unique method names (Lox will ignore anything after the __ and
// will use only the parameters to match the production).

func (p *myParser) on_goal(v float64) float64 {
	p.result = v
	return v
}

func (p *myParser) on_expr__binary(left float64, op Token, right float64) float64 {
	switch op.Type {
	case ADD:
		return left + right
	case SUB:
		return left - right
	case MUL:
		return left * right
	case DIV:
		return left / right
	default:
		panic("unreached")
	}
}

func (p *myParser) on_expr__number(v float64) float64 {
	return v
}

func (p *myParser) on_expr__paren(_ Token, v float64, _ Token) float64 {
	return v
}

func (p *myParser) on_number__positive(n Token) float64 {
	v, err := strconv.ParseFloat(string(n.Str), 64)
	if err != nil {
		panic("invalid number") // For illustration purposes; you can do better
	}
	return v
}

func (p *myParser) on_number__negative(_ Token, n Token) float64 {
	v, err := strconv.ParseFloat(string(n.Str), 64)
	if err != nil {
		panic("invalid number")
	}
	return -v
}
```
### What next?
* Check out the lexer and parser section references.
* Explore more complex examples.
