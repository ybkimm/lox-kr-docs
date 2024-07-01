package parser

import (
	_i0 "github.com/dcaiafa/lox/internal/ast"
	_i1 "github.com/dcaiafa/loxlex/simplelexer"
)

var _rules = []int32{
	0, 1, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 8, 8,
	8, 9, 10, 10, 10, 11, 11, 12, 13, 13, 14, 14, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 22, 22, 23, 23, 23, 23, 24,
	24, 25, 26, 26, 27, 27, 27, 27, 28, 29, 30, 31, 32, 32,
	33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39,
	40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 45, 45, 46, 46,
	47, 47, 48, 48, 49, 49, 50, 50, 51, 51, 52, 52,
}

var _termCounts = []int32{
	1, 1, 1, 1, 1, 2, 1, 5, 2, 1, 2, 1, 1, 1,
	1, 6, 1, 1, 1, 4, 4, 2, 1, 1, 1, 1, 1, 5,
	5, 4, 5, 1, 1, 2, 1, 1, 1, 1, 1, 1, 3, 3,
	1, 4, 1, 1, 1, 1, 1, 1, 1, 4, 1, 4, 1, 0,
	2, 1, 1, 0, 2, 1, 1, 0, 3, 1, 2, 1, 1, 0,
	1, 0, 1, 0, 2, 1, 1, 0, 2, 1, 1, 0, 2, 1,
	3, 1, 2, 1, 1, 0, 1, 0, 2, 1, 1, 0,
}

var _actions = []int32{
	135, 144, 147, 162, 173, 180, 187, 194, 197, 204, 207, 210, 213, 224,
	235, 242, 253, 264, 267, 270, 273, 290, 305, 320, 327, 342, 359, 374,
	391, 398, 401, 253, 404, 419, 450, 253, 481, 484, 487, 520, 551, 562,
	577, 608, 633, 658, 661, 672, 687, 698, 551, 707, 718, 721, 732, 735,
	746, 749, 760, 771, 782, 793, 253, 253, 804, 829, 854, 879, 904, 929,
	954, 957, 962, 967, 972, 977, 1004, 1031, 1034, 1061, 1088, 1093, 1116, 1133,
	1150, 1159, 1162, 1171, 1174, 1191, 1194, 1225, 1256, 1263, 1270, 1277, 1284, 1289,
	1292, 1303, 687, 1318, 1329, 1332, 1335, 1340, 1345, 1362, 1379, 1396, 1413, 1430,
	1447, 1456, 1471, 1488, 1505, 1538, 1545, 1548, 1551, 1554, 1563, 1566, 1569, 1572,
	1583, 1594, 1447, 1599, 1602, 1605, 1608, 1613, 1618, 8, 0, -55, 1, 1,
	18, 2, 17, 3, 2, 0, -2, 14, 0, -73, 22, 16, 33, 17,
	18, -73, 21, 18, 23, 19, 17, -73, 10, 0, -59, 33, -63, 18,
	-59, 17, -59, 19, 10, 6, 0, -4, 18, -4, 17, -4, 6, 0,
	-3, 18, -3, 17, -3, 6, 0, -57, 18, -57, 17, -57, 2, 0,
	-1, 6, 0, -54, 18, 2, 17, 3, 2, 0, 2147483647, 2, 33, -62,
	2, 33, 29, 10, 0, -6, 33, -6, 18, -6, 17, -6, 19, -6,
	10, 0, -61, 33, -61, 18, -61, 17, -61, 19, -61, 6, 0, -5,
	18, -5, 17, -5, 10, 0, -58, 33, -63, 18, -58, 17, -58, 19,
	10, 10, 33, 33, 35, 34, 36, -91, 9, 35, 8, 36, 2, 4,
	31, 2, 33, 45, 2, 33, 30, 16, 7, -25, 0, -25, 22, -25,
	33, -25, 18, -25, 21, -25, 23, -25, 17, -25, 14, 0, -23, 22,
	-23, 33, -23, 18, -23, 21, -23, 23, -23, 17, -23, 14, 0, -75,
	22, -75, 33, -75, 18, -75, 21, -75, 23, -75, 17, -75, 6, 0,
	-21, 18, -21, 17, -21, 14, 0, -72, 22, 16, 33, 17, 18, -72,
	21, 18, 23, 19, 17, -72, 16, 7, -26, 0, -26, 22, -26, 33,
	-26, 18, -26, 21, -26, 23, -26, 17, -26, 14, 0, -22, 22, -22,
	33, -22, 18, -22, 21, -22, 23, -22, 17, -22, 16, 7, -24, 0,
	-24, 22, -24, 33, -24, 18, -24, 21, -24, 23, -24, 17, -24, 6,
	0, -56, 18, -56, 17, -56, 2, 4, 48, 2, 6, 49, 14, 10,
	-31, 20, -31, 30, -31, 5, 63, 25, -31, 24, -31, 2, -31, 30,
	10, -38, 20, -38, 30, -38, 33, -38, 35, -38, 36, -38, 14, -38,
	9, -38, 5, -38, 25, -38, 24, -38, 2, -38, 8, -38, 13, -38,
	12, -38, 30, 10, -37, 20, -37, 30, -37, 33, -37, 35, -37, 36,
	-37, 14, -37, 9, -37, 5, -37, 25, -37, 24, -37, 2, -37, 8,
	-37, 13, -37, 12, -37, 2, 36, -90, 2, 36, 72, 32, 10, -42,
	20, -42, 30, -42, 33, -42, 35, -42, 36, -42, 14, -42, 9, -42,
	5, -42, 25, -42, 24, -42, 2, -42, 11, 71, 8, -42, 13, -42,
	12, -42, 30, 10, -39, 20, -39, 30, -39, 33, -39, 35, -39, 36,
	-39, 14, -39, 9, -39, 5, -39, 25, -39, 24, -39, 2, -39, 8,
	-39, 13, -39, 12, -39, 10, 20, 51, 30, 52, 25, 53, 24, 54,
	2, -81, 14, 10, -85, 20, -85, 30, -85, 5, -85, 25, -85, 24,
	-85, 2, -85, 30, 10, -89, 20, -89, 30, -89, 33, -89, 35, -89,
	36, -89, 14, 65, 9, -89, 5, -89, 25, -89, 24, -89, 2, -89,
	8, -89, 13, 66, 12, 67, 24, 10, -87, 20, -87, 30, -87, 33,
	-87, 35, -87, 36, -87, 9, -87, 5, -87, 25, -87, 24, -87, 2,
	-87, 8, -87, 24, 10, -32, 20, -32, 30, -32, 33, 33, 35, 34,
	36, -91, 9, 35, 5, -32, 25, -32, 24, -32, 2, -32, 8, 36,
	2, 4, 62, 10, 0, -60, 33, -60, 18, -60, 17, -60, 19, -60,
	14, 0, -74, 22, -74, 33, -74, 18, -74, 21, -74, 23, -74, 17,
	-74, 10, 31, 74, 26, 75, 33, 76, 28, 77, 35, 78, 8, 7,
	-77, 22, 16, 33, 17, 21, 18, 10, 20, -50, 30, -50, 25, -50,
	24, -50, 2, -50, 2, 9, 97, 10, 20, -52, 30, -52, 25, -52,
	24, -52, 2, -52, 2, 9, 96, 10, 20, -83, 30, -83, 25, -83,
	24, -83, 2, -83, 2, 2, 88, 10, 20, 51, 30, 52, 25, 53,
	24, 54, 2, -80, 10, 20, -46, 30, -46, 25, -46, 24, -46, 2,
	-46, 10, 20, -49, 30, -49, 25, -49, 24, -49, 2, -49, 10, 20,
	-48, 30, -48, 25, -48, 24, -48, 2, -48, 10, 20, -47, 30, -47,
	25, -47, 24, -47, 2, -47, 24, 10, -86, 20, -86, 30, -86, 33,
	-86, 35, -86, 36, -86, 9, -86, 5, -86, 25, -86, 24, -86, 2,
	-86, 8, -86, 24, 10, -36, 20, -36, 30, -36, 33, -36, 35, -36,
	36, -36, 9, -36, 5, -36, 25, -36, 24, -36, 2, -36, 8, -36,
	24, 10, -35, 20, -35, 30, -35, 33, -35, 35, -35, 36, -35, 9,
	-35, 5, -35, 25, -35, 24, -35, 2, -35, 8, -35, 24, 10, -34,
	20, -34, 30, -34, 33, -34, 35, -34, 36, -34, 9, -34, 5, -34,
	25, -34, 24, -34, 2, -34, 8, -34, 24, 10, -88, 20, -88, 30,
	-88, 33, -88, 35, -88, 36, -88, 9, -88, 5, -88, 25, -88, 24,
	-88, 2, -88, 8, -88, 24, 10, -33, 20, -33, 30, -33, 33, -33,
	35, -33, 36, -33, 9, -33, 5, -33, 25, -33, 24, -33, 2, -33,
	8, -33, 2, 10, 90, 4, 36, -91, 8, 36, 4, 39, 92, 38,
	93, 4, 5, 100, 2, 101, 4, 5, -9, 2, -9, 26, 3, -13,
	10, -13, 26, -13, 33, -13, 27, -13, 28, -13, 35, -13, 14, -13,
	5, -13, 29, -13, 2, -13, 13, -13, 12, -13, 26, 3, -11, 10,
	-11, 26, -11, 33, -11, 27, -11, 28, -11, 35, -11, 14, -11, 5,
	-11, 29, -11, 2, -11, 13, -11, 12, -11, 2, 9, 112, 26, 3,
	-12, 10, -12, 26, -12, 33, -12, 27, -12, 28, -12, 35, -12, 14,
	-12, 5, -12, 29, -12, 2, -12, 13, -12, 12, -12, 26, 3, -14,
	10, -14, 26, -14, 33, -14, 27, -14, 28, -14, 35, -14, 14, -14,
	5, -14, 29, -14, 2, -14, 13, -14, 12, -14, 4, 5, -65, 2,
	-65, 22, 26, -71, 33, -71, 27, -71, 28, -71, 35, -71, 14, 107,
	5, -71, 29, -71, 2, -71, 13, 108, 12, 109, 16, 26, -67, 33,
	-67, 27, -67, 28, -67, 35, -67, 5, -67, 29, -67, 2, -67, 16,
	26, 75, 33, 76, 27, 102, 28, 77, 35, 78, 5, -69, 29, 103,
	2, -69, 8, 7, -79, 22, -79, 33, -79, 21, -79, 2, 7, 113,
	8, 7, -76, 22, 16, 33, 17, 21, 18, 2, 2, 114, 16, 7,
	-29, 0, -29, 22, -29, 33, -29, 18, -29, 21, -29, 23, -29, 17,
	-29, 2, 2, 115, 30, 10, -40, 20, -40, 30, -40, 33, -40, 35,
	-40, 36, -40, 14, -40, 9, -40, 5, -40, 25, -40, 24, -40, 2,
	-40, 8, -40, 13, -40, 12, -40, 30, 10, -41, 20, -41, 30, -41,
	33, -41, 35, -41, 36, -41, 14, -41, 9, -41, 5, -41, 25, -41,
	24, -41, 2, -41, 8, -41, 13, -41, 12, -41, 6, 37, -44, 39,
	-44, 38, -44, 6, 37, -45, 39, -45, 38, -45, 6, 37, -93, 39,
	-93, 38, -93, 6, 37, 116, 39, 92, 38, 93, 4, 10, -95, 33,
	118, 2, 33, 120, 10, 20, -82, 30, -82, 25, -82, 24, -82, 2,
	-82, 14, 10, -84, 20, -84, 30, -84, 5, -84, 25, -84, 24, -84,
	2, -84, 10, 0, -7, 33, -7, 18, -7, 17, -7, 19, -7, 2,
	9, 123, 2, 9, 124, 4, 5, -68, 2, -68, 4, 5, -8, 2,
	-8, 16, 26, -66, 33, -66, 27, -66, 28, -66, 35, -66, 5, -66,
	29, -66, 2, -66, 16, 26, -17, 33, -17, 27, -17, 28, -17, 35,
	-17, 5, -17, 29, -17, 2, -17, 16, 26, -16, 33, -16, 27, -16,
	28, -16, 35, -16, 5, -16, 29, -16, 2, -16, 16, 26, -18, 33,
	-18, 27, -18, 28, -18, 35, -18, 5, -18, 29, -18, 2, -18, 16,
	26, -70, 33, -70, 27, -70, 28, -70, 35, -70, 5, -70, 29, -70,
	2, -70, 16, 26, -10, 33, -10, 27, -10, 28, -10, 35, -10, 5,
	-10, 29, -10, 2, -10, 8, 26, 75, 33, 76, 28, 77, 35, 78,
	14, 0, -27, 22, -27, 33, -27, 18, -27, 21, -27, 23, -27, 17,
	-27, 16, 7, -28, 0, -28, 22, -28, 33, -28, 18, -28, 21, -28,
	23, -28, 17, -28, 16, 7, -30, 0, -30, 22, -30, 33, -30, 18,
	-30, 21, -30, 23, -30, 17, -30, 32, 10, -43, 20, -43, 30, -43,
	33, -43, 35, -43, 36, -43, 14, -43, 9, -43, 5, -43, 25, -43,
	24, -43, 2, -43, 11, -43, 8, -43, 13, -43, 12, -43, 6, 37,
	-92, 39, -92, 38, -92, 2, 10, -94, 2, 10, 125, 2, 10, 126,
	8, 7, -78, 22, -78, 33, -78, 21, -78, 2, 3, 128, 2, 34,
	129, 2, 34, 130, 10, 20, -51, 30, -51, 25, -51, 24, -51, 2,
	-51, 10, 20, -53, 30, -53, 25, -53, 24, -53, 2, -53, 4, 5,
	-64, 2, -64, 2, 10, 132, 2, 10, 133, 2, 10, 134, 4, 5,
	-19, 2, -19, 4, 5, -20, 2, -20, 26, 3, -15, 10, -15, 26,
	-15, 33, -15, 27, -15, 28, -15, 35, -15, 14, -15, 5, -15, 29,
	-15, 2, -15, 13, -15, 12, -15,
}

var _goto = []int32{
	135, 148, 149, 166, 148, 148, 148, 148, 177, 148, 148, 148, 148, 148,
	148, 184, 191, 148, 148, 148, 148, 148, 148, 148, 210, 148, 148, 148,
	148, 148, 148, 223, 148, 148, 148, 242, 148, 148, 148, 148, 261, 148,
	276, 148, 281, 148, 148, 148, 292, 305, 318, 148, 148, 148, 148, 148,
	148, 333, 148, 148, 148, 148, 344, 363, 148, 148, 148, 148, 148, 148,
	148, 378, 383, 148, 148, 148, 148, 148, 148, 148, 148, 388, 148, 393,
	148, 148, 404, 148, 148, 148, 148, 148, 148, 148, 148, 413, 416, 148,
	148, 148, 419, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148,
	430, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148,
	148, 148, 435, 148, 148, 148, 148, 148, 148, 12, 12, 4, 3, 5,
	2, 6, 32, 7, 33, 8, 1, 9, 0, 16, 17, 20, 14, 21,
	13, 22, 41, 23, 42, 24, 18, 25, 15, 26, 16, 27, 10, 36,
	11, 5, 12, 4, 13, 34, 14, 35, 15, 6, 12, 4, 3, 5,
	2, 28, 6, 36, 11, 5, 12, 4, 46, 18, 47, 32, 50, 37,
	25, 38, 24, 39, 19, 40, 20, 41, 23, 42, 21, 43, 48, 44,
	12, 17, 20, 14, 21, 13, 47, 18, 25, 15, 26, 16, 27, 18,
	47, 32, 50, 37, 25, 38, 24, 39, 19, 50, 20, 41, 23, 42,
	21, 43, 48, 44, 18, 47, 32, 50, 37, 25, 38, 24, 39, 19,
	70, 20, 41, 23, 42, 21, 43, 48, 44, 14, 27, 55, 45, 56,
	46, 57, 28, 58, 31, 59, 30, 60, 29, 61, 4, 22, 68, 49,
	69, 10, 50, 37, 25, 38, 24, 39, 23, 42, 21, 64, 12, 37,
	73, 9, 79, 6, 80, 8, 81, 7, 82, 38, 83, 12, 17, 20,
	14, 84, 43, 85, 44, 86, 18, 25, 16, 27, 14, 27, 55, 45,
	87, 46, 57, 28, 58, 31, 59, 30, 60, 29, 61, 10, 27, 98,
	28, 58, 31, 59, 30, 60, 29, 61, 18, 47, 32, 50, 37, 25,
	38, 24, 39, 19, 89, 20, 41, 23, 42, 21, 43, 48, 44, 14,
	50, 37, 25, 38, 24, 39, 20, 99, 23, 42, 21, 43, 48, 44,
	4, 50, 37, 25, 91, 4, 26, 94, 51, 95, 4, 10, 110, 40,
	111, 10, 9, 79, 11, 104, 39, 105, 8, 81, 7, 106, 8, 17,
	20, 14, 121, 18, 25, 16, 27, 2, 26, 117, 2, 52, 119, 10,
	9, 79, 6, 127, 8, 81, 7, 82, 38, 83, 4, 9, 79, 8,
	122, 4, 9, 79, 8, 131,
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
			latok, ok := p._lasym.(Token)
			if !ok {
				latok = p._lasym.(Error).Token
			}
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
		return p.on_parser_prod__empty(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 10:
		return p.on_parser_term_card(
			_cast[*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[_i0.ParserTermType](p._stack.Peek(0).Sym),
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
		return p.on_parser_term__token(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 14:
		return p.on_parser_term__list(
			_cast[*_i0.ParserTerm](p._stack.Peek(0).Sym),
		)
	case 15:
		return p.on_parser_list(
			_cast[_i1.Token](p._stack.Peek(5).Sym),
			_cast[_i1.Token](p._stack.Peek(4).Sym),
			_cast[*_i0.ParserTerm](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[*_i0.ParserTerm](p._stack.Peek(1).Sym),
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
		return p.on_parser_card(
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
		return p.on_parser_qualif(
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[_i1.Token](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 21:
		return p.on_lexer_section(
			_cast[_i1.Token](p._stack.Peek(1).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 22:
		return p.on_lexer_statement(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 23:
		return p.on_lexer_statement(
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
		return p.on_lexer_rule(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 27:
		return p.on_mode(
			_cast[_i1.Token](p._stack.Peek(4).Sym),
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 28:
		return p.on_token_rule(
			_cast[_i1.Token](p._stack.Peek(4).Sym),
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(2).Sym),
			_cast[[]_i0.Action](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 29:
		return p.on_frag_rule(
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(2).Sym),
			_cast[[]_i0.Action](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 30:
		return p.on_macro_rule(
			_cast[_i1.Token](p._stack.Peek(4).Sym),
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 31:
		return p.on_lexer_expr(
			_cast[[]*_i0.LexerFactor](p._stack.Peek(0).Sym),
		)
	case 32:
		return p.on_lexer_factor(
			_cast[[]*_i0.LexerTermCard](p._stack.Peek(0).Sym),
		)
	case 33:
		return p.on_lexer_term_card(
			_cast[_i0.LexerTerm](p._stack.Peek(1).Sym),
			_cast[_i0.Card](p._stack.Peek(0).Sym),
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
		return p.on_lexer_card(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 37:
		return p.on_lexer_term__tok(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 38:
		return p.on_lexer_term__tok(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 39:
		return p.on_lexer_term__char_class_expr(
			_cast[_i0.CharClassExpr](p._stack.Peek(0).Sym),
		)
	case 40:
		return p.on_lexer_term__expr(
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 41:
		return p.on_char_class_expr__binary(
			_cast[_i0.CharClassExpr](p._stack.Peek(2).Sym),
			_cast[_i1.Token](p._stack.Peek(1).Sym),
			_cast[_i0.CharClassExpr](p._stack.Peek(0).Sym),
		)
	case 42:
		return p.on_char_class_expr__char_class(
			_cast[*_i0.CharClass](p._stack.Peek(0).Sym),
		)
	case 43:
		return p.on_char_class(
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[[]_i1.Token](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 44:
		return p.on_char_class_item(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 45:
		return p.on_char_class_item(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
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
		return p.on_action(
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 50:
		return p.on_action_discard(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 51:
		return p.on_action_push_mode(
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[_i1.Token](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 52:
		return p.on_action_pop_mode(
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 53:
		return p.on_action_emit(
			_cast[_i1.Token](p._stack.Peek(3).Sym),
			_cast[_i1.Token](p._stack.Peek(2).Sym),
			_cast[_i1.Token](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 54: // ZeroOrMore
		return _cast[[][]_i0.Statement](p._stack.Peek(0).Sym)
	case 55: // ZeroOrMore
		{
			var zero [][]_i0.Statement
			return zero
		}
	case 56: // OneOrMore
		return append(
			_cast[[][]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 57: // OneOrMore
		return [][]_i0.Statement{
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 58: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 59: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 60: // OneOrMore
		return append(
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 61: // OneOrMore
		return []_i0.Statement{
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 62: // ZeroOrOne
		return _cast[_i1.Token](p._stack.Peek(0).Sym)
	case 63: // ZeroOrOne
		{
			var zero _i1.Token
			return zero
		}
	case 64: // List
		return append(
			_cast[[]*_i0.ParserProd](p._stack.Peek(2).Sym),
			_cast[*_i0.ParserProd](p._stack.Peek(0).Sym),
		)
	case 65: // List
		return []*_i0.ParserProd{
			_cast[*_i0.ParserProd](p._stack.Peek(0).Sym),
		}
	case 66: // OneOrMore
		return append(
			_cast[[]*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[*_i0.ParserTerm](p._stack.Peek(0).Sym),
		)
	case 67: // OneOrMore
		return []*_i0.ParserTerm{
			_cast[*_i0.ParserTerm](p._stack.Peek(0).Sym),
		}
	case 68: // ZeroOrOne
		return _cast[*_i0.ProdQualifier](p._stack.Peek(0).Sym)
	case 69: // ZeroOrOne
		{
			var zero *_i0.ProdQualifier
			return zero
		}
	case 70: // ZeroOrOne
		return _cast[_i0.ParserTermType](p._stack.Peek(0).Sym)
	case 71: // ZeroOrOne
		{
			var zero _i0.ParserTermType
			return zero
		}
	case 72: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 73: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 74: // OneOrMore
		return append(
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 75: // OneOrMore
		return []_i0.Statement{
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 76: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 77: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 78: // OneOrMore
		return append(
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 79: // OneOrMore
		return []_i0.Statement{
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 80: // ZeroOrMore
		return _cast[[]_i0.Action](p._stack.Peek(0).Sym)
	case 81: // ZeroOrMore
		{
			var zero []_i0.Action
			return zero
		}
	case 82: // OneOrMore
		return append(
			_cast[[]_i0.Action](p._stack.Peek(1).Sym),
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 83: // OneOrMore
		return []_i0.Action{
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		}
	case 84: // List
		return append(
			_cast[[]*_i0.LexerFactor](p._stack.Peek(2).Sym),
			_cast[*_i0.LexerFactor](p._stack.Peek(0).Sym),
		)
	case 85: // List
		return []*_i0.LexerFactor{
			_cast[*_i0.LexerFactor](p._stack.Peek(0).Sym),
		}
	case 86: // OneOrMore
		return append(
			_cast[[]*_i0.LexerTermCard](p._stack.Peek(1).Sym),
			_cast[*_i0.LexerTermCard](p._stack.Peek(0).Sym),
		)
	case 87: // OneOrMore
		return []*_i0.LexerTermCard{
			_cast[*_i0.LexerTermCard](p._stack.Peek(0).Sym),
		}
	case 88: // ZeroOrOne
		return _cast[_i0.Card](p._stack.Peek(0).Sym)
	case 89: // ZeroOrOne
		{
			var zero _i0.Card
			return zero
		}
	case 90: // ZeroOrOne
		return _cast[_i1.Token](p._stack.Peek(0).Sym)
	case 91: // ZeroOrOne
		{
			var zero _i1.Token
			return zero
		}
	case 92: // OneOrMore
		return append(
			_cast[[]_i1.Token](p._stack.Peek(1).Sym),
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		)
	case 93: // OneOrMore
		return []_i1.Token{
			_cast[_i1.Token](p._stack.Peek(0).Sym),
		}
	case 94: // ZeroOrOne
		return _cast[_i1.Token](p._stack.Peek(0).Sym)
	case 95: // ZeroOrOne
		{
			var zero _i1.Token
			return zero
		}
	default:
		panic("unreachable")
	}
}
