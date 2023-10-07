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
	48, 48, 49, 49, 50, 50, 51, 51,
}

var _TermCounts = []int32{
	1, 1, 1, 1, 2, 1, 5, 2, 2, 1, 1, 1, 1, 6,
	1, 1, 1, 4, 4, 2, 1, 1, 1, 1, 1, 5, 5, 4,
	5, 1, 1, 2, 1, 1, 1, 1, 1, 1, 3, 4, 1, 1,
	1, 1, 1, 1, 4, 1, 1, 0, 1, 0, 1, 0, 3, 1,
	2, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
	3, 1, 2, 1, 1, 0, 1, 0, 2, 1, 2, 1, 2, 1,
	2, 1, 2, 1, 2, 1, 2, 1,
}

var _Actions = []int32{
	127, 134, 149, 160, 167, 174, 181, 184, 187, 194, 197, 208, 211, 218,
	229, 240, 251, 254, 257, 260, 277, 292, 299, 314, 329, 346, 361, 378,
	385, 388, 240, 391, 420, 240, 449, 452, 481, 484, 493, 506, 519, 542,
	571, 594, 597, 608, 623, 632, 641, 650, 659, 662, 671, 680, 689, 698,
	707, 710, 240, 240, 719, 742, 765, 788, 811, 834, 857, 860, 865, 892,
	919, 922, 949, 976, 981, 998, 1003, 1026, 1043, 1052, 1055, 1064, 1073, 1076,
	1085, 1102, 1105, 1134, 1141, 1148, 1155, 1162, 1165, 1174, 623, 1187, 1198, 1201,
	1204, 1209, 1214, 1231, 1248, 1265, 1282, 1299, 623, 1316, 1331, 1348, 1365, 1394,
	1401, 1404, 1413, 1422, 1425, 1428, 1431, 1440, 623, 1445, 1448, 1451, 1454, 1459,
	1464, 6, 0, -49, 25, 1, 24, 2, 14, 0, -63, 29, 15, 2,
	16, 25, -63, 28, 17, 30, 18, 24, -63, 10, 0, -51, 2, -53,
	25, -51, 24, -51, 26, 9, 6, 0, -3, 25, -3, 24, -3, 6,
	0, -2, 25, -2, 24, -2, 6, 0, -81, 25, -81, 24, -81, 2,
	0, 2147483647, 2, 0, -1, 6, 0, -48, 25, 1, 24, 2, 2, 2,
	-52, 10, 0, -5, 2, -5, 25, -5, 24, -5, 26, -5, 2, 2,
	28, 6, 0, -4, 25, -4, 24, -4, 10, 0, -50, 2, -53, 25,
	-50, 24, -50, 26, 9, 10, 0, -83, 2, -83, 25, -83, 24, -83,
	26, -83, 10, 2, 31, 3, 32, 15, -77, 17, 33, 14, 34, 2,
	8, 30, 2, 2, 43, 2, 2, 29, 16, 12, -23, 0, -23, 29,
	-23, 2, -23, 25, -23, 28, -23, 30, -23, 24, -23, 14, 0, -21,
	29, -21, 2, -21, 25, -21, 28, -21, 30, -21, 24, -21, 6, 0,
	-19, 25, -19, 24, -19, 14, 0, -62, 29, 15, 2, 16, 25, -62,
	28, 17, 30, 18, 24, -62, 14, 0, -85, 29, -85, 2, -85, 25,
	-85, 28, -85, 30, -85, 24, -85, 16, 12, -24, 0, -24, 29, -24,
	2, -24, 25, -24, 28, -24, 30, -24, 24, -24, 14, 0, -20, 29,
	-20, 2, -20, 25, -20, 28, -20, 30, -20, 24, -20, 16, 12, -22,
	0, -22, 29, -22, 2, -22, 25, -22, 28, -22, 30, -22, 24, -22,
	6, 0, -80, 25, -80, 24, -80, 2, 8, 46, 2, 11, 47, 28,
	18, -36, 2, -36, 3, -36, 15, -36, 21, -36, 17, -36, 9, -36,
	32, -36, 31, -36, 6, -36, 27, -36, 14, -36, 20, -36, 19, -36,
	28, 18, -35, 2, -35, 3, -35, 15, -35, 21, -35, 17, -35, 9,
	-35, 32, -35, 31, -35, 6, -35, 27, -35, 14, -35, 20, -35, 19,
	-35, 2, 15, -76, 28, 18, -37, 2, -37, 3, -37, 15, -37, 21,
	-37, 17, -37, 9, -37, 32, -37, 31, -37, 6, -37, 27, -37, 14,
	-37, 20, -37, 19, -37, 2, 15, 67, 8, 32, 49, 31, 50, 6,
	-69, 27, 51, 12, 18, -29, 9, 59, 32, -29, 31, -29, 6, -29,
	27, -29, 12, 18, -71, 9, -71, 32, -71, 31, -71, 6, -71, 27,
	-71, 22, 18, -30, 2, 31, 3, 32, 15, -77, 17, 33, 9, -30,
	32, -30, 31, -30, 6, -30, 27, -30, 14, 34, 28, 18, -75, 2,
	-75, 3, -75, 15, -75, 21, 61, 17, -75, 9, -75, 32, -75, 31,
	-75, 6, -75, 27, -75, 14, -75, 20, 62, 19, 63, 22, 18, -73,
	2, -73, 3, -73, 15, -73, 17, -73, 9, -73, 32, -73, 31, -73,
	6, -73, 27, -73, 14, -73, 2, 8, 58, 10, 0, -82, 2, -82,
	25, -82, 24, -82, 26, -82, 14, 0, -84, 29, -84, 2, -84, 25,
	-84, 28, -84, 30, -84, 24, -84, 8, 33, 68, 2, 69, 35, 70,
	3, 71, 8, 12, -65, 29, 15, 2, 16, 28, 17, 8, 32, 49,
	31, 50, 6, -67, 27, 51, 8, 32, -47, 31, -47, 6, -47, 27,
	-47, 2, 17, 91, 8, 32, -45, 31, -45, 6, -45, 27, -45, 8,
	32, -91, 31, -91, 6, -91, 27, -91, 8, 32, -44, 31, -44, 6,
	-44, 27, -44, 8, 32, -43, 31, -43, 6, -43, 27, -43, 8, 32,
	-42, 31, -42, 6, -42, 27, -42, 2, 6, 84, 8, 32, 49, 31,
	50, 6, -68, 27, 51, 22, 18, -72, 2, -72, 3, -72, 15, -72,
	17, -72, 9, -72, 32, -72, 31, -72, 6, -72, 27, -72, 14, -72,
	22, 18, -34, 2, -34, 3, -34, 15, -34, 17, -34, 9, -34, 32,
	-34, 31, -34, 6, -34, 27, -34, 14, -34, 22, 18, -33, 2, -33,
	3, -33, 15, -33, 17, -33, 9, -33, 32, -33, 31, -33, 6, -33,
	27, -33, 14, -33, 22, 18, -32, 2, -32, 3, -32, 15, -32, 17,
	-32, 9, -32, 32, -32, 31, -32, 6, -32, 27, -32, 14, -32, 22,
	18, -74, 2, -74, 3, -74, 15, -74, 17, -74, 9, -74, 32, -74,
	31, -74, 6, -74, 27, -74, 14, -74, 22, 18, -31, 2, -31, 3,
	-31, 15, -31, 17, -31, 9, -31, 32, -31, 31, -31, 6, -31, 27,
	-31, 14, -31, 2, 18, 86, 4, 4, 87, 13, 88, 26, 7, -11,
	18, -11, 33, -11, 2, -11, 34, -11, 35, -11, 3, -11, 21, -11,
	9, -11, 36, -11, 6, -11, 20, -11, 19, -11, 26, 7, -9, 18,
	-9, 33, -9, 2, -9, 34, -9, 35, -9, 3, -9, 21, -9, 9,
	-9, 36, -9, 6, -9, 20, -9, 19, -9, 2, 17, 106, 26, 7,
	-10, 18, -10, 33, -10, 2, -10, 34, -10, 35, -10, 3, -10, 21,
	-10, 9, -10, 36, -10, 6, -10, 20, -10, 19, -10, 26, 7, -12,
	18, -12, 33, -12, 2, -12, 34, -12, 35, -12, 3, -12, 21, -12,
	9, -12, 36, -12, 6, -12, 20, -12, 19, -12, 4, 9, -55, 6,
	-55, 16, 33, 68, 2, 69, 34, 96, 35, 70, 3, 71, 9, -59,
	36, 97, 6, -59, 4, 9, 94, 6, 95, 22, 33, -61, 2, -61,
	34, -61, 35, -61, 3, -61, 21, 101, 9, -61, 36, -61, 6, -61,
	20, 102, 19, 103, 16, 33, -57, 2, -57, 34, -57, 35, -57, 3,
	-57, 9, -57, 36, -57, 6, -57, 8, 12, -87, 29, -87, 2, -87,
	28, -87, 2, 12, 107, 8, 12, -64, 29, 15, 2, 16, 28, 17,
	8, 32, -89, 31, -89, 6, -89, 27, -89, 2, 6, 108, 8, 32,
	49, 31, 50, 6, -66, 27, 51, 16, 12, -27, 0, -27, 29, -27,
	2, -27, 25, -27, 28, -27, 30, -27, 24, -27, 2, 6, 109, 28,
	18, -38, 2, -38, 3, -38, 15, -38, 21, -38, 17, -38, 9, -38,
	32, -38, 31, -38, 6, -38, 27, -38, 14, -38, 20, -38, 19, -38,
	6, 16, -40, 4, -40, 13, -40, 6, 16, -41, 4, -41, 13, -41,
	6, 16, 110, 4, 87, 13, 88, 6, 16, -79, 4, -79, 13, -79,
	2, 2, 112, 8, 32, -90, 31, -90, 6, -90, 27, -90, 12, 18,
	-70, 9, -70, 32, -70, 31, -70, 6, -70, 27, -70, 10, 0, -6,
	2, -6, 25, -6, 24, -6, 26, -6, 2, 17, 116, 2, 17, 117,
	4, 9, -7, 6, -7, 4, 9, -58, 6, -58, 16, 33, -56, 2,
	-56, 34, -56, 35, -56, 3, -56, 9, -56, 36, -56, 6, -56, 16,
	33, -15, 2, -15, 34, -15, 35, -15, 3, -15, 9, -15, 36, -15,
	6, -15, 16, 33, -14, 2, -14, 34, -14, 35, -14, 3, -14, 9,
	-14, 36, -14, 6, -14, 16, 33, -16, 2, -16, 34, -16, 35, -16,
	3, -16, 9, -16, 36, -16, 6, -16, 16, 33, -60, 2, -60, 34,
	-60, 35, -60, 3, -60, 9, -60, 36, -60, 6, -60, 16, 33, -8,
	2, -8, 34, -8, 35, -8, 3, -8, 9, -8, 36, -8, 6, -8,
	14, 0, -25, 29, -25, 2, -25, 25, -25, 28, -25, 30, -25, 24,
	-25, 16, 12, -26, 0, -26, 29, -26, 2, -26, 25, -26, 28, -26,
	30, -26, 24, -26, 16, 12, -28, 0, -28, 29, -28, 2, -28, 25,
	-28, 28, -28, 30, -28, 24, -28, 28, 18, -39, 2, -39, 3, -39,
	15, -39, 21, -39, 17, -39, 9, -39, 32, -39, 31, -39, 6, -39,
	27, -39, 14, -39, 20, -39, 19, -39, 6, 16, -78, 4, -78, 13,
	-78, 2, 18, 118, 8, 12, -86, 29, -86, 2, -86, 28, -86, 8,
	32, -88, 31, -88, 6, -88, 27, -88, 2, 7, 120, 2, 5, 121,
	2, 5, 122, 8, 32, -46, 31, -46, 6, -46, 27, -46, 4, 9,
	-54, 6, -54, 2, 18, 124, 2, 18, 125, 2, 18, 126, 4, 9,
	-17, 6, -17, 4, 9, -18, 6, -18, 26, 7, -13, 18, -13, 33,
	-13, 2, -13, 34, -13, 35, -13, 3, -13, 21, -13, 9, -13, 36,
	-13, 6, -13, 20, -13, 19, -13,
}

var _Goto = []int32{
	127, 140, 157, 168, 168, 168, 168, 168, 169, 168, 168, 168, 168, 176,
	168, 183, 168, 168, 168, 168, 168, 168, 200, 168, 168, 168, 168, 168,
	168, 168, 213, 168, 168, 230, 168, 168, 168, 247, 168, 168, 260, 269,
	168, 168, 168, 168, 274, 287, 300, 168, 168, 168, 168, 168, 168, 168,
	168, 313, 322, 339, 168, 168, 168, 168, 168, 168, 168, 352, 168, 168,
	168, 168, 168, 168, 357, 168, 368, 168, 168, 168, 373, 168, 168, 382,
	168, 168, 168, 168, 168, 391, 168, 168, 168, 168, 394, 168, 168, 168,
	168, 168, 168, 168, 168, 168, 168, 168, 405, 168, 168, 168, 168, 168,
	168, 168, 168, 168, 168, 168, 168, 168, 410, 168, 168, 168, 168, 168,
	168, 12, 12, 3, 3, 4, 2, 5, 1, 6, 30, 7, 46, 8,
	16, 17, 19, 14, 20, 37, 21, 48, 22, 13, 23, 18, 24, 15,
	25, 16, 26, 10, 5, 10, 32, 11, 31, 12, 47, 13, 4, 14,
	0, 6, 12, 3, 3, 4, 2, 27, 6, 5, 10, 32, 11, 4,
	44, 16, 24, 35, 44, 36, 19, 37, 41, 38, 20, 39, 42, 40,
	23, 41, 21, 42, 12, 17, 19, 14, 20, 13, 45, 18, 24, 15,
	25, 16, 26, 16, 24, 35, 44, 36, 19, 48, 41, 38, 20, 39,
	42, 40, 23, 41, 21, 42, 16, 24, 35, 44, 36, 19, 66, 41,
	38, 20, 39, 42, 40, 23, 41, 21, 42, 12, 26, 52, 29, 53,
	28, 54, 27, 55, 40, 56, 51, 57, 8, 24, 35, 44, 36, 23,
	41, 21, 60, 4, 22, 64, 43, 65, 12, 9, 72, 6, 73, 34,
	74, 33, 75, 8, 76, 7, 77, 12, 17, 19, 14, 78, 18, 24,
	38, 79, 49, 80, 16, 26, 12, 26, 81, 29, 53, 28, 54, 27,
	55, 39, 82, 50, 83, 8, 26, 92, 29, 53, 28, 54, 27, 55,
	16, 24, 35, 44, 36, 19, 85, 41, 38, 20, 39, 42, 40, 23,
	41, 21, 42, 12, 24, 35, 44, 36, 20, 93, 42, 40, 23, 41,
	21, 42, 4, 45, 89, 25, 90, 10, 9, 72, 35, 98, 11, 99,
	8, 76, 7, 100, 4, 10, 104, 36, 105, 8, 17, 19, 14, 113,
	18, 24, 16, 26, 8, 26, 114, 29, 53, 28, 54, 27, 55, 2,
	25, 111, 10, 9, 72, 6, 119, 34, 74, 8, 76, 7, 77, 4,
	9, 72, 8, 115, 4, 9, 72, 8, 123,
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
			_cast[Token](p._sym.Peek(4)),
			_cast[Token](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[[]*_i1.ParserProd](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 7:
		return p.on_parser_prod(
			_cast[[]*_i1.ParserTerm](p._sym.Peek(1)),
			_cast[*_i1.ProdQualifier](p._sym.Peek(0)),
		)
	case 8:
		return p.on_parser_term_card(
			_cast[*_i1.ParserTerm](p._sym.Peek(1)),
			_cast[_i1.ParserTermType](p._sym.Peek(0)),
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
			_cast[*_i1.ParserTerm](p._sym.Peek(0)),
		)
	case 13:
		return p.on_parser_list(
			_cast[Token](p._sym.Peek(5)),
			_cast[Token](p._sym.Peek(4)),
			_cast[*_i1.ParserTerm](p._sym.Peek(3)),
			_cast[Token](p._sym.Peek(2)),
			_cast[*_i1.ParserTerm](p._sym.Peek(1)),
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
	case 52: // ZeroOrOne
		return _cast[Token](p._sym.Peek(0))
	case 53: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 54: // List
		return append(
			_cast[[]*_i1.ParserProd](p._sym.Peek(2)),
			_cast[*_i1.ParserProd](p._sym.Peek(0)),
		)
	case 55: // List
		return []*_i1.ParserProd{
			_cast[*_i1.ParserProd](p._sym.Peek(0)),
		}
	case 56: // OneOrMore
		return append(
			_cast[[]*_i1.ParserTerm](p._sym.Peek(1)),
			_cast[*_i1.ParserTerm](p._sym.Peek(0)),
		)
	case 57: // OneOrMore
		return []*_i1.ParserTerm{
			_cast[*_i1.ParserTerm](p._sym.Peek(0)),
		}
	case 58: // ZeroOrOne
		return _cast[*_i1.ProdQualifier](p._sym.Peek(0))
	case 59: // ZeroOrOne
		{
			var zero *_i1.ProdQualifier
			return zero
		}
	case 60: // ZeroOrOne
		return _cast[_i1.ParserTermType](p._sym.Peek(0))
	case 61: // ZeroOrOne
		{
			var zero _i1.ParserTermType
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
		return _cast[[]_i1.Statement](p._sym.Peek(0))
	case 65: // ZeroOrOne
		{
			var zero []_i1.Statement
			return zero
		}
	case 66: // ZeroOrOne
		return _cast[[]_i1.Action](p._sym.Peek(0))
	case 67: // ZeroOrOne
		{
			var zero []_i1.Action
			return zero
		}
	case 68: // ZeroOrOne
		return _cast[[]_i1.Action](p._sym.Peek(0))
	case 69: // ZeroOrOne
		{
			var zero []_i1.Action
			return zero
		}
	case 70: // List
		return append(
			_cast[[]*_i1.LexerFactor](p._sym.Peek(2)),
			_cast[*_i1.LexerFactor](p._sym.Peek(0)),
		)
	case 71: // List
		return []*_i1.LexerFactor{
			_cast[*_i1.LexerFactor](p._sym.Peek(0)),
		}
	case 72: // OneOrMore
		return append(
			_cast[[]*_i1.LexerTermCard](p._sym.Peek(1)),
			_cast[*_i1.LexerTermCard](p._sym.Peek(0)),
		)
	case 73: // OneOrMore
		return []*_i1.LexerTermCard{
			_cast[*_i1.LexerTermCard](p._sym.Peek(0)),
		}
	case 74: // ZeroOrOne
		return _cast[_i1.Card](p._sym.Peek(0))
	case 75: // ZeroOrOne
		{
			var zero _i1.Card
			return zero
		}
	case 76: // ZeroOrOne
		return _cast[Token](p._sym.Peek(0))
	case 77: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 78: // OneOrMore
		return append(
			_cast[[]Token](p._sym.Peek(1)),
			_cast[Token](p._sym.Peek(0)),
		)
	case 79: // OneOrMore
		return []Token{
			_cast[Token](p._sym.Peek(0)),
		}
	case 80: // OneOrMore
		return append(
			_cast[[][]_i1.Statement](p._sym.Peek(1)),
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		)
	case 81: // OneOrMore
		return [][]_i1.Statement{
			_cast[[]_i1.Statement](p._sym.Peek(0)),
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
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 87: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._sym.Peek(0)),
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
	case 90: // OneOrMore
		return append(
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 91: // OneOrMore
		return []_i1.Action{
			_cast[_i1.Action](p._sym.Peek(0)),
		}
	default:
		panic("unreachable")
	}
}
