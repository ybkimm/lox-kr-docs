package parser

import (
	_i0 "errors"
)

var _LHS = []int32{
	0, 1, 2, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 11, 11, 11, 11, 12, 12, 12, 12, 13, 14, 14, 14, 14,
	14, 15, 16, 16, 16, 17, 18, 19, 20, 20, 21, 21, 22, 22,
	23, 23, 24, 24, 25, 25, 26, 26, 27, 27, 28, 28, 29, 29,
	30, 30, 31, 31,
}

var _TermCounts = []int32{
	1, 1, 1, 1, 1, 1, 1, 5, 5, 4, 5, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 4, 3, 3, 3, 3,
	1, 2, 1, 1, 1, 1, 4, 1, 1, 0, 1, 0, 1, 0,
	1, 0, 3, 1, 2, 1, 1, 0, 1, 0, 2, 1, 2, 1,
	2, 1, 2, 1,
}

var _Actions = []int32{
	84, 95, 106, 109, 112, 115, 128, 141, 152, 163, 166, 169, 180, 191,
	204, 95, 207, 236, 95, 265, 268, 297, 300, 305, 314, 323, 342, 371,
	390, 393, 404, 413, 418, 425, 428, 95, 95, 431, 450, 469, 488, 507,
	526, 545, 564, 583, 586, 591, 594, 603, 612, 615, 618, 631, 634, 663,
	670, 673, 680, 687, 696, 699, 708, 717, 726, 735, 744, 753, 762, 773,
	786, 799, 828, 835, 840, 845, 854, 857, 866, 873, 880, 887, 894, 897,
	10, 0, -37, 25, 1, 2, 2, 24, 3, 26, 4, 10, 2, 16,
	3, 17, 14, -51, 16, 18, 13, 19, 2, 6, 15, 2, 2, 28,
	2, 2, 14, 12, 10, -5, 0, -5, 25, -5, 2, -5, 24, -5,
	26, -5, 12, 10, -6, 0, -6, 25, -6, 2, -6, 24, -6, 26,
	-6, 10, 0, -3, 25, -3, 2, -3, 24, -3, 26, -3, 10, 0,
	-2, 25, -2, 2, -2, 24, -2, 26, -2, 2, 0, 2147483647, 2, 0,
	-1, 10, 0, -36, 25, 1, 2, 2, 24, 3, 26, 4, 10, 0,
	-57, 25, -57, 2, -57, 24, -57, 26, -57, 12, 10, -4, 0, -4,
	25, -4, 2, -4, 24, -4, 26, -4, 2, 9, 30, 28, 8, -20,
	17, -20, 2, -20, 3, -20, 14, -20, 20, -20, 22, -20, 16, -20,
	7, -20, 5, -20, 13, -20, 19, -20, 21, -20, 18, -20, 28, 8,
	-19, 17, -19, 2, -19, 3, -19, 14, -19, 20, -19, 22, -19, 16,
	-19, 7, -19, 5, -19, 13, -19, 19, -19, 21, -19, 18, -19, 2,
	14, -50, 28, 8, -21, 17, -21, 2, -21, 3, -21, 14, -21, 20,
	-21, 22, -21, 16, -21, 7, -21, 5, -21, 13, -21, 19, -21, 21,
	-21, 18, -21, 2, 14, 46, 4, 8, 32, 5, -43, 8, 8, -11,
	17, -11, 7, 36, 5, -11, 8, 8, -45, 17, -45, 7, -45, 5,
	-45, 18, 8, -12, 17, -12, 2, 16, 3, 17, 14, -51, 16, 18,
	7, -12, 5, -12, 13, 19, 28, 8, -49, 17, -49, 2, -49, 3,
	-49, 14, -49, 20, 38, 22, 39, 16, -49, 7, -49, 5, -49, 13,
	-49, 19, 40, 21, 41, 18, 42, 18, 8, -47, 17, -47, 2, -47,
	3, -47, 14, -47, 16, -47, 7, -47, 5, -47, 13, -47, 2, 6,
	35, 10, 0, -56, 25, -56, 2, -56, 24, -56, 26, -56, 8, 10,
	-39, 25, 1, 2, 2, 24, 3, 4, 8, 32, 5, -41, 6, 28,
	59, 27, 60, 23, 61, 2, 5, -42, 2, 5, 52, 18, 8, -46,
	17, -46, 2, -46, 3, -46, 14, -46, 16, -46, 7, -46, 5, -46,
	13, -46, 18, 8, -16, 17, -16, 2, -16, 3, -16, 14, -16, 16,
	-16, 7, -16, 5, -16, 13, -16, 18, 8, -18, 17, -18, 2, -18,
	3, -18, 14, -18, 16, -18, 7, -18, 5, -18, 13, -18, 18, 8,
	-15, 17, -15, 2, -15, 3, -15, 14, -15, 16, -15, 7, -15, 5,
	-15, 13, -15, 18, 8, -17, 17, -17, 2, -17, 3, -17, 14, -17,
	16, -17, 7, -17, 5, -17, 13, -17, 18, 8, -14, 17, -14, 2,
	-14, 3, -14, 14, -14, 16, -14, 7, -14, 5, -14, 13, -14, 18,
	8, -48, 17, -48, 2, -48, 3, -48, 14, -48, 16, -48, 7, -48,
	5, -48, 13, -48, 18, 8, -13, 17, -13, 2, -13, 3, -13, 14,
	-13, 16, -13, 7, -13, 5, -13, 13, -13, 2, 17, 54, 4, 4,
	55, 11, 56, 2, 10, 68, 8, 10, -38, 25, 1, 2, 2, 24,
	3, 8, 10, -59, 25, -59, 2, -59, 24, -59, 2, 5, -40, 2,
	5, 69, 12, 10, -9, 0, -9, 25, -9, 2, -9, 24, -9, 26,
	-9, 2, 5, 70, 28, 8, -22, 17, -22, 2, -22, 3, -22, 14,
	-22, 20, -22, 22, -22, 16, -22, 7, -22, 5, -22, 13, -22, 19,
	-22, 21, -22, 18, -22, 6, 15, -28, 4, -28, 11, -28, 2, 11,
	74, 6, 15, 71, 4, 55, 11, 56, 6, 15, -53, 4, -53, 11,
	-53, 8, 28, -35, 27, -35, 5, -35, 23, -35, 2, 16, 76, 8,
	28, -33, 27, -33, 5, -33, 23, -33, 8, 28, -55, 27, -55, 5,
	-55, 23, -55, 8, 28, -32, 27, -32, 5, -32, 23, -32, 8, 28,
	-31, 27, -31, 5, -31, 23, -31, 8, 28, -30, 27, -30, 5, -30,
	23, -30, 8, 28, 59, 27, 60, 5, -29, 23, 61, 8, 8, -44,
	17, -44, 7, -44, 5, -44, 10, 0, -7, 25, -7, 2, -7, 24,
	-7, 26, -7, 12, 10, -8, 0, -8, 25, -8, 2, -8, 24, -8,
	26, -8, 12, 10, -10, 0, -10, 25, -10, 2, -10, 24, -10, 26,
	-10, 28, 8, -23, 17, -23, 2, -23, 3, -23, 14, -23, 20, -23,
	22, -23, 16, -23, 7, -23, 5, -23, 13, -23, 19, -23, 21, -23,
	18, -23, 6, 15, -52, 4, -52, 11, -52, 4, 4, 78, 11, 79,
	4, 4, 80, 11, 81, 8, 28, -54, 27, -54, 5, -54, 23, -54,
	2, 2, 82, 8, 10, -58, 25, -58, 2, -58, 24, -58, 6, 15,
	-24, 4, -24, 11, -24, 6, 15, -25, 4, -25, 11, -25, 6, 15,
	-26, 4, -26, 11, -26, 6, 15, -27, 4, -27, 11, -27, 2, 17,
	83, 8, 28, -34, 27, -34, 5, -34, 23, -34,
}

var _Goto = []int32{
	84, 103, 120, 120, 120, 120, 120, 120, 120, 120, 120, 121, 120, 120,
	120, 134, 120, 120, 151, 120, 120, 120, 168, 120, 120, 173, 182, 120,
	120, 120, 187, 200, 205, 120, 120, 216, 233, 120, 120, 120, 120, 120,
	120, 120, 120, 120, 246, 120, 251, 120, 120, 120, 120, 120, 120, 120,
	120, 260, 120, 120, 120, 120, 120, 120, 120, 120, 263, 120, 120, 120,
	120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120,
	18, 6, 5, 7, 6, 4, 7, 3, 8, 1, 9, 20, 10, 30,
	11, 2, 12, 5, 13, 16, 13, 20, 27, 21, 8, 22, 24, 23,
	9, 24, 25, 25, 12, 26, 10, 27, 0, 12, 6, 5, 7, 6,
	4, 7, 3, 8, 2, 29, 5, 13, 16, 13, 20, 27, 21, 8,
	31, 24, 23, 9, 24, 25, 25, 12, 26, 10, 27, 16, 13, 20,
	27, 21, 8, 45, 24, 23, 9, 24, 25, 25, 12, 26, 10, 27,
	4, 15, 33, 23, 34, 8, 13, 20, 27, 21, 12, 26, 10, 37,
	4, 11, 43, 26, 44, 12, 6, 5, 7, 6, 21, 47, 31, 48,
	3, 49, 5, 13, 4, 15, 50, 22, 51, 10, 16, 62, 19, 63,
	18, 64, 17, 65, 29, 66, 16, 13, 20, 27, 21, 8, 53, 24,
	23, 9, 24, 25, 25, 12, 26, 10, 27, 12, 13, 20, 27, 21,
	9, 67, 25, 25, 12, 26, 10, 27, 4, 28, 57, 14, 58, 8,
	6, 5, 7, 6, 3, 77, 5, 13, 2, 14, 72, 8, 16, 75,
	19, 63, 18, 64, 17, 65,
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
	_lex   lexer
	_state _Stack[int32]
	_sym   _Stack[any]

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
			p._ReadToken()
		} else { // reduce
			prod := -action
			termCount := _TermCounts[int(prod)]
			rule := _LHS[int(prod)]
			res := p._Act(prod)
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

		for len(p._state) > 1 {
			topState := p._state.Peek(0)
			action, ok := _Find(_Actions, topState, int32(ERROR))
			if ok {
				action2, ok := _Find(
					_Actions, action, int32(p._lookaheadType))
				if ok {
					p._state.Push(action)
					p._sym.Push(_errorPlaceholder)
					return action2, true
				}
			}
			p._state.Pop(1)
			p._sym.Pop(1)
		}

		if p._lookaheadType == EOF {
			p.onError()
			return 0, false
		}

		p._ReadToken()
		p._state = saveState
		p._sym = saveSym
	}
}

func (p *parser) _Act(prod int32) any {
	switch prod {
	case 1:
		return p.on_spec(
			_cast[[]any](p._sym.Peek(0)),
		)
	case 2:
		return p.on_statement(
			_cast[any](p._sym.Peek(0)),
		)
	case 3:
		return p.on_statement(
			_cast[any](p._sym.Peek(0)),
		)
	case 4:
		return p.on_rule(
			_cast[any](p._sym.Peek(0)),
		)
	case 5:
		return p.on_rule(
			_cast[any](p._sym.Peek(0)),
		)
	case 6:
		return p.on_rule(
			_cast[any](p._sym.Peek(0)),
		)
	case 7:
		return p.on_mode(
			_cast[Token](p._sym.Peek(4)),
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[[]any](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 8:
		return p.on_token_rule(
			_cast[Token](p._sym.Peek(4)),
			_cast[Token](p._sym.Peek(3)),
			_cast[any](p._sym.Peek(2)),
			_cast[any](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 9:
		return p.on_frag_rule(
			_cast[Token](p._sym.Peek(3)),
			_cast[any](p._sym.Peek(2)),
			_cast[any](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 10:
		return p.on_macro_rule(
			_cast[Token](p._sym.Peek(4)),
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[any](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 11:
		return p.on_expr(
			_cast[[]any](p._sym.Peek(0)),
		)
	case 12:
		return p.on_factor(
			_cast[[]any](p._sym.Peek(0)),
		)
	case 13:
		return p.on_term_card(
			_cast[any](p._sym.Peek(1)),
			_cast[any](p._sym.Peek(0)),
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
		return p.on_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 18:
		return p.on_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 19:
		return p.on_term__tok(
			_cast[Token](p._sym.Peek(0)),
		)
	case 20:
		return p.on_term__tok(
			_cast[Token](p._sym.Peek(0)),
		)
	case 21:
		return p.on_term__char_class(
			_cast[any](p._sym.Peek(0)),
		)
	case 22:
		return p.on_term__expr(
			_cast[Token](p._sym.Peek(2)),
			_cast[any](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 23:
		return p.on_char_class(
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[[]any](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 24:
		return p.on_char_class_item__range(
			_cast[Token](p._sym.Peek(2)),
			_cast[Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 25:
		return p.on_char_class_item__range(
			_cast[Token](p._sym.Peek(2)),
			_cast[Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 26:
		return p.on_char_class_item__range(
			_cast[Token](p._sym.Peek(2)),
			_cast[Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 27:
		return p.on_char_class_item__range(
			_cast[Token](p._sym.Peek(2)),
			_cast[Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 28:
		return p.on_char_class_item__char(
			_cast[Token](p._sym.Peek(0)),
		)
	case 29:
		return p.on_actions(
			_cast[Token](p._sym.Peek(1)),
			_cast[[]any](p._sym.Peek(0)),
		)
	case 30:
		return p.on_action(
			_cast[any](p._sym.Peek(0)),
		)
	case 31:
		return p.on_action(
			_cast[any](p._sym.Peek(0)),
		)
	case 32:
		return p.on_action(
			_cast[any](p._sym.Peek(0)),
		)
	case 33:
		return p.on_action_skip(
			_cast[Token](p._sym.Peek(0)),
		)
	case 34:
		return p.on_action_push_mode(
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 35:
		return p.on_action_pop_mode(
			_cast[Token](p._sym.Peek(0)),
		)
	case 36: // ZeroOrOne
		return _cast[[]any](p._sym.Peek(0))
	case 37: // ZeroOrOne
		{
			var zero []any
			return zero
		}
	case 38: // ZeroOrOne
		return _cast[[]any](p._sym.Peek(0))
	case 39: // ZeroOrOne
		{
			var zero []any
			return zero
		}
	case 40: // ZeroOrOne
		return _cast[any](p._sym.Peek(0))
	case 41: // ZeroOrOne
		{
			var zero any
			return zero
		}
	case 42: // ZeroOrOne
		return _cast[any](p._sym.Peek(0))
	case 43: // ZeroOrOne
		{
			var zero any
			return zero
		}
	case 44: // List
		return append(
			_cast[[]any](p._sym.Peek(2)),
			_cast[any](p._sym.Peek(0)),
		)
	case 45: // List
		return []any{
			_cast[any](p._sym.Peek(0)),
		}
	case 46: // OneOrMore
		return append(
			_cast[[]any](p._sym.Peek(1)),
			_cast[any](p._sym.Peek(0)),
		)
	case 47: // OneOrMore
		return []any{
			_cast[any](p._sym.Peek(0)),
		}
	case 48: // ZeroOrOne
		return _cast[any](p._sym.Peek(0))
	case 49: // ZeroOrOne
		{
			var zero any
			return zero
		}
	case 50: // ZeroOrOne
		return _cast[Token](p._sym.Peek(0))
	case 51: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 52: // OneOrMore
		return append(
			_cast[[]any](p._sym.Peek(1)),
			_cast[any](p._sym.Peek(0)),
		)
	case 53: // OneOrMore
		return []any{
			_cast[any](p._sym.Peek(0)),
		}
	case 54: // OneOrMore
		return append(
			_cast[[]any](p._sym.Peek(1)),
			_cast[any](p._sym.Peek(0)),
		)
	case 55: // OneOrMore
		return []any{
			_cast[any](p._sym.Peek(0)),
		}
	case 56: // OneOrMore
		return append(
			_cast[[]any](p._sym.Peek(1)),
			_cast[any](p._sym.Peek(0)),
		)
	case 57: // OneOrMore
		return []any{
			_cast[any](p._sym.Peek(0)),
		}
	case 58: // OneOrMore
		return append(
			_cast[[]any](p._sym.Peek(1)),
			_cast[any](p._sym.Peek(0)),
		)
	case 59: // OneOrMore
		return []any{
			_cast[any](p._sym.Peek(0)),
		}
	default:
		panic("unreachable")
	}
}
