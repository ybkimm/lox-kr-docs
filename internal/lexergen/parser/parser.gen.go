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
	76, 87, 98, 101, 104, 107, 120, 131, 144, 155, 158, 161, 172, 183,
	196, 87, 199, 228, 87, 257, 260, 289, 292, 301, 314, 327, 350, 379,
	402, 405, 416, 425, 434, 443, 446, 455, 464, 473, 482, 491, 494, 87,
	87, 503, 526, 549, 572, 595, 618, 641, 644, 649, 658, 661, 670, 679,
	682, 691, 704, 707, 736, 743, 750, 757, 764, 767, 776, 789, 800, 813,
	826, 855, 862, 865, 874, 883, 10, 0, -31, 24, 1, 2, 2, 23,
	3, 25, 4, 10, 2, 16, 3, 17, 13, -45, 15, 18, 12, 19,
	2, 6, 15, 2, 2, 28, 2, 2, 14, 12, 10, -5, 0, -5,
	24, -5, 2, -5, 23, -5, 25, -5, 10, 0, -3, 24, -3, 2,
	-3, 23, -3, 25, -3, 12, 10, -6, 0, -6, 24, -6, 2, -6,
	23, -6, 25, -6, 10, 0, -2, 24, -2, 2, -2, 23, -2, 25,
	-2, 2, 0, 2147483647, 2, 0, -1, 10, 0, -30, 24, 1, 2, 2,
	23, 3, 25, 4, 10, 0, -49, 24, -49, 2, -49, 23, -49, 25,
	-49, 12, 10, -4, 0, -4, 24, -4, 2, -4, 23, -4, 25, -4,
	2, 9, 30, 28, 16, -18, 2, -18, 3, -18, 13, -18, 19, -18,
	15, -18, 7, -18, 27, -18, 26, -18, 5, -18, 22, -18, 12, -18,
	18, -18, 17, -18, 28, 16, -17, 2, -17, 3, -17, 13, -17, 19,
	-17, 15, -17, 7, -17, 27, -17, 26, -17, 5, -17, 22, -17, 12,
	-17, 18, -17, 17, -17, 2, 13, -44, 28, 16, -19, 2, -19, 3,
	-19, 13, -19, 19, -19, 15, -19, 7, -19, 27, -19, 26, -19, 5,
	-19, 22, -19, 12, -19, 18, -19, 17, -19, 2, 13, 50, 8, 27,
	32, 26, 33, 5, -37, 22, 34, 12, 16, -11, 7, 42, 27, -11,
	26, -11, 5, -11, 22, -11, 12, 16, -39, 7, -39, 27, -39, 26,
	-39, 5, -39, 22, -39, 22, 16, -12, 2, 16, 3, 17, 13, -45,
	15, 18, 7, -12, 27, -12, 26, -12, 5, -12, 22, -12, 12, 19,
	28, 16, -43, 2, -43, 3, -43, 13, -43, 19, 44, 15, -43, 7,
	-43, 27, -43, 26, -43, 5, -43, 22, -43, 12, -43, 18, 45, 17,
	46, 22, 16, -41, 2, -41, 3, -41, 13, -41, 15, -41, 7, -41,
	27, -41, 26, -41, 5, -41, 22, -41, 12, -41, 2, 6, 41, 10,
	0, -48, 24, -48, 2, -48, 23, -48, 25, -48, 8, 10, -33, 24,
	1, 2, 2, 23, 3, 8, 27, 32, 26, 33, 5, -35, 22, 34,
	8, 27, -29, 26, -29, 5, -29, 22, -29, 2, 15, 64, 8, 27,
	-27, 26, -27, 5, -27, 22, -27, 8, 27, -55, 26, -55, 5, -55,
	22, -55, 8, 27, -26, 26, -26, 5, -26, 22, -26, 8, 27, -25,
	26, -25, 5, -25, 22, -25, 8, 27, -24, 26, -24, 5, -24, 22,
	-24, 2, 5, 57, 8, 27, 32, 26, 33, 5, -36, 22, 34, 22,
	16, -40, 2, -40, 3, -40, 13, -40, 15, -40, 7, -40, 27, -40,
	26, -40, 5, -40, 22, -40, 12, -40, 22, 16, -16, 2, -16, 3,
	-16, 13, -16, 15, -16, 7, -16, 27, -16, 26, -16, 5, -16, 22,
	-16, 12, -16, 22, 16, -15, 2, -15, 3, -15, 13, -15, 15, -15,
	7, -15, 27, -15, 26, -15, 5, -15, 22, -15, 12, -15, 22, 16,
	-14, 2, -14, 3, -14, 13, -14, 15, -14, 7, -14, 27, -14, 26,
	-14, 5, -14, 22, -14, 12, -14, 22, 16, -42, 2, -42, 3, -42,
	13, -42, 15, -42, 7, -42, 27, -42, 26, -42, 5, -42, 22, -42,
	12, -42, 22, 16, -13, 2, -13, 3, -13, 13, -13, 15, -13, 7,
	-13, 27, -13, 26, -13, 5, -13, 22, -13, 12, -13, 2, 16, 59,
	4, 4, 60, 11, 61, 8, 10, -51, 24, -51, 2, -51, 23, -51,
	2, 10, 67, 8, 10, -32, 24, 1, 2, 2, 23, 3, 8, 27,
	-53, 26, -53, 5, -53, 22, -53, 2, 5, 68, 8, 27, 32, 26,
	33, 5, -34, 22, 34, 12, 10, -9, 0, -9, 24, -9, 2, -9,
	23, -9, 25, -9, 2, 5, 69, 28, 16, -20, 2, -20, 3, -20,
	13, -20, 19, -20, 15, -20, 7, -20, 27, -20, 26, -20, 5, -20,
	22, -20, 12, -20, 18, -20, 17, -20, 6, 14, -22, 4, -22, 11,
	-22, 6, 14, -23, 4, -23, 11, -23, 6, 14, 70, 4, 60, 11,
	61, 6, 14, -47, 4, -47, 11, -47, 2, 2, 72, 8, 27, -54,
	26, -54, 5, -54, 22, -54, 12, 16, -38, 7, -38, 27, -38, 26,
	-38, 5, -38, 22, -38, 10, 0, -7, 24, -7, 2, -7, 23, -7,
	25, -7, 12, 10, -8, 0, -8, 24, -8, 2, -8, 23, -8, 25,
	-8, 12, 10, -10, 0, -10, 24, -10, 2, -10, 23, -10, 25, -10,
	28, 16, -21, 2, -21, 3, -21, 13, -21, 19, -21, 15, -21, 7,
	-21, 27, -21, 26, -21, 5, -21, 22, -21, 12, -21, 18, -21, 17,
	-21, 6, 14, -46, 4, -46, 11, -46, 2, 16, 75, 8, 10, -50,
	24, -50, 2, -50, 23, -50, 8, 27, -52, 26, -52, 5, -52, 22,
	-52, 8, 27, -28, 26, -28, 5, -28, 22, -28,
}

var _Goto = []int32{
	76, 95, 112, 112, 112, 112, 112, 112, 112, 112, 112, 113, 112, 112,
	112, 126, 112, 112, 143, 112, 112, 112, 160, 112, 112, 173, 182, 112,
	112, 112, 187, 200, 112, 112, 112, 112, 112, 112, 112, 112, 213, 222,
	239, 112, 112, 112, 112, 112, 112, 112, 252, 112, 112, 257, 112, 112,
	266, 112, 112, 112, 112, 112, 275, 112, 112, 112, 112, 112, 112, 112,
	112, 112, 112, 112, 112, 112, 18, 6, 5, 3, 6, 7, 7, 4,
	8, 1, 9, 19, 10, 28, 11, 2, 12, 5, 13, 16, 13, 20,
	26, 21, 8, 22, 23, 23, 9, 24, 24, 25, 12, 26, 10, 27,
	0, 12, 6, 5, 3, 6, 7, 7, 4, 8, 2, 29, 5, 13,
	16, 13, 20, 26, 21, 8, 31, 23, 23, 9, 24, 24, 25, 12,
	26, 10, 27, 16, 13, 20, 26, 21, 8, 49, 23, 23, 9, 24,
	24, 25, 12, 26, 10, 27, 12, 15, 35, 18, 36, 17, 37, 16,
	38, 22, 39, 31, 40, 8, 13, 20, 26, 21, 12, 26, 10, 43,
	4, 11, 47, 25, 48, 12, 6, 5, 3, 51, 7, 7, 20, 52,
	29, 53, 5, 13, 12, 15, 54, 18, 36, 17, 37, 16, 38, 21,
	55, 30, 56, 8, 15, 65, 18, 36, 17, 37, 16, 38, 16, 13,
	20, 26, 21, 8, 58, 23, 23, 9, 24, 24, 25, 12, 26, 10,
	27, 12, 13, 20, 26, 21, 9, 66, 24, 25, 12, 26, 10, 27,
	4, 27, 62, 14, 63, 8, 6, 5, 3, 73, 7, 7, 5, 13,
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
		return p.on_lexer_rule(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 5:
		return p.on_lexer_rule(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 6:
		return p.on_lexer_rule(
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
			_cast[*_i1.LexerExpr](p._sym.Peek(2)),
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 9:
		return p.on_frag_rule(
			_cast[Token](p._sym.Peek(3)),
			_cast[*_i1.LexerExpr](p._sym.Peek(2)),
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 10:
		return p.on_macro_rule(
			_cast[Token](p._sym.Peek(4)),
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[*_i1.LexerExpr](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 11:
		return p.on_lexer_expr(
			_cast[[]*_i1.LexerFactor](p._sym.Peek(0)),
		)
	case 12:
		return p.on_lexer_factor(
			_cast[[]*_i1.LexerTermCard](p._sym.Peek(0)),
		)
	case 13:
		return p.on_lexer_term_card(
			_cast[_i1.LexerTerm](p._sym.Peek(1)),
			_cast[_i1.Card](p._sym.Peek(0)),
		)
	case 14:
		return p.on_lexer_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 15:
		return p.on_lexer_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 16:
		return p.on_lexer_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 17:
		return p.on_lexer_term__tok(
			_cast[Token](p._sym.Peek(0)),
		)
	case 18:
		return p.on_lexer_term__tok(
			_cast[Token](p._sym.Peek(0)),
		)
	case 19:
		return p.on_lexer_term__char_class(
			_cast[*_i1.LexerTermCharClass](p._sym.Peek(0)),
		)
	case 20:
		return p.on_lexer_term__expr(
			_cast[Token](p._sym.Peek(2)),
			_cast[*_i1.LexerExpr](p._sym.Peek(1)),
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
			_cast[[]*_i1.LexerFactor](p._sym.Peek(2)),
			_cast[*_i1.LexerFactor](p._sym.Peek(0)),
		)
	case 39: // List
		return []*_i1.LexerFactor{
			_cast[*_i1.LexerFactor](p._sym.Peek(0)),
		}
	case 40: // OneOrMore
		return append(
			_cast[[]*_i1.LexerTermCard](p._sym.Peek(1)),
			_cast[*_i1.LexerTermCard](p._sym.Peek(0)),
		)
	case 41: // OneOrMore
		return []*_i1.LexerTermCard{
			_cast[*_i1.LexerTermCard](p._sym.Peek(0)),
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
