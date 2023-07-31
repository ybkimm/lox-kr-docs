package parser

import (
  _i0 "github.com/dcaiafa/lox/internal/ast"
)

var _lxLHS = []int32 {
	0, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 9, 9, 9, 
10, 10, 11, 12, 13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 
18, 19, 19, 20, 20, 21, 21, 22, 22, 
}

var _lxTermCounts = []int32 {
	1, 1, 1, 1, 2, 1, 4, 3, 1, 2, 2, 1, 1, 1, 
4, 4, 2, 1, 3, 2, 1, 1, 0, 2, 1, 1, 0, 1, 
0, 1, 0, 2, 1, 2, 1, 2, 1, 	
}

var _lxActions = []int32 {
	50, 55, 64, 71, 80, 87, 94, 97, 104, 113, 120, 129, 138, 141, 
144, 151, 160, 169, 178, 185, 194, 199, 204, 207, 216, 221, 230, 247, 
252, 263, 268, 279, 290, 301, 312, 323, 334, 337, 342, 347, 358, 204, 
361, 370, 373, 376, 381, 384, 387, 392, 4, 14, 1, 13, 3, 8, 
0, -30, 14, -30, 13, -30, 15, 12, 6, 0, -3, 14, -3, 13, 
-3, 8, 0, -22, 2, 13, 14, -22, 13, -22, 6, 0, -2, 14, 
-2, 13, -2, 6, 0, -20, 14, -20, 13, -20, 2, 0, 2147483647, 6, 
0, -1, 14, 1, 13, 3, 8, 0, -36, 14, -36, 13, -36, 15, 
-36, 6, 0, -16, 14, -16, 13, -16, 8, 0, -29, 14, -29, 13, 
-29, 15, 12, 8, 0, -17, 14, -17, 13, -17, 15, -17, 2, 2, 
20, 2, 8, 22, 6, 0, -4, 14, -4, 13, -4, 8, 0, -21, 
2, 13, 14, -21, 13, -21, 8, 0, -34, 2, -34, 14, -34, 13, 
-34, 8, 0, -5, 2, -5, 14, -5, 13, -5, 6, 0, -19, 14, 
-19, 13, -19, 8, 0, -35, 14, -35, 13, -35, 15, -35, 4, 2, 
-32, 10, -32, 4, 2, 24, 10, 25, 2, 2, 26, 8, 0, -33, 
2, -33, 14, -33, 13, -33, 4, 2, -31, 10, -31, 8, 0, -18, 
14, -18, 13, -18, 15, -18, 16, 2, -28, 16, -28, 6, 31, 9, 
-28, 17, -28, 10, -28, 5, 34, 7, 35, 4, 9, -8, 10, -8, 
10, 2, 26, 16, 36, 9, -26, 17, 40, 10, -26, 4, 9, 41, 
10, 42, 10, 2, -24, 16, -24, 9, -24, 17, -24, 10, -24, 10, 
2, -12, 16, -12, 9, -12, 17, -12, 10, -12, 10, 2, -27, 16, 
-27, 9, -27, 17, -27, 10, -27, 10, 2, -10, 16, -10, 9, -10, 
17, -10, 10, -10, 10, 2, -11, 16, -11, 9, -11, 17, -11, 10, 
-11, 10, 2, -13, 16, -13, 9, -13, 17, -13, 10, -13, 2, 11, 
43, 4, 9, -9, 10, -9, 4, 9, -25, 10, -25, 10, 2, -23, 
16, -23, 9, -23, 17, -23, 10, -23, 2, 11, 44, 8, 0, -6, 
2, -6, 14, -6, 13, -6, 2, 4, 46, 2, 4, 47, 4, 9, 
-7, 10, -7, 2, 12, 48, 2, 12, 49, 4, 9, -14, 10, -14, 
4, 9, -15, 10, -15, 
}

var _lxGoto = []int32 {
	50, 61, 70, 71, 70, 70, 70, 80, 70, 70, 87, 70, 92, 70, 
70, 95, 70, 70, 70, 70, 70, 70, 100, 70, 70, 70, 109, 70, 
114, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 121, 
70, 70, 70, 70, 70, 70, 70, 70, 10, 11, 2, 3, 4, 2, 
5, 1, 6, 14, 7, 8, 12, 8, 19, 9, 22, 10, 13, 11, 
0, 8, 15, 14, 21, 15, 4, 16, 5, 17, 6, 11, 2, 3, 
4, 2, 18, 4, 12, 19, 13, 11, 2, 20, 21, 4, 4, 23, 
5, 17, 8, 7, 27, 16, 28, 6, 29, 8, 30, 4, 9, 32, 
18, 33, 6, 17, 37, 10, 38, 8, 39, 6, 7, 45, 16, 28, 
8, 30, 
}

type _lxStack[T any] []T

func (s *_lxStack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *_lxStack[T]) Pop(n int) {
	*s = (*s)[:len(*s)-n]
}

func (s _lxStack[T]) Peek(n int) T {
	return s[len(s)-n-1]
}

func (s _lxStack[T]) Slice(n int) []T {
	return s[len(s)-n:]
}

func _lxFind(table []int32, y, x int32) (int32, bool) {
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

type loxParser struct {
	state _lxStack[int32]
	sym   _lxStack[any]
}

func (p *parser) parse(lex _lxLexer, errLogger _lxErrorLogger) bool {
  const accept = 2147483647

	p.loxParser.state.Push(0)
	tok := lex.NextToken()

	for {
		lookahead := int32(tok.Type)
		topState := p.loxParser.state.Peek(0)
		action, ok := _lxFind(_lxActions, topState, lookahead)
		if !ok {
			errLogger.Error(tok.Pos, &_lxUnexpectedTokenError{Token: tok})
			return false
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			p.loxParser.state.Push(action)
			p.loxParser.sym.Push(tok)
			tok = lex.NextToken()
		} else { // reduce
			prod := -action
			termCount := _lxTermCounts[int(prod)]
			rule := _lxLHS[int(prod)]
			res := p._lxAct(prod)
			p.onReduce(res, p.loxParser.sym.Slice(int(termCount)))
			p.loxParser.state.Pop(int(termCount))
			p.loxParser.sym.Pop(int(termCount))
			topState = p.loxParser.state.Peek(0)
			nextState, _ := _lxFind(_lxGoto, topState, rule)
			p.loxParser.state.Push(nextState)
			p.loxParser.sym.Push(res)
		}
	}

	return true
}

func (p *parser) _lxAct(prod int32) any {
	switch prod {
			case 1:
				return p.reduceSpec(
					p.sym.Peek(0).([]_i0.Section),
		    )
			case 2:
				return p.reduceSection(
					p.sym.Peek(0).(_i0.Section),
		    )
			case 3:
				return p.reduceSection(
					p.sym.Peek(0).(_i0.Section),
		    )
			case 4:
				return p.reduceParser(
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).([]_i0.ParserDecl),
		    )
			case 5:
				return p.reducePdecl(
					p.sym.Peek(0).(*_i0.Rule),
		    )
			case 6:
				return p.reducePrule(
					p.sym.Peek(3).(Token),
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).([]*_i0.Prod),
					p.sym.Peek(0).(Token),
		    )
			case 7:
				return p.reducePprods(
					p.sym.Peek(2).([]*_i0.Prod),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(*_i0.Prod),
		    )
			case 8:
				return p.reducePprods_1(
					p.sym.Peek(0).(*_i0.Prod),
		    )
			case 9:
				return p.reducePprod(
					p.sym.Peek(1).([]*_i0.Term),
					p.sym.Peek(0).(*_i0.ProdQualifier),
		    )
			case 10:
				return p.reducePterm(
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(_i0.Qualifier),
		    )
			case 11:
				return p.reducePcard(
					p.sym.Peek(0).(Token),
		    )
			case 12:
				return p.reducePcard(
					p.sym.Peek(0).(Token),
		    )
			case 13:
				return p.reducePcard(
					p.sym.Peek(0).(Token),
		    )
			case 14:
				return p.reducePqualif(
					p.sym.Peek(3).(Token),
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(Token),
		    )
			case 15:
				return p.reducePqualif(
					p.sym.Peek(3).(Token),
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(Token),
		    )
			case 16:
				return p.reduceLexer(
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).([]_i0.LexerDecl),
		    )
			case 17:
				return p.reduceLdecl(
					p.sym.Peek(0).(_i0.LexerDecl),
		    )
			case 18:
				return p.reduceLtoken(
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).([]Token),
					p.sym.Peek(0).(Token),
		    )
  case 19:  // OneOrMore
			return append(
				p.sym.Peek(1).([]_i0.Section),
				p.sym.Peek(0).(_i0.Section),
			)
  case 20:  // OneOrMore
		  return []_i0.Section{
				p.sym.Peek(0).(_i0.Section),
			}
  case 21:  // ZeroOrOne
			return p.sym.Peek(0).([]_i0.ParserDecl)
  case 22:  // ZeroOrOne
			{
				var zero []_i0.ParserDecl
				return zero
			}
  case 23:  // OneOrMore
			return append(
				p.sym.Peek(1).([]*_i0.Term),
				p.sym.Peek(0).(*_i0.Term),
			)
  case 24:  // OneOrMore
		  return []*_i0.Term{
				p.sym.Peek(0).(*_i0.Term),
			}
  case 25:  // ZeroOrOne
			return p.sym.Peek(0).(*_i0.ProdQualifier)
  case 26:  // ZeroOrOne
			{
				var zero *_i0.ProdQualifier
				return zero
			}
  case 27:  // ZeroOrOne
			return p.sym.Peek(0).(_i0.Qualifier)
  case 28:  // ZeroOrOne
			{
				var zero _i0.Qualifier
				return zero
			}
  case 29:  // ZeroOrOne
			return p.sym.Peek(0).([]_i0.LexerDecl)
  case 30:  // ZeroOrOne
			{
				var zero []_i0.LexerDecl
				return zero
			}
  case 31:  // OneOrMore
			return append(
				p.sym.Peek(1).([]Token),
				p.sym.Peek(0).(Token),
			)
  case 32:  // OneOrMore
		  return []Token{
				p.sym.Peek(0).(Token),
			}
  case 33:  // OneOrMore
			return append(
				p.sym.Peek(1).([]_i0.ParserDecl),
				p.sym.Peek(0).(_i0.ParserDecl),
			)
  case 34:  // OneOrMore
		  return []_i0.ParserDecl{
				p.sym.Peek(0).(_i0.ParserDecl),
			}
  case 35:  // OneOrMore
			return append(
				p.sym.Peek(1).([]_i0.LexerDecl),
				p.sym.Peek(0).(_i0.LexerDecl),
			)
  case 36:  // OneOrMore
		  return []_i0.LexerDecl{
				p.sym.Peek(0).(_i0.LexerDecl),
			}
	default:
		panic("unreachable")
	}
}
