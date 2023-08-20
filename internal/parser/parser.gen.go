package parser

import (
  _i0 "github.com/dcaiafa/lox/internal/ast"
)

var _lxLHS = []int32 {
	0, 1, 2, 2, 3, 4, 4, 5, 6, 7, 7, 7, 8, 8, 
9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 
}

var _lxTermCounts = []int32 {
	1, 1, 1, 1, 4, 3, 1, 2, 2, 1, 1, 1, 4, 4, 
3, 1, 0, 2, 1, 1, 0, 1, 0, 2, 1, 2, 1, 	
}

var _lxActions = []int32 {
	39, 46, 49, 56, 59, 62, 69, 76, 83, 86, 89, 96, 101, 106, 
123, 128, 139, 144, 155, 160, 167, 178, 189, 200, 211, 222, 225, 230, 
235, 246, 86, 249, 256, 259, 262, 267, 270, 273, 278, 6, 0, -16, 
2, 1, 13, 8, 2, 8, 9, 6, 0, -3, 2, -3, 13, -3, 
2, 0, 2147483647, 2, 0, -1, 6, 0, -15, 2, 1, 13, 8, 6, 
0, -26, 2, -26, 13, -26, 6, 0, -2, 2, -2, 13, -2, 2, 
2, 11, 2, 2, 13, 6, 0, -25, 2, -25, 13, -25, 4, 2, 
-24, 10, -24, 4, 2, 18, 10, 19, 16, 2, -22, 14, -22, 6, 
20, 9, -22, 15, -22, 10, -22, 5, 23, 7, 24, 4, 9, -6, 
10, -6, 10, 2, 13, 14, 25, 9, -20, 15, 29, 10, -20, 4, 
9, 30, 10, 31, 10, 2, -18, 14, -18, 9, -18, 15, -18, 10, 
-18, 4, 2, -23, 10, -23, 6, 0, -14, 2, -14, 13, -14, 10, 
2, -10, 14, -10, 9, -10, 15, -10, 10, -10, 10, 2, -21, 14, 
-21, 9, -21, 15, -21, 10, -21, 10, 2, -8, 14, -8, 9, -8, 
15, -8, 10, -8, 10, 2, -9, 14, -9, 9, -9, 15, -9, 10, 
-9, 10, 2, -11, 14, -11, 9, -11, 15, -11, 10, -11, 2, 11, 
32, 4, 9, -7, 10, -7, 4, 9, -19, 10, -19, 10, 2, -17, 
14, -17, 9, -17, 15, -17, 10, -17, 2, 11, 33, 6, 0, -4, 
2, -4, 13, -4, 2, 4, 35, 2, 4, 36, 4, 9, -5, 10, 
-5, 2, 12, 37, 2, 12, 38, 4, 9, -12, 10, -12, 4, 9, 
-13, 10, -13, 
}

var _lxGoto = []int32 {
	39, 52, 52, 52, 52, 53, 52, 52, 60, 63, 52, 52, 52, 72, 
52, 77, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 
52, 52, 84, 52, 52, 52, 52, 52, 52, 52, 52, 12, 9, 2, 
1, 3, 10, 4, 15, 5, 2, 6, 3, 7, 0, 6, 9, 2, 
2, 10, 3, 7, 2, 14, 12, 8, 5, 14, 11, 15, 4, 16, 
6, 17, 4, 7, 21, 13, 22, 6, 12, 26, 8, 27, 6, 28, 
6, 5, 34, 11, 15, 6, 17, 
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
	bounds _lxStack[_lxBounds]
}

func (p *parser) parse(lex _lxLexer, errLogger _lxErrorLogger) bool {
  const accept = 2147483647

	p.loxParser.state.Push(0)
	tok, tokType := lex.NextToken()

	for {
		topState := p.loxParser.state.Peek(0)
		action, ok := _lxFind(_lxActions, topState, int32(tokType))
		if !ok {
			errLogger.ParserError(&_lxUnexpectedTokenError{Token: tok})
			return false
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			p.loxParser.state.Push(action)
			p.loxParser.sym.Push(tok)
			p.loxParser.bounds.Push(_lxBounds{Begin: tok, End: tok})
			tok, tokType = lex.NextToken()
		} else { // reduce
			prod := -action
			termCount := _lxTermCounts[int(prod)]
			rule := _lxLHS[int(prod)]
			res := p._lxAct(prod)
			if termCount > 0 {
				bounds := _lxBounds{
					Begin: p.loxParser.bounds.Peek(int(termCount-1)).Begin,
					End: p.loxParser.bounds.Peek(0).End,
				}
				p.onReduce(res, bounds.Begin, bounds.End)
				p.loxParser.bounds.Pop(int(termCount))
				p.loxParser.bounds.Push(bounds)
			} else {
				bounds := p.loxParser.bounds.Peek(0)
				bounds.Begin = bounds.End
				p.loxParser.bounds.Push(bounds)
			}
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
				return p.reduceParser(
					p.sym.Peek(0).([]_i0.ParserDecl),
		    )
			case 2:
				return p.reducePdecl(
					p.sym.Peek(0).(_i0.ParserDecl),
		    )
			case 3:
				return p.reducePdecl(
					p.sym.Peek(0).(_i0.ParserDecl),
		    )
			case 4:
				return p.reducePrule(
					p.sym.Peek(3).(Token),
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).([]*_i0.Prod),
					p.sym.Peek(0).(Token),
		    )
			case 5:
				return p.reducePprods(
					p.sym.Peek(2).([]*_i0.Prod),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(*_i0.Prod),
		    )
			case 6:
				return p.reducePprods_1(
					p.sym.Peek(0).(*_i0.Prod),
		    )
			case 7:
				return p.reducePprod(
					p.sym.Peek(1).([]*_i0.Term),
					p.sym.Peek(0).(*_i0.ProdQualifier),
		    )
			case 8:
				return p.reducePterm(
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(_i0.Qualifier),
		    )
			case 9:
				return p.reducePcard(
					p.sym.Peek(0).(Token),
		    )
			case 10:
				return p.reducePcard(
					p.sym.Peek(0).(Token),
		    )
			case 11:
				return p.reducePcard(
					p.sym.Peek(0).(Token),
		    )
			case 12:
				return p.reducePqualif(
					p.sym.Peek(3).(Token),
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(Token),
		    )
			case 13:
				return p.reducePqualif(
					p.sym.Peek(3).(Token),
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(Token),
		    )
			case 14:
				return p.reduceLtoken(
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).([]Token),
					p.sym.Peek(0).(Token),
		    )
  case 15:  // ZeroOrOne
			return p.sym.Peek(0).([]_i0.ParserDecl)
  case 16:  // ZeroOrOne
			{
				var zero []_i0.ParserDecl
				return zero
			}
  case 17:  // OneOrMore
			return append(
				p.sym.Peek(1).([]*_i0.Term),
				p.sym.Peek(0).(*_i0.Term),
			)
  case 18:  // OneOrMore
		  return []*_i0.Term{
				p.sym.Peek(0).(*_i0.Term),
			}
  case 19:  // ZeroOrOne
			return p.sym.Peek(0).(*_i0.ProdQualifier)
  case 20:  // ZeroOrOne
			{
				var zero *_i0.ProdQualifier
				return zero
			}
  case 21:  // ZeroOrOne
			return p.sym.Peek(0).(_i0.Qualifier)
  case 22:  // ZeroOrOne
			{
				var zero _i0.Qualifier
				return zero
			}
  case 23:  // OneOrMore
			return append(
				p.sym.Peek(1).([]Token),
				p.sym.Peek(0).(Token),
			)
  case 24:  // OneOrMore
		  return []Token{
				p.sym.Peek(0).(Token),
			}
  case 25:  // OneOrMore
			return append(
				p.sym.Peek(1).([]_i0.ParserDecl),
				p.sym.Peek(0).(_i0.ParserDecl),
			)
  case 26:  // OneOrMore
		  return []_i0.ParserDecl{
				p.sym.Peek(0).(_i0.ParserDecl),
			}
	default:
		panic("unreachable")
	}
}
