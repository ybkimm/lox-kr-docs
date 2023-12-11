package parser

import (
	_i0 "errors"
	_i1 "github.com/dcaiafa/lox/internal/ast"
	_i2 "github.com/dcaiafa/lox/internal/base/baselexer"
)

var _LHS = []int32{
	0, 1, 2, 2, 3, 4, 5, 6, 7, 8, 8, 8, 8, 9,
	10, 10, 10, 11, 11, 12, 13, 13, 14, 14, 14, 15, 16, 17,
	18, 19, 20, 21, 22, 22, 22, 23, 23, 23, 23, 24, 24, 25,
	26, 26, 27, 27, 27, 28, 29, 30, 31, 31, 32, 32, 33, 33,
	34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40,
	41, 41, 42, 42, 43, 43, 44, 44, 45, 45, 46, 46, 47, 47,
	48, 48, 49, 49, 50, 50,
}

var _TermCounts = []int32{
	1, 1, 1, 1, 2, 1, 5, 2, 2, 1, 1, 1, 1, 6,
	1, 1, 1, 4, 4, 2, 1, 1, 1, 1, 1, 5, 5, 4,
	5, 1, 1, 2, 1, 1, 1, 1, 1, 1, 3, 3, 1, 4,
	1, 1, 1, 1, 1, 1, 4, 1, 1, 0, 2, 1, 1, 0,
	2, 1, 1, 0, 3, 1, 2, 1, 1, 0, 1, 0, 1, 0,
	2, 1, 1, 0, 2, 1, 1, 0, 2, 1, 3, 1, 2, 1,
	1, 0, 1, 0, 2, 1,
}

var _Actions = []int32{
	127, 134, 149, 160, 167, 174, 181, 184, 191, 194, 197, 200, 211, 222,
	229, 240, 251, 254, 257, 260, 277, 292, 307, 314, 329, 346, 361, 378,
	385, 388, 240, 391, 404, 433, 240, 462, 465, 468, 499, 528, 537, 550,
	579, 602, 625, 628, 639, 654, 663, 528, 672, 681, 690, 693, 702, 705,
	714, 723, 732, 240, 240, 741, 764, 787, 810, 833, 856, 879, 882, 887,
	892, 897, 924, 951, 954, 981, 1008, 1013, 1036, 1053, 1070, 1079, 1082, 1091,
	1094, 1111, 1114, 1143, 1172, 1179, 1186, 1193, 1200, 1203, 1212, 654, 1225, 1236,
	1239, 1242, 1247, 1252, 1269, 1286, 1303, 1320, 1337, 654, 1354, 1369, 1386, 1403,
	1434, 1441, 1444, 1453, 1456, 1459, 1462, 1471, 654, 1476, 1479, 1482, 1485, 1490,
	1495, 6, 0, -51, 18, 1, 17, 2, 14, 0, -69, 22, 15, 31,
	16, 18, -69, 21, 17, 23, 18, 17, -69, 10, 0, -55, 31, -59,
	18, -55, 17, -55, 19, 9, 6, 0, -3, 18, -3, 17, -3, 6,
	0, -2, 18, -2, 17, -2, 6, 0, -53, 18, -53, 17, -53, 2,
	0, -1, 6, 0, -50, 18, 1, 17, 2, 2, 0, 2147483647, 2, 31,
	-58, 2, 31, 28, 10, 0, -5, 31, -5, 18, -5, 17, -5, 19,
	-5, 10, 0, -57, 31, -57, 18, -57, 17, -57, 19, -57, 6, 0,
	-4, 18, -4, 17, -4, 10, 0, -54, 31, -59, 18, -54, 17, -54,
	19, 9, 10, 31, 32, 33, 33, 34, -87, 9, 34, 8, 35, 2,
	4, 30, 2, 31, 44, 2, 31, 29, 16, 7, -23, 0, -23, 22,
	-23, 31, -23, 18, -23, 21, -23, 23, -23, 17, -23, 14, 0, -21,
	22, -21, 31, -21, 18, -21, 21, -21, 23, -21, 17, -21, 14, 0,
	-71, 22, -71, 31, -71, 18, -71, 21, -71, 23, -71, 17, -71, 6,
	0, -19, 18, -19, 17, -19, 14, 0, -68, 22, 15, 31, 16, 18,
	-68, 21, 17, 23, 18, 17, -68, 16, 7, -24, 0, -24, 22, -24,
	31, -24, 18, -24, 21, -24, 23, -24, 17, -24, 14, 0, -20, 22,
	-20, 31, -20, 18, -20, 21, -20, 23, -20, 17, -20, 16, 7, -22,
	0, -22, 22, -22, 31, -22, 18, -22, 21, -22, 23, -22, 17, -22,
	6, 0, -52, 18, -52, 17, -52, 2, 4, 47, 2, 6, 48, 12,
	10, -29, 20, -29, 5, 60, 25, -29, 24, -29, 2, -29, 28, 10,
	-36, 20, -36, 31, -36, 33, -36, 34, -36, 14, -36, 9, -36, 5,
	-36, 25, -36, 24, -36, 2, -36, 8, -36, 13, -36, 12, -36, 28,
	10, -35, 20, -35, 31, -35, 33, -35, 34, -35, 14, -35, 9, -35,
	5, -35, 25, -35, 24, -35, 2, -35, 8, -35, 13, -35, 12, -35,
	2, 34, -86, 2, 34, 69, 30, 10, -40, 20, -40, 31, -40, 33,
	-40, 34, -40, 14, -40, 9, -40, 5, -40, 25, -40, 24, -40, 2,
	-40, 11, 68, 8, -40, 13, -40, 12, -40, 28, 10, -37, 20, -37,
	31, -37, 33, -37, 34, -37, 14, -37, 9, -37, 5, -37, 25, -37,
	24, -37, 2, -37, 8, -37, 13, -37, 12, -37, 8, 20, 50, 25,
	51, 24, 52, 2, -77, 12, 10, -81, 20, -81, 5, -81, 25, -81,
	24, -81, 2, -81, 28, 10, -85, 20, -85, 31, -85, 33, -85, 34,
	-85, 14, 62, 9, -85, 5, -85, 25, -85, 24, -85, 2, -85, 8,
	-85, 13, 63, 12, 64, 22, 10, -83, 20, -83, 31, -83, 33, -83,
	34, -83, 9, -83, 5, -83, 25, -83, 24, -83, 2, -83, 8, -83,
	22, 10, -30, 20, -30, 31, 32, 33, 33, 34, -87, 9, 34, 5,
	-30, 25, -30, 24, -30, 2, -30, 8, 35, 2, 4, 59, 10, 0,
	-56, 31, -56, 18, -56, 17, -56, 19, -56, 14, 0, -70, 22, -70,
	31, -70, 18, -70, 21, -70, 23, -70, 17, -70, 8, 26, 71, 31,
	72, 28, 73, 33, 74, 8, 7, -73, 22, 15, 31, 16, 21, 17,
	8, 20, -47, 25, -47, 24, -47, 2, -47, 8, 20, -49, 25, -49,
	24, -49, 2, -49, 2, 9, 92, 8, 20, -79, 25, -79, 24, -79,
	2, -79, 2, 2, 84, 8, 20, 50, 25, 51, 24, 52, 2, -76,
	8, 20, -44, 25, -44, 24, -44, 2, -44, 8, 20, -46, 25, -46,
	24, -46, 2, -46, 8, 20, -45, 25, -45, 24, -45, 2, -45, 22,
	10, -82, 20, -82, 31, -82, 33, -82, 34, -82, 9, -82, 5, -82,
	25, -82, 24, -82, 2, -82, 8, -82, 22, 10, -34, 20, -34, 31,
	-34, 33, -34, 34, -34, 9, -34, 5, -34, 25, -34, 24, -34, 2,
	-34, 8, -34, 22, 10, -33, 20, -33, 31, -33, 33, -33, 34, -33,
	9, -33, 5, -33, 25, -33, 24, -33, 2, -33, 8, -33, 22, 10,
	-32, 20, -32, 31, -32, 33, -32, 34, -32, 9, -32, 5, -32, 25,
	-32, 24, -32, 2, -32, 8, -32, 22, 10, -84, 20, -84, 31, -84,
	33, -84, 34, -84, 9, -84, 5, -84, 25, -84, 24, -84, 2, -84,
	8, -84, 22, 10, -31, 20, -31, 31, -31, 33, -31, 34, -31, 9,
	-31, 5, -31, 25, -31, 24, -31, 2, -31, 8, -31, 2, 10, 86,
	4, 34, -87, 8, 35, 4, 37, 88, 36, 89, 4, 5, 95, 2,
	96, 26, 3, -11, 10, -11, 26, -11, 31, -11, 27, -11, 28, -11,
	33, -11, 14, -11, 5, -11, 29, -11, 2, -11, 13, -11, 12, -11,
	26, 3, -9, 10, -9, 26, -9, 31, -9, 27, -9, 28, -9, 33,
	-9, 14, -9, 5, -9, 29, -9, 2, -9, 13, -9, 12, -9, 2,
	9, 107, 26, 3, -10, 10, -10, 26, -10, 31, -10, 27, -10, 28,
	-10, 33, -10, 14, -10, 5, -10, 29, -10, 2, -10, 13, -10, 12,
	-10, 26, 3, -12, 10, -12, 26, -12, 31, -12, 27, -12, 28, -12,
	33, -12, 14, -12, 5, -12, 29, -12, 2, -12, 13, -12, 12, -12,
	4, 5, -61, 2, -61, 22, 26, -67, 31, -67, 27, -67, 28, -67,
	33, -67, 14, 102, 5, -67, 29, -67, 2, -67, 13, 103, 12, 104,
	16, 26, -63, 31, -63, 27, -63, 28, -63, 33, -63, 5, -63, 29,
	-63, 2, -63, 16, 26, 71, 31, 72, 27, 97, 28, 73, 33, 74,
	5, -65, 29, 98, 2, -65, 8, 7, -75, 22, -75, 31, -75, 21,
	-75, 2, 7, 108, 8, 7, -72, 22, 15, 31, 16, 21, 17, 2,
	2, 109, 16, 7, -27, 0, -27, 22, -27, 31, -27, 18, -27, 21,
	-27, 23, -27, 17, -27, 2, 2, 110, 28, 10, -38, 20, -38, 31,
	-38, 33, -38, 34, -38, 14, -38, 9, -38, 5, -38, 25, -38, 24,
	-38, 2, -38, 8, -38, 13, -38, 12, -38, 28, 10, -39, 20, -39,
	31, -39, 33, -39, 34, -39, 14, -39, 9, -39, 5, -39, 25, -39,
	24, -39, 2, -39, 8, -39, 13, -39, 12, -39, 6, 35, -42, 37,
	-42, 36, -42, 6, 35, -43, 37, -43, 36, -43, 6, 35, -89, 37,
	-89, 36, -89, 6, 35, 111, 37, 88, 36, 89, 2, 31, 113, 8,
	20, -78, 25, -78, 24, -78, 2, -78, 12, 10, -80, 20, -80, 5,
	-80, 25, -80, 24, -80, 2, -80, 10, 0, -6, 31, -6, 18, -6,
	17, -6, 19, -6, 2, 9, 116, 2, 9, 117, 4, 5, -64, 2,
	-64, 4, 5, -7, 2, -7, 16, 26, -62, 31, -62, 27, -62, 28,
	-62, 33, -62, 5, -62, 29, -62, 2, -62, 16, 26, -15, 31, -15,
	27, -15, 28, -15, 33, -15, 5, -15, 29, -15, 2, -15, 16, 26,
	-14, 31, -14, 27, -14, 28, -14, 33, -14, 5, -14, 29, -14, 2,
	-14, 16, 26, -16, 31, -16, 27, -16, 28, -16, 33, -16, 5, -16,
	29, -16, 2, -16, 16, 26, -66, 31, -66, 27, -66, 28, -66, 33,
	-66, 5, -66, 29, -66, 2, -66, 16, 26, -8, 31, -8, 27, -8,
	28, -8, 33, -8, 5, -8, 29, -8, 2, -8, 14, 0, -25, 22,
	-25, 31, -25, 18, -25, 21, -25, 23, -25, 17, -25, 16, 7, -26,
	0, -26, 22, -26, 31, -26, 18, -26, 21, -26, 23, -26, 17, -26,
	16, 7, -28, 0, -28, 22, -28, 31, -28, 18, -28, 21, -28, 23,
	-28, 17, -28, 30, 10, -41, 20, -41, 31, -41, 33, -41, 34, -41,
	14, -41, 9, -41, 5, -41, 25, -41, 24, -41, 2, -41, 11, -41,
	8, -41, 13, -41, 12, -41, 6, 35, -88, 37, -88, 36, -88, 2,
	10, 118, 8, 7, -74, 22, -74, 31, -74, 21, -74, 2, 3, 120,
	2, 32, 121, 2, 32, 122, 8, 20, -48, 25, -48, 24, -48, 2,
	-48, 4, 5, -60, 2, -60, 2, 10, 124, 2, 10, 125, 2, 10,
	126, 4, 5, -17, 2, -17, 4, 5, -18, 2, -18, 26, 3, -13,
	10, -13, 26, -13, 31, -13, 27, -13, 28, -13, 33, -13, 14, -13,
	5, -13, 29, -13, 2, -13, 13, -13, 12, -13,
}

var _Goto = []int32{
	127, 140, 157, 168, 168, 168, 168, 169, 168, 168, 168, 168, 168, 168,
	176, 183, 168, 168, 168, 168, 168, 168, 168, 202, 168, 168, 168, 168,
	168, 168, 215, 168, 168, 168, 234, 168, 168, 168, 168, 253, 168, 266,
	168, 271, 168, 168, 168, 282, 295, 308, 168, 168, 168, 168, 168, 321,
	168, 168, 168, 330, 349, 168, 168, 168, 168, 168, 168, 168, 364, 369,
	168, 168, 168, 168, 168, 168, 168, 374, 168, 379, 168, 168, 390, 168,
	168, 168, 168, 168, 168, 168, 168, 399, 168, 168, 168, 402, 168, 168,
	168, 168, 168, 168, 168, 168, 168, 168, 168, 413, 168, 168, 168, 168,
	168, 168, 168, 168, 168, 168, 168, 168, 418, 168, 168, 168, 168, 168,
	168, 12, 12, 3, 3, 4, 2, 5, 31, 6, 32, 7, 1, 8,
	16, 17, 19, 14, 20, 13, 21, 40, 22, 41, 23, 18, 24, 15,
	25, 16, 26, 10, 35, 10, 5, 11, 4, 12, 33, 13, 34, 14,
	0, 6, 12, 3, 3, 4, 2, 27, 6, 35, 10, 5, 11, 4,
	45, 18, 46, 31, 49, 36, 25, 37, 24, 38, 19, 39, 20, 40,
	23, 41, 21, 42, 47, 43, 12, 17, 19, 14, 20, 13, 46, 18,
	24, 15, 25, 16, 26, 18, 46, 31, 49, 36, 25, 37, 24, 38,
	19, 49, 20, 40, 23, 41, 21, 42, 47, 43, 18, 46, 31, 49,
	36, 25, 37, 24, 38, 19, 67, 20, 40, 23, 41, 21, 42, 47,
	43, 12, 27, 53, 44, 54, 45, 55, 28, 56, 30, 57, 29, 58,
	4, 22, 65, 48, 66, 10, 49, 36, 25, 37, 24, 38, 23, 41,
	21, 61, 12, 36, 70, 9, 75, 6, 76, 8, 77, 7, 78, 37,
	79, 12, 17, 19, 14, 80, 42, 81, 43, 82, 18, 24, 16, 26,
	12, 27, 53, 44, 83, 45, 55, 28, 56, 30, 57, 29, 58, 8,
	27, 93, 28, 56, 30, 57, 29, 58, 18, 46, 31, 49, 36, 25,
	37, 24, 38, 19, 85, 20, 40, 23, 41, 21, 42, 47, 43, 14,
	49, 36, 25, 37, 24, 38, 20, 94, 23, 41, 21, 42, 47, 43,
	4, 49, 36, 25, 87, 4, 26, 90, 50, 91, 4, 10, 105, 39,
	106, 10, 9, 75, 11, 99, 38, 100, 8, 77, 7, 101, 8, 17,
	19, 14, 114, 18, 24, 16, 26, 2, 26, 112, 10, 9, 75, 6,
	119, 8, 77, 7, 78, 37, 79, 4, 9, 75, 8, 115, 4, 9,
	75, 8, 123,
}

type _Bounds struct {
	Begin Token
	End   Token
	Empty bool
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

type _Lexer interface {
	ReadToken() (Token, int)
}

type lox struct {
	_lex    _Lexer
	_state  _Stack[int32]
	_sym    _Stack[any]
	_bounds _Stack[_Bounds]

	_lookahead     Token
	_lookaheadType int
	_errorToken    Token
}

func (p *parser) parse(lex _Lexer) bool {
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

			// Compute reduction token bounds.
			// Trim leading and trailing empty bounds.
			boundSlice := p._bounds.PeekSlice(int(termCount))
			for len(boundSlice) > 0 && boundSlice[0].Empty {
				boundSlice = boundSlice[1:]
			}
			for len(boundSlice) > 0 && boundSlice[len(boundSlice)-1].Empty {
				boundSlice = boundSlice[:len(boundSlice)-1]
			}
			var bounds _Bounds
			if len(boundSlice) > 0 {
				bounds.Begin = boundSlice[0].Begin
				bounds.End = boundSlice[len(boundSlice)-1].End
			} else {
				bounds.Empty = true
			}
			if !bounds.Empty {
				p._onBounds(res, bounds.Begin, bounds.End)
			}
			p._bounds.Pop(int(termCount))
			p._bounds.Push(bounds)
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
	p._onError()

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
			_cast[_i2.Token](p._sym.Peek(1)),
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		)
	case 5:
		return p.on_parser_statement(
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 6:
		return p.on_parser_rule(
			_cast[_i2.Token](p._sym.Peek(4)),
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[[]*_i1.ParserProd](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
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
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 10:
		return p.on_parser_term__token(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 11:
		return p.on_parser_term__token(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 12:
		return p.on_parser_term__list(
			_cast[*_i1.ParserTerm](p._sym.Peek(0)),
		)
	case 13:
		return p.on_parser_list(
			_cast[_i2.Token](p._sym.Peek(5)),
			_cast[_i2.Token](p._sym.Peek(4)),
			_cast[*_i1.ParserTerm](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[*_i1.ParserTerm](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 14:
		return p.on_parser_card(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 15:
		return p.on_parser_card(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 16:
		return p.on_parser_card(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 17:
		return p.on_parser_qualif(
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[_i2.Token](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 18:
		return p.on_parser_qualif(
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[_i2.Token](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 19:
		return p.on_lexer_section(
			_cast[_i2.Token](p._sym.Peek(1)),
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
			_cast[_i2.Token](p._sym.Peek(4)),
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 26:
		return p.on_token_rule(
			_cast[_i2.Token](p._sym.Peek(4)),
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[*_i1.LexerExpr](p._sym.Peek(2)),
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 27:
		return p.on_frag_rule(
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[*_i1.LexerExpr](p._sym.Peek(2)),
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 28:
		return p.on_macro_rule(
			_cast[_i2.Token](p._sym.Peek(4)),
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[*_i1.LexerExpr](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
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
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 33:
		return p.on_lexer_card(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 34:
		return p.on_lexer_card(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 35:
		return p.on_lexer_term__tok(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 36:
		return p.on_lexer_term__tok(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 37:
		return p.on_lexer_term__char_class_expr(
			_cast[_i1.CharClassExpr](p._sym.Peek(0)),
		)
	case 38:
		return p.on_lexer_term__expr(
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[*_i1.LexerExpr](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 39:
		return p.on_char_class_expr__binary(
			_cast[_i1.CharClassExpr](p._sym.Peek(2)),
			_cast[_i2.Token](p._sym.Peek(1)),
			_cast[_i1.CharClassExpr](p._sym.Peek(0)),
		)
	case 40:
		return p.on_char_class_expr__char_class(
			_cast[*_i1.CharClass](p._sym.Peek(0)),
		)
	case 41:
		return p.on_char_class(
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[[]_i2.Token](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 42:
		return p.on_char_class_item(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 43:
		return p.on_char_class_item(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 44:
		return p.on_action(
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 45:
		return p.on_action(
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 46:
		return p.on_action(
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 47:
		return p.on_action_discard(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 48:
		return p.on_action_push_mode(
			_cast[_i2.Token](p._sym.Peek(3)),
			_cast[_i2.Token](p._sym.Peek(2)),
			_cast[_i2.Token](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 49:
		return p.on_action_pop_mode(
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 50: // ZeroOrMore
		return _cast[[][]_i1.Statement](p._sym.Peek(0))
	case 51: // ZeroOrMore
		{
			var zero [][]_i1.Statement
			return zero
		}
	case 52: // OneOrMore
		return append(
			_cast[[][]_i1.Statement](p._sym.Peek(1)),
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		)
	case 53: // OneOrMore
		return [][]_i1.Statement{
			_cast[[]_i1.Statement](p._sym.Peek(0)),
		}
	case 54: // ZeroOrMore
		return _cast[[]_i1.Statement](p._sym.Peek(0))
	case 55: // ZeroOrMore
		{
			var zero []_i1.Statement
			return zero
		}
	case 56: // OneOrMore
		return append(
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 57: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._sym.Peek(0)),
		}
	case 58: // ZeroOrOne
		return _cast[_i2.Token](p._sym.Peek(0))
	case 59: // ZeroOrOne
		{
			var zero _i2.Token
			return zero
		}
	case 60: // List
		return append(
			_cast[[]*_i1.ParserProd](p._sym.Peek(2)),
			_cast[*_i1.ParserProd](p._sym.Peek(0)),
		)
	case 61: // List
		return []*_i1.ParserProd{
			_cast[*_i1.ParserProd](p._sym.Peek(0)),
		}
	case 62: // OneOrMore
		return append(
			_cast[[]*_i1.ParserTerm](p._sym.Peek(1)),
			_cast[*_i1.ParserTerm](p._sym.Peek(0)),
		)
	case 63: // OneOrMore
		return []*_i1.ParserTerm{
			_cast[*_i1.ParserTerm](p._sym.Peek(0)),
		}
	case 64: // ZeroOrOne
		return _cast[*_i1.ProdQualifier](p._sym.Peek(0))
	case 65: // ZeroOrOne
		{
			var zero *_i1.ProdQualifier
			return zero
		}
	case 66: // ZeroOrOne
		return _cast[_i1.ParserTermType](p._sym.Peek(0))
	case 67: // ZeroOrOne
		{
			var zero _i1.ParserTermType
			return zero
		}
	case 68: // ZeroOrMore
		return _cast[[]_i1.Statement](p._sym.Peek(0))
	case 69: // ZeroOrMore
		{
			var zero []_i1.Statement
			return zero
		}
	case 70: // OneOrMore
		return append(
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 71: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._sym.Peek(0)),
		}
	case 72: // ZeroOrMore
		return _cast[[]_i1.Statement](p._sym.Peek(0))
	case 73: // ZeroOrMore
		{
			var zero []_i1.Statement
			return zero
		}
	case 74: // OneOrMore
		return append(
			_cast[[]_i1.Statement](p._sym.Peek(1)),
			_cast[_i1.Statement](p._sym.Peek(0)),
		)
	case 75: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._sym.Peek(0)),
		}
	case 76: // ZeroOrMore
		return _cast[[]_i1.Action](p._sym.Peek(0))
	case 77: // ZeroOrMore
		{
			var zero []_i1.Action
			return zero
		}
	case 78: // OneOrMore
		return append(
			_cast[[]_i1.Action](p._sym.Peek(1)),
			_cast[_i1.Action](p._sym.Peek(0)),
		)
	case 79: // OneOrMore
		return []_i1.Action{
			_cast[_i1.Action](p._sym.Peek(0)),
		}
	case 80: // List
		return append(
			_cast[[]*_i1.LexerFactor](p._sym.Peek(2)),
			_cast[*_i1.LexerFactor](p._sym.Peek(0)),
		)
	case 81: // List
		return []*_i1.LexerFactor{
			_cast[*_i1.LexerFactor](p._sym.Peek(0)),
		}
	case 82: // OneOrMore
		return append(
			_cast[[]*_i1.LexerTermCard](p._sym.Peek(1)),
			_cast[*_i1.LexerTermCard](p._sym.Peek(0)),
		)
	case 83: // OneOrMore
		return []*_i1.LexerTermCard{
			_cast[*_i1.LexerTermCard](p._sym.Peek(0)),
		}
	case 84: // ZeroOrOne
		return _cast[_i1.Card](p._sym.Peek(0))
	case 85: // ZeroOrOne
		{
			var zero _i1.Card
			return zero
		}
	case 86: // ZeroOrOne
		return _cast[_i2.Token](p._sym.Peek(0))
	case 87: // ZeroOrOne
		{
			var zero _i2.Token
			return zero
		}
	case 88: // OneOrMore
		return append(
			_cast[[]_i2.Token](p._sym.Peek(1)),
			_cast[_i2.Token](p._sym.Peek(0)),
		)
	case 89: // OneOrMore
		return []_i2.Token{
			_cast[_i2.Token](p._sym.Peek(0)),
		}
	default:
		panic("unreachable")
	}
}
