package parser

import (
	_i0 "errors"
	_i1 "github.com/dcaiafa/lox/internal/lexergen/ast"
)

var _LHS = []int32{
	0, 1, 2, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 11, 11, 12, 12, 12, 12, 13, 14, 14, 15, 15, 15, 16,
	17, 18, 19, 19, 20, 20, 21, 21, 22, 22, 23, 23, 24, 24,
	25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31,
}

var _TermCounts = []int32{
	1, 1, 1, 1, 1, 1, 1, 5, 5, 4, 5, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 3, 4, 1, 1, 1, 1, 1, 1,
	4, 1, 1, 0, 1, 0, 1, 0, 1, 0, 3, 1, 2, 1,
	1, 0, 1, 0, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
}

var _Actions = []int32{
	76, 87, 98, 101, 104, 107, 120, 133, 144, 155, 158, 161, 172, 183,
	196, 87, 199, 228, 87, 257, 260, 289, 292, 301, 314, 327, 350, 379,
	402, 405, 416, 425, 434, 443, 446, 455, 464, 473, 482, 491, 494, 87,
	87, 503, 526, 549, 572, 595, 618, 641, 644, 649, 652, 661, 670, 679,
	682, 691, 704, 707, 736, 743, 750, 757, 764, 767, 776, 789, 800, 813,
	826, 855, 862, 865, 874, 883, 10, 0, -31, 25, 1, 2, 2, 24,
	3, 26, 4, 10, 2, 16, 3, 17, 14, -45, 16, 18, 13, 19,
	2, 6, 15, 2, 2, 28, 2, 2, 14, 12, 10, -5, 0, -5,
	25, -5, 2, -5, 24, -5, 26, -5, 12, 10, -6, 0, -6, 25,
	-6, 2, -6, 24, -6, 26, -6, 10, 0, -2, 25, -2, 2, -2,
	24, -2, 26, -2, 10, 0, -3, 25, -3, 2, -3, 24, -3, 26,
	-3, 2, 0, 2147483647, 2, 0, -1, 10, 0, -30, 25, 1, 2, 2,
	24, 3, 26, 4, 10, 0, -49, 25, -49, 2, -49, 24, -49, 26,
	-49, 12, 10, -4, 0, -4, 25, -4, 2, -4, 24, -4, 26, -4,
	2, 9, 30, 28, 17, -18, 2, -18, 3, -18, 14, -18, 20, -18,
	16, -18, 7, -18, 28, -18, 27, -18, 5, -18, 23, -18, 13, -18,
	19, -18, 18, -18, 28, 17, -17, 2, -17, 3, -17, 14, -17, 20,
	-17, 16, -17, 7, -17, 28, -17, 27, -17, 5, -17, 23, -17, 13,
	-17, 19, -17, 18, -17, 2, 14, -44, 28, 17, -19, 2, -19, 3,
	-19, 14, -19, 20, -19, 16, -19, 7, -19, 28, -19, 27, -19, 5,
	-19, 23, -19, 13, -19, 19, -19, 18, -19, 2, 14, 50, 8, 28,
	32, 27, 33, 5, -37, 23, 34, 12, 17, -11, 7, 42, 28, -11,
	27, -11, 5, -11, 23, -11, 12, 17, -39, 7, -39, 28, -39, 27,
	-39, 5, -39, 23, -39, 22, 17, -12, 2, 16, 3, 17, 14, -45,
	16, 18, 7, -12, 28, -12, 27, -12, 5, -12, 23, -12, 13, 19,
	28, 17, -43, 2, -43, 3, -43, 14, -43, 20, 44, 16, -43, 7,
	-43, 28, -43, 27, -43, 5, -43, 23, -43, 13, -43, 19, 45, 18,
	46, 22, 17, -41, 2, -41, 3, -41, 14, -41, 16, -41, 7, -41,
	28, -41, 27, -41, 5, -41, 23, -41, 13, -41, 2, 6, 41, 10,
	0, -48, 25, -48, 2, -48, 24, -48, 26, -48, 8, 10, -33, 25,
	1, 2, 2, 24, 3, 8, 28, 32, 27, 33, 5, -35, 23, 34,
	8, 28, -29, 27, -29, 5, -29, 23, -29, 2, 16, 64, 8, 28,
	-27, 27, -27, 5, -27, 23, -27, 8, 28, -55, 27, -55, 5, -55,
	23, -55, 8, 28, -26, 27, -26, 5, -26, 23, -26, 8, 28, -25,
	27, -25, 5, -25, 23, -25, 8, 28, -24, 27, -24, 5, -24, 23,
	-24, 2, 5, 57, 8, 28, 32, 27, 33, 5, -36, 23, 34, 22,
	17, -40, 2, -40, 3, -40, 14, -40, 16, -40, 7, -40, 28, -40,
	27, -40, 5, -40, 23, -40, 13, -40, 22, 17, -16, 2, -16, 3,
	-16, 14, -16, 16, -16, 7, -16, 28, -16, 27, -16, 5, -16, 23,
	-16, 13, -16, 22, 17, -15, 2, -15, 3, -15, 14, -15, 16, -15,
	7, -15, 28, -15, 27, -15, 5, -15, 23, -15, 13, -15, 22, 17,
	-14, 2, -14, 3, -14, 14, -14, 16, -14, 7, -14, 28, -14, 27,
	-14, 5, -14, 23, -14, 13, -14, 22, 17, -42, 2, -42, 3, -42,
	14, -42, 16, -42, 7, -42, 28, -42, 27, -42, 5, -42, 23, -42,
	13, -42, 22, 17, -13, 2, -13, 3, -13, 14, -13, 16, -13, 7,
	-13, 28, -13, 27, -13, 5, -13, 23, -13, 13, -13, 2, 17, 59,
	4, 4, 60, 11, 61, 2, 10, 67, 8, 10, -32, 25, 1, 2,
	2, 24, 3, 8, 10, -51, 25, -51, 2, -51, 24, -51, 8, 28,
	-53, 27, -53, 5, -53, 23, -53, 2, 5, 68, 8, 28, 32, 27,
	33, 5, -34, 23, 34, 12, 10, -9, 0, -9, 25, -9, 2, -9,
	24, -9, 26, -9, 2, 5, 69, 28, 17, -20, 2, -20, 3, -20,
	14, -20, 20, -20, 16, -20, 7, -20, 28, -20, 27, -20, 5, -20,
	23, -20, 13, -20, 19, -20, 18, -20, 6, 15, -22, 4, -22, 11,
	-22, 6, 15, -23, 4, -23, 11, -23, 6, 15, 70, 4, 60, 11,
	61, 6, 15, -47, 4, -47, 11, -47, 2, 2, 72, 8, 28, -54,
	27, -54, 5, -54, 23, -54, 12, 17, -38, 7, -38, 28, -38, 27,
	-38, 5, -38, 23, -38, 10, 0, -7, 25, -7, 2, -7, 24, -7,
	26, -7, 12, 10, -8, 0, -8, 25, -8, 2, -8, 24, -8, 26,
	-8, 12, 10, -10, 0, -10, 25, -10, 2, -10, 24, -10, 26, -10,
	28, 17, -21, 2, -21, 3, -21, 14, -21, 20, -21, 16, -21, 7,
	-21, 28, -21, 27, -21, 5, -21, 23, -21, 13, -21, 19, -21, 18,
	-21, 6, 15, -46, 4, -46, 11, -46, 2, 17, 75, 8, 10, -50,
	25, -50, 2, -50, 24, -50, 8, 28, -52, 27, -52, 5, -52, 23,
	-52, 8, 28, -28, 27, -28, 5, -28, 23, -28,
}

var _Goto = []int32{
	76, 95, 112, 112, 112, 112, 112, 112, 112, 112, 112, 113, 112, 112,
	112, 126, 112, 112, 143, 112, 112, 112, 160, 112, 112, 173, 182, 112,
	112, 112, 187, 200, 112, 112, 112, 112, 112, 112, 112, 112, 213, 222,
	239, 112, 112, 112, 112, 112, 112, 112, 252, 112, 257, 112, 112, 112,
	266, 112, 112, 112, 112, 112, 275, 112, 112, 112, 112, 112, 112, 112,
	112, 112, 112, 112, 112, 112, 18, 6, 5, 7, 6, 4, 7, 3,
	8, 1, 9, 19, 10, 28, 11, 2, 12, 5, 13, 16, 13, 20,
	26, 21, 8, 22, 23, 23, 9, 24, 24, 25, 12, 26, 10, 27,
	0, 12, 6, 5, 7, 6, 4, 7, 3, 8, 2, 29, 5, 13,
	16, 13, 20, 26, 21, 8, 31, 23, 23, 9, 24, 24, 25, 12,
	26, 10, 27, 16, 13, 20, 26, 21, 8, 49, 23, 23, 9, 24,
	24, 25, 12, 26, 10, 27, 12, 15, 35, 18, 36, 17, 37, 16,
	38, 22, 39, 31, 40, 8, 13, 20, 26, 21, 12, 26, 10, 43,
	4, 11, 47, 25, 48, 12, 6, 5, 7, 6, 20, 51, 29, 52,
	3, 53, 5, 13, 12, 15, 54, 18, 36, 17, 37, 16, 38, 21,
	55, 30, 56, 8, 15, 65, 18, 36, 17, 37, 16, 38, 16, 13,
	20, 26, 21, 8, 58, 23, 23, 9, 24, 24, 25, 12, 26, 10,
	27, 12, 13, 20, 26, 21, 9, 66, 24, 25, 12, 26, 10, 27,
	4, 27, 62, 14, 63, 8, 6, 5, 7, 6, 3, 73, 5, 13,
	8, 15, 74, 18, 36, 17, 37, 16, 38, 2, 14, 71,
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
		return p.on_spec(
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		)
	case 2:
		return p.on_statement(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 3:
		return p.on_statement(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 4:
		return p.on_rule(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 5:
		return p.on_rule(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 6:
		return p.on_rule(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 7:
		return p.on_mode(
			_cast[Token](p._sym.Peek(4)),
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 8:
		return p.on_token_rule(
			_cast[Token](p._sym.Peek(4)),
			_cast[Token](p._sym.Peek(3)),
			_cast[*_i1.Expr](p._sym.Peek(2)),
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 9:
		return p.on_frag_rule(
			_cast[Token](p._sym.Peek(3)),
			_cast[*_i1.Expr](p._sym.Peek(2)),
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 10:
		return p.on_macro_rule(
			_cast[Token](p._sym.Peek(4)),
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[*_i1.Expr](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 11:
		return p.on_expr(
			_cast[[]*_i1.Factor](p._sym.Peek(0)),
		)
	case 12:
		return p.on_factor(
			_cast[[]*_i1.TermCard](p._sym.Peek(0)),
		)
	case 13:
		return p.on_term_card(
			_cast[_i1.Term](p._sym.Peek(1)),
			_cast[_i1.Card](p._sym.Peek(0)),
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
		return p.on_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 17:
		return p.on_term__tok(
			_cast[Token](p._sym.Peek(0)),
		)
	case 18:
		return p.on_term__tok(
			_cast[Token](p._sym.Peek(0)),
		)
	case 19:
		return p.on_term__char_class(
			_cast[*_i1.TermCharClass](p._sym.Peek(0)),
		)
	case 20:
		return p.on_term__expr(
			_cast[Token](p._sym.Peek(2)),
			_cast[*_i1.Expr](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 21:
		return p.on_char_class(
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[[]Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 22:
		return p.on_char_class_item(
			_cast[Token](p._sym.Peek(0)),
		)
	case 23:
		return p.on_char_class_item(
			_cast[Token](p._sym.Peek(0)),
		)
	case 24:
		return p.on_action(
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 25:
		return p.on_action(
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 26:
		return p.on_action(
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 27:
		return p.on_action_skip(
			_cast[Token](p._sym.Peek(0)),
		)
	case 28:
		return p.on_action_push_mode(
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 29:
		return p.on_action_pop_mode(
			_cast[Token](p._sym.Peek(0)),
		)
	case 30: // ZeroOrOne
		return _cast[[]_i1.Statement](p._sym.Peek(0))
	case 31: // ZeroOrOne
		{
			var zero []_i1.Statement
			return zero
		}
	case 32: // ZeroOrOne
		return _cast[[]_i1.Statement](p._sym.Peek(0))
	case 33: // ZeroOrOne
		{
			var zero []_i1.Statement
			return zero
		}
	case 34: // ZeroOrOne
		return _cast[[]_i1.Action](p._sym.Peek(0))
	case 35: // ZeroOrOne
		{
			var zero []_i1.Action
			return zero
		}
	case 36: // ZeroOrOne
		return _cast[[]_i1.Action](p._sym.Peek(0))
	case 37: // ZeroOrOne
		{
			var zero []_i1.Action
			return zero
		}
	case 38: // List
		return append(
			_cast[[]*_i1.Factor](p._sym.Peek(2)),
			_cast[*_i1.Factor](p._sym.Peek(0)),
		)
	case 39: // List
		return []*_i1.Factor{
			_cast[*_i1.Factor](p._sym.Peek(0)),
		}
	case 40: // OneOrMore
		return append(
			_cast[[]*_i1.TermCard](p._sym.Peek(1)),
			_cast[*_i1.TermCard](p._sym.Peek(0)),
		)
	case 41: // OneOrMore
		return []*_i1.TermCard{
			_cast[*_i1.TermCard](p._sym.Peek(0)),
		}
	case 42: // ZeroOrOne
		return _cast[_i1.Card](p._sym.Peek(0))
	case 43: // ZeroOrOne
		{
			var zero _i1.Card
			return zero
		}
	case 44: // ZeroOrOne
		return _cast[Token](p._sym.Peek(0))
	case 45: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 46: // OneOrMore
		return append(
			_cast[[]Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 47: // OneOrMore
		return []Token{
			_cast[Token](p._sym.Peek(0)),
		}
	case 48: // OneOrMore
		return append(
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 49: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._sym.Peek(0)),
		}
	case 50: // OneOrMore
		return append(
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 51: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._sym.Peek(0)),
		}
	case 52: // OneOrMore
		return append(
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 53: // OneOrMore
		return []_i1.Action{
			_cast[_i1.Action](p._sym.Peek(0)),
		}
	case 54: // OneOrMore
		return append(
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 55: // OneOrMore
		return []_i1.Action{
			_cast[_i1.Action](p._sym.Peek(0)),
		}
	default:
		panic("unreachable")
	}
}
