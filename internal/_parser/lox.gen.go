package parser2

import (
	_i0 "fmt"
	_i1 "os"
)

const (
	ID           = 1
	LITERAL      = 2
	LABEL        = 3
	ZERO_OR_MANY = 4
	ONE_OR_MANY  = 5
	ZERO_OR_ONE  = 6
	DEFINE       = 7
	SEMICOLON    = 8
	PARSER       = 9
	LEXER        = 10
	CUSTOM       = 11
)

var _lxLHS = []int32{
	0, 1, 2, 2, 3, 4, 5, 6, 7, 8, 8, 8, 9, 10,
	11, 12, 13, 13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18,
	19, 19, 20, 20, 21, 21, 22, 22,
}

var _lxReduction = []int32{
	1, 1, 1, 1, 2, 1, 4, 2, 2, 1, 1, 1, 1, 2,
	1, 3, 2, 1, 1, 0, 2, 1, 2, 1, 1, 0, 1, 0,
	1, 0, 2, 1, 2, 1, 2, 1,
}

var _lxAction = []int32{
	42, 47, 56, 63, 72, 79, 86, 89, 96, 99, 108, 117, 124, 133,
	136, 143, 152, 161, 170, 177, 182, 187, 196, 199, 208, 213, 222, 235,
	240, 247, 252, 259, 266, 273, 280, 287, 294, 299, 304, 309, 316, 321,
	4, 10, 1, 9, 3, 8, 0, -29, 11, 8, 10, -29, 9, -29,
	6, 0, -3, 10, -3, 9, -3, 8, 0, -19, 1, 13, 10, -19,
	9, -19, 6, 0, -2, 10, -2, 9, -2, 6, 0, -17, 10, -17,
	9, -17, 2, 0, 2147483647, 6, 0, -1, 10, 1, 9, 3, 2, 1,
	19, 8, 0, -14, 11, -14, 10, -14, 9, -14, 8, 0, -35, 11,
	-35, 10, -35, 9, -35, 6, 0, -13, 10, -13, 9, -13, 8, 0,
	-28, 11, 8, 10, -28, 9, -28, 2, 7, 22, 6, 0, -4, 10,
	-4, 9, -4, 8, 0, -18, 1, 13, 10, -18, 9, -18, 8, 0,
	-33, 1, -33, 10, -33, 9, -33, 8, 0, -5, 1, -5, 10, -5,
	9, -5, 6, 0, -16, 10, -16, 9, -16, 4, 1, -31, 8, -31,
	4, 1, 24, 8, 25, 8, 0, -34, 11, -34, 10, -34, 9, -34,
	2, 1, 26, 8, 0, -32, 1, -32, 10, -32, 9, -32, 4, 1,
	-30, 8, -30, 8, 0, -15, 11, -15, 10, -15, 9, -15, 12, 1,
	-27, 3, -27, 5, 31, 8, -27, 4, 34, 6, 35, 4, 1, -21,
	8, -21, 6, 1, 26, 3, 36, 8, -25, 4, 1, 26, 8, 41,
	6, 1, -23, 3, -23, 8, -23, 6, 1, -10, 3, -10, 8, -10,
	6, 1, -26, 3, -26, 8, -26, 6, 1, -8, 3, -8, 8, -8,
	6, 1, -9, 3, -9, 8, -9, 6, 1, -11, 3, -11, 8, -11,
	4, 1, -12, 8, -12, 4, 1, -24, 8, -24, 4, 1, -7, 8,
	-7, 6, 1, -22, 3, -22, 8, -22, 4, 1, -20, 8, -20, 8,
	0, -6, 1, -6, 10, -6, 9, -6,
}

var _lxGoto = []int32{
	42, 53, 62, 63, 62, 62, 62, 72, 79, 62, 62, 62, 82, 62,
	62, 87, 62, 62, 62, 62, 62, 62, 92, 62, 62, 62, 101, 62,
	106, 113, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62,
	10, 10, 2, 3, 4, 2, 5, 1, 6, 13, 7, 8, 12, 9,
	11, 10, 19, 11, 22, 12, 0, 8, 14, 14, 21, 15, 4, 16,
	5, 17, 6, 10, 2, 3, 4, 2, 18, 2, 20, 20, 4, 12,
	9, 11, 21, 4, 4, 23, 5, 17, 8, 6, 27, 16, 28, 15,
	29, 7, 30, 4, 8, 32, 18, 33, 6, 9, 37, 17, 38, 7,
	39, 6, 6, 40, 16, 28, 7, 30,
}

type _lxStack[T any] []T

func (s *_lxStack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *_lxStack[T]) Pop(n int) {
	*s = (*s)[:len(*s)-n]
}

func (s _lxStack[T]) Peek(n int) T {
	return s[len(s)-n-1]
}

func _lxFind(table []int32, y, x int32) (int32, bool) {
	i := int(table[int(x)])
	count := int(table[i])
	i++
	end := i + count
	for ; i < end; i += 2 {
		if table[i] == x {
			return table[i+1], true
		}
	}
	return 0, false
}

type _lxLexer interface {
	Token() (int, Token)
}

type loxParser struct {
	state _lxStack[int32]
	sym   _lxStack[any]
}

func (p *Parser) parse(lex _lxLexer) {
	const accept = 2147483647

	p.loxParser.state.Push(0)
	lookahead, tok := lex.Token()

	for {
		topState := p.loxParser.state.Peek(0)
		action, ok := _lxFind(_lxAction, topState, int32(lookahead))
		if !ok {
			p.onError(tok, "boom")
			return
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			p.loxParser.state.Push(action)
			p.loxParser.sym.Push(tok)
		} else { // reduce
			prod := -action
			termCount := _lxReduction[int(prod)]
			rule := _lxLHS[int(prod)]
			p.loxParser.state.Pop(int(termCount))
			p.loxParser.sym.Pop(int(termCount))
			topState = p.loxParser.state.Peek(0)
			nextState, _ := _lxFind(_lxGoto, topState, rule)
			p.loxParser.state.Push(nextState)
			p.loxParser.sym.Push(nil)
		}
	}
}

func (p *Parser) onError(tok Token, err string) {
	_i0.Println("ERROR:", err)
	_i1.Exit(1)
}
