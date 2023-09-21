package parser

import (
	_i0 "errors"
)

var _LHS = []int32{
	0, 1, 2, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 11, 11, 11, 11, 12, 12, 12, 12, 13, 14, 14, 15, 16,
	16, 16, 17, 18, 19, 20, 20, 21, 21, 22, 22, 23, 23, 24,
	24, 25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31,
	31,
}

var _TermCounts = []int32{
	1, 1, 1, 1, 1, 1, 1, 5, 5, 4, 5, 1, 1, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 4, 3, 1, 2, 1,
	1, 1, 1, 4, 1, 1, 0, 1, 0, 1, 0, 1, 0, 3,
	1, 2, 1, 1, 0, 1, 0, 2, 1, 2, 1, 2, 1, 2,
	1,
}

var _Actions = []int32{
	79, 90, 101, 104, 107, 110, 123, 136, 147, 158, 161, 164, 175, 186,
	199, 90, 202, 231, 90, 260, 263, 292, 295, 300, 309, 318, 337, 366,
	385, 388, 399, 408, 413, 420, 423, 90, 90, 426, 445, 464, 483, 502,
	521, 540, 559, 578, 581, 584, 587, 596, 605, 608, 611, 624, 627, 656,
	663, 668, 673, 682, 685, 694, 703, 712, 721, 730, 739, 748, 759, 772,
	785, 814, 819, 822, 831, 834, 843, 848, 851, 10, 0, -34, 25, 1,
	2, 2, 24, 3, 26, 4, 10, 2, 16, 3, 17, 14, -48, 16,
	18, 13, 19, 2, 6, 15, 2, 2, 28, 2, 2, 14, 12, 10,
	-5, 0, -5, 25, -5, 2, -5, 24, -5, 26, -5, 12, 10, -6,
	0, -6, 25, -6, 2, -6, 24, -6, 26, -6, 10, 0, -3, 25,
	-3, 2, -3, 24, -3, 26, -3, 10, 0, -2, 25, -2, 2, -2,
	24, -2, 26, -2, 2, 0, 2147483647, 2, 0, -1, 10, 0, -33, 25,
	1, 2, 2, 24, 3, 26, 4, 10, 0, -54, 25, -54, 2, -54,
	24, -54, 26, -54, 12, 10, -4, 0, -4, 25, -4, 2, -4, 24,
	-4, 26, -4, 2, 9, 30, 28, 8, -20, 17, -20, 2, -20, 3,
	-20, 14, -20, 20, -20, 22, -20, 16, -20, 7, -20, 5, -20, 13,
	-20, 19, -20, 21, -20, 18, -20, 28, 8, -19, 17, -19, 2, -19,
	3, -19, 14, -19, 20, -19, 22, -19, 16, -19, 7, -19, 5, -19,
	13, -19, 19, -19, 21, -19, 18, -19, 2, 14, -47, 28, 8, -21,
	17, -21, 2, -21, 3, -21, 14, -21, 20, -21, 22, -21, 16, -21,
	7, -21, 5, -21, 13, -21, 19, -21, 21, -21, 18, -21, 2, 14,
	46, 4, 8, 32, 5, -40, 8, 8, -11, 17, -11, 7, 36, 5,
	-11, 8, 8, -42, 17, -42, 7, -42, 5, -42, 18, 8, -12, 17,
	-12, 2, 16, 3, 17, 14, -48, 16, 18, 7, -12, 5, -12, 13,
	19, 28, 8, -46, 17, -46, 2, -46, 3, -46, 14, -46, 20, 38,
	22, 39, 16, -46, 7, -46, 5, -46, 13, -46, 19, 40, 21, 41,
	18, 42, 18, 8, -44, 17, -44, 2, -44, 3, -44, 14, -44, 16,
	-44, 7, -44, 5, -44, 13, -44, 2, 6, 35, 10, 0, -53, 25,
	-53, 2, -53, 24, -53, 26, -53, 8, 10, -36, 25, 1, 2, 2,
	24, 3, 4, 8, 32, 5, -38, 6, 28, 58, 27, 59, 23, 60,
	2, 5, -39, 2, 5, 52, 18, 8, -43, 17, -43, 2, -43, 3,
	-43, 14, -43, 16, -43, 7, -43, 5, -43, 13, -43, 18, 8, -16,
	17, -16, 2, -16, 3, -16, 14, -16, 16, -16, 7, -16, 5, -16,
	13, -16, 18, 8, -18, 17, -18, 2, -18, 3, -18, 14, -18, 16,
	-18, 7, -18, 5, -18, 13, -18, 18, 8, -15, 17, -15, 2, -15,
	3, -15, 14, -15, 16, -15, 7, -15, 5, -15, 13, -15, 18, 8,
	-17, 17, -17, 2, -17, 3, -17, 14, -17, 16, -17, 7, -17, 5,
	-17, 13, -17, 18, 8, -14, 17, -14, 2, -14, 3, -14, 14, -14,
	16, -14, 7, -14, 5, -14, 13, -14, 18, 8, -45, 17, -45, 2,
	-45, 3, -45, 14, -45, 16, -45, 7, -45, 5, -45, 13, -45, 18,
	8, -13, 17, -13, 2, -13, 3, -13, 14, -13, 16, -13, 7, -13,
	5, -13, 13, -13, 2, 17, 54, 2, 4, 55, 2, 10, 67, 8,
	10, -35, 25, 1, 2, 2, 24, 3, 8, 10, -56, 25, -56, 2,
	-56, 24, -56, 2, 5, -37, 2, 5, 68, 12, 10, -9, 0, -9,
	25, -9, 2, -9, 24, -9, 26, -9, 2, 5, 69, 28, 8, -22,
	17, -22, 2, -22, 3, -22, 14, -22, 20, -22, 22, -22, 16, -22,
	7, -22, 5, -22, 13, -22, 19, -22, 21, -22, 18, -22, 6, 15,
	-25, 4, -25, 11, 72, 4, 15, 70, 4, 55, 4, 15, -50, 4,
	-50, 8, 28, -32, 27, -32, 5, -32, 23, -32, 2, 16, 74, 8,
	28, -30, 27, -30, 5, -30, 23, -30, 8, 28, -52, 27, -52, 5,
	-52, 23, -52, 8, 28, -29, 27, -29, 5, -29, 23, -29, 8, 28,
	-28, 27, -28, 5, -28, 23, -28, 8, 28, -27, 27, -27, 5, -27,
	23, -27, 8, 28, 58, 27, 59, 5, -26, 23, 60, 8, 8, -41,
	17, -41, 7, -41, 5, -41, 10, 0, -7, 25, -7, 2, -7, 24,
	-7, 26, -7, 12, 10, -8, 0, -8, 25, -8, 2, -8, 24, -8,
	26, -8, 12, 10, -10, 0, -10, 25, -10, 2, -10, 24, -10, 26,
	-10, 28, 8, -23, 17, -23, 2, -23, 3, -23, 14, -23, 20, -23,
	22, -23, 16, -23, 7, -23, 5, -23, 13, -23, 19, -23, 21, -23,
	18, -23, 4, 15, -49, 4, -49, 2, 4, 76, 8, 28, -51, 27,
	-51, 5, -51, 23, -51, 2, 2, 77, 8, 10, -55, 25, -55, 2,
	-55, 24, -55, 4, 15, -24, 4, -24, 2, 17, 78, 8, 28, -31,
	27, -31, 5, -31, 23, -31,
}

var _Goto = []int32{
	79, 98, 115, 115, 115, 115, 115, 115, 115, 115, 115, 116, 115, 115,
	115, 129, 115, 115, 146, 115, 115, 115, 163, 115, 115, 168, 177, 115,
	115, 115, 182, 195, 200, 115, 115, 211, 228, 115, 115, 115, 115, 115,
	115, 115, 115, 115, 241, 115, 246, 115, 115, 115, 115, 115, 115, 115,
	255, 115, 115, 115, 115, 115, 115, 115, 115, 258, 115, 115, 115, 115,
	115, 115, 115, 115, 115, 115, 115, 115, 115, 18, 6, 5, 7, 6,
	4, 7, 3, 8, 1, 9, 20, 10, 30, 11, 2, 12, 5, 13,
	16, 13, 20, 27, 21, 8, 22, 24, 23, 9, 24, 25, 25, 12,
	26, 10, 27, 0, 12, 6, 5, 7, 6, 4, 7, 3, 8, 2,
	29, 5, 13, 16, 13, 20, 27, 21, 8, 31, 24, 23, 9, 24,
	25, 25, 12, 26, 10, 27, 16, 13, 20, 27, 21, 8, 45, 24,
	23, 9, 24, 25, 25, 12, 26, 10, 27, 4, 15, 33, 23, 34,
	8, 13, 20, 27, 21, 12, 26, 10, 37, 4, 11, 43, 26, 44,
	12, 6, 5, 7, 6, 21, 47, 31, 48, 3, 49, 5, 13, 4,
	15, 50, 22, 51, 10, 16, 61, 19, 62, 18, 63, 17, 64, 29,
	65, 16, 13, 20, 27, 21, 8, 53, 24, 23, 9, 24, 25, 25,
	12, 26, 10, 27, 12, 13, 20, 27, 21, 9, 66, 25, 25, 12,
	26, 10, 27, 4, 28, 56, 14, 57, 8, 6, 5, 7, 6, 3,
	75, 5, 13, 2, 14, 71, 8, 16, 73, 19, 62, 18, 63, 17,
	64,
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
		return p.on_char_class_item__char(
			_cast[Token](p._sym.Peek(0)),
		)
	case 26:
		return p.on_actions(
			_cast[Token](p._sym.Peek(1)),
			_cast[[]any](p._sym.Peek(0)),
		)
	case 27:
		return p.on_action(
			_cast[any](p._sym.Peek(0)),
		)
	case 28:
		return p.on_action(
			_cast[any](p._sym.Peek(0)),
		)
	case 29:
		return p.on_action(
			_cast[any](p._sym.Peek(0)),
		)
	case 30:
		return p.on_action_skip(
			_cast[Token](p._sym.Peek(0)),
		)
	case 31:
		return p.on_action_push_mode(
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 32:
		return p.on_action_pop_mode(
			_cast[Token](p._sym.Peek(0)),
		)
	case 33: // ZeroOrOne
		return _cast[[]any](p._sym.Peek(0))
	case 34: // ZeroOrOne
		{
			var zero []any
			return zero
		}
	case 35: // ZeroOrOne
		return _cast[[]any](p._sym.Peek(0))
	case 36: // ZeroOrOne
		{
			var zero []any
			return zero
		}
	case 37: // ZeroOrOne
		return _cast[any](p._sym.Peek(0))
	case 38: // ZeroOrOne
		{
			var zero any
			return zero
		}
	case 39: // ZeroOrOne
		return _cast[any](p._sym.Peek(0))
	case 40: // ZeroOrOne
		{
			var zero any
			return zero
		}
	case 41: // List
		return append(
			_cast[[]any](p._sym.Peek(2)),
			_cast[any](p._sym.Peek(0)),
		)
	case 42: // List
		return []any{
			_cast[any](p._sym.Peek(0)),
		}
	case 43: // OneOrMore
		return append(
			_cast[[]any](p._sym.Peek(1)),
			_cast[any](p._sym.Peek(0)),
		)
	case 44: // OneOrMore
		return []any{
			_cast[any](p._sym.Peek(0)),
		}
	case 45: // ZeroOrOne
		return _cast[any](p._sym.Peek(0))
	case 46: // ZeroOrOne
		{
			var zero any
			return zero
		}
	case 47: // ZeroOrOne
		return _cast[Token](p._sym.Peek(0))
	case 48: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 49: // OneOrMore
		return append(
			_cast[[]any](p._sym.Peek(1)),
			_cast[any](p._sym.Peek(0)),
		)
	case 50: // OneOrMore
		return []any{
			_cast[any](p._sym.Peek(0)),
		}
	case 51: // OneOrMore
		return append(
			_cast[[]any](p._sym.Peek(1)),
			_cast[any](p._sym.Peek(0)),
		)
	case 52: // OneOrMore
		return []any{
			_cast[any](p._sym.Peek(0)),
		}
	case 53: // OneOrMore
		return append(
			_cast[[]any](p._sym.Peek(1)),
			_cast[any](p._sym.Peek(0)),
		)
	case 54: // OneOrMore
		return []any{
			_cast[any](p._sym.Peek(0)),
		}
	case 55: // OneOrMore
		return append(
			_cast[[]any](p._sym.Peek(1)),
			_cast[any](p._sym.Peek(0)),
		)
	case 56: // OneOrMore
		return []any{
			_cast[any](p._sym.Peek(0)),
		}
	default:
		panic("unreachable")
	}
}
