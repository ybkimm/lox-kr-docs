package parser2

import (
  _i0 "fmt"
  _i2 "github.com/dcaiafa/lox/internal/ast"
  _i3 "github.com/dcaiafa/lox/internal/token"
  _i1 "os"
)

var _lxLHS = []int32 {
	0, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 9, 9, 9, 
10, 11, 12, 13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18, 
19, 19, 20, 20, 21, 21, 22, 22, 
}

var _lxTermCounts = []int32 {
	1, 1, 1, 1, 2, 1, 4, 3, 1, 2, 2, 1, 1, 1, 
1, 2, 1, 3, 2, 1, 1, 0, 2, 1, 1, 0, 1, 0, 
1, 0, 2, 1, 2, 1, 2, 1, 	
}

var _lxActions = []int32 {
	43, 48, 57, 64, 73, 80, 87, 90, 97, 100, 109, 118, 125, 134, 
137, 144, 153, 162, 171, 178, 183, 188, 197, 200, 209, 214, 223, 238, 
243, 252, 257, 266, 275, 284, 293, 302, 311, 316, 321, 326, 197, 335, 
344, 4, 11, 1, 10, 3, 8, 12, 8, 0, -29, 11, -29, 10, 
-29, 6, 0, -3, 11, -3, 10, -3, 8, 0, -21, 1, 13, 11, 
-21, 10, -21, 6, 0, -2, 11, -2, 10, -2, 6, 0, -19, 11, 
-19, 10, -19, 2, 0, 2147483647, 6, 0, -1, 11, 1, 10, 3, 2, 
1, 19, 8, 12, -16, 0, -16, 11, -16, 10, -16, 8, 12, -35, 
0, -35, 11, -35, 10, -35, 6, 0, -15, 11, -15, 10, -15, 8, 
12, 8, 0, -28, 11, -28, 10, -28, 2, 7, 22, 6, 0, -4, 
11, -4, 10, -4, 8, 0, -20, 1, 13, 11, -20, 10, -20, 8, 
0, -33, 1, -33, 11, -33, 10, -33, 8, 0, -5, 1, -5, 11, 
-5, 10, -5, 6, 0, -18, 11, -18, 10, -18, 4, 1, -31, 9, 
-31, 4, 1, 24, 9, 25, 8, 12, -34, 0, -34, 11, -34, 10, 
-34, 2, 1, 26, 8, 0, -32, 1, -32, 11, -32, 10, -32, 4, 
1, -30, 9, -30, 8, 12, -17, 0, -17, 11, -17, 10, -17, 14, 
1, -27, 3, -27, 5, 31, 8, -27, 9, -27, 4, 34, 6, 35, 
4, 8, -8, 9, -8, 8, 1, 26, 3, 36, 8, -25, 9, -25, 
4, 8, 40, 9, 41, 8, 1, -23, 3, -23, 8, -23, 9, -23, 
8, 1, -12, 3, -12, 8, -12, 9, -12, 8, 1, -26, 3, -26, 
8, -26, 9, -26, 8, 1, -10, 3, -10, 8, -10, 9, -10, 8, 
1, -11, 3, -11, 8, -11, 9, -11, 8, 1, -13, 3, -13, 8, 
-13, 9, -13, 4, 8, -14, 9, -14, 4, 8, -24, 9, -24, 4, 
8, -9, 9, -9, 8, 1, -22, 3, -22, 8, -22, 9, -22, 8, 
0, -6, 1, -6, 11, -6, 10, -6, 4, 8, -7, 9, -7, 
}

var _lxGoto = []int32 {
	43, 54, 63, 64, 63, 63, 63, 73, 80, 63, 63, 63, 83, 63, 
63, 88, 63, 63, 63, 63, 63, 63, 93, 63, 63, 63, 102, 63, 
107, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 114, 63, 
63, 10, 11, 2, 3, 4, 2, 5, 1, 6, 14, 7, 8, 13, 
9, 12, 10, 19, 11, 22, 12, 0, 8, 15, 14, 21, 15, 4, 
16, 5, 17, 6, 11, 2, 3, 4, 2, 18, 2, 20, 20, 4, 
13, 9, 12, 21, 4, 4, 23, 5, 17, 8, 7, 27, 16, 28, 
6, 29, 8, 30, 4, 9, 32, 18, 33, 6, 10, 37, 17, 38, 
8, 39, 6, 7, 42, 16, 28, 8, 30, 
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

type _lxLexer interface {
	NextToken() (int, Token)
}

type loxParser struct {
	state _lxStack[int32]
	sym   _lxStack[any]
}

func (p *parser) parse(lex _lxLexer) {
  const accept = 2147483647

	p.loxParser.state.Push(0)
	lookahead, tok := lex.NextToken()

	for {
		topState := p.loxParser.state.Peek(0)
		action, ok := _lxFind(_lxActions, topState, int32(lookahead))
		if !ok {
			p._lxRecover(tok, "boom")
			return
		}
		if action == accept {
    	break
		} else if action >= 0 { // shift
			p.loxParser.state.Push(action)
			p.loxParser.sym.Push(tok)
			lookahead, tok = lex.NextToken()
		} else { // reduce
			prod := -action
			termCount := _lxTermCounts[int(prod)]
			rule := _lxLHS[int(prod)]
			res := p._lxAct(prod)
			p.loxParser.state.Pop(int(termCount))
			p.loxParser.sym.Pop(int(termCount))
			topState = p.loxParser.state.Peek(0)
			nextState, _ := _lxFind(_lxGoto, topState, rule)
			p.loxParser.state.Push(nextState)
			p.loxParser.sym.Push(res)
		}
	}
}

func (p *parser) _lxRecover(tok Token, err string) {
	_i0.Println("ERROR:", err)
	_i1.Exit(1)
}

func (p *parser) _lxAct(prod int32) any {
	switch prod {
			case 1:
				return p.reduceSpec(
					p.sym.Peek(0).([]_i2.Section),
		    )
			case 2:
				return p.reduceSection(
					p.sym.Peek(0).(_i2.Section),
		    )
			case 3:
				return p.reduceSection(
					p.sym.Peek(0).(_i2.Section),
		    )
			case 4:
				return p.reduceParser(
					p.sym.Peek(1).(_i3.Token),
					p.sym.Peek(0).([]_i2.ParserDecl),
		    )
			case 5:
				return p.reducePdecl(
					p.sym.Peek(0).(*_i2.Rule),
		    )
			case 6:
				return p.reducePrule(
					p.sym.Peek(3).(_i3.Token),
					p.sym.Peek(2).(_i3.Token),
					p.sym.Peek(1).([]*_i2.Prod),
					p.sym.Peek(0).(_i3.Token),
		    )
			case 7:
				return p.reducePprods(
					p.sym.Peek(2).([]*_i2.Prod),
					p.sym.Peek(1).(_i3.Token),
					p.sym.Peek(0).(*_i2.Prod),
		    )
			case 8:
				return p.reducePprods_1(
					p.sym.Peek(0).(*_i2.Prod),
		    )
			case 9:
				return p.reducePprod(
					p.sym.Peek(1).([]*_i2.Term),
					p.sym.Peek(0).(*_i2.Label),
		    )
			case 10:
				return p.reducePterm(
					p.sym.Peek(1).(_i3.Token),
					p.sym.Peek(0).(_i2.Qualifier),
		    )
			case 11:
				return p.reducePcard(
					p.sym.Peek(0).(_i3.Token),
		    )
			case 12:
				return p.reducePcard(
					p.sym.Peek(0).(_i3.Token),
		    )
			case 13:
				return p.reducePcard(
					p.sym.Peek(0).(_i3.Token),
		    )
			case 14:
				return p.reduceLabel(
					p.sym.Peek(0).(_i3.Token),
		    )
			case 15:
				return p.reduceLexer(
					p.sym.Peek(1).(_i3.Token),
					p.sym.Peek(0).([]_i2.LexerDecl),
		    )
			case 16:
				return p.reduceLdecl(
					p.sym.Peek(0).(_i2.LexerDecl),
		    )
			case 17:
				return p.reduceLcustom(
					p.sym.Peek(2).(_i3.Token),
					p.sym.Peek(1).([]_i3.Token),
					p.sym.Peek(0).(_i3.Token),
		    )
  case 18:  // OneOrMore
			return append(
				p.sym.Peek(1).([]_i2.Section),
				p.sym.Peek(0).(_i2.Section),
			)
  case 19:  // OneOrMore
		  return []_i2.Section{
				p.sym.Peek(0).(_i2.Section),
			}
  case 20:  // ZeroOrOne
			return p.sym.Peek(0).([]_i2.ParserDecl)
  case 21:  // ZeroOrOne
			{
				var zero []_i2.ParserDecl
				return zero
			}
  case 22:  // OneOrMore
			return append(
				p.sym.Peek(1).([]*_i2.Term),
				p.sym.Peek(0).(*_i2.Term),
			)
  case 23:  // OneOrMore
		  return []*_i2.Term{
				p.sym.Peek(0).(*_i2.Term),
			}
  case 24:  // ZeroOrOne
			return p.sym.Peek(0).(*_i2.Label)
  case 25:  // ZeroOrOne
			{
				var zero *_i2.Label
				return zero
			}
  case 26:  // ZeroOrOne
			return p.sym.Peek(0).(_i2.Qualifier)
  case 27:  // ZeroOrOne
			{
				var zero _i2.Qualifier
				return zero
			}
  case 28:  // ZeroOrOne
			return p.sym.Peek(0).([]_i2.LexerDecl)
  case 29:  // ZeroOrOne
			{
				var zero []_i2.LexerDecl
				return zero
			}
  case 30:  // OneOrMore
			return append(
				p.sym.Peek(1).([]_i3.Token),
				p.sym.Peek(0).(_i3.Token),
			)
  case 31:  // OneOrMore
		  return []_i3.Token{
				p.sym.Peek(0).(_i3.Token),
			}
  case 32:  // OneOrMore
			return append(
				p.sym.Peek(1).([]_i2.ParserDecl),
				p.sym.Peek(0).(_i2.ParserDecl),
			)
  case 33:  // OneOrMore
		  return []_i2.ParserDecl{
				p.sym.Peek(0).(_i2.ParserDecl),
			}
  case 34:  // OneOrMore
			return append(
				p.sym.Peek(1).([]_i2.LexerDecl),
				p.sym.Peek(0).(_i2.LexerDecl),
			)
  case 35:  // OneOrMore
		  return []_i2.LexerDecl{
				p.sym.Peek(0).(_i2.LexerDecl),
			}
	default:
		panic("unreachable")
	}
}
