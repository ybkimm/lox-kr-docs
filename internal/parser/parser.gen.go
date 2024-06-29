package parser

import (
	_i0 "github.com/dcaiafa/lox/internal/ast"
	_i1 "github.com/dcaiafa/loxlex/simplelexer"
)

var _rules = []int32{
	0, 1, 1, 2, 2, 3, 4, 5, 6, 7, 8, 8, 8, 8,
	9, 10, 10, 10, 11, 11, 12, 13, 13, 14, 14, 14, 15, 16,
	17, 18, 19, 20, 21, 22, 22, 22, 23, 23, 23, 23, 24, 24,
	25, 26, 26, 27, 27, 27, 27, 28, 29, 30, 31, 32, 32, 33,
	33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40,
	40, 41, 41, 42, 42, 43, 43, 44, 44, 45, 45, 46, 46, 47,
	47, 48, 48, 49, 49, 50, 50, 51, 51, 52, 52,
}

var _termCounts = []int32{
	1, 1, 1, 1, 1, 2, 1, 5, 2, 2, 1, 1, 1, 1,
	6, 1, 1, 1, 4, 4, 2, 1, 1, 1, 1, 1, 5, 5,
	4, 5, 1, 1, 2, 1, 1, 1, 1, 1, 1, 3, 3, 1,
	4, 1, 1, 1, 1, 1, 1, 1, 4, 1, 4, 1, 0, 2,
	1, 1, 0, 2, 1, 1, 0, 3, 1, 2, 1, 1, 0, 1,
	0, 1, 0, 2, 1, 1, 0, 2, 1, 1, 0, 2, 1, 3,
	1, 2, 1, 1, 0, 1, 0, 2, 1, 1, 0,
}

var _actions = []int32{
	134, 143, 146, 161, 172, 179, 186, 193, 196, 203, 206, 209, 212, 223,
	234, 241, 252, 263, 266, 269, 272, 289, 304, 319, 326, 341, 358, 373,
	390, 397, 400, 252, 403, 418, 449, 252, 480, 483, 486, 519, 550, 561,
	576, 607, 632, 657, 660, 671, 686, 695, 550, 704, 715, 718, 729, 732,
	743, 746, 757, 768, 779, 790, 252, 252, 801, 826, 851, 876, 901, 926,
	951, 954, 959, 964, 969, 996, 1023, 1026, 1053, 1080, 1085, 1108, 1125, 1142,
	1151, 1154, 1163, 1166, 1183, 1186, 1217, 1248, 1255, 1262, 1269, 1276, 1281, 1284,
	1295, 686, 1310, 1321, 1324, 1327, 1332, 1337, 1354, 1371, 1388, 1405, 1422, 686,
	1439, 1454, 1471, 1488, 1521, 1528, 1531, 1534, 1537, 1546, 1549, 1552, 1555, 1566,
	1577, 686, 1582, 1585, 1588, 1591, 1596, 1601, 8, 0, -54, 1, 1, 18,
	2, 17, 3, 2, 0, -2, 14, 0, -72, 22, 16, 32, 17, 18,
	-72, 21, 18, 23, 19, 17, -72, 10, 0, -58, 32, -62, 18, -58,
	17, -58, 19, 10, 6, 0, -4, 18, -4, 17, -4, 6, 0, -3,
	18, -3, 17, -3, 6, 0, -56, 18, -56, 17, -56, 2, 0, -1,
	6, 0, -53, 18, 2, 17, 3, 2, 0, 2147483647, 2, 32, -61, 2,
	32, 29, 10, 0, -6, 32, -6, 18, -6, 17, -6, 19, -6, 10,
	0, -60, 32, -60, 18, -60, 17, -60, 19, -60, 6, 0, -5, 18,
	-5, 17, -5, 10, 0, -57, 32, -62, 18, -57, 17, -57, 19, 10,
	10, 32, 33, 34, 34, 35, -90, 9, 35, 8, 36, 2, 4, 31,
	2, 32, 45, 2, 32, 30, 16, 7, -24, 0, -24, 22, -24, 32,
	-24, 18, -24, 21, -24, 23, -24, 17, -24, 14, 0, -22, 22, -22,
	32, -22, 18, -22, 21, -22, 23, -22, 17, -22, 14, 0, -74, 22,
	-74, 32, -74, 18, -74, 21, -74, 23, -74, 17, -74, 6, 0, -20,
	18, -20, 17, -20, 14, 0, -71, 22, 16, 32, 17, 18, -71, 21,
	18, 23, 19, 17, -71, 16, 7, -25, 0, -25, 22, -25, 32, -25,
	18, -25, 21, -25, 23, -25, 17, -25, 14, 0, -21, 22, -21, 32,
	-21, 18, -21, 21, -21, 23, -21, 17, -21, 16, 7, -23, 0, -23,
	22, -23, 32, -23, 18, -23, 21, -23, 23, -23, 17, -23, 6, 0,
	-55, 18, -55, 17, -55, 2, 4, 48, 2, 6, 49, 14, 10, -30,
	20, -30, 30, -30, 5, 63, 25, -30, 24, -30, 2, -30, 30, 10,
	-37, 20, -37, 30, -37, 32, -37, 34, -37, 35, -37, 14, -37, 9,
	-37, 5, -37, 25, -37, 24, -37, 2, -37, 8, -37, 13, -37, 12,
	-37, 30, 10, -36, 20, -36, 30, -36, 32, -36, 34, -36, 35, -36,
	14, -36, 9, -36, 5, -36, 25, -36, 24, -36, 2, -36, 8, -36,
	13, -36, 12, -36, 2, 35, -89, 2, 35, 72, 32, 10, -41, 20,
	-41, 30, -41, 32, -41, 34, -41, 35, -41, 14, -41, 9, -41, 5,
	-41, 25, -41, 24, -41, 2, -41, 11, 71, 8, -41, 13, -41, 12,
	-41, 30, 10, -38, 20, -38, 30, -38, 32, -38, 34, -38, 35, -38,
	14, -38, 9, -38, 5, -38, 25, -38, 24, -38, 2, -38, 8, -38,
	13, -38, 12, -38, 10, 20, 51, 30, 52, 25, 53, 24, 54, 2,
	-80, 14, 10, -84, 20, -84, 30, -84, 5, -84, 25, -84, 24, -84,
	2, -84, 30, 10, -88, 20, -88, 30, -88, 32, -88, 34, -88, 35,
	-88, 14, 65, 9, -88, 5, -88, 25, -88, 24, -88, 2, -88, 8,
	-88, 13, 66, 12, 67, 24, 10, -86, 20, -86, 30, -86, 32, -86,
	34, -86, 35, -86, 9, -86, 5, -86, 25, -86, 24, -86, 2, -86,
	8, -86, 24, 10, -31, 20, -31, 30, -31, 32, 33, 34, 34, 35,
	-90, 9, 35, 5, -31, 25, -31, 24, -31, 2, -31, 8, 36, 2,
	4, 62, 10, 0, -59, 32, -59, 18, -59, 17, -59, 19, -59, 14,
	0, -73, 22, -73, 32, -73, 18, -73, 21, -73, 23, -73, 17, -73,
	8, 26, 74, 32, 75, 28, 76, 34, 77, 8, 7, -76, 22, 16,
	32, 17, 21, 18, 10, 20, -49, 30, -49, 25, -49, 24, -49, 2,
	-49, 2, 9, 96, 10, 20, -51, 30, -51, 25, -51, 24, -51, 2,
	-51, 2, 9, 95, 10, 20, -82, 30, -82, 25, -82, 24, -82, 2,
	-82, 2, 2, 87, 10, 20, 51, 30, 52, 25, 53, 24, 54, 2,
	-79, 10, 20, -45, 30, -45, 25, -45, 24, -45, 2, -45, 10, 20,
	-48, 30, -48, 25, -48, 24, -48, 2, -48, 10, 20, -47, 30, -47,
	25, -47, 24, -47, 2, -47, 10, 20, -46, 30, -46, 25, -46, 24,
	-46, 2, -46, 24, 10, -85, 20, -85, 30, -85, 32, -85, 34, -85,
	35, -85, 9, -85, 5, -85, 25, -85, 24, -85, 2, -85, 8, -85,
	24, 10, -35, 20, -35, 30, -35, 32, -35, 34, -35, 35, -35, 9,
	-35, 5, -35, 25, -35, 24, -35, 2, -35, 8, -35, 24, 10, -34,
	20, -34, 30, -34, 32, -34, 34, -34, 35, -34, 9, -34, 5, -34,
	25, -34, 24, -34, 2, -34, 8, -34, 24, 10, -33, 20, -33, 30,
	-33, 32, -33, 34, -33, 35, -33, 9, -33, 5, -33, 25, -33, 24,
	-33, 2, -33, 8, -33, 24, 10, -87, 20, -87, 30, -87, 32, -87,
	34, -87, 35, -87, 9, -87, 5, -87, 25, -87, 24, -87, 2, -87,
	8, -87, 24, 10, -32, 20, -32, 30, -32, 32, -32, 34, -32, 35,
	-32, 9, -32, 5, -32, 25, -32, 24, -32, 2, -32, 8, -32, 2,
	10, 89, 4, 35, -90, 8, 36, 4, 38, 91, 37, 92, 4, 5,
	99, 2, 100, 26, 3, -12, 10, -12, 26, -12, 32, -12, 27, -12,
	28, -12, 34, -12, 14, -12, 5, -12, 29, -12, 2, -12, 13, -12,
	12, -12, 26, 3, -10, 10, -10, 26, -10, 32, -10, 27, -10, 28,
	-10, 34, -10, 14, -10, 5, -10, 29, -10, 2, -10, 13, -10, 12,
	-10, 2, 9, 111, 26, 3, -11, 10, -11, 26, -11, 32, -11, 27,
	-11, 28, -11, 34, -11, 14, -11, 5, -11, 29, -11, 2, -11, 13,
	-11, 12, -11, 26, 3, -13, 10, -13, 26, -13, 32, -13, 27, -13,
	28, -13, 34, -13, 14, -13, 5, -13, 29, -13, 2, -13, 13, -13,
	12, -13, 4, 5, -64, 2, -64, 22, 26, -70, 32, -70, 27, -70,
	28, -70, 34, -70, 14, 106, 5, -70, 29, -70, 2, -70, 13, 107,
	12, 108, 16, 26, -66, 32, -66, 27, -66, 28, -66, 34, -66, 5,
	-66, 29, -66, 2, -66, 16, 26, 74, 32, 75, 27, 101, 28, 76,
	34, 77, 5, -68, 29, 102, 2, -68, 8, 7, -78, 22, -78, 32,
	-78, 21, -78, 2, 7, 112, 8, 7, -75, 22, 16, 32, 17, 21,
	18, 2, 2, 113, 16, 7, -28, 0, -28, 22, -28, 32, -28, 18,
	-28, 21, -28, 23, -28, 17, -28, 2, 2, 114, 30, 10, -39, 20,
	-39, 30, -39, 32, -39, 34, -39, 35, -39, 14, -39, 9, -39, 5,
	-39, 25, -39, 24, -39, 2, -39, 8, -39, 13, -39, 12, -39, 30,
	10, -40, 20, -40, 30, -40, 32, -40, 34, -40, 35, -40, 14, -40,
	9, -40, 5, -40, 25, -40, 24, -40, 2, -40, 8, -40, 13, -40,
	12, -40, 6, 36, -43, 38, -43, 37, -43, 6, 36, -44, 38, -44,
	37, -44, 6, 36, -92, 38, -92, 37, -92, 6, 36, 115, 38, 91,
	37, 92, 4, 10, -94, 32, 117, 2, 32, 119, 10, 20, -81, 30,
	-81, 25, -81, 24, -81, 2, -81, 14, 10, -83, 20, -83, 30, -83,
	5, -83, 25, -83, 24, -83, 2, -83, 10, 0, -7, 32, -7, 18,
	-7, 17, -7, 19, -7, 2, 9, 122, 2, 9, 123, 4, 5, -67,
	2, -67, 4, 5, -8, 2, -8, 16, 26, -65, 32, -65, 27, -65,
	28, -65, 34, -65, 5, -65, 29, -65, 2, -65, 16, 26, -16, 32,
	-16, 27, -16, 28, -16, 34, -16, 5, -16, 29, -16, 2, -16, 16,
	26, -15, 32, -15, 27, -15, 28, -15, 34, -15, 5, -15, 29, -15,
	2, -15, 16, 26, -17, 32, -17, 27, -17, 28, -17, 34, -17, 5,
	-17, 29, -17, 2, -17, 16, 26, -69, 32, -69, 27, -69, 28, -69,
	34, -69, 5, -69, 29, -69, 2, -69, 16, 26, -9, 32, -9, 27,
	-9, 28, -9, 34, -9, 5, -9, 29, -9, 2, -9, 14, 0, -26,
	22, -26, 32, -26, 18, -26, 21, -26, 23, -26, 17, -26, 16, 7,
	-27, 0, -27, 22, -27, 32, -27, 18, -27, 21, -27, 23, -27, 17,
	-27, 16, 7, -29, 0, -29, 22, -29, 32, -29, 18, -29, 21, -29,
	23, -29, 17, -29, 32, 10, -42, 20, -42, 30, -42, 32, -42, 34,
	-42, 35, -42, 14, -42, 9, -42, 5, -42, 25, -42, 24, -42, 2,
	-42, 11, -42, 8, -42, 13, -42, 12, -42, 6, 36, -91, 38, -91,
	37, -91, 2, 10, -93, 2, 10, 124, 2, 10, 125, 8, 7, -77,
	22, -77, 32, -77, 21, -77, 2, 3, 127, 2, 33, 128, 2, 33,
	129, 10, 20, -50, 30, -50, 25, -50, 24, -50, 2, -50, 10, 20,
	-52, 30, -52, 25, -52, 24, -52, 2, -52, 4, 5, -63, 2, -63,
	2, 10, 131, 2, 10, 132, 2, 10, 133, 4, 5, -18, 2, -18,
	4, 5, -19, 2, -19, 26, 3, -14, 10, -14, 26, -14, 32, -14,
	27, -14, 28, -14, 34, -14, 14, -14, 5, -14, 29, -14, 2, -14,
	13, -14, 12, -14,
}

var _goto = []int32{
	134, 147, 148, 165, 147, 147, 147, 147, 176, 147, 147, 147, 147, 147,
	147, 183, 190, 147, 147, 147, 147, 147, 147, 147, 209, 147, 147, 147,
	147, 147, 147, 222, 147, 147, 147, 241, 147, 147, 147, 147, 260, 147,
	275, 147, 280, 147, 147, 147, 291, 304, 317, 147, 147, 147, 147, 147,
	147, 332, 147, 147, 147, 147, 343, 362, 147, 147, 147, 147, 147, 147,
	147, 377, 382, 147, 147, 147, 147, 147, 147, 147, 387, 147, 392, 147,
	147, 403, 147, 147, 147, 147, 147, 147, 147, 147, 412, 415, 147, 147,
	147, 418, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 429,
	147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147,
	147, 434, 147, 147, 147, 147, 147, 147, 12, 12, 4, 3, 5, 2,
	6, 32, 7, 33, 8, 1, 9, 0, 16, 17, 20, 14, 21, 13,
	22, 41, 23, 42, 24, 18, 25, 15, 26, 16, 27, 10, 36, 11,
	5, 12, 4, 13, 34, 14, 35, 15, 6, 12, 4, 3, 5, 2,
	28, 6, 36, 11, 5, 12, 4, 46, 18, 47, 32, 50, 37, 25,
	38, 24, 39, 19, 40, 20, 41, 23, 42, 21, 43, 48, 44, 12,
	17, 20, 14, 21, 13, 47, 18, 25, 15, 26, 16, 27, 18, 47,
	32, 50, 37, 25, 38, 24, 39, 19, 50, 20, 41, 23, 42, 21,
	43, 48, 44, 18, 47, 32, 50, 37, 25, 38, 24, 39, 19, 70,
	20, 41, 23, 42, 21, 43, 48, 44, 14, 27, 55, 45, 56, 46,
	57, 28, 58, 31, 59, 30, 60, 29, 61, 4, 22, 68, 49, 69,
	10, 50, 37, 25, 38, 24, 39, 23, 42, 21, 64, 12, 37, 73,
	9, 78, 6, 79, 8, 80, 7, 81, 38, 82, 12, 17, 20, 14,
	83, 43, 84, 44, 85, 18, 25, 16, 27, 14, 27, 55, 45, 86,
	46, 57, 28, 58, 31, 59, 30, 60, 29, 61, 10, 27, 97, 28,
	58, 31, 59, 30, 60, 29, 61, 18, 47, 32, 50, 37, 25, 38,
	24, 39, 19, 88, 20, 41, 23, 42, 21, 43, 48, 44, 14, 50,
	37, 25, 38, 24, 39, 20, 98, 23, 42, 21, 43, 48, 44, 4,
	50, 37, 25, 90, 4, 26, 93, 51, 94, 4, 10, 109, 40, 110,
	10, 9, 78, 11, 103, 39, 104, 8, 80, 7, 105, 8, 17, 20,
	14, 120, 18, 25, 16, 27, 2, 26, 116, 2, 52, 118, 10, 9,
	78, 6, 126, 8, 80, 7, 81, 38, 82, 4, 9, 78, 8, 121,
	4, 9, 78, 8, 130,
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

type Error struct {
	Token    Token
	Expected []int
}

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

type _item struct {
	State  int32
	Sym    any
	Bounds _Bounds
}

type lox struct {
	_lex   _Lexer
	_stack _Stack[_item]

	_la    int
	_lasym any

	_qla    int
	_qlasym any
}

func (p *parser) parse(lex _Lexer) bool {
	const accept = 2147483647

	p._lex = lex
	p._qla = -1
	p._stack.Push(_item{})

	p._readToken()

	for {
		topState := p._stack.Peek(0).State
		action, ok := _Find(_actions, topState, int32(p._la))
		if !ok {
			if !p._recover() {
				return false
			}
			continue
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			latok := p._lasym.(Token)
			p._stack.Push(_item{
				State: action,
				Sym:   p._lasym,
				Bounds: _Bounds{
					Begin: latok,
					End:   latok,
				},
			})
			p._readToken()
		} else { // reduce
			prod := -action
			termCount := _termCounts[int(prod)]
			rule := _rules[int(prod)]
			res := p._act(prod)

			// Compute reduction token bounds.
			// Trim leading and trailing empty bounds.
			boundSlice := p._stack.PeekSlice(int(termCount))
			for len(boundSlice) > 0 && boundSlice[0].Bounds.Empty {
				boundSlice = boundSlice[1:]
			}
			for len(boundSlice) > 0 && boundSlice[len(boundSlice)-1].Bounds.Empty {
				boundSlice = boundSlice[:len(boundSlice)-1]
			}
			var bounds _Bounds
			if len(boundSlice) > 0 {
				bounds.Begin = boundSlice[0].Bounds.Begin
				bounds.End = boundSlice[len(boundSlice)-1].Bounds.End
			} else {
				bounds.Empty = true
			}
			if !bounds.Empty {
				p._onBounds(res, bounds.Begin, bounds.End)
			}
			p._stack.Pop(int(termCount))
			topState = p._stack.Peek(0).State
			nextState, _ := _Find(_goto, topState, rule)
			p._stack.Push(_item{
				State:  nextState,
				Sym:    res,
				Bounds: bounds,
			})
		}
	}

	return true
}

// recoverLookahead can be called during an error production action (an action
// for a production that has a @error term) to recover the lookahead that was
// possibly lost in the process of reducing the error production.
func (p *parser) recoverLookahead(typ int, tok Token) {
	if p._qla != -1 {
		panic("recovered lookahead already pending")
	}

	p._qla = p._la
	p._qlasym = p._lasym
	p._la = typ
	p._lasym = tok
}

func (p *parser) _readToken() {
	if p._qla != -1 {
		p._la = p._qla
		p._lasym = p._qlasym
		p._qla = -1
		p._qlasym = nil
		return
	}

	p._lasym, p._la = p._lex.ReadToken()
	if p._la == ERROR {
		p._lasym = p._makeError()
	}
}

func (p *parser) _recover() bool {
	errSym, ok := p._lasym.(Error)
	if !ok {
		errSym = p._makeError()
	}

	for p._la == ERROR {
		p._readToken()
	}

	for {
		save := p._stack

		for len(p._stack) >= 1 {
			state := p._stack.Peek(0).State

			for {
				action, ok := _Find(_actions, state, int32(ERROR))
				if !ok {
					break
				}

				if action < 0 {
					prod := -action
					rule := _rules[int(prod)]
					state, _ = _Find(_goto, state, rule)
					continue
				}

				state = action

				_, ok = _Find(_actions, state, int32(p._la))
				if !ok {
					break
				}

				p._qla = p._la
				p._qlasym = p._lasym
				p._la = ERROR
				p._lasym = errSym
				return true
			}

			p._stack.Pop(1)
		}

		if p._la == EOF {
			return false
		}

		p._stack = save
		p._readToken()
	}
}

func (p *parser) _makeError() Error {
	e := Error{
		Token: p._lasym.(Token),
	}

	// Compile list of allowed tokens at this state.
	// See _Find for the format of the _actions table.
	s := p._stack.Peek(0).State
	i := int(_actions[int(s)])
	count := int(_actions[i])
	i++
	end := i + count
	for ; i < end; i += 2 {
		e.Expected = append(e.Expected, int(_actions[i]))
	}

	return e
}

func (p *parser) _act(prod int32) any {
	switch prod {
	case 1:
		return p.on_spec(
			_cast[[][]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 2:
		return p.on_spec__error(
			_cast[Error](p._stack.Peek(0).Sym),
		)
	case 3:
		return p.on_section(
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 4:
		return p.on_section(
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 5:
		return p.on_parser_section(
			_cast[_i1.Token](p._stack.Peek(1).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 6:
		return p.on_parser_statement(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 7:
		return p.on_parser_rule(
			_cast[_i1.Token](p._stack.Peek(4).Sym),
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[[]*_i0.ParserProd](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 8:
		return p.on_parser_prod(
			_cast[[]*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[*_i0.ProdQualifier](p._stack.Peek(0).Sym),
		)
	case 9:
		return p.on_parser_term_card(
			_cast[*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[_i0.ParserTermType](p._stack.Peek(0).Sym),
		)
	case 10:
		return p.on_parser_term__token(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 11:
		return p.on_parser_term__token(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 12:
		return p.on_parser_term__token(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 13:
		return p.on_parser_term__list(
			_cast[*_i0.ParserTerm](p._stack.Peek(0).Sym),
		)
	case 14:
		return p.on_parser_list(
			_cast[_i1.Token](p._stack.Peek(5).Sym),
			_cast[_i1.Token](p._stack.Peek(4).Sym),
			_cast[*_i0.ParserTerm](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 15:
		return p.on_parser_card(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 16:
		return p.on_parser_card(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 17:
		return p.on_parser_card(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 18:
		return p.on_parser_qualif(
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[_i1.Token](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 19:
		return p.on_parser_qualif(
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[_i1.Token](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 20:
		return p.on_lexer_section(
			_cast[_i1.Token](p._stack.Peek(1).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 21:
		return p.on_lexer_statement(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 22:
		return p.on_lexer_statement(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 23:
		return p.on_lexer_rule(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 24:
		return p.on_lexer_rule(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 25:
		return p.on_lexer_rule(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 26:
		return p.on_mode(
			_cast[_i1.Token](p._stack.Peek(4).Sym),
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 27:
		return p.on_token_rule(
			_cast[_i1.Token](p._stack.Peek(4).Sym),
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(2).Sym),
			_cast[[]_i0.Action](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 28:
		return p.on_frag_rule(
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(2).Sym),
			_cast[[]_i0.Action](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 29:
		return p.on_macro_rule(
			_cast[_i1.Token](p._stack.Peek(4).Sym),
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 30:
		return p.on_lexer_expr(
			_cast[[]*_i0.LexerFactor](p._stack.Peek(0).Sym),
		)
	case 31:
		return p.on_lexer_factor(
			_cast[[]*_i0.LexerTermCard](p._stack.Peek(0).Sym),
		)
	case 32:
		return p.on_lexer_term_card(
			_cast[_i0.LexerTerm](p._stack.Peek(1).Sym),
			_cast[_i0.Card](p._stack.Peek(0).Sym),
		)
	case 33:
		return p.on_lexer_card(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 34:
		return p.on_lexer_card(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 35:
		return p.on_lexer_card(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 36:
		return p.on_lexer_term__tok(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 37:
		return p.on_lexer_term__tok(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 38:
		return p.on_lexer_term__char_class_expr(
			_cast[_i0.CharClassExpr](p._stack.Peek(0).Sym),
		)
	case 39:
		return p.on_lexer_term__expr(
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 40:
		return p.on_char_class_expr__binary(
			_cast[_i0.CharClassExpr](p._stack.Peek(2).Sym),
			_cast[_i1.Token](p._stack.Peek(1).Sym),
			_cast[_i0.CharClassExpr](p._stack.Peek(0).Sym),
		)
	case 41:
		return p.on_char_class_expr__char_class(
			_cast[*_i0.CharClass](p._stack.Peek(0).Sym),
		)
	case 42:
		return p.on_char_class(
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[[]_i1.Token](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 43:
		return p.on_char_class_item(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 44:
		return p.on_char_class_item(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 45:
		return p.on_action(
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 46:
		return p.on_action(
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 47:
		return p.on_action(
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 48:
		return p.on_action(
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 49:
		return p.on_action_discard(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 50:
		return p.on_action_push_mode(
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[_i1.Token](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 51:
		return p.on_action_pop_mode(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 52:
		return p.on_action_emit(
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[_i1.Token](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 53: // ZeroOrMore
		return _cast[[][]_i0.Statement](p._stack.Peek(0).Sym)
	case 54: // ZeroOrMore
		{
			var zero [][]_i0.Statement
			return zero
		}
	case 55: // OneOrMore
		return append(
			_cast[[][]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 56: // OneOrMore
		return [][]_i0.Statement{
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 57: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 58: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 59: // OneOrMore
		return append(
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 60: // OneOrMore
		return []_i0.Statement{
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 61: // ZeroOrOne
		return _cast[_i1.Token](p._stack.Peek(0).Sym)
	case 62: // ZeroOrOne
		{
			var zero _i1.Token
			return zero
		}
	case 63: // List
		return append(
			_cast[[]*_i0.ParserProd](p._stack.Peek(2).Sym),
			_cast[*_i0.ParserProd](p._stack.Peek(0).Sym),
		)
	case 64: // List
		return []*_i0.ParserProd{
			_cast[*_i0.ParserProd](p._stack.Peek(0).Sym),
		}
	case 65: // OneOrMore
		return append(
			_cast[[]*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[*_i0.ParserTerm](p._stack.Peek(0).Sym),
		)
	case 66: // OneOrMore
		return []*_i0.ParserTerm{
			_cast[*_i0.ParserTerm](p._stack.Peek(0).Sym),
		}
	case 67: // ZeroOrOne
		return _cast[*_i0.ProdQualifier](p._stack.Peek(0).Sym)
	case 68: // ZeroOrOne
		{
			var zero *_i0.ProdQualifier
			return zero
		}
	case 69: // ZeroOrOne
		return _cast[_i0.ParserTermType](p._stack.Peek(0).Sym)
	case 70: // ZeroOrOne
		{
			var zero _i0.ParserTermType
			return zero
		}
	case 71: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 72: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 73: // OneOrMore
		return append(
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 74: // OneOrMore
		return []_i0.Statement{
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 75: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 76: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 77: // OneOrMore
		return append(
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 78: // OneOrMore
		return []_i0.Statement{
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 79: // ZeroOrMore
		return _cast[[]_i0.Action](p._stack.Peek(0).Sym)
	case 80: // ZeroOrMore
		{
			var zero []_i0.Action
			return zero
		}
	case 81: // OneOrMore
		return append(
			_cast[[]_i0.Action](p._stack.Peek(1).Sym),
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 82: // OneOrMore
		return []_i0.Action{
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		}
	case 83: // List
		return append(
			_cast[[]*_i0.LexerFactor](p._stack.Peek(2).Sym),
			_cast[*_i0.LexerFactor](p._stack.Peek(0).Sym),
		)
	case 84: // List
		return []*_i0.LexerFactor{
			_cast[*_i0.LexerFactor](p._stack.Peek(0).Sym),
		}
	case 85: // OneOrMore
		return append(
			_cast[[]*_i0.LexerTermCard](p._stack.Peek(1).Sym),
			_cast[*_i0.LexerTermCard](p._stack.Peek(0).Sym),
		)
	case 86: // OneOrMore
		return []*_i0.LexerTermCard{
			_cast[*_i0.LexerTermCard](p._stack.Peek(0).Sym),
		}
	case 87: // ZeroOrOne
		return _cast[_i0.Card](p._stack.Peek(0).Sym)
	case 88: // ZeroOrOne
		{
			var zero _i0.Card
			return zero
		}
	case 89: // ZeroOrOne
		return _cast[_i1.Token](p._stack.Peek(0).Sym)
	case 90: // ZeroOrOne
		{
			var zero _i1.Token
			return zero
		}
	case 91: // OneOrMore
		return append(
			_cast[[]_i1.Token](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 92: // OneOrMore
		return []_i1.Token{
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		}
	case 93: // ZeroOrOne
		return _cast[_i1.Token](p._stack.Peek(0).Sym)
	case 94: // ZeroOrOne
		{
			var zero _i1.Token
			return zero
		}
	default:
		panic("unreachable")
	}
}
