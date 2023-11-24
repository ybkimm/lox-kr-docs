package parser

import (
  _i0 "errors"
  _i1 "github.com/dcaiafa/lox/internal/ast"
  _i2 "github.com/dcaiafa/lox/internal/base/baselexer"
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
-49, 18, 1, 17, 2, 14, 0, -67, 22, 15, 31, 16, 18, -67, 
21, 17, 23, 18, 17, -67, 10, 0, -53, 31, -57, 18, -53, 17, 
-53, 19, 9, 6, 0, -3, 18, -3, 17, -3, 6, 0, -2, 18, 
-2, 17, -2, 6, 0, -51, 18, -51, 17, -51, 2, 0, -1, 6, 
0, -48, 18, 1, 17, 2, 2, 0, 2147483647, 2, 31, -56, 2, 31, 
28, 10, 0, -5, 31, -5, 18, -5, 17, -5, 19, -5, 10, 0, 
-55, 31, -55, 18, -55, 17, -55, 19, -55, 6, 0, -4, 18, -4, 
17, -4, 10, 0, -52, 31, -57, 18, -52, 17, -52, 19, 9, 10, 
31, 32, 33, 33, 34, -85, 9, 34, 8, 35, 2, 4, 30, 2, 
31, 43, 2, 31, 29, 16, 7, -23, 0, -23, 22, -23, 31, -23, 
18, -23, 21, -23, 23, -23, 17, -23, 14, 0, -21, 22, -21, 31, 
-21, 18, -21, 21, -21, 23, -21, 17, -21, 14, 0, -69, 22, -69, 
31, -69, 18, -69, 21, -69, 23, -69, 17, -69, 6, 0, -19, 18, 
-19, 17, -19, 14, 0, -66, 22, 15, 31, 16, 18, -66, 21, 17, 
23, 18, 17, -66, 16, 7, -24, 0, -24, 22, -24, 31, -24, 18, 
-24, 21, -24, 23, -24, 17, -24, 14, 0, -20, 22, -20, 31, -20, 
18, -20, 21, -20, 23, -20, 17, -20, 16, 7, -22, 0, -22, 22, 
-22, 31, -22, 18, -22, 21, -22, 23, -22, 17, -22, 6, 0, -50, 
18, -50, 17, -50, 2, 4, 46, 2, 6, 47, 12, 10, -29, 20, 
-29, 5, 59, 25, -29, 24, -29, 2, -29, 28, 10, -36, 20, -36, 
31, -36, 33, -36, 34, -36, 14, -36, 9, -36, 5, -36, 25, -36, 
24, -36, 2, -36, 8, -36, 13, -36, 12, -36, 28, 10, -35, 20, 
-35, 31, -35, 33, -35, 34, -35, 14, -35, 9, -35, 5, -35, 25, 
-35, 24, -35, 2, -35, 8, -35, 13, -35, 12, -35, 2, 34, -84, 
2, 34, 67, 28, 10, -37, 20, -37, 31, -37, 33, -37, 34, -37, 
14, -37, 9, -37, 5, -37, 25, -37, 24, -37, 2, -37, 8, -37, 
13, -37, 12, -37, 8, 20, 49, 25, 50, 24, 51, 2, -75, 12, 
10, -79, 20, -79, 5, -79, 25, -79, 24, -79, 2, -79, 28, 10, 
-83, 20, -83, 31, -83, 33, -83, 34, -83, 14, 61, 9, -83, 5, 
-83, 25, -83, 24, -83, 2, -83, 8, -83, 13, 62, 12, 63, 22, 
10, -81, 20, -81, 31, -81, 33, -81, 34, -81, 9, -81, 5, -81, 
25, -81, 24, -81, 2, -81, 8, -81, 22, 10, -30, 20, -30, 31, 
32, 33, 33, 34, -85, 9, 34, 5, -30, 25, -30, 24, -30, 2, 
-30, 8, 35, 2, 4, 58, 10, 0, -54, 31, -54, 18, -54, 17, 
-54, 19, -54, 14, 0, -68, 22, -68, 31, -68, 18, -68, 21, -68, 
23, -68, 17, -68, 8, 26, 69, 31, 70, 28, 71, 33, 72, 8, 
7, -71, 22, 15, 31, 16, 21, 17, 8, 20, -45, 25, -45, 24, 
-45, 2, -45, 8, 20, -47, 25, -47, 24, -47, 2, -47, 2, 9, 
89, 8, 20, -77, 25, -77, 24, -77, 2, -77, 2, 2, 82, 8, 
20, 49, 25, 50, 24, 51, 2, -74, 8, 20, -42, 25, -42, 24, 
-42, 2, -42, 8, 20, -44, 25, -44, 24, -44, 2, -44, 8, 20, 
-43, 25, -43, 24, -43, 2, -43, 22, 10, -80, 20, -80, 31, -80, 
33, -80, 34, -80, 9, -80, 5, -80, 25, -80, 24, -80, 2, -80, 
8, -80, 22, 10, -34, 20, -34, 31, -34, 33, -34, 34, -34, 9, 
-34, 5, -34, 25, -34, 24, -34, 2, -34, 8, -34, 22, 10, -33, 
20, -33, 31, -33, 33, -33, 34, -33, 9, -33, 5, -33, 25, -33, 
24, -33, 2, -33, 8, -33, 22, 10, -32, 20, -32, 31, -32, 33, 
-32, 34, -32, 9, -32, 5, -32, 25, -32, 24, -32, 2, -32, 8, 
-32, 22, 10, -82, 20, -82, 31, -82, 33, -82, 34, -82, 9, -82, 
5, -82, 25, -82, 24, -82, 2, -82, 8, -82, 22, 10, -31, 20, 
-31, 31, -31, 33, -31, 34, -31, 9, -31, 5, -31, 25, -31, 24, 
-31, 2, -31, 8, -31, 2, 10, 84, 4, 37, 85, 36, 86, 4, 
5, 92, 2, 93, 26, 3, -11, 10, -11, 26, -11, 31, -11, 27, 
-11, 28, -11, 33, -11, 14, -11, 5, -11, 29, -11, 2, -11, 13, 
-11, 12, -11, 26, 3, -9, 10, -9, 26, -9, 31, -9, 27, -9, 
28, -9, 33, -9, 14, -9, 5, -9, 29, -9, 2, -9, 13, -9, 
12, -9, 2, 9, 104, 26, 3, -10, 10, -10, 26, -10, 31, -10, 
27, -10, 28, -10, 33, -10, 14, -10, 5, -10, 29, -10, 2, -10, 
13, -10, 12, -10, 26, 3, -12, 10, -12, 26, -12, 31, -12, 27, 
-12, 28, -12, 33, -12, 14, -12, 5, -12, 29, -12, 2, -12, 13, 
-12, 12, -12, 4, 5, -59, 2, -59, 22, 26, -65, 31, -65, 27, 
-65, 28, -65, 33, -65, 14, 99, 5, -65, 29, -65, 2, -65, 13, 
100, 12, 101, 16, 26, -61, 31, -61, 27, -61, 28, -61, 33, -61, 
5, -61, 29, -61, 2, -61, 16, 26, 69, 31, 70, 27, 94, 28, 
71, 33, 72, 5, -63, 29, 95, 2, -63, 8, 7, -73, 22, -73, 
31, -73, 21, -73, 2, 7, 105, 8, 7, -70, 22, 15, 31, 16, 
21, 17, 2, 2, 106, 16, 7, -27, 0, -27, 22, -27, 31, -27, 
18, -27, 21, -27, 23, -27, 17, -27, 2, 2, 107, 28, 10, -38, 
20, -38, 31, -38, 33, -38, 34, -38, 14, -38, 9, -38, 5, -38, 
25, -38, 24, -38, 2, -38, 8, -38, 13, -38, 12, -38, 6, 35, 
-40, 37, -40, 36, -40, 6, 35, -41, 37, -41, 36, -41, 6, 35, 
-87, 37, -87, 36, -87, 6, 35, 108, 37, 85, 36, 86, 2, 31, 
110, 8, 20, -76, 25, -76, 24, -76, 2, -76, 12, 10, -78, 20, 
-78, 5, -78, 25, -78, 24, -78, 2, -78, 10, 0, -6, 31, -6, 
18, -6, 17, -6, 19, -6, 2, 9, 113, 2, 9, 114, 4, 5, 
-62, 2, -62, 4, 5, -7, 2, -7, 16, 26, -60, 31, -60, 27, 
-60, 28, -60, 33, -60, 5, -60, 29, -60, 2, -60, 16, 26, -15, 
31, -15, 27, -15, 28, -15, 33, -15, 5, -15, 29, -15, 2, -15, 
16, 26, -14, 31, -14, 27, -14, 28, -14, 33, -14, 5, -14, 29, 
-14, 2, -14, 16, 26, -16, 31, -16, 27, -16, 28, -16, 33, -16, 
5, -16, 29, -16, 2, -16, 16, 26, -64, 31, -64, 27, -64, 28, 
-64, 33, -64, 5, -64, 29, -64, 2, -64, 16, 26, -8, 31, -8, 
27, -8, 28, -8, 33, -8, 5, -8, 29, -8, 2, -8, 14, 0, 
-25, 22, -25, 31, -25, 18, -25, 21, -25, 23, -25, 17, -25, 16, 
7, -26, 0, -26, 22, -26, 31, -26, 18, -26, 21, -26, 23, -26, 
17, -26, 16, 7, -28, 0, -28, 22, -28, 31, -28, 18, -28, 21, 
-28, 23, -28, 17, -28, 28, 10, -39, 20, -39, 31, -39, 33, -39, 
34, -39, 14, -39, 9, -39, 5, -39, 25, -39, 24, -39, 2, -39, 
8, -39, 13, -39, 12, -39, 6, 35, -86, 37, -86, 36, -86, 2, 
10, 115, 8, 7, -72, 22, -72, 31, -72, 21, -72, 2, 3, 117, 
2, 32, 118, 2, 32, 119, 8, 20, -46, 25, -46, 24, -46, 2, 
-46, 4, 5, -58, 2, -58, 2, 10, 121, 2, 10, 122, 2, 10, 
123, 4, 5, -17, 2, -17, 4, 5, -18, 2, -18, 26, 3, -13, 
10, -13, 26, -13, 31, -13, 27, -13, 28, -13, 33, -13, 14, -13, 
5, -13, 29, -13, 2, -13, 13, -13, 12, -13, 
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
