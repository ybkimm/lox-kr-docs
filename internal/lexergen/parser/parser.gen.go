package parser

import (
  _i0 "errors"
  _i1 "github.com/dcaiafa/lox/internal/lexergen/ast"
  _i2 "github.com/dcaiafa/lox/internal/util/baselexer"
)

var _LHS = []int32 {
	0, 1, 2, 2, 3, 4, 5, 6, 7, 8, 8, 8, 8, 9, 
10, 10, 10, 11, 11, 12, 13, 13, 14, 14, 14, 15, 16, 17, 
18, 19, 20, 21, 22, 22, 22, 23, 23, 23, 23, 24, 25, 25, 
26, 26, 26, 27, 28, 29, 30, 30, 31, 31, 32, 32, 33, 33, 
34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 
41, 41, 42, 42, 43, 43, 44, 44, 45, 45, 46, 46, 47, 47, 
48, 48, 49, 49, 
}

var _TermCounts = []int32 {
	1, 1, 1, 1, 2, 1, 5, 2, 2, 1, 1, 1, 1, 6, 
1, 1, 1, 4, 4, 2, 1, 1, 1, 1, 1, 5, 5, 4, 
5, 1, 1, 2, 1, 1, 1, 1, 1, 1, 3, 4, 1, 1, 
1, 1, 1, 1, 4, 1, 1, 0, 2, 1, 1, 0, 2, 1, 
1, 0, 3, 1, 2, 1, 1, 0, 1, 0, 1, 0, 2, 1, 
1, 0, 2, 1, 1, 0, 2, 1, 3, 1, 2, 1, 1, 0, 
1, 0, 2, 1, 	
}

var _Actions = []int32 {
	124, 131, 146, 157, 164, 171, 178, 181, 188, 191, 194, 197, 208, 219, 
226, 237, 248, 251, 254, 257, 274, 289, 304, 311, 326, 343, 358, 375, 
382, 385, 237, 388, 401, 430, 237, 459, 462, 465, 494, 503, 516, 545, 
568, 591, 594, 605, 620, 629, 494, 638, 647, 656, 659, 668, 671, 680, 
689, 698, 237, 237, 707, 730, 753, 776, 799, 822, 845, 848, 853, 858, 
885, 912, 915, 942, 969, 974, 997, 1014, 1031, 1040, 1043, 1052, 1055, 1072, 
1075, 1104, 1111, 1118, 1125, 1132, 1135, 1144, 620, 1157, 1168, 1171, 1174, 1179, 
1184, 1201, 1218, 1235, 1252, 1269, 620, 1286, 1301, 1318, 1335, 1364, 1371, 1374, 
1383, 1386, 1389, 1392, 1401, 620, 1406, 1409, 1412, 1415, 1420, 1425, 6, 0, 
-49, 17, 1, 16, 2, 14, 0, -67, 21, 15, 30, 16, 17, -67, 
20, 17, 22, 18, 16, -67, 10, 0, -53, 30, -57, 17, -53, 16, 
-53, 18, 9, 6, 0, -3, 17, -3, 16, -3, 6, 0, -2, 17, 
-2, 16, -2, 6, 0, -51, 17, -51, 16, -51, 2, 0, -1, 6, 
0, -48, 17, 1, 16, 2, 2, 0, 2147483647, 2, 30, -56, 2, 30, 
28, 10, 0, -5, 30, -5, 17, -5, 16, -5, 18, -5, 10, 0, 
-55, 30, -55, 17, -55, 16, -55, 18, -55, 6, 0, -4, 17, -4, 
16, -4, 10, 0, -52, 30, -57, 17, -52, 16, -52, 18, 9, 10, 
30, 32, 32, 33, 33, -85, 9, 34, 8, 35, 2, 4, 30, 2, 
30, 43, 2, 30, 29, 16, 7, -23, 0, -23, 21, -23, 30, -23, 
17, -23, 20, -23, 22, -23, 16, -23, 14, 0, -21, 21, -21, 30, 
-21, 17, -21, 20, -21, 22, -21, 16, -21, 14, 0, -69, 21, -69, 
30, -69, 17, -69, 20, -69, 22, -69, 16, -69, 6, 0, -19, 17, 
-19, 16, -19, 14, 0, -66, 21, 15, 30, 16, 17, -66, 20, 17, 
22, 18, 16, -66, 16, 7, -24, 0, -24, 21, -24, 30, -24, 17, 
-24, 20, -24, 22, -24, 16, -24, 14, 0, -20, 21, -20, 30, -20, 
17, -20, 20, -20, 22, -20, 16, -20, 16, 7, -22, 0, -22, 21, 
-22, 30, -22, 17, -22, 20, -22, 22, -22, 16, -22, 6, 0, -50, 
17, -50, 16, -50, 2, 4, 46, 2, 6, 47, 12, 10, -29, 19, 
-29, 5, 59, 24, -29, 23, -29, 2, -29, 28, 10, -36, 19, -36, 
30, -36, 32, -36, 33, -36, 13, -36, 9, -36, 5, -36, 24, -36, 
23, -36, 2, -36, 8, -36, 12, -36, 11, -36, 28, 10, -35, 19, 
-35, 30, -35, 32, -35, 33, -35, 13, -35, 9, -35, 5, -35, 24, 
-35, 23, -35, 2, -35, 8, -35, 12, -35, 11, -35, 2, 33, -84, 
2, 33, 67, 28, 10, -37, 19, -37, 30, -37, 32, -37, 33, -37, 
13, -37, 9, -37, 5, -37, 24, -37, 23, -37, 2, -37, 8, -37, 
12, -37, 11, -37, 8, 19, 49, 24, 50, 23, 51, 2, -75, 12, 
10, -79, 19, -79, 5, -79, 24, -79, 23, -79, 2, -79, 28, 10, 
-83, 19, -83, 30, -83, 32, -83, 33, -83, 13, 61, 9, -83, 5, 
-83, 24, -83, 23, -83, 2, -83, 8, -83, 12, 62, 11, 63, 22, 
10, -81, 19, -81, 30, -81, 32, -81, 33, -81, 9, -81, 5, -81, 
24, -81, 23, -81, 2, -81, 8, -81, 22, 10, -30, 19, -30, 30, 
32, 32, 33, 33, -85, 9, 34, 5, -30, 24, -30, 23, -30, 2, 
-30, 8, 35, 2, 4, 58, 10, 0, -54, 30, -54, 17, -54, 16, 
-54, 18, -54, 14, 0, -68, 21, -68, 30, -68, 17, -68, 20, -68, 
22, -68, 16, -68, 8, 25, 69, 30, 70, 27, 71, 32, 72, 8, 
7, -71, 21, 15, 30, 16, 20, 17, 8, 19, -45, 24, -45, 23, 
-45, 2, -45, 8, 19, -47, 24, -47, 23, -47, 2, -47, 2, 9, 
89, 8, 19, -77, 24, -77, 23, -77, 2, -77, 2, 2, 82, 8, 
19, 49, 24, 50, 23, 51, 2, -74, 8, 19, -42, 24, -42, 23, 
-42, 2, -42, 8, 19, -44, 24, -44, 23, -44, 2, -44, 8, 19, 
-43, 24, -43, 23, -43, 2, -43, 22, 10, -80, 19, -80, 30, -80, 
32, -80, 33, -80, 9, -80, 5, -80, 24, -80, 23, -80, 2, -80, 
8, -80, 22, 10, -34, 19, -34, 30, -34, 32, -34, 33, -34, 9, 
-34, 5, -34, 24, -34, 23, -34, 2, -34, 8, -34, 22, 10, -33, 
19, -33, 30, -33, 32, -33, 33, -33, 9, -33, 5, -33, 24, -33, 
23, -33, 2, -33, 8, -33, 22, 10, -32, 19, -32, 30, -32, 32, 
-32, 33, -32, 9, -32, 5, -32, 24, -32, 23, -32, 2, -32, 8, 
-32, 22, 10, -82, 19, -82, 30, -82, 32, -82, 33, -82, 9, -82, 
5, -82, 24, -82, 23, -82, 2, -82, 8, -82, 22, 10, -31, 19, 
-31, 30, -31, 32, -31, 33, -31, 9, -31, 5, -31, 24, -31, 23, 
-31, 2, -31, 8, -31, 2, 10, 84, 4, 36, 85, 35, 86, 4, 
5, 92, 2, 93, 26, 3, -11, 10, -11, 25, -11, 30, -11, 26, 
-11, 27, -11, 32, -11, 13, -11, 5, -11, 28, -11, 2, -11, 12, 
-11, 11, -11, 26, 3, -9, 10, -9, 25, -9, 30, -9, 26, -9, 
27, -9, 32, -9, 13, -9, 5, -9, 28, -9, 2, -9, 12, -9, 
11, -9, 2, 9, 104, 26, 3, -10, 10, -10, 25, -10, 30, -10, 
26, -10, 27, -10, 32, -10, 13, -10, 5, -10, 28, -10, 2, -10, 
12, -10, 11, -10, 26, 3, -12, 10, -12, 25, -12, 30, -12, 26, 
-12, 27, -12, 32, -12, 13, -12, 5, -12, 28, -12, 2, -12, 12, 
-12, 11, -12, 4, 5, -59, 2, -59, 22, 25, -65, 30, -65, 26, 
-65, 27, -65, 32, -65, 13, 99, 5, -65, 28, -65, 2, -65, 12, 
100, 11, 101, 16, 25, -61, 30, -61, 26, -61, 27, -61, 32, -61, 
5, -61, 28, -61, 2, -61, 16, 25, 69, 30, 70, 26, 94, 27, 
71, 32, 72, 5, -63, 28, 95, 2, -63, 8, 7, -73, 21, -73, 
30, -73, 20, -73, 2, 7, 105, 8, 7, -70, 21, 15, 30, 16, 
20, 17, 2, 2, 106, 16, 7, -27, 0, -27, 21, -27, 30, -27, 
17, -27, 20, -27, 22, -27, 16, -27, 2, 2, 107, 28, 10, -38, 
19, -38, 30, -38, 32, -38, 33, -38, 13, -38, 9, -38, 5, -38, 
24, -38, 23, -38, 2, -38, 8, -38, 12, -38, 11, -38, 6, 34, 
-40, 36, -40, 35, -40, 6, 34, -41, 36, -41, 35, -41, 6, 34, 
-87, 36, -87, 35, -87, 6, 34, 108, 36, 85, 35, 86, 2, 30, 
110, 8, 19, -76, 24, -76, 23, -76, 2, -76, 12, 10, -78, 19, 
-78, 5, -78, 24, -78, 23, -78, 2, -78, 10, 0, -6, 30, -6, 
17, -6, 16, -6, 18, -6, 2, 9, 113, 2, 9, 114, 4, 5, 
-62, 2, -62, 4, 5, -7, 2, -7, 16, 25, -60, 30, -60, 26, 
-60, 27, -60, 32, -60, 5, -60, 28, -60, 2, -60, 16, 25, -15, 
30, -15, 26, -15, 27, -15, 32, -15, 5, -15, 28, -15, 2, -15, 
16, 25, -14, 30, -14, 26, -14, 27, -14, 32, -14, 5, -14, 28, 
-14, 2, -14, 16, 25, -16, 30, -16, 26, -16, 27, -16, 32, -16, 
5, -16, 28, -16, 2, -16, 16, 25, -64, 30, -64, 26, -64, 27, 
-64, 32, -64, 5, -64, 28, -64, 2, -64, 16, 25, -8, 30, -8, 
26, -8, 27, -8, 32, -8, 5, -8, 28, -8, 2, -8, 14, 0, 
-25, 21, -25, 30, -25, 17, -25, 20, -25, 22, -25, 16, -25, 16, 
7, -26, 0, -26, 21, -26, 30, -26, 17, -26, 20, -26, 22, -26, 
16, -26, 16, 7, -28, 0, -28, 21, -28, 30, -28, 17, -28, 20, 
-28, 22, -28, 16, -28, 28, 10, -39, 19, -39, 30, -39, 32, -39, 
33, -39, 13, -39, 9, -39, 5, -39, 24, -39, 23, -39, 2, -39, 
8, -39, 12, -39, 11, -39, 6, 34, -86, 36, -86, 35, -86, 2, 
10, 115, 8, 7, -72, 21, -72, 30, -72, 20, -72, 2, 3, 117, 
2, 31, 118, 2, 31, 119, 8, 19, -46, 24, -46, 23, -46, 2, 
-46, 4, 5, -58, 2, -58, 2, 10, 121, 2, 10, 122, 2, 10, 
123, 4, 5, -17, 2, -17, 4, 5, -18, 2, -18, 26, 3, -13, 
10, -13, 25, -13, 30, -13, 26, -13, 27, -13, 32, -13, 13, -13, 
5, -13, 28, -13, 2, -13, 12, -13, 11, -13, 
}

var _Goto = []int32 {
	124, 137, 154, 165, 165, 165, 165, 166, 165, 165, 165, 165, 165, 165, 
173, 180, 165, 165, 165, 165, 165, 165, 165, 197, 165, 165, 165, 165, 
165, 165, 210, 165, 165, 165, 227, 165, 165, 165, 244, 165, 257, 165, 
262, 165, 165, 165, 271, 284, 297, 165, 165, 165, 165, 165, 310, 165, 
165, 165, 319, 336, 165, 165, 165, 165, 165, 165, 165, 349, 165, 165, 
165, 165, 165, 165, 165, 354, 165, 359, 165, 165, 370, 165, 165, 165, 
165, 165, 165, 165, 379, 165, 165, 165, 382, 165, 165, 165, 165, 165, 
165, 165, 165, 165, 165, 165, 393, 165, 165, 165, 165, 165, 165, 165, 
165, 165, 165, 165, 165, 398, 165, 165, 165, 165, 165, 165, 12, 12, 
3, 3, 4, 2, 5, 30, 6, 31, 7, 1, 8, 16, 17, 19, 
14, 20, 13, 21, 39, 22, 40, 23, 18, 24, 15, 25, 16, 26, 
10, 34, 10, 5, 11, 4, 12, 32, 13, 33, 14, 0, 6, 12, 
3, 3, 4, 2, 27, 6, 34, 10, 5, 11, 4, 44, 16, 45, 
31, 48, 36, 24, 37, 19, 38, 20, 39, 23, 40, 21, 41, 46, 
42, 12, 17, 19, 14, 20, 13, 45, 18, 24, 15, 25, 16, 26, 
16, 45, 31, 48, 36, 24, 37, 19, 48, 20, 39, 23, 40, 21, 
41, 46, 42, 16, 45, 31, 48, 36, 24, 37, 19, 66, 20, 39, 
23, 40, 21, 41, 46, 42, 12, 26, 52, 43, 53, 44, 54, 27, 
55, 29, 56, 28, 57, 4, 22, 64, 47, 65, 8, 48, 36, 24, 
37, 23, 40, 21, 60, 12, 35, 68, 9, 73, 6, 74, 8, 75, 
7, 76, 36, 77, 12, 17, 19, 14, 78, 41, 79, 42, 80, 18, 
24, 16, 26, 12, 26, 52, 43, 81, 44, 54, 27, 55, 29, 56, 
28, 57, 8, 26, 90, 27, 55, 29, 56, 28, 57, 16, 45, 31, 
48, 36, 24, 37, 19, 83, 20, 39, 23, 40, 21, 41, 46, 42, 
12, 48, 36, 24, 37, 20, 91, 23, 40, 21, 41, 46, 42, 4, 
25, 87, 49, 88, 4, 10, 102, 38, 103, 10, 9, 73, 11, 96, 
37, 97, 8, 75, 7, 98, 8, 17, 19, 14, 111, 18, 24, 16, 
26, 2, 25, 109, 10, 9, 73, 6, 116, 8, 75, 7, 76, 36, 
77, 4, 9, 73, 8, 112, 4, 9, 73, 8, 120, 
}

type _Bounds struct {
	Begin Token
	End   Token
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
	for ; i < end; i+=2 {
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
	_lex   _Lexer
	_state _Stack[int32]
	_sym   _Stack[any]
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
			if termCount > 0 {
				bounds := _Bounds{
					Begin: p._bounds.Peek(int(termCount-1)).Begin,
					End: p._bounds.Peek(0).End,
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
				return p.on_lexer_term__char_class(
				  _cast[*_i1.LexerTermCharClass](p._sym.Peek(0)),
		    )
			case 38:
				return p.on_lexer_term__expr(
				  _cast[_i2.Token](p._sym.Peek(2)),
				  _cast[*_i1.LexerExpr](p._sym.Peek(1)),
				  _cast[_i2.Token](p._sym.Peek(0)),
		    )
			case 39:
				return p.on_char_class(
				  _cast[_i2.Token](p._sym.Peek(3)),
				  _cast[_i2.Token](p._sym.Peek(2)),
				  _cast[[]_i2.Token](p._sym.Peek(1)),
				  _cast[_i2.Token](p._sym.Peek(0)),
		    )
			case 40:
				return p.on_char_class_item(
				  _cast[_i2.Token](p._sym.Peek(0)),
		    )
			case 41:
				return p.on_char_class_item(
				  _cast[_i2.Token](p._sym.Peek(0)),
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
				return p.on_action_discard(
				  _cast[_i2.Token](p._sym.Peek(0)),
		    )
			case 46:
				return p.on_action_push_mode(
				  _cast[_i2.Token](p._sym.Peek(3)),
				  _cast[_i2.Token](p._sym.Peek(2)),
				  _cast[_i2.Token](p._sym.Peek(1)),
				  _cast[_i2.Token](p._sym.Peek(0)),
		    )
			case 47:
				return p.on_action_pop_mode(
				  _cast[_i2.Token](p._sym.Peek(0)),
		    )
  case 48:  // ZeroOrMore
			return _cast[[][]_i1.Statement](p._sym.Peek(0))
  case 49:  // ZeroOrMore
			{
				var zero [][]_i1.Statement
				return zero
			}
	case 50:  // OneOrMore
			return append(
				_cast[[][]_i1.Statement](p._sym.Peek(1)),
				_cast[[]_i1.Statement](p._sym.Peek(0)),
			)
	case 51:  // OneOrMore
		  return [][]_i1.Statement{
				_cast[[]_i1.Statement](p._sym.Peek(0)),
			}
  case 52:  // ZeroOrMore
			return _cast[[]_i1.Statement](p._sym.Peek(0))
  case 53:  // ZeroOrMore
			{
				var zero []_i1.Statement
				return zero
			}
	case 54:  // OneOrMore
			return append(
				_cast[[]_i1.Statement](p._sym.Peek(1)),
				_cast[_i1.Statement](p._sym.Peek(0)),
			)
	case 55:  // OneOrMore
		  return []_i1.Statement{
				_cast[_i1.Statement](p._sym.Peek(0)),
			}
  case 56:  // ZeroOrOne
			return _cast[_i2.Token](p._sym.Peek(0))
  case 57:  // ZeroOrOne
			{
				var zero _i2.Token
				return zero
			}
	case 58:  // List
			return append(
				_cast[[]*_i1.ParserProd](p._sym.Peek(2)),
				_cast[*_i1.ParserProd](p._sym.Peek(0)),
			)
	case 59:  // List
		  return []*_i1.ParserProd{
				_cast[*_i1.ParserProd](p._sym.Peek(0)),
			}
	case 60:  // OneOrMore
			return append(
				_cast[[]*_i1.ParserTerm](p._sym.Peek(1)),
				_cast[*_i1.ParserTerm](p._sym.Peek(0)),
			)
	case 61:  // OneOrMore
		  return []*_i1.ParserTerm{
				_cast[*_i1.ParserTerm](p._sym.Peek(0)),
			}
  case 62:  // ZeroOrOne
			return _cast[*_i1.ProdQualifier](p._sym.Peek(0))
  case 63:  // ZeroOrOne
			{
				var zero *_i1.ProdQualifier
				return zero
			}
  case 64:  // ZeroOrOne
			return _cast[_i1.ParserTermType](p._sym.Peek(0))
  case 65:  // ZeroOrOne
			{
				var zero _i1.ParserTermType
				return zero
			}
  case 66:  // ZeroOrMore
			return _cast[[]_i1.Statement](p._sym.Peek(0))
  case 67:  // ZeroOrMore
			{
				var zero []_i1.Statement
				return zero
			}
	case 68:  // OneOrMore
			return append(
				_cast[[]_i1.Statement](p._sym.Peek(1)),
				_cast[_i1.Statement](p._sym.Peek(0)),
			)
	case 69:  // OneOrMore
		  return []_i1.Statement{
				_cast[_i1.Statement](p._sym.Peek(0)),
			}
  case 70:  // ZeroOrMore
			return _cast[[]_i1.Statement](p._sym.Peek(0))
  case 71:  // ZeroOrMore
			{
				var zero []_i1.Statement
				return zero
			}
	case 72:  // OneOrMore
			return append(
				_cast[[]_i1.Statement](p._sym.Peek(1)),
				_cast[_i1.Statement](p._sym.Peek(0)),
			)
	case 73:  // OneOrMore
		  return []_i1.Statement{
				_cast[_i1.Statement](p._sym.Peek(0)),
			}
  case 74:  // ZeroOrMore
			return _cast[[]_i1.Action](p._sym.Peek(0))
  case 75:  // ZeroOrMore
			{
				var zero []_i1.Action
				return zero
			}
	case 76:  // OneOrMore
			return append(
				_cast[[]_i1.Action](p._sym.Peek(1)),
				_cast[_i1.Action](p._sym.Peek(0)),
			)
	case 77:  // OneOrMore
		  return []_i1.Action{
				_cast[_i1.Action](p._sym.Peek(0)),
			}
	case 78:  // List
			return append(
				_cast[[]*_i1.LexerFactor](p._sym.Peek(2)),
				_cast[*_i1.LexerFactor](p._sym.Peek(0)),
			)
	case 79:  // List
		  return []*_i1.LexerFactor{
				_cast[*_i1.LexerFactor](p._sym.Peek(0)),
			}
	case 80:  // OneOrMore
			return append(
				_cast[[]*_i1.LexerTermCard](p._sym.Peek(1)),
				_cast[*_i1.LexerTermCard](p._sym.Peek(0)),
			)
	case 81:  // OneOrMore
		  return []*_i1.LexerTermCard{
				_cast[*_i1.LexerTermCard](p._sym.Peek(0)),
			}
  case 82:  // ZeroOrOne
			return _cast[_i1.Card](p._sym.Peek(0))
  case 83:  // ZeroOrOne
			{
				var zero _i1.Card
				return zero
			}
  case 84:  // ZeroOrOne
			return _cast[_i2.Token](p._sym.Peek(0))
  case 85:  // ZeroOrOne
			{
				var zero _i2.Token
				return zero
			}
	case 86:  // OneOrMore
			return append(
				_cast[[]_i2.Token](p._sym.Peek(1)),
				_cast[_i2.Token](p._sym.Peek(0)),
			)
	case 87:  // OneOrMore
		  return []_i2.Token{
				_cast[_i2.Token](p._sym.Peek(0)),
			}
	default:
		panic("unreachable")
	}
}
