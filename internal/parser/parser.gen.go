package parser

import (
  _i0 "github.com/dcaiafa/lox/internal/ast"
)

var _lxLHS = []int32 {
	0, 1, 2, 2, 3, 4, 4, 5, 6, 7, 8, 8, 8, 9, 
9, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 16, 
}

var _lxTermCounts = []int32 {
	1, 1, 1, 1, 4, 3, 1, 2, 2, 1, 1, 1, 1, 4, 
4, 3, 1, 0, 2, 1, 1, 0, 1, 0, 2, 1, 2, 1, 	
}

var _lxActions = []int32 {
	40, 47, 50, 53, 60, 63, 66, 73, 80, 87, 90, 95, 100, 107, 
124, 129, 140, 145, 162, 173, 178, 185, 188, 191, 196, 201, 87, 212, 
219, 230, 241, 252, 263, 274, 277, 280, 285, 288, 291, 296, 6, 0, 
-17, 2, 1, 13, 2, 2, 8, 9, 2, 2, 10, 6, 0, -27, 
2, -27, 13, -27, 2, 0, 2147483647, 2, 0, -1, 6, 0, -16, 2, 
1, 13, 2, 6, 0, -2, 2, -2, 13, -2, 6, 0, -3, 2, 
-3, 13, -3, 2, 2, 13, 4, 2, -25, 10, -25, 4, 2, 19, 
10, 20, 6, 0, -26, 2, -26, 13, -26, 16, 2, -9, 14, -9, 
6, -9, 9, -9, 15, -9, 10, -9, 5, -9, 7, -9, 4, 9, 
-6, 10, -6, 10, 2, 13, 14, 21, 9, -21, 15, 22, 10, -21, 
4, 9, 26, 10, 27, 16, 2, -23, 14, -23, 6, 28, 9, -23, 
15, -23, 10, -23, 5, 29, 7, 30, 10, 2, -19, 14, -19, 9, 
-19, 15, -19, 10, -19, 4, 2, -24, 10, -24, 6, 0, -15, 2, 
-15, 13, -15, 2, 11, 33, 2, 11, 34, 4, 9, -7, 10, -7, 
4, 9, -20, 10, -20, 10, 2, -18, 14, -18, 9, -18, 15, -18, 
10, -18, 6, 0, -4, 2, -4, 13, -4, 10, 2, -11, 14, -11, 
9, -11, 15, -11, 10, -11, 10, 2, -10, 14, -10, 9, -10, 15, 
-10, 10, -10, 10, 2, -12, 14, -12, 9, -12, 15, -12, 10, -12, 
10, 2, -22, 14, -22, 9, -22, 15, -22, 10, -22, 10, 2, -8, 
14, -8, 9, -8, 15, -8, 10, -8, 2, 4, 36, 2, 4, 37, 
4, 9, -5, 10, -5, 2, 12, 38, 2, 12, 39, 4, 9, -13, 
10, -13, 4, 9, -14, 10, -14, 
}

var _lxGoto = []int32 {
	40, 53, 54, 53, 53, 53, 57, 53, 53, 64, 53, 53, 53, 53, 
53, 75, 53, 84, 53, 53, 53, 53, 53, 53, 53, 53, 89, 53, 
53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 12, 2, 
3, 1, 4, 11, 5, 16, 6, 3, 7, 10, 8, 0, 2, 15, 
11, 6, 2, 12, 3, 7, 10, 8, 10, 5, 14, 12, 15, 4, 
16, 7, 17, 6, 18, 8, 13, 23, 9, 24, 7, 17, 6, 25, 
4, 8, 31, 14, 32, 8, 5, 35, 12, 15, 7, 17, 6, 18, 
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
				return p.on_term(
					p.sym.Peek(0).(Token),
		    )
			case 10:
				return p.on_card(
					p.sym.Peek(0).(Token),
		    )
			case 11:
				return p.on_card(
					p.sym.Peek(0).(Token),
		    )
			case 12:
				return p.on_card(
					p.sym.Peek(0).(Token),
		    )
			case 13:
				return p.on_qualif(
					p.sym.Peek(3).(Token),
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(Token),
		    )
			case 14:
				return p.on_qualif(
					p.sym.Peek(3).(Token),
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(Token),
		    )
			case 15:
				return p.on_token(
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).([]Token),
					p.sym.Peek(0).(Token),
		    )
  case 16:  // ZeroOrOne
			return p.sym.Peek(0).([]_i0.ParserDecl)
  case 17:  // ZeroOrOne
			{
				var zero []_i0.ParserDecl
				return zero
			}
  case 18:  // OneOrMore
			return append(
				p.sym.Peek(1).([]*_i0.Term),
				p.sym.Peek(0).(*_i0.Term),
			)
  case 19:  // OneOrMore
		  return []*_i0.Term{
				p.sym.Peek(0).(*_i0.Term),
			}
  case 20:  // ZeroOrOne
			return p.sym.Peek(0).(*_i0.ProdQualifier)
  case 21:  // ZeroOrOne
			{
				var zero *_i0.ProdQualifier
				return zero
			}
  case 22:  // ZeroOrOne
			return p.sym.Peek(0).(_i0.Qualifier)
  case 23:  // ZeroOrOne
			{
				var zero _i0.Qualifier
				return zero
			}
  case 24:  // OneOrMore
			return append(
				p.sym.Peek(1).([]Token),
				p.sym.Peek(0).(Token),
			)
  case 25:  // OneOrMore
		  return []Token{
				p.sym.Peek(0).(Token),
			}
  case 26:  // OneOrMore
			return append(
				p.sym.Peek(1).([]_i0.ParserDecl),
				p.sym.Peek(0).(_i0.ParserDecl),
			)
  case 27:  // OneOrMore
		  return []_i0.ParserDecl{
				p.sym.Peek(0).(_i0.ParserDecl),
			}
	default:
		panic("unreachable")
	}
}
