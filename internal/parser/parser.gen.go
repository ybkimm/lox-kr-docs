package parser

import (
	_i0 "errors"
	_i1 "github.com/dcaiafa/lox/internal/ast"
)

var _LHS = []int32{
	0, 1, 2, 2, 2, 3, 4, 5, 6, 6, 6, 6, 7, 8,
	8, 8, 9, 9, 10, 11, 12, 12, 13, 13, 14, 14, 15, 15,
	16, 16, 17, 17, 18, 18, 19, 19,
}

var _TermCounts = []int32{
	1, 1, 1, 1, 2, 4, 2, 2, 1, 1, 1, 1, 6, 1,
	1, 1, 4, 4, 3, 2, 1, 0, 3, 1, 2, 1, 1, 0,
	1, 0, 2, 1, 1, 0, 2, 1,
}

var _Actions = []int32{
	54, 63, 66, 69, 72, 81, 84, 87, 96, 105, 114, 123, 132, 139,
	144, 149, 158, 185, 212, 215, 242, 269, 274, 291, 296, 319, 336, 341,
	346, 355, 123, 360, 363, 366, 371, 376, 123, 393, 402, 419, 436, 453,
	470, 487, 490, 493, 496, 123, 501, 504, 507, 510, 515, 520, 8, 0,
	-21, 1, 1, 2, 2, 18, 3, 2, 13, 10, 2, 10, 11, 2,
	2, 12, 8, 0, -35, 1, -35, 2, -35, 18, -35, 2, 0, 2147483647,
	2, 0, -1, 8, 0, -20, 1, 1, 2, 2, 18, 3, 8, 0,
	-2, 1, -2, 2, -2, 18, -2, 8, 0, -3, 1, -3, 2, -3,
	18, -3, 8, 0, -4, 1, -4, 2, -4, 18, -4, 8, 14, 16,
	2, 17, 16, 18, 3, 19, 6, 2, -33, 3, 26, 13, -33, 4,
	2, -31, 13, -31, 4, 2, 12, 13, 28, 8, 0, -34, 1, -34,
	2, -34, 18, -34, 26, 8, -11, 9, -11, 14, -11, 2, -11, 15,
	-11, 16, -11, 3, -11, 6, -11, 12, -11, 17, -11, 13, -11, 5,
	-11, 7, -11, 26, 8, -8, 9, -8, 14, -8, 2, -8, 15, -8,
	16, -8, 3, -8, 6, -8, 12, -8, 17, -8, 13, -8, 5, -8,
	7, -8, 2, 11, 30, 26, 8, -9, 9, -9, 14, -9, 2, -9,
	15, -9, 16, -9, 3, -9, 6, -9, 12, -9, 17, -9, 13, -9,
	5, -9, 7, -9, 26, 8, -10, 9, -10, 14, -10, 2, -10, 15,
	-10, 16, -10, 3, -10, 6, -10, 12, -10, 17, -10, 13, -10, 5,
	-10, 7, -10, 4, 12, -23, 13, -23, 16, 14, 16, 2, 17, 15,
	31, 16, 18, 3, 19, 12, -27, 17, 32, 13, -27, 4, 12, 36,
	13, 37, 22, 14, -29, 2, -29, 15, -29, 16, -29, 3, -29, 6,
	38, 12, -29, 17, -29, 13, -29, 5, 39, 7, 40, 16, 14, -25,
	2, -25, 15, -25, 16, -25, 3, -25, 12, -25, 17, -25, 13, -25,
	4, 2, -32, 13, -32, 4, 2, -19, 13, -19, 8, 0, -18, 1,
	-18, 2, -18, 18, -18, 4, 2, -30, 13, -30, 2, 11, 44, 2,
	11, 45, 4, 12, -6, 13, -6, 4, 12, -26, 13, -26, 16, 14,
	-24, 2, -24, 15, -24, 16, -24, 3, -24, 12, -24, 17, -24, 13,
	-24, 8, 0, -5, 1, -5, 2, -5, 18, -5, 16, 14, -14, 2,
	-14, 15, -14, 16, -14, 3, -14, 12, -14, 17, -14, 13, -14, 16,
	14, -13, 2, -13, 15, -13, 16, -13, 3, -13, 12, -13, 17, -13,
	13, -13, 16, 14, -15, 2, -15, 15, -15, 16, -15, 3, -15, 12,
	-15, 17, -15, 13, -15, 16, 14, -28, 2, -28, 15, -28, 16, -28,
	3, -28, 12, -28, 17, -28, 13, -28, 16, 14, -7, 2, -7, 15,
	-7, 16, -7, 3, -7, 12, -7, 17, -7, 13, -7, 2, 8, 47,
	2, 4, 48, 2, 4, 49, 4, 12, -22, 13, -22, 2, 9, 51,
	2, 9, 52, 2, 9, 53, 4, 12, -16, 13, -16, 4, 12, -17,
	13, -17, 26, 8, -12, 9, -12, 14, -12, 2, -12, 15, -12, 16,
	-12, 3, -12, 6, -12, 12, -12, 17, -12, 13, -12, 5, -12, 7,
	-12,
}

var _Goto = []int32{
	54, 67, 67, 68, 67, 67, 67, 73, 67, 67, 67, 80, 93, 67,
	96, 67, 67, 67, 67, 67, 67, 67, 99, 67, 110, 67, 67, 67,
	67, 67, 115, 67, 67, 67, 67, 67, 120, 67, 67, 67, 67, 67,
	67, 67, 67, 67, 67, 131, 67, 67, 67, 67, 67, 67, 12, 2,
	4, 1, 5, 12, 6, 19, 7, 3, 8, 10, 9, 0, 4, 11,
	13, 17, 14, 6, 2, 15, 3, 8, 10, 9, 12, 7, 20, 4,
	21, 14, 22, 13, 23, 6, 24, 5, 25, 2, 18, 27, 2, 11,
	29, 10, 7, 20, 15, 33, 9, 34, 6, 24, 5, 35, 4, 8,
	41, 16, 42, 4, 7, 20, 6, 43, 10, 7, 20, 4, 46, 14,
	22, 6, 24, 5, 25, 4, 7, 20, 6, 50,
}

type _Stack[T any] []T

func (s *_Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *_Stack[T]) Pop(n int) {
	*s = (*s)[:len(*s)-n]
}

func (s _Stack[T]) Peek(n int) T {
	return s[len(s)-n-1]
}

func _cast[T any](v any) T {
	cv, _ := v.(T)
	return cv
}

var _errorPlaceholder = _i0.New("error placeholder")

func _Find(table []int32, y, x int32) (int32, bool) {
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

type lox struct {
	_lex    lexer
	_state  _Stack[int32]
	_sym    _Stack[any]
	_bounds _Stack[_Bounds]

	_lookahead     Token
	_lookaheadType TokenType
	_errorToken    Token
}

func (p *parser) parse(lex lexer) bool {
	const accept = 2147483647

	p._lex = lex

	p._state.Push(0)
	p._ReadToken()

	for {
		if p._lookaheadType == ERROR {
			_, ok := p._Recover()
			if !ok {
				return false
			}
		}
		topState := p._state.Peek(0)
		action, ok := _Find(
			_Actions, topState, int32(p._lookaheadType))
		if !ok {
			action, ok = p._Recover()
			if !ok {
				return false
			}
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			p._state.Push(action)
			p._sym.Push(p._lookahead)
			p._bounds.Push(
				_Bounds{Begin: p._lookahead,
					End: p._lookahead})
			p._ReadToken()
		} else { // reduce
			prod := -action
			termCount := _TermCounts[int(prod)]
			rule := _LHS[int(prod)]
			res := p._Act(prod)
			if termCount > 0 {
				bounds := _Bounds{
					Begin: p._bounds.Peek(int(termCount - 1)).Begin,
					End:   p._bounds.Peek(0).End,
				}
				p.onReduce(res, bounds.Begin, bounds.End)
				p._bounds.Pop(int(termCount))
				p._bounds.Push(bounds)
			} else {
				bounds := p._bounds.Peek(0)
				bounds.Begin = bounds.End
				p._bounds.Push(bounds)
			}
			p._state.Pop(int(termCount))
			p._sym.Pop(int(termCount))
			topState = p._state.Peek(0)
			nextState, _ := _Find(_Goto, topState, rule)
			p._state.Push(nextState)
			p._sym.Push(res)
		}
	}

	return true
}

func (p *parser) errorToken() Token {
	return p._errorToken
}

func (p *parser) _ReadToken() {
	p._lookahead, p._lookaheadType = p._lex.ReadToken()
}

func (p *parser) _Recover() (int32, bool) {
	p._errorToken = p._lookahead

	for {
		for p._lookaheadType == ERROR {
			p._ReadToken()
		}

		saveState := p._state
		saveSym := p._sym
		saveBounds := p._bounds

		for len(p._state) > 1 {
			topState := p._state.Peek(0)
			action, ok := _Find(_Actions, topState, int32(ERROR))
			if ok {
				action2, ok := _Find(
					_Actions, action, int32(p._lookaheadType))
				if ok {
					p._state.Push(action)
					p._sym.Push(_errorPlaceholder)
					p._bounds.Push(_Bounds{})
					return action2, true
				}
			}
			p._state.Pop(1)
			p._sym.Pop(1)
			p._bounds.Pop(1)
		}

		if p._lookaheadType == EOF {
			p.onError()
			return 0, false
		}

		p._ReadToken()
		p._state = saveState
		p._sym = saveSym
		p._bounds = saveBounds
	}
}

func (p *parser) _Act(prod int32) any {
	switch prod {
	case 1:
		return p.on_parser(
			_cast[[]_i1.ParserDecl](p._sym.Peek(0)),
		)
	case 2:
		return p.on_decl(
			_cast[_i1.ParserDecl](p._sym.Peek(0)),
		)
	case 3:
		return p.on_decl(
			_cast[_i1.ParserDecl](p._sym.Peek(0)),
		)
	case 4:
		return p.on_decl__error(
			_cast[error](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 5:
		return p.on_rule(
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[[]*_i1.Prod](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 6:
		return p.on_prod(
			_cast[[]*_i1.Term](p._sym.Peek(1)),
			_cast[*_i1.ProdQualifier](p._sym.Peek(0)),
		)
	case 7:
		return p.on_term_card(
			_cast[*_i1.Term](p._sym.Peek(1)),
			_cast[_i1.TermType](p._sym.Peek(0)),
		)
	case 8:
		return p.on_term__token(
			_cast[Token](p._sym.Peek(0)),
		)
	case 9:
		return p.on_term__token(
			_cast[Token](p._sym.Peek(0)),
		)
	case 10:
		return p.on_term__list(
			_cast[*_i1.Term](p._sym.Peek(0)),
		)
	case 11:
		return p.on_term__token(
			_cast[Token](p._sym.Peek(0)),
		)
	case 12:
		return p.on_list(
			_cast[Token](p._sym.Peek(5)),
			_cast[Token](p._sym.Peek(4)),
			_cast[*_i1.Term](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[*_i1.Term](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 13:
		return p.on_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 14:
		return p.on_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 15:
		return p.on_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 16:
		return p.on_qualif(
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 17:
		return p.on_qualif(
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 18:
		return p.on_token_decl(
			_cast[Token](p._sym.Peek(2)),
			_cast[[]*_i1.CustomToken](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 19:
		return p.on_token(
			_cast[Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 20: // ZeroOrOne
		return _cast[[]_i1.ParserDecl](p._sym.Peek(0))
	case 21: // ZeroOrOne
		{
			var zero []_i1.ParserDecl
			return zero
		}
	case 22: // List
		return append(
			_cast[[]*_i1.Prod](p._sym.Peek(2)),
			_cast[*_i1.Prod](p._sym.Peek(0)),
		)
	case 23: // List
		return []*_i1.Prod{
			_cast[*_i1.Prod](p._sym.Peek(0)),
		}
	case 24: // OneOrMore
		return append(
			_cast[[]*_i1.Term](p._sym.Peek(1)),
			_cast[*_i1.Term](p._sym.Peek(0)),
		)
	case 25: // OneOrMore
		return []*_i1.Term{
			_cast[*_i1.Term](p._sym.Peek(0)),
		}
	case 26: // ZeroOrOne
		return _cast[*_i1.ProdQualifier](p._sym.Peek(0))
	case 27: // ZeroOrOne
		{
			var zero *_i1.ProdQualifier
			return zero
		}
	case 28: // ZeroOrOne
		return _cast[_i1.TermType](p._sym.Peek(0))
	case 29: // ZeroOrOne
		{
			var zero _i1.TermType
			return zero
		}
	case 30: // OneOrMore
		return append(
			_cast[[]*_i1.CustomToken](p._sym.Peek(1)),
			_cast[*_i1.CustomToken](p._sym.Peek(0)),
		)
	case 31: // OneOrMore
		return []*_i1.CustomToken{
			_cast[*_i1.CustomToken](p._sym.Peek(0)),
		}
	case 32: // ZeroOrOne
		return _cast[Token](p._sym.Peek(0))
	case 33: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 34: // OneOrMore
		return append(
			_cast[[]_i1.ParserDecl](p._sym.Peek(1)),
			_cast[_i1.ParserDecl](p._sym.Peek(0)),
		)
	case 35: // OneOrMore
		return []_i1.ParserDecl{
			_cast[_i1.ParserDecl](p._sym.Peek(0)),
		}
	default:
		panic("unreachable")
	}
}
