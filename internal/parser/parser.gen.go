package parser

import (
  _i0 "fmt"
  _i2 "github.com/dcaiafa/lox/internal/ast"
  _i1 "os"
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
361, 370, 373, 376, 381, 384, 387, 392, 4, 13, 1, 12, 3, 8, 
0, -30, 13, -30, 12, -30, 14, 12, 6, 0, -3, 13, -3, 12, 
-3, 8, 0, -22, 1, 13, 13, -22, 12, -22, 6, 0, -2, 13, 
-2, 12, -2, 6, 0, -20, 13, -20, 12, -20, 2, 0, 2147483647, 6, 
0, -1, 13, 1, 12, 3, 8, 0, -36, 13, -36, 12, -36, 14, 
-36, 6, 0, -16, 13, -16, 12, -16, 8, 0, -29, 13, -29, 12, 
-29, 14, 12, 8, 0, -17, 13, -17, 12, -17, 14, -17, 2, 1, 
20, 2, 7, 22, 6, 0, -4, 13, -4, 12, -4, 8, 0, -21, 
1, 13, 13, -21, 12, -21, 8, 0, -34, 1, -34, 13, -34, 12, 
-34, 8, 0, -5, 1, -5, 13, -5, 12, -5, 6, 0, -19, 13, 
-19, 12, -19, 8, 0, -35, 13, -35, 12, -35, 14, -35, 4, 1, 
-32, 9, -32, 4, 1, 24, 9, 25, 2, 1, 26, 8, 0, -33, 
1, -33, 13, -33, 12, -33, 4, 1, -31, 9, -31, 8, 0, -18, 
13, -18, 12, -18, 14, -18, 16, 1, -28, 15, -28, 5, 31, 8, 
-28, 16, -28, 9, -28, 4, 34, 6, 35, 4, 8, -8, 9, -8, 
10, 1, 26, 15, 36, 8, -26, 16, 40, 9, -26, 4, 8, 41, 
9, 42, 10, 1, -24, 15, -24, 8, -24, 16, -24, 9, -24, 10, 
1, -12, 15, -12, 8, -12, 16, -12, 9, -12, 10, 1, -27, 15, 
-27, 8, -27, 16, -27, 9, -27, 10, 1, -10, 15, -10, 8, -10, 
16, -10, 9, -10, 10, 1, -11, 15, -11, 8, -11, 16, -11, 9, 
-11, 10, 1, -13, 15, -13, 8, -13, 16, -13, 9, -13, 2, 10, 
43, 4, 8, -9, 9, -9, 4, 8, -25, 9, -25, 10, 1, -23, 
15, -23, 8, -23, 16, -23, 9, -23, 2, 10, 44, 8, 0, -6, 
1, -6, 13, -6, 12, -6, 2, 3, 46, 2, 3, 47, 4, 8, 
-7, 9, -7, 2, 11, 48, 2, 11, 49, 4, 8, -14, 9, -14, 
4, 8, -15, 9, -15, 
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

type _lxUnexpectedTokenError struct {
	Token Token
}

func (e _lxUnexpectedTokenError) Error() string {
	return _i0.Sprintf("unexpected token: %v", e.Token)
}

type _lxLexer interface {
	NextToken() (Token, error)
}

type loxParser struct {
	state _lxStack[int32]
	sym   _lxStack[any]
}

func (p *parser) parse(lex _lxLexer) error {
  const accept = 2147483647

	p.loxParser.state.Push(0)
	tok, err := lex.NextToken()
	if err != nil {
		return err
	}

	for {
		lookahead := int32(tok.Type)
		topState := p.loxParser.state.Peek(0)
		action, ok := _lxFind(_lxActions, topState, lookahead)
		if !ok {
			return &_lxUnexpectedTokenError{Token: tok}
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			p.loxParser.state.Push(action)
			p.loxParser.sym.Push(tok)
			tok, err = lex.NextToken()
			if err != nil {
				return err
			}
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

	return nil
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
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).([]_i2.ParserDecl),
		    )
			case 5:
				return p.reducePdecl(
					p.sym.Peek(0).(*_i2.Rule),
		    )
			case 6:
				return p.reducePrule(
					p.sym.Peek(3).(Token),
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).([]*_i2.Prod),
					p.sym.Peek(0).(Token),
		    )
			case 7:
				return p.reducePprods(
					p.sym.Peek(2).([]*_i2.Prod),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(*_i2.Prod),
		    )
			case 8:
				return p.reducePprods_1(
					p.sym.Peek(0).(*_i2.Prod),
		    )
			case 9:
				return p.reducePprod(
					p.sym.Peek(1).([]*_i2.Term),
					p.sym.Peek(0).(*_i2.ProdQualifier),
		    )
			case 10:
				return p.reducePterm(
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(_i2.Qualifier),
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
					p.sym.Peek(0).([]_i2.LexerDecl),
		    )
			case 17:
				return p.reduceLdecl(
					p.sym.Peek(0).(_i2.LexerDecl),
		    )
			case 18:
				return p.reduceLtoken(
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).([]Token),
					p.sym.Peek(0).(Token),
		    )
  case 19:  // OneOrMore
			return append(
				p.sym.Peek(1).([]_i2.Section),
				p.sym.Peek(0).(_i2.Section),
			)
  case 20:  // OneOrMore
		  return []_i2.Section{
				p.sym.Peek(0).(_i2.Section),
			}
  case 21:  // ZeroOrOne
			return p.sym.Peek(0).([]_i2.ParserDecl)
  case 22:  // ZeroOrOne
			{
				var zero []_i2.ParserDecl
				return zero
			}
  case 23:  // OneOrMore
			return append(
				p.sym.Peek(1).([]*_i2.Term),
				p.sym.Peek(0).(*_i2.Term),
			)
  case 24:  // OneOrMore
		  return []*_i2.Term{
				p.sym.Peek(0).(*_i2.Term),
			}
  case 25:  // ZeroOrOne
			return p.sym.Peek(0).(*_i2.ProdQualifier)
  case 26:  // ZeroOrOne
			{
				var zero *_i2.ProdQualifier
				return zero
			}
  case 27:  // ZeroOrOne
			return p.sym.Peek(0).(_i2.Qualifier)
  case 28:  // ZeroOrOne
			{
				var zero _i2.Qualifier
				return zero
			}
  case 29:  // ZeroOrOne
			return p.sym.Peek(0).([]_i2.LexerDecl)
  case 30:  // ZeroOrOne
			{
				var zero []_i2.LexerDecl
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
				p.sym.Peek(1).([]_i2.ParserDecl),
				p.sym.Peek(0).(_i2.ParserDecl),
			)
  case 34:  // OneOrMore
		  return []_i2.ParserDecl{
				p.sym.Peek(0).(_i2.ParserDecl),
			}
  case 35:  // OneOrMore
			return append(
				p.sym.Peek(1).([]_i2.LexerDecl),
				p.sym.Peek(0).(_i2.LexerDecl),
			)
  case 36:  // OneOrMore
		  return []_i2.LexerDecl{
				p.sym.Peek(0).(_i2.LexerDecl),
			}
	default:
		panic("unreachable")
	}
}
