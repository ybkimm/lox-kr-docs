package parser

import (
  _i0 "github.com/dcaiafa/lox/internal/ast"
)

var _lxLHS = []int32 {
	0, 1, 2, 2, 3, 4, 4, 5, 6, 7, 7, 8, 9, 9, 
9, 10, 10, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 16, 
17, 17, 
}

var _lxTermCounts = []int32 {
	1, 1, 1, 1, 4, 3, 1, 2, 2, 1, 1, 5, 1, 1, 
1, 4, 4, 3, 1, 0, 2, 1, 1, 0, 1, 0, 2, 1, 
2, 1, 	
}

var _lxActions = []int32 {
	46, 53, 56, 59, 66, 69, 72, 79, 86, 93, 98, 103, 108, 115, 
136, 139, 160, 165, 178, 183, 202, 215, 220, 93, 227, 230, 233, 238, 
243, 93, 256, 263, 276, 289, 302, 315, 328, 331, 334, 337, 93, 342, 
345, 348, 369, 374, 6, 0, -19, 2, 1, 14, 2, 2, 10, 9, 
2, 2, 10, 6, 0, -29, 2, -29, 14, -29, 2, 0, 2147483647, 2, 
0, -1, 6, 0, -18, 2, 1, 14, 2, 6, 0, -2, 2, -2, 
14, -2, 6, 0, -3, 2, -3, 14, -3, 4, 2, 13, 17, 14, 
4, 2, -27, 13, -27, 4, 2, 21, 13, 22, 6, 0, -28, 2, 
-28, 14, -28, 20, 8, -9, 2, -9, 15, -9, 17, -9, 6, -9, 
12, -9, 16, -9, 13, -9, 5, -9, 7, -9, 2, 11, 23, 20, 
8, -10, 2, -10, 15, -10, 17, -10, 6, -10, 12, -10, 16, -10, 
13, -10, 5, -10, 7, -10, 4, 12, -6, 13, -6, 12, 2, 13, 
15, 24, 17, 14, 12, -23, 16, 25, 13, -23, 4, 12, 29, 13, 
30, 18, 2, -25, 15, -25, 17, -25, 6, 31, 12, -25, 16, -25, 
13, -25, 5, 32, 7, 33, 12, 2, -21, 15, -21, 17, -21, 12, 
-21, 16, -21, 13, -21, 4, 2, -26, 13, -26, 6, 0, -17, 2, 
-17, 14, -17, 2, 11, 37, 2, 11, 38, 4, 12, -7, 13, -7, 
4, 12, -22, 13, -22, 12, 2, -20, 15, -20, 17, -20, 12, -20, 
16, -20, 13, -20, 6, 0, -4, 2, -4, 14, -4, 12, 2, -13, 
15, -13, 17, -13, 12, -13, 16, -13, 13, -13, 12, 2, -12, 15, 
-12, 17, -12, 12, -12, 16, -12, 13, -12, 12, 2, -14, 15, -14, 
17, -14, 12, -14, 16, -14, 13, -14, 12, 2, -24, 15, -24, 17, 
-24, 12, -24, 16, -24, 13, -24, 12, 2, -8, 15, -8, 17, -8, 
12, -8, 16, -8, 13, -8, 2, 8, 40, 2, 4, 41, 2, 4, 
42, 4, 12, -5, 13, -5, 2, 9, 44, 2, 9, 45, 20, 8, 
-11, 2, -11, 15, -11, 17, -11, 6, -11, 12, -11, 16, -11, 13, 
-11, 5, -11, 7, -11, 4, 12, -15, 13, -15, 4, 12, -16, 13, 
-16, 
}

var _lxGoto = []int32 {
	46, 59, 60, 59, 59, 59, 63, 59, 59, 70, 59, 59, 59, 59, 
59, 59, 59, 83, 59, 94, 59, 59, 59, 99, 59, 59, 59, 59, 
59, 104, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 115, 59, 
59, 59, 59, 59, 12, 2, 3, 1, 4, 12, 5, 17, 6, 3, 
7, 11, 8, 0, 2, 16, 11, 6, 2, 12, 3, 7, 11, 8, 
12, 8, 15, 5, 16, 13, 17, 4, 18, 7, 19, 6, 20, 10, 
8, 15, 14, 26, 10, 27, 7, 19, 6, 28, 4, 9, 34, 15, 
35, 4, 8, 15, 7, 36, 10, 8, 15, 5, 39, 13, 17, 7, 
19, 6, 20, 4, 8, 15, 7, 43, 
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
				return p.on_parser(
					p.sym.Peek(0).([]_i0.ParserDecl),
		    )
			case 2:
				return p.on_decl(
					p.sym.Peek(0).(_i0.ParserDecl),
		    )
			case 3:
				return p.on_decl(
					p.sym.Peek(0).(_i0.ParserDecl),
		    )
			case 4:
				return p.on_rule(
					p.sym.Peek(3).(Token),
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).([]*_i0.Prod),
					p.sym.Peek(0).(Token),
		    )
			case 5:
				return p.on_prods(
					p.sym.Peek(2).([]*_i0.Prod),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(*_i0.Prod),
		    )
			case 6:
				return p.on_prods__1(
					p.sym.Peek(0).(*_i0.Prod),
		    )
			case 7:
				return p.on_prod(
					p.sym.Peek(1).([]*_i0.Term),
					p.sym.Peek(0).(*_i0.ProdQualifier),
		    )
			case 8:
				return p.on_term_card(
					p.sym.Peek(1).(*_i0.Term),
					p.sym.Peek(0).(_i0.Qualifier),
		    )
			case 9:
				return p.on_term__id(
					p.sym.Peek(0).(Token),
		    )
			case 10:
				return p.on_term__list(
					p.sym.Peek(0).(*_i0.Term),
		    )
			case 11:
				return p.on_list(
					p.sym.Peek(4).(Token),
					p.sym.Peek(3).(Token),
					p.sym.Peek(2).(*_i0.Term),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(*_i0.Term),
		    )
			case 12:
				return p.on_card(
					p.sym.Peek(0).(Token),
		    )
			case 13:
				return p.on_card(
					p.sym.Peek(0).(Token),
		    )
			case 14:
				return p.on_card(
					p.sym.Peek(0).(Token),
		    )
			case 15:
				return p.on_qualif(
					p.sym.Peek(3).(Token),
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(Token),
		    )
			case 16:
				return p.on_qualif(
					p.sym.Peek(3).(Token),
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(Token),
		    )
			case 17:
				return p.on_token(
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).([]Token),
					p.sym.Peek(0).(Token),
		    )
  case 18:  // ZeroOrOne
			return p.sym.Peek(0).([]_i0.ParserDecl)
  case 19:  // ZeroOrOne
			{
				var zero []_i0.ParserDecl
				return zero
			}
  case 20:  // OneOrMore
			return append(
				p.sym.Peek(1).([]*_i0.Term),
				p.sym.Peek(0).(*_i0.Term),
			)
  case 21:  // OneOrMore
		  return []*_i0.Term{
				p.sym.Peek(0).(*_i0.Term),
			}
  case 22:  // ZeroOrOne
			return p.sym.Peek(0).(*_i0.ProdQualifier)
  case 23:  // ZeroOrOne
			{
				var zero *_i0.ProdQualifier
				return zero
			}
  case 24:  // ZeroOrOne
			return p.sym.Peek(0).(_i0.Qualifier)
  case 25:  // ZeroOrOne
			{
				var zero _i0.Qualifier
				return zero
			}
  case 26:  // OneOrMore
			return append(
				p.sym.Peek(1).([]Token),
				p.sym.Peek(0).(Token),
			)
  case 27:  // OneOrMore
		  return []Token{
				p.sym.Peek(0).(Token),
			}
  case 28:  // OneOrMore
			return append(
				p.sym.Peek(1).([]_i0.ParserDecl),
				p.sym.Peek(0).(_i0.ParserDecl),
			)
  case 29:  // OneOrMore
		  return []_i0.ParserDecl{
				p.sym.Peek(0).(_i0.ParserDecl),
			}
	default:
		panic("unreachable")
	}
}
