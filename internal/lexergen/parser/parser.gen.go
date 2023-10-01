package parser

import (
	_i0 "errors"
	_i1 "github.com/dcaiafa/lox/internal/lexergen/ast"
)

var _LHS = []int32{
	0, 1, 2, 2, 3, 4, 5, 6, 7, 8, 8, 8, 8, 9,
	10, 10, 10, 11, 11, 12, 13, 13, 14, 14, 14, 15, 16, 17,
	18, 19, 20, 21, 22, 22, 22, 23, 23, 23, 23, 24, 25, 25,
	26, 26, 26, 27, 28, 29, 30, 30, 31, 31, 32, 32, 33, 33,
	34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40,
	41, 41, 42, 42, 43, 43, 44, 44, 45, 45, 46, 46, 47, 47,
	48, 48, 49, 49, 50, 50,
}

var _TermCounts = []int32{
	1, 1, 1, 1, 2, 1, 4, 2, 2, 1, 1, 1, 1, 6,
	1, 1, 1, 4, 4, 2, 1, 1, 1, 1, 1, 5, 5, 4,
	5, 1, 1, 2, 1, 1, 1, 1, 1, 1, 3, 4, 1, 1,
	1, 1, 1, 1, 4, 1, 1, 0, 1, 0, 3, 1, 2, 1,
	1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 1,
	2, 1, 1, 0, 1, 0, 2, 1, 2, 1, 2, 1, 2, 1,
	2, 1, 2, 1, 2, 1,
}

var _Actions = []int32{
	125, 132, 147, 156, 163, 170, 177, 180, 183, 190, 193, 202, 209, 218,
	227, 238, 241, 244, 247, 264, 279, 286, 301, 316, 333, 348, 365, 372,
	381, 227, 384, 413, 227, 442, 445, 474, 477, 486, 499, 512, 535, 564,
	587, 590, 599, 614, 641, 668, 671, 698, 725, 730, 747, 752, 775, 792,
	801, 810, 819, 822, 831, 840, 849, 858, 867, 870, 227, 227, 879, 902,
	925, 948, 971, 994, 1017, 1020, 372, 1025, 1034, 1037, 1040, 1045, 1050, 1067,
	1084, 1101, 1118, 1135, 372, 1152, 1161, 1164, 1173, 1182, 1185, 1194, 1211, 1214,
	1243, 1250, 1257, 1264, 1271, 1274, 1283, 1296, 1299, 1302, 1305, 1320, 1337, 1354,
	1383, 1390, 1393, 1398, 1407, 372, 1416, 1419, 1422, 1431, 1434, 1439, 1444, 6,
	0, -49, 25, 1, 24, 2, 14, 0, -61, 28, 14, 2, 15, 25,
	-61, 27, 16, 29, 17, 24, -61, 8, 0, -51, 2, 9, 25, -51,
	24, -51, 6, 0, -3, 25, -3, 24, -3, 6, 0, -2, 25, -2,
	24, -2, 6, 0, -79, 25, -79, 24, -79, 2, 0, 2147483647, 2, 0,
	-1, 6, 0, -48, 25, 1, 24, 2, 2, 8, 27, 8, 0, -5,
	2, -5, 25, -5, 24, -5, 6, 0, -4, 25, -4, 24, -4, 8,
	0, -50, 2, 9, 25, -50, 24, -50, 8, 0, -81, 2, -81, 25,
	-81, 24, -81, 10, 2, 30, 3, 31, 15, -75, 17, 32, 14, 33,
	2, 8, 29, 2, 2, 42, 2, 2, 28, 16, 12, -23, 0, -23,
	28, -23, 2, -23, 25, -23, 27, -23, 29, -23, 24, -23, 14, 0,
	-21, 28, -21, 2, -21, 25, -21, 27, -21, 29, -21, 24, -21, 6,
	0, -19, 25, -19, 24, -19, 14, 0, -60, 28, 14, 2, 15, 25,
	-60, 27, 16, 29, 17, 24, -60, 14, 0, -83, 28, -83, 2, -83,
	25, -83, 27, -83, 29, -83, 24, -83, 16, 12, -24, 0, -24, 28,
	-24, 2, -24, 25, -24, 27, -24, 29, -24, 24, -24, 14, 0, -20,
	28, -20, 2, -20, 25, -20, 27, -20, 29, -20, 24, -20, 16, 12,
	-22, 0, -22, 28, -22, 2, -22, 25, -22, 27, -22, 29, -22, 24,
	-22, 6, 0, -78, 25, -78, 24, -78, 8, 32, 45, 2, 46, 34,
	47, 3, 48, 2, 11, 55, 28, 18, -36, 2, -36, 3, -36, 15,
	-36, 21, -36, 17, -36, 9, -36, 31, -36, 30, -36, 6, -36, 26,
	-36, 14, -36, 20, -36, 19, -36, 28, 18, -35, 2, -35, 3, -35,
	15, -35, 21, -35, 17, -35, 9, -35, 31, -35, 30, -35, 6, -35,
	26, -35, 14, -35, 20, -35, 19, -35, 2, 15, -74, 28, 18, -37,
	2, -37, 3, -37, 15, -37, 21, -37, 17, -37, 9, -37, 31, -37,
	30, -37, 6, -37, 26, -37, 14, -37, 20, -37, 19, -37, 2, 15,
	75, 8, 31, 57, 30, 58, 6, -67, 26, 59, 12, 18, -29, 9,
	67, 31, -29, 30, -29, 6, -29, 26, -29, 12, 18, -69, 9, -69,
	31, -69, 30, -69, 6, -69, 26, -69, 22, 18, -30, 2, 30, 3,
	31, 15, -75, 17, 32, 9, -30, 31, -30, 30, -30, 6, -30, 26,
	-30, 14, 33, 28, 18, -73, 2, -73, 3, -73, 15, -73, 21, 69,
	17, -73, 9, -73, 31, -73, 30, -73, 6, -73, 26, -73, 14, -73,
	20, 70, 19, 71, 22, 18, -71, 2, -71, 3, -71, 15, -71, 17,
	-71, 9, -71, 31, -71, 30, -71, 6, -71, 26, -71, 14, -71, 2,
	8, 66, 8, 0, -80, 2, -80, 25, -80, 24, -80, 14, 0, -82,
	28, -82, 2, -82, 25, -82, 27, -82, 29, -82, 24, -82, 26, 7,
	-11, 18, -11, 32, -11, 2, -11, 33, -11, 34, -11, 3, -11, 21,
	-11, 9, -11, 35, -11, 6, -11, 20, -11, 19, -11, 26, 7, -9,
	18, -9, 32, -9, 2, -9, 33, -9, 34, -9, 3, -9, 21, -9,
	9, -9, 35, -9, 6, -9, 20, -9, 19, -9, 2, 17, 88, 26,
	7, -10, 18, -10, 32, -10, 2, -10, 33, -10, 34, -10, 3, -10,
	21, -10, 9, -10, 35, -10, 6, -10, 20, -10, 19, -10, 26, 7,
	-12, 18, -12, 32, -12, 2, -12, 33, -12, 34, -12, 3, -12, 21,
	-12, 9, -12, 35, -12, 6, -12, 20, -12, 19, -12, 4, 9, -53,
	6, -53, 16, 32, 45, 2, 46, 33, 78, 34, 47, 3, 48, 9,
	-57, 35, 79, 6, -57, 4, 9, 76, 6, 77, 22, 32, -59, 2,
	-59, 33, -59, 34, -59, 3, -59, 21, 83, 9, -59, 35, -59, 6,
	-59, 20, 84, 19, 85, 16, 32, -55, 2, -55, 33, -55, 34, -55,
	3, -55, 9, -55, 35, -55, 6, -55, 8, 12, -63, 28, 14, 2,
	15, 27, 16, 8, 31, 57, 30, 58, 6, -65, 26, 59, 8, 31,
	-47, 30, -47, 6, -47, 26, -47, 2, 17, 102, 8, 31, -45, 30,
	-45, 6, -45, 26, -45, 8, 31, -89, 30, -89, 6, -89, 26, -89,
	8, 31, -44, 30, -44, 6, -44, 26, -44, 8, 31, -43, 30, -43,
	6, -43, 26, -43, 8, 31, -42, 30, -42, 6, -42, 26, -42, 2,
	6, 95, 8, 31, 57, 30, 58, 6, -66, 26, 59, 22, 18, -70,
	2, -70, 3, -70, 15, -70, 17, -70, 9, -70, 31, -70, 30, -70,
	6, -70, 26, -70, 14, -70, 22, 18, -34, 2, -34, 3, -34, 15,
	-34, 17, -34, 9, -34, 31, -34, 30, -34, 6, -34, 26, -34, 14,
	-34, 22, 18, -33, 2, -33, 3, -33, 15, -33, 17, -33, 9, -33,
	31, -33, 30, -33, 6, -33, 26, -33, 14, -33, 22, 18, -32, 2,
	-32, 3, -32, 15, -32, 17, -32, 9, -32, 31, -32, 30, -32, 6,
	-32, 26, -32, 14, -32, 22, 18, -72, 2, -72, 3, -72, 15, -72,
	17, -72, 9, -72, 31, -72, 30, -72, 6, -72, 26, -72, 14, -72,
	22, 18, -31, 2, -31, 3, -31, 15, -31, 17, -31, 9, -31, 31,
	-31, 30, -31, 6, -31, 26, -31, 14, -31, 2, 18, 97, 4, 4,
	98, 13, 99, 8, 0, -6, 2, -6, 25, -6, 24, -6, 2, 17,
	106, 2, 17, 107, 4, 9, -7, 6, -7, 4, 9, -56, 6, -56,
	16, 32, -54, 2, -54, 33, -54, 34, -54, 3, -54, 9, -54, 35,
	-54, 6, -54, 16, 32, -15, 2, -15, 33, -15, 34, -15, 3, -15,
	9, -15, 35, -15, 6, -15, 16, 32, -14, 2, -14, 33, -14, 34,
	-14, 3, -14, 9, -14, 35, -14, 6, -14, 16, 32, -16, 2, -16,
	33, -16, 34, -16, 3, -16, 9, -16, 35, -16, 6, -16, 16, 32,
	-58, 2, -58, 33, -58, 34, -58, 3, -58, 9, -58, 35, -58, 6,
	-58, 16, 32, -8, 2, -8, 33, -8, 34, -8, 3, -8, 9, -8,
	35, -8, 6, -8, 8, 12, -85, 28, -85, 2, -85, 27, -85, 2,
	12, 108, 8, 12, -62, 28, 14, 2, 15, 27, 16, 8, 31, -87,
	30, -87, 6, -87, 26, -87, 2, 6, 109, 8, 31, 57, 30, 58,
	6, -64, 26, 59, 16, 12, -27, 0, -27, 28, -27, 2, -27, 25,
	-27, 27, -27, 29, -27, 24, -27, 2, 6, 110, 28, 18, -38, 2,
	-38, 3, -38, 15, -38, 21, -38, 17, -38, 9, -38, 31, -38, 30,
	-38, 6, -38, 26, -38, 14, -38, 20, -38, 19, -38, 6, 16, -40,
	4, -40, 13, -40, 6, 16, -41, 4, -41, 13, -41, 6, 16, 111,
	4, 98, 13, 99, 6, 16, -77, 4, -77, 13, -77, 2, 2, 113,
	8, 31, -88, 30, -88, 6, -88, 26, -88, 12, 18, -68, 9, -68,
	31, -68, 30, -68, 6, -68, 26, -68, 2, 7, 117, 2, 5, 118,
	2, 5, 119, 14, 0, -25, 28, -25, 2, -25, 25, -25, 27, -25,
	29, -25, 24, -25, 16, 12, -26, 0, -26, 28, -26, 2, -26, 25,
	-26, 27, -26, 29, -26, 24, -26, 16, 12, -28, 0, -28, 28, -28,
	2, -28, 25, -28, 27, -28, 29, -28, 24, -28, 28, 18, -39, 2,
	-39, 3, -39, 15, -39, 21, -39, 17, -39, 9, -39, 31, -39, 30,
	-39, 6, -39, 26, -39, 14, -39, 20, -39, 19, -39, 6, 16, -76,
	4, -76, 13, -76, 2, 18, 120, 4, 9, -52, 6, -52, 8, 12,
	-84, 28, -84, 2, -84, 27, -84, 8, 31, -86, 30, -86, 6, -86,
	26, -86, 2, 18, 122, 2, 18, 123, 8, 31, -46, 30, -46, 6,
	-46, 26, -46, 2, 18, 124, 4, 9, -17, 6, -17, 4, 9, -18,
	6, -18, 26, 7, -13, 18, -13, 32, -13, 2, -13, 33, -13, 34,
	-13, 3, -13, 21, -13, 9, -13, 35, -13, 6, -13, 20, -13, 19,
	-13,
}

var _Goto = []int32{
	125, 138, 155, 164, 164, 164, 164, 164, 165, 164, 164, 164, 172, 164,
	177, 164, 164, 164, 164, 164, 164, 194, 164, 164, 164, 164, 164, 207,
	164, 220, 164, 164, 237, 164, 164, 164, 254, 164, 164, 267, 276, 164,
	164, 164, 164, 164, 164, 164, 164, 164, 164, 281, 164, 292, 164, 297,
	310, 164, 164, 164, 164, 164, 164, 164, 164, 323, 332, 349, 164, 164,
	164, 164, 164, 164, 164, 362, 367, 164, 164, 164, 164, 164, 164, 164,
	164, 164, 164, 164, 378, 164, 164, 383, 164, 164, 392, 164, 164, 164,
	164, 164, 401, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164,
	164, 164, 164, 164, 164, 404, 164, 164, 164, 164, 164, 164, 164, 12,
	12, 3, 3, 4, 2, 5, 1, 6, 30, 7, 45, 8, 16, 17,
	18, 14, 19, 36, 20, 47, 21, 13, 22, 18, 23, 15, 24, 16,
	25, 8, 5, 10, 31, 11, 46, 12, 4, 13, 0, 6, 12, 3,
	3, 4, 2, 26, 4, 5, 10, 4, 43, 16, 24, 34, 43, 35,
	19, 36, 40, 37, 20, 38, 41, 39, 23, 40, 21, 41, 12, 17,
	18, 14, 19, 13, 44, 18, 23, 15, 24, 16, 25, 12, 9, 49,
	6, 50, 33, 51, 32, 52, 8, 53, 7, 54, 16, 24, 34, 43,
	35, 19, 56, 40, 37, 20, 38, 41, 39, 23, 40, 21, 41, 16,
	24, 34, 43, 35, 19, 74, 40, 37, 20, 38, 41, 39, 23, 40,
	21, 41, 12, 26, 60, 29, 61, 28, 62, 27, 63, 39, 64, 50,
	65, 8, 24, 34, 43, 35, 23, 40, 21, 68, 4, 22, 72, 42,
	73, 10, 9, 49, 34, 80, 11, 81, 8, 53, 7, 82, 4, 10,
	86, 35, 87, 12, 17, 18, 14, 89, 18, 23, 37, 90, 48, 91,
	16, 25, 12, 26, 92, 29, 61, 28, 62, 27, 63, 38, 93, 49,
	94, 8, 26, 103, 29, 61, 28, 62, 27, 63, 16, 24, 34, 43,
	35, 19, 96, 40, 37, 20, 38, 41, 39, 23, 40, 21, 41, 12,
	24, 34, 43, 35, 20, 104, 41, 39, 23, 40, 21, 41, 4, 44,
	100, 25, 101, 10, 9, 49, 6, 114, 33, 51, 8, 53, 7, 54,
	4, 9, 49, 8, 105, 8, 17, 18, 14, 115, 18, 23, 16, 25,
	8, 26, 116, 29, 61, 28, 62, 27, 63, 2, 25, 112, 4, 9,
	49, 8, 121,
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
			_cast[[][]_i1.Statement](p._sym.Peek(0)),
		)
	case 2:
		return p.on_section(
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		)
	case 3:
		return p.on_section(
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		)
	case 4:
		return p.on_parser_section(
			_cast[Token](p._sym.Peek(1)),
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		)
	case 5:
		return p.on_parser_statement(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 6:
		return p.on_parser_rule(
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[[]*_i1.Prod](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 7:
		return p.on_parser_prod(
			_cast[[]*_i1.Term](p._sym.Peek(1)),
			_cast[*_i1.ProdQualifier](p._sym.Peek(0)),
		)
	case 8:
		return p.on_parser_term_card(
			_cast[*_i1.Term](p._sym.Peek(1)),
			_cast[_i1.TermType](p._sym.Peek(0)),
		)
	case 9:
		return p.on_parser_term__token(
			_cast[Token](p._sym.Peek(0)),
		)
	case 10:
		return p.on_parser_term__token(
			_cast[Token](p._sym.Peek(0)),
		)
	case 11:
		return p.on_parser_term__token(
			_cast[Token](p._sym.Peek(0)),
		)
	case 12:
		return p.on_parser_term__list(
			_cast[*_i1.Term](p._sym.Peek(0)),
		)
	case 13:
		return p.on_parser_list(
			_cast[Token](p._sym.Peek(5)),
			_cast[Token](p._sym.Peek(4)),
			_cast[*_i1.Term](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[*_i1.Term](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 14:
		return p.on_parser_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 15:
		return p.on_parser_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 16:
		return p.on_parser_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 17:
		return p.on_parser_qualif(
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 18:
		return p.on_parser_qualif(
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 19:
		return p.on_lexer_section(
			_cast[Token](p._sym.Peek(1)),
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		)
	case 20:
		return p.on_lexer_statement(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 21:
		return p.on_lexer_statement(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 22:
		return p.on_lexer_rule(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 23:
		return p.on_lexer_rule(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 24:
		return p.on_lexer_rule(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 25:
		return p.on_mode(
			_cast[Token](p._sym.Peek(4)),
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 26:
		return p.on_token_rule(
			_cast[Token](p._sym.Peek(4)),
			_cast[Token](p._sym.Peek(3)),
			_cast[*_i1.LexerExpr](p._sym.Peek(2)),
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 27:
		return p.on_frag_rule(
			_cast[Token](p._sym.Peek(3)),
			_cast[*_i1.LexerExpr](p._sym.Peek(2)),
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 28:
		return p.on_macro_rule(
			_cast[Token](p._sym.Peek(4)),
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[*_i1.LexerExpr](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 29:
		return p.on_lexer_expr(
			_cast[[]*_i1.LexerFactor](p._sym.Peek(0)),
		)
	case 30:
		return p.on_lexer_factor(
			_cast[[]*_i1.LexerTermCard](p._sym.Peek(0)),
		)
	case 31:
		return p.on_lexer_term_card(
			_cast[_i1.LexerTerm](p._sym.Peek(1)),
			_cast[_i1.Card](p._sym.Peek(0)),
		)
	case 32:
		return p.on_lexer_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 33:
		return p.on_lexer_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 34:
		return p.on_lexer_card(
			_cast[Token](p._sym.Peek(0)),
		)
	case 35:
		return p.on_lexer_term__tok(
			_cast[Token](p._sym.Peek(0)),
		)
	case 36:
		return p.on_lexer_term__tok(
			_cast[Token](p._sym.Peek(0)),
		)
	case 37:
		return p.on_lexer_term__char_class(
			_cast[*_i1.LexerTermCharClass](p._sym.Peek(0)),
		)
	case 38:
		return p.on_lexer_term__expr(
			_cast[Token](p._sym.Peek(2)),
			_cast[*_i1.LexerExpr](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 39:
		return p.on_char_class(
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[[]Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 40:
		return p.on_char_class_item(
			_cast[Token](p._sym.Peek(0)),
		)
	case 41:
		return p.on_char_class_item(
			_cast[Token](p._sym.Peek(0)),
		)
	case 42:
		return p.on_action(
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 43:
		return p.on_action(
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 44:
		return p.on_action(
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 45:
		return p.on_action_skip(
			_cast[Token](p._sym.Peek(0)),
		)
	case 46:
		return p.on_action_push_mode(
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 47:
		return p.on_action_pop_mode(
			_cast[Token](p._sym.Peek(0)),
		)
	case 48: // ZeroOrOne
		return _cast[[][]_i1.Statement](p._sym.Peek(0))
	case 49: // ZeroOrOne
		{
			var zero [][]_i1.Statement
			return zero
		}
	case 50: // ZeroOrOne
		return _cast[[]_i1.Statement](p._sym.Peek(0))
	case 51: // ZeroOrOne
		{
			var zero []_i1.Statement
			return zero
		}
	case 52: // List
		return append(
			_cast[[]*_i1.Prod](p._sym.Peek(2)),
			_cast[*_i1.Prod](p._sym.Peek(0)),
		)
	case 53: // List
		return []*_i1.Prod{
			_cast[*_i1.Prod](p._sym.Peek(0)),
		}
	case 54: // OneOrMore
		return append(
			_cast[[]*_i1.Term](p._sym.Peek(1)),
			_cast[*_i1.Term](p._sym.Peek(0)),
		)
	case 55: // OneOrMore
		return []*_i1.Term{
			_cast[*_i1.Term](p._sym.Peek(0)),
		}
	case 56: // ZeroOrOne
		return _cast[*_i1.ProdQualifier](p._sym.Peek(0))
	case 57: // ZeroOrOne
		{
			var zero *_i1.ProdQualifier
			return zero
		}
	case 58: // ZeroOrOne
		return _cast[_i1.TermType](p._sym.Peek(0))
	case 59: // ZeroOrOne
		{
			var zero _i1.TermType
			return zero
		}
	case 60: // ZeroOrOne
		return _cast[[]_i1.Statement](p._sym.Peek(0))
	case 61: // ZeroOrOne
		{
			var zero []_i1.Statement
			return zero
		}
	case 62: // ZeroOrOne
		return _cast[[]_i1.Statement](p._sym.Peek(0))
	case 63: // ZeroOrOne
		{
			var zero []_i1.Statement
			return zero
		}
	case 64: // ZeroOrOne
		return _cast[[]_i1.Action](p._sym.Peek(0))
	case 65: // ZeroOrOne
		{
			var zero []_i1.Action
			return zero
		}
	case 66: // ZeroOrOne
		return _cast[[]_i1.Action](p._sym.Peek(0))
	case 67: // ZeroOrOne
		{
			var zero []_i1.Action
			return zero
		}
	case 68: // List
		return append(
			_cast[[]*_i1.LexerFactor](p._sym.Peek(2)),
			_cast[*_i1.LexerFactor](p._sym.Peek(0)),
		)
	case 69: // List
		return []*_i1.LexerFactor{
			_cast[*_i1.LexerFactor](p._sym.Peek(0)),
		}
	case 70: // OneOrMore
		return append(
			_cast[[]*_i1.LexerTermCard](p._sym.Peek(1)),
			_cast[*_i1.LexerTermCard](p._sym.Peek(0)),
		)
	case 71: // OneOrMore
		return []*_i1.LexerTermCard{
			_cast[*_i1.LexerTermCard](p._sym.Peek(0)),
		}
	case 72: // ZeroOrOne
		return _cast[_i1.Card](p._sym.Peek(0))
	case 73: // ZeroOrOne
		{
			var zero _i1.Card
			return zero
		}
	case 74: // ZeroOrOne
		return _cast[Token](p._sym.Peek(0))
	case 75: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 76: // OneOrMore
		return append(
			_cast[[]Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 77: // OneOrMore
		return []Token{
			_cast[Token](p._sym.Peek(0)),
		}
	case 78: // OneOrMore
		return append(
			_cast[[][]_i1.Statement](p._sym.Peek(1)),
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		)
	case 79: // OneOrMore
		return [][]_i1.Statement{
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		}
	case 80: // OneOrMore
		return append(
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 81: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._sym.Peek(0)),
		}
	case 82: // OneOrMore
		return append(
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 83: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._sym.Peek(0)),
		}
	case 84: // OneOrMore
		return append(
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 85: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._sym.Peek(0)),
		}
	case 86: // OneOrMore
		return append(
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 87: // OneOrMore
		return []_i1.Action{
			_cast[_i1.Action](p._sym.Peek(0)),
		}
	case 88: // OneOrMore
		return append(
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 89: // OneOrMore
		return []_i1.Action{
			_cast[_i1.Action](p._sym.Peek(0)),
		}
	default:
		panic("unreachable")
	}
}
