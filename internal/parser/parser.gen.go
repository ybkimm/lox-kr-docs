package parser

import (
	_i0 "github.com/dcaiafa/lox/internal/ast"
)

var _lxLHS = []int32{
	0, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 8, 8, 9,
	9, 10, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 16, 17,
	17, 18, 18, 19, 19,
}

var _lxTermCounts = []int32{
	1, 1, 1, 1, 4, 2, 2, 1, 1, 6, 1, 1, 1, 4,
	4, 3, 2, 1, 0, 3, 1, 2, 1, 1, 0, 1, 0, 2,
	1, 1, 0, 2, 1,
}

var _lxActions = []int32{
	50, 57, 60, 63, 70, 73, 76, 83, 90, 97, 102, 109, 114, 119,
	126, 149, 152, 175, 180, 193, 198, 217, 230, 235, 240, 247, 97, 252,
	255, 258, 263, 268, 97, 281, 288, 301, 314, 327, 340, 353, 356, 359,
	362, 97, 367, 370, 373, 376, 381, 386, 6, 0, -18, 2, 1, 14,
	2, 2, 10, 9, 2, 2, 10, 6, 0, -32, 2, -32, 14, -32,
	2, 0, 2147483647, 2, 0, -1, 6, 0, -17, 2, 1, 14, 2, 6,
	0, -2, 2, -2, 14, -2, 6, 0, -3, 2, -3, 14, -3, 4,
	2, 14, 17, 15, 6, 2, -30, 3, 22, 13, -30, 4, 2, -28,
	13, -28, 4, 2, 10, 13, 24, 6, 0, -31, 2, -31, 14, -31,
	22, 8, -7, 9, -7, 2, -7, 15, -7, 17, -7, 6, -7, 12,
	-7, 16, -7, 13, -7, 5, -7, 7, -7, 2, 11, 26, 22, 8,
	-8, 9, -8, 2, -8, 15, -8, 17, -8, 6, -8, 12, -8, 16,
	-8, 13, -8, 5, -8, 7, -8, 4, 12, -20, 13, -20, 12, 2,
	14, 15, 27, 17, 15, 12, -24, 16, 28, 13, -24, 4, 12, 32,
	13, 33, 18, 2, -26, 15, -26, 17, -26, 6, 34, 12, -26, 16,
	-26, 13, -26, 5, 35, 7, 36, 12, 2, -22, 15, -22, 17, -22,
	12, -22, 16, -22, 13, -22, 4, 2, -29, 13, -29, 4, 2, -16,
	13, -16, 6, 0, -15, 2, -15, 14, -15, 4, 2, -27, 13, -27,
	2, 11, 40, 2, 11, 41, 4, 12, -5, 13, -5, 4, 12, -23,
	13, -23, 12, 2, -21, 15, -21, 17, -21, 12, -21, 16, -21, 13,
	-21, 6, 0, -4, 2, -4, 14, -4, 12, 2, -11, 15, -11, 17,
	-11, 12, -11, 16, -11, 13, -11, 12, 2, -10, 15, -10, 17, -10,
	12, -10, 16, -10, 13, -10, 12, 2, -12, 15, -12, 17, -12, 12,
	-12, 16, -12, 13, -12, 12, 2, -25, 15, -25, 17, -25, 12, -25,
	16, -25, 13, -25, 12, 2, -6, 15, -6, 17, -6, 12, -6, 16,
	-6, 13, -6, 2, 8, 43, 2, 4, 44, 2, 4, 45, 4, 12,
	-19, 13, -19, 2, 9, 47, 2, 9, 48, 2, 9, 49, 4, 12,
	-13, 13, -13, 4, 12, -14, 13, -14, 22, 8, -9, 9, -9, 2,
	-9, 15, -9, 17, -9, 6, -9, 12, -9, 16, -9, 13, -9, 5,
	-9, 7, -9,
}

var _lxGoto = []int32{
	50, 63, 64, 63, 63, 63, 69, 63, 63, 76, 89, 63, 92, 63,
	63, 63, 63, 63, 95, 63, 106, 63, 63, 63, 63, 63, 111, 63,
	63, 63, 63, 63, 116, 63, 63, 63, 63, 63, 63, 63, 63, 63,
	63, 127, 63, 63, 63, 63, 63, 63, 12, 2, 3, 1, 4, 12,
	5, 19, 6, 3, 7, 10, 8, 0, 4, 11, 11, 17, 12, 6,
	2, 13, 3, 7, 10, 8, 12, 7, 16, 4, 17, 14, 18, 13,
	19, 6, 20, 5, 21, 2, 18, 23, 2, 11, 25, 10, 7, 16,
	15, 29, 9, 30, 6, 20, 5, 31, 4, 8, 37, 16, 38, 4,
	7, 16, 6, 39, 10, 7, 16, 4, 42, 14, 18, 6, 20, 5,
	21, 4, 7, 16, 6, 46,
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

func (s _lxStack[T]) Slice(n int) []T {
	return s[len(s)-n:]
}

func _lxFind(table []int32, y, x int32) (int32, bool) {
	i := int(table[int(y)])
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

type loxParser struct {
	state  _lxStack[int32]
	sym    _lxStack[any]
	bounds _lxStack[_lxBounds]
}

func (p *parser) parse(lex _lxLexer, errLogger _lxErrorLogger) bool {
	const accept = 2147483647

	p.loxParser.state.Push(0)
	tok, tokType := lex.NextToken()

	for {
		topState := p.loxParser.state.Peek(0)
		action, ok := _lxFind(_lxActions, topState, int32(tokType))
		if !ok {
			errLogger.ParserError(&_lxUnexpectedTokenError{Token: tok})
			return false
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			p.loxParser.state.Push(action)
			p.loxParser.sym.Push(tok)
			p.loxParser.bounds.Push(_lxBounds{Begin: tok, End: tok})
			tok, tokType = lex.NextToken()
		} else { // reduce
			prod := -action
			termCount := _lxTermCounts[int(prod)]
			rule := _lxLHS[int(prod)]
			res := p._lxAct(prod)
			if termCount > 0 {
				bounds := _lxBounds{
					Begin: p.loxParser.bounds.Peek(int(termCount - 1)).Begin,
					End:   p.loxParser.bounds.Peek(0).End,
				}
				p.onReduce(res, bounds.Begin, bounds.End)
				p.loxParser.bounds.Pop(int(termCount))
				p.loxParser.bounds.Push(bounds)
			} else {
				bounds := p.loxParser.bounds.Peek(0)
				bounds.Begin = bounds.End
				p.loxParser.bounds.Push(bounds)
			}
			p.loxParser.state.Pop(int(termCount))
			p.loxParser.sym.Pop(int(termCount))
			topState = p.loxParser.state.Peek(0)
			nextState, _ := _lxFind(_lxGoto, topState, rule)
			p.loxParser.state.Push(nextState)
			p.loxParser.sym.Push(res)
		}
	}

	return true
}

func (p *parser) _lxAct(prod int32) any {
	switch prod {
	case 1:
		return p.on_parser(
			p.sym.Peek(0).([]_i0.ParserDecl),
		)
	case 2:
		return p.on_decl(
			p.sym.Peek(0).(_i0.ParserDecl),
		)
	case 3:
		return p.on_decl(
			p.sym.Peek(0).(_i0.ParserDecl),
		)
	case 4:
		return p.on_rule(
			p.sym.Peek(3).(Token),
			p.sym.Peek(2).(Token),
			p.sym.Peek(1).([]*_i0.Prod),
			p.sym.Peek(0).(Token),
		)
	case 5:
		return p.on_prod(
			p.sym.Peek(1).([]*_i0.Term),
			p.sym.Peek(0).(*_i0.ProdQualifier),
		)
	case 6:
		return p.on_term_card(
			p.sym.Peek(1).(*_i0.Term),
			p.sym.Peek(0).(_i0.TermType),
		)
	case 7:
		return p.on_term__id(
			p.sym.Peek(0).(Token),
		)
	case 8:
		return p.on_term__list(
			p.sym.Peek(0).(*_i0.Term),
		)
	case 9:
		return p.on_list(
			p.sym.Peek(5).(Token),
			p.sym.Peek(4).(Token),
			p.sym.Peek(3).(*_i0.Term),
			p.sym.Peek(2).(Token),
			p.sym.Peek(1).(*_i0.Term),
			p.sym.Peek(0).(Token),
		)
	case 10:
		return p.on_card(
			p.sym.Peek(0).(Token),
		)
	case 11:
		return p.on_card(
			p.sym.Peek(0).(Token),
		)
	case 12:
		return p.on_card(
			p.sym.Peek(0).(Token),
		)
	case 13:
		return p.on_qualif(
			p.sym.Peek(3).(Token),
			p.sym.Peek(2).(Token),
			p.sym.Peek(1).(Token),
			p.sym.Peek(0).(Token),
		)
	case 14:
		return p.on_qualif(
			p.sym.Peek(3).(Token),
			p.sym.Peek(2).(Token),
			p.sym.Peek(1).(Token),
			p.sym.Peek(0).(Token),
		)
	case 15:
		return p.on_token_decl(
			p.sym.Peek(2).(Token),
			p.sym.Peek(1).([]*_i0.CustomToken),
			p.sym.Peek(0).(Token),
		)
	case 16:
		return p.on_token(
			p.sym.Peek(1).(Token),
			p.sym.Peek(0).(Token),
		)
	case 17: // ZeroOrOne
		return p.sym.Peek(0).([]_i0.ParserDecl)
	case 18: // ZeroOrOne
		{
			var zero []_i0.ParserDecl
			return zero
		}
	case 19: // List
		return append(
			p.sym.Peek(2).([]*_i0.Prod),
			p.sym.Peek(0).(*_i0.Prod),
		)
	case 20: // List
		return []*_i0.Prod{
			p.sym.Peek(0).(*_i0.Prod),
		}
	case 21: // OneOrMore
		return append(
			p.sym.Peek(1).([]*_i0.Term),
			p.sym.Peek(0).(*_i0.Term),
		)
	case 22: // OneOrMore
		return []*_i0.Term{
			p.sym.Peek(0).(*_i0.Term),
		}
	case 23: // ZeroOrOne
		return p.sym.Peek(0).(*_i0.ProdQualifier)
	case 24: // ZeroOrOne
		{
			var zero *_i0.ProdQualifier
			return zero
		}
	case 25: // ZeroOrOne
		return p.sym.Peek(0).(_i0.TermType)
	case 26: // ZeroOrOne
		{
			var zero _i0.TermType
			return zero
		}
	case 27: // OneOrMore
		return append(
			p.sym.Peek(1).([]*_i0.CustomToken),
			p.sym.Peek(0).(*_i0.CustomToken),
		)
	case 28: // OneOrMore
		return []*_i0.CustomToken{
			p.sym.Peek(0).(*_i0.CustomToken),
		}
	case 29: // ZeroOrOne
		return p.sym.Peek(0).(Token)
	case 30: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 31: // OneOrMore
		return append(
			p.sym.Peek(1).([]_i0.ParserDecl),
			p.sym.Peek(0).(_i0.ParserDecl),
		)
	case 32: // OneOrMore
		return []_i0.ParserDecl{
			p.sym.Peek(0).(_i0.ParserDecl),
		}
	default:
		panic("unreachable")
	}
}
