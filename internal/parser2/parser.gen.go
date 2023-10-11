package parser

import (
  _i0 "errors"
  _i1 "github.com/dcaiafa/lox/internal/lexergen/ast"
)

var _LHS = []int32 {
	0, 1, 2, 2, 3, 4, 5, 6, 7, 8, 8, 8, 8, 9, 
10, 10, 10, 11, 11, 12, 13, 13, 14, 14, 14, 15, 16, 17, 
18, 19, 20, 21, 22, 22, 22, 23, 23, 23, 23, 24, 25, 25, 
26, 26, 26, 27, 28, 29, 30, 30, 31, 31, 32, 32, 33, 33, 
34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 
41, 41, 42, 42, 43, 43, 44, 44, 45, 45, 46, 46, 47, 47, 
48, 48, 
}

var _TermCounts = []int32 {
	1, 1, 1, 1, 2, 1, 5, 2, 2, 1, 1, 1, 1, 6, 
1, 1, 1, 4, 4, 2, 1, 1, 1, 1, 1, 5, 5, 4, 
5, 1, 1, 2, 1, 1, 1, 1, 1, 1, 3, 4, 1, 1, 
1, 1, 1, 1, 4, 1, 1, 0, 2, 1, 1, 0, 2, 1, 
1, 0, 3, 1, 2, 1, 1, 0, 1, 0, 1, 0, 2, 1, 
1, 0, 2, 1, 1, 0, 2, 1, 3, 1, 2, 1, 1, 0, 
2, 1, 	
}

var _Actions = []int32 {
	123, 130, 145, 156, 163, 170, 177, 180, 187, 190, 193, 198, 209, 220, 
227, 238, 249, 252, 255, 258, 275, 290, 305, 312, 327, 344, 359, 376, 
383, 386, 238, 389, 392, 405, 434, 238, 463, 492, 501, 514, 543, 566, 
589, 592, 603, 618, 627, 492, 636, 645, 648, 657, 666, 669, 678, 687, 
696, 238, 238, 705, 728, 751, 774, 797, 820, 843, 846, 851, 856, 883, 
910, 913, 940, 967, 972, 995, 1012, 1029, 1038, 1041, 1050, 1053, 1070, 1073, 
1102, 1109, 1116, 1123, 1130, 1133, 1142, 618, 1155, 1166, 1169, 1172, 1177, 1182, 
1199, 1216, 1233, 1250, 1267, 618, 1284, 1299, 1316, 1333, 1362, 1369, 1372, 1381, 
1384, 1387, 1390, 1399, 618, 1404, 1407, 1410, 1413, 1418, 1423, 6, 0, -49, 
23, 1, 22, 2, 14, 0, -67, 27, 15, 2, 16, 23, -67, 26, 
17, 28, 18, 22, -67, 10, 0, -53, 2, -57, 23, -53, 22, -53, 
24, 10, 6, 0, -3, 23, -3, 22, -3, 6, 0, -2, 23, -2, 
22, -2, 6, 0, -51, 23, -51, 22, -51, 2, 0, -1, 6, 0, 
-48, 23, 1, 22, 2, 2, 0, 2147483647, 2, 2, 28, 4, 2, -56, 
15, -56, 10, 0, -5, 2, -5, 23, -5, 22, -5, 24, -5, 10, 
0, -55, 2, -55, 23, -55, 22, -55, 24, -55, 6, 0, -4, 23, 
-4, 22, -4, 10, 0, -52, 2, -57, 23, -52, 22, -52, 24, 10, 
10, 2, 33, 3, 34, 15, -57, 17, 35, 24, 10, 2, 8, 30, 
2, 2, 42, 2, 2, 29, 16, 12, -23, 0, -23, 27, -23, 2, 
-23, 23, -23, 26, -23, 28, -23, 22, -23, 14, 0, -21, 27, -21, 
2, -21, 23, -21, 26, -21, 28, -21, 22, -21, 14, 0, -69, 27, 
-69, 2, -69, 23, -69, 26, -69, 28, -69, 22, -69, 6, 0, -19, 
23, -19, 22, -19, 14, 0, -66, 27, 15, 2, 16, 23, -66, 26, 
17, 28, 18, 22, -66, 16, 12, -24, 0, -24, 27, -24, 2, -24, 
23, -24, 26, -24, 28, -24, 22, -24, 14, 0, -20, 27, -20, 2, 
-20, 23, -20, 26, -20, 28, -20, 22, -20, 16, 12, -22, 0, -22, 
27, -22, 2, -22, 23, -22, 26, -22, 28, -22, 22, -22, 6, 0, 
-50, 23, -50, 22, -50, 2, 8, 45, 2, 11, 46, 2, 15, 66, 
12, 18, -29, 9, 58, 30, -29, 29, -29, 6, -29, 25, -29, 28, 
18, -36, 2, -36, 3, -36, 15, -36, 21, -36, 17, -36, 9, -36, 
30, -36, 29, -36, 6, -36, 25, -36, 24, -36, 20, -36, 19, -36, 
28, 18, -35, 2, -35, 3, -35, 15, -35, 21, -35, 17, -35, 9, 
-35, 30, -35, 29, -35, 6, -35, 25, -35, 24, -35, 20, -35, 19, 
-35, 28, 18, -37, 2, -37, 3, -37, 15, -37, 21, -37, 17, -37, 
9, -37, 30, -37, 29, -37, 6, -37, 25, -37, 24, -37, 20, -37, 
19, -37, 8, 30, 48, 29, 49, 6, -75, 25, 50, 12, 18, -79, 
9, -79, 30, -79, 29, -79, 6, -79, 25, -79, 28, 18, -83, 2, 
-83, 3, -83, 15, -83, 21, 60, 17, -83, 9, -83, 30, -83, 29, 
-83, 6, -83, 25, -83, 24, -83, 20, 61, 19, 62, 22, 18, -81, 
2, -81, 3, -81, 15, -81, 17, -81, 9, -81, 30, -81, 29, -81, 
6, -81, 25, -81, 24, -81, 22, 18, -30, 2, 33, 3, 34, 15, 
-57, 17, 35, 9, -30, 30, -30, 29, -30, 6, -30, 25, -30, 24, 
10, 2, 8, 57, 10, 0, -54, 2, -54, 23, -54, 22, -54, 24, 
-54, 14, 0, -68, 27, -68, 2, -68, 23, -68, 26, -68, 28, -68, 
22, -68, 8, 31, 68, 2, 69, 33, 70, 3, 71, 8, 12, -71, 
27, 15, 2, 16, 26, 17, 8, 30, -47, 29, -47, 6, -47, 25, 
-47, 2, 17, 88, 8, 30, -45, 29, -45, 6, -45, 25, -45, 8, 
30, -77, 29, -77, 6, -77, 25, -77, 2, 6, 81, 8, 30, 48, 
29, 49, 6, -74, 25, 50, 8, 30, -44, 29, -44, 6, -44, 25, 
-44, 8, 30, -43, 29, -43, 6, -43, 25, -43, 8, 30, -42, 29, 
-42, 6, -42, 25, -42, 22, 18, -80, 2, -80, 3, -80, 15, -80, 
17, -80, 9, -80, 30, -80, 29, -80, 6, -80, 25, -80, 24, -80, 
22, 18, -34, 2, -34, 3, -34, 15, -34, 17, -34, 9, -34, 30, 
-34, 29, -34, 6, -34, 25, -34, 24, -34, 22, 18, -33, 2, -33, 
3, -33, 15, -33, 17, -33, 9, -33, 30, -33, 29, -33, 6, -33, 
25, -33, 24, -33, 22, 18, -32, 2, -32, 3, -32, 15, -32, 17, 
-32, 9, -32, 30, -32, 29, -32, 6, -32, 25, -32, 24, -32, 22, 
18, -82, 2, -82, 3, -82, 15, -82, 17, -82, 9, -82, 30, -82, 
29, -82, 6, -82, 25, -82, 24, -82, 22, 18, -31, 2, -31, 3, 
-31, 15, -31, 17, -31, 9, -31, 30, -31, 29, -31, 6, -31, 25, 
-31, 24, -31, 2, 18, 83, 4, 4, 84, 13, 85, 4, 9, 91, 
6, 92, 26, 7, -11, 18, -11, 31, -11, 2, -11, 32, -11, 33, 
-11, 3, -11, 21, -11, 9, -11, 34, -11, 6, -11, 20, -11, 19, 
-11, 26, 7, -9, 18, -9, 31, -9, 2, -9, 32, -9, 33, -9, 
3, -9, 21, -9, 9, -9, 34, -9, 6, -9, 20, -9, 19, -9, 
2, 17, 103, 26, 7, -10, 18, -10, 31, -10, 2, -10, 32, -10, 
33, -10, 3, -10, 21, -10, 9, -10, 34, -10, 6, -10, 20, -10, 
19, -10, 26, 7, -12, 18, -12, 31, -12, 2, -12, 32, -12, 33, 
-12, 3, -12, 21, -12, 9, -12, 34, -12, 6, -12, 20, -12, 19, 
-12, 4, 9, -59, 6, -59, 22, 31, -65, 2, -65, 32, -65, 33, 
-65, 3, -65, 21, 98, 9, -65, 34, -65, 6, -65, 20, 99, 19, 
100, 16, 31, -61, 2, -61, 32, -61, 33, -61, 3, -61, 9, -61, 
34, -61, 6, -61, 16, 31, 68, 2, 69, 32, 93, 33, 70, 3, 
71, 9, -63, 34, 94, 6, -63, 8, 12, -73, 27, -73, 2, -73, 
26, -73, 2, 12, 104, 8, 12, -70, 27, 15, 2, 16, 26, 17, 
2, 6, 105, 16, 12, -27, 0, -27, 27, -27, 2, -27, 23, -27, 
26, -27, 28, -27, 22, -27, 2, 6, 106, 28, 18, -38, 2, -38, 
3, -38, 15, -38, 21, -38, 17, -38, 9, -38, 30, -38, 29, -38, 
6, -38, 25, -38, 24, -38, 20, -38, 19, -38, 6, 16, -40, 4, 
-40, 13, -40, 6, 16, -41, 4, -41, 13, -41, 6, 16, -85, 4, 
-85, 13, -85, 6, 16, 107, 4, 84, 13, 85, 2, 2, 109, 8, 
30, -76, 29, -76, 6, -76, 25, -76, 12, 18, -78, 9, -78, 30, 
-78, 29, -78, 6, -78, 25, -78, 10, 0, -6, 2, -6, 23, -6, 
22, -6, 24, -6, 2, 17, 112, 2, 17, 113, 4, 9, -62, 6, 
-62, 4, 9, -7, 6, -7, 16, 31, -60, 2, -60, 32, -60, 33, 
-60, 3, -60, 9, -60, 34, -60, 6, -60, 16, 31, -15, 2, -15, 
32, -15, 33, -15, 3, -15, 9, -15, 34, -15, 6, -15, 16, 31, 
-14, 2, -14, 32, -14, 33, -14, 3, -14, 9, -14, 34, -14, 6, 
-14, 16, 31, -16, 2, -16, 32, -16, 33, -16, 3, -16, 9, -16, 
34, -16, 6, -16, 16, 31, -64, 2, -64, 32, -64, 33, -64, 3, 
-64, 9, -64, 34, -64, 6, -64, 16, 31, -8, 2, -8, 32, -8, 
33, -8, 3, -8, 9, -8, 34, -8, 6, -8, 14, 0, -25, 27, 
-25, 2, -25, 23, -25, 26, -25, 28, -25, 22, -25, 16, 12, -26, 
0, -26, 27, -26, 2, -26, 23, -26, 26, -26, 28, -26, 22, -26, 
16, 12, -28, 0, -28, 27, -28, 2, -28, 23, -28, 26, -28, 28, 
-28, 22, -28, 28, 18, -39, 2, -39, 3, -39, 15, -39, 21, -39, 
17, -39, 9, -39, 30, -39, 29, -39, 6, -39, 25, -39, 24, -39, 
20, -39, 19, -39, 6, 16, -84, 4, -84, 13, -84, 2, 18, 114, 
8, 12, -72, 27, -72, 2, -72, 26, -72, 2, 7, 116, 2, 5, 
117, 2, 5, 118, 8, 30, -46, 29, -46, 6, -46, 25, -46, 4, 
9, -58, 6, -58, 2, 18, 120, 2, 18, 121, 2, 18, 122, 4, 
9, -17, 6, -17, 4, 9, -18, 6, -18, 26, 7, -13, 18, -13, 
31, -13, 2, -13, 32, -13, 33, -13, 3, -13, 21, -13, 9, -13, 
34, -13, 6, -13, 20, -13, 19, -13, 
}

var _Goto = []int32 {
	123, 136, 153, 164, 164, 164, 164, 165, 164, 164, 164, 164, 164, 164, 
172, 179, 164, 164, 164, 164, 164, 164, 164, 196, 164, 164, 164, 164, 
164, 164, 209, 164, 164, 164, 164, 226, 164, 243, 164, 256, 164, 261, 
164, 164, 164, 270, 283, 296, 164, 164, 164, 164, 164, 309, 164, 164, 
164, 318, 335, 164, 164, 164, 164, 164, 164, 164, 348, 164, 164, 164, 
164, 164, 164, 164, 353, 164, 358, 164, 164, 369, 164, 164, 164, 164, 
164, 164, 164, 378, 164, 164, 164, 381, 164, 164, 164, 164, 164, 164, 
164, 164, 164, 164, 164, 392, 164, 164, 164, 164, 164, 164, 164, 164, 
164, 164, 164, 164, 397, 164, 164, 164, 164, 164, 164, 12, 12, 3, 
3, 4, 2, 5, 30, 6, 31, 7, 1, 8, 16, 17, 19, 14, 
20, 13, 21, 39, 22, 40, 23, 18, 24, 15, 25, 16, 26, 10, 
34, 9, 5, 11, 4, 12, 32, 13, 33, 14, 0, 6, 12, 3, 
3, 4, 2, 27, 6, 34, 9, 5, 11, 4, 43, 16, 34, 31, 
45, 32, 24, 36, 19, 37, 20, 38, 23, 39, 21, 40, 46, 41, 
12, 17, 19, 14, 20, 13, 44, 18, 24, 15, 25, 16, 26, 16, 
34, 31, 45, 32, 24, 36, 19, 47, 20, 38, 23, 39, 21, 40, 
46, 41, 16, 34, 31, 45, 32, 24, 36, 19, 65, 20, 38, 23, 
39, 21, 40, 46, 41, 12, 26, 51, 43, 52, 44, 53, 29, 54, 
28, 55, 27, 56, 4, 22, 63, 47, 64, 8, 34, 31, 24, 36, 
23, 39, 21, 59, 12, 35, 67, 9, 72, 6, 73, 8, 74, 7, 
75, 36, 76, 12, 17, 19, 14, 77, 41, 78, 42, 79, 18, 24, 
16, 26, 12, 26, 51, 43, 80, 44, 53, 29, 54, 28, 55, 27, 
56, 8, 26, 89, 29, 54, 28, 55, 27, 56, 16, 34, 31, 45, 
32, 24, 36, 19, 82, 20, 38, 23, 39, 21, 40, 46, 41, 12, 
34, 31, 24, 36, 20, 90, 23, 39, 21, 40, 46, 41, 4, 25, 
86, 48, 87, 4, 10, 101, 38, 102, 10, 9, 72, 11, 95, 37, 
96, 8, 74, 7, 97, 8, 17, 19, 14, 110, 18, 24, 16, 26, 
2, 25, 108, 10, 9, 72, 6, 115, 8, 74, 7, 75, 36, 76, 
4, 9, 72, 8, 111, 4, 9, 72, 8, 119, 
}

type _Bounds struct {
	Begin Token
	End   Token
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
	for ; i < end; i+=2 {
		if table[i] == x {
			return table[i+1], true
		}
	}
	return 0, false
}

type _Lexer interface {
	ReadToken() (Token, TokenType)
}

type lox struct {
	_lex   _Lexer
	_state _Stack[int32]
	_sym   _Stack[any]
	_bounds _Stack[_Bounds]

	_lookahead     Token
	_lookaheadType TokenType
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
			return _cast[Token](p._sym.Peek(0))
  case 57:  // ZeroOrOne
			{
				var zero Token
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
	case 84:  // OneOrMore
			return append(
				_cast[[]Token](p._sym.Peek(1)),
				_cast[Token](p._sym.Peek(0)),
			)
	case 85:  // OneOrMore
		  return []Token{
				_cast[Token](p._sym.Peek(0)),
			}
	default:
		panic("unreachable")
	}
}
