package parser

import (
	_i0 "github.com/dcaiafa/lox/internal/ast"
)

var _lxLHS = []int32{
	0, 1, 2, 2, 3, 4, 5, 6, 6, 6, 7, 8, 8, 8,
	9, 9, 10, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 16,
	17, 17, 18, 18, 19, 19,
}

var _lxTermCounts = []int32{
	1, 1, 1, 1, 4, 2, 2, 1, 1, 1, 6, 1, 1, 1,
	4, 4, 3, 2, 1, 0, 3, 1, 2, 1, 1, 0, 1, 0,
	2, 1, 1, 0, 2, 1,
}

var _lxActions = []int32{
	51, 58, 61, 64, 71, 74, 77, 84, 91, 98, 105, 112, 117, 122,
	129, 154, 157, 182, 207, 212, 227, 232, 253, 268, 273, 278, 285, 98,
	290, 293, 296, 301, 306, 98, 321, 328, 343, 358, 373, 388, 403, 406,
	409, 412, 98, 417, 420, 423, 426, 431, 436, 6, 0, -19, 2, 1,
	14, 2, 2, 10, 9, 2, 2, 10, 6, 0, -33, 2, -33, 14,
	-33, 2, 0, 2147483647, 2, 0, -1, 6, 0, -18, 2, 1, 14, 2,
	6, 0, -2, 2, -2, 14, -2, 6, 0, -3, 2, -3, 14, -3,
	6, 2, 14, 17, 15, 3, 16, 6, 2, -31, 3, 23, 13, -31,
	4, 2, -29, 13, -29, 4, 2, 10, 13, 25, 6, 0, -32, 2,
	-32, 14, -32, 24, 8, -7, 9, -7, 2, -7, 15, -7, 17, -7,
	3, -7, 6, -7, 12, -7, 16, -7, 13, -7, 5, -7, 7, -7,
	2, 11, 27, 24, 8, -8, 9, -8, 2, -8, 15, -8, 17, -8,
	3, -8, 6, -8, 12, -8, 16, -8, 13, -8, 5, -8, 7, -8,
	24, 8, -9, 9, -9, 2, -9, 15, -9, 17, -9, 3, -9, 6,
	-9, 12, -9, 16, -9, 13, -9, 5, -9, 7, -9, 4, 12, -21,
	13, -21, 14, 2, 14, 15, 28, 17, 15, 3, 16, 12, -25, 16,
	29, 13, -25, 4, 12, 33, 13, 34, 20, 2, -27, 15, -27, 17,
	-27, 3, -27, 6, 35, 12, -27, 16, -27, 13, -27, 5, 36, 7,
	37, 14, 2, -23, 15, -23, 17, -23, 3, -23, 12, -23, 16, -23,
	13, -23, 4, 2, -30, 13, -30, 4, 2, -17, 13, -17, 6, 0,
	-16, 2, -16, 14, -16, 4, 2, -28, 13, -28, 2, 11, 41, 2,
	11, 42, 4, 12, -5, 13, -5, 4, 12, -24, 13, -24, 14, 2,
	-22, 15, -22, 17, -22, 3, -22, 12, -22, 16, -22, 13, -22, 6,
	0, -4, 2, -4, 14, -4, 14, 2, -12, 15, -12, 17, -12, 3,
	-12, 12, -12, 16, -12, 13, -12, 14, 2, -11, 15, -11, 17, -11,
	3, -11, 12, -11, 16, -11, 13, -11, 14, 2, -13, 15, -13, 17,
	-13, 3, -13, 12, -13, 16, -13, 13, -13, 14, 2, -26, 15, -26,
	17, -26, 3, -26, 12, -26, 16, -26, 13, -26, 14, 2, -6, 15,
	-6, 17, -6, 3, -6, 12, -6, 16, -6, 13, -6, 2, 8, 44,
	2, 4, 45, 2, 4, 46, 4, 12, -20, 13, -20, 2, 9, 48,
	2, 9, 49, 2, 9, 50, 4, 12, -14, 13, -14, 4, 12, -15,
	13, -15, 24, 8, -10, 9, -10, 2, -10, 15, -10, 17, -10, 3,
	-10, 6, -10, 12, -10, 16, -10, 13, -10, 5, -10, 7, -10,
}

var _lxGoto = []int32{
	51, 64, 65, 64, 64, 64, 70, 64, 64, 77, 90, 64, 93, 64,
	64, 64, 64, 64, 64, 96, 64, 107, 64, 64, 64, 64, 64, 112,
	64, 64, 64, 64, 64, 117, 64, 64, 64, 64, 64, 64, 64, 64,
	64, 64, 128, 64, 64, 64, 64, 64, 64, 12, 2, 3, 1, 4,
	12, 5, 19, 6, 3, 7, 10, 8, 0, 4, 11, 11, 17, 12,
	6, 2, 13, 3, 7, 10, 8, 12, 7, 17, 4, 18, 14, 19,
	13, 20, 6, 21, 5, 22, 2, 18, 24, 2, 11, 26, 10, 7,
	17, 15, 30, 9, 31, 6, 21, 5, 32, 4, 8, 38, 16, 39,
	4, 7, 17, 6, 40, 10, 7, 17, 4, 43, 14, 19, 6, 21,
	5, 22, 4, 7, 17, 6, 47,
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
		return p.on_term__token(
			p.sym.Peek(0).(Token),
		)
	case 8:
		return p.on_term__token(
			p.sym.Peek(0).(Token),
		)
	case 9:
		return p.on_term__list(
			p.sym.Peek(0).(*_i0.Term),
		)
	case 10:
		return p.on_list(
			p.sym.Peek(5).(Token),
			p.sym.Peek(4).(Token),
			p.sym.Peek(3).(*_i0.Term),
			p.sym.Peek(2).(Token),
			p.sym.Peek(1).(*_i0.Term),
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
		return p.on_card(
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
		return p.on_qualif(
			p.sym.Peek(3).(Token),
			p.sym.Peek(2).(Token),
			p.sym.Peek(1).(Token),
			p.sym.Peek(0).(Token),
		)
	case 16:
		return p.on_token_decl(
			p.sym.Peek(2).(Token),
			p.sym.Peek(1).([]*_i0.CustomToken),
			p.sym.Peek(0).(Token),
		)
	case 17:
		return p.on_token(
			p.sym.Peek(1).(Token),
			p.sym.Peek(0).(Token),
		)
	case 18: // ZeroOrOne
		return p.sym.Peek(0).([]_i0.ParserDecl)
	case 19: // ZeroOrOne
		{
			var zero []_i0.ParserDecl
			return zero
		}
	case 20: // List
		return append(
			p.sym.Peek(2).([]*_i0.Prod),
			p.sym.Peek(0).(*_i0.Prod),
		)
	case 21: // List
		return []*_i0.Prod{
			p.sym.Peek(0).(*_i0.Prod),
		}
	case 22: // OneOrMore
		return append(
			p.sym.Peek(1).([]*_i0.Term),
			p.sym.Peek(0).(*_i0.Term),
		)
	case 23: // OneOrMore
		return []*_i0.Term{
			p.sym.Peek(0).(*_i0.Term),
		}
	case 24: // ZeroOrOne
		return p.sym.Peek(0).(*_i0.ProdQualifier)
	case 25: // ZeroOrOne
		{
			var zero *_i0.ProdQualifier
			return zero
		}
	case 26: // ZeroOrOne
		return p.sym.Peek(0).(_i0.TermType)
	case 27: // ZeroOrOne
		{
			var zero _i0.TermType
			return zero
		}
	case 28: // OneOrMore
		return append(
			p.sym.Peek(1).([]*_i0.CustomToken),
			p.sym.Peek(0).(*_i0.CustomToken),
		)
	case 29: // OneOrMore
		return []*_i0.CustomToken{
			p.sym.Peek(0).(*_i0.CustomToken),
		}
	case 30: // ZeroOrOne
		return p.sym.Peek(0).(Token)
	case 31: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 32: // OneOrMore
		return append(
			p.sym.Peek(1).([]_i0.ParserDecl),
			p.sym.Peek(0).(_i0.ParserDecl),
		)
	case 33: // OneOrMore
		return []_i0.ParserDecl{
			p.sym.Peek(0).(_i0.ParserDecl),
		}
	default:
		panic("unreachable")
	}
}
