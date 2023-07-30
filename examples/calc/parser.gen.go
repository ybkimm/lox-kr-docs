package main

import (
)

var _lxLHS = []int32 {
	0, 1, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 
}

var _lxTermCounts = []int32 {
	1, 1, 3, 3, 3, 3, 3, 3, 3, 1, 1, 2, 	
}

var _lxActions = []int32 {
	22, 29, 44, 47, 64, 22, 81, 22, 22, 22, 22, 22, 22, 84, 
101, 116, 133, 150, 167, 184, 201, 218, 6, 4, 2, 2, 3, 9, 
5, 14, 6, 7, 0, -1, 4, 8, 5, 9, 3, 10, 8, 11, 
7, 12, 2, 2, 13, 16, 10, -10, 6, -10, 0, -10, 4, -10, 
5, -10, 3, -10, 8, -10, 7, -10, 16, 10, -9, 6, -9, 0, 
-9, 4, -9, 5, -9, 3, -9, 8, -9, 7, -9, 2, 0, 2147483647, 
16, 10, -11, 6, -11, 0, -11, 4, -11, 5, -11, 3, -11, 8, 
-11, 7, -11, 14, 10, 21, 6, 7, 4, 8, 5, 9, 3, 10, 
8, 11, 7, 12, 16, 10, -5, 6, -5, 0, -5, 4, -5, 5, 
-5, 3, -5, 8, 11, 7, -5, 16, 10, -3, 6, 7, 0, -3, 
4, -3, 5, 9, 3, -3, 8, 11, 7, 12, 16, 10, -4, 6, 
-4, 0, -4, 4, -4, 5, -4, 3, -4, 8, 11, 7, -4, 16, 
10, -2, 6, 7, 0, -2, 4, -2, 5, 9, 3, -2, 8, 11, 
7, 12, 16, 10, -7, 6, -7, 0, -7, 4, -7, 5, -7, 3, 
-7, 8, 11, 7, -7, 16, 10, -6, 6, -6, 0, -6, 4, -6, 
5, -6, 3, -6, 8, 11, 7, -6, 16, 10, -8, 6, -8, 0, 
-8, 4, -8, 5, -8, 3, -8, 8, -8, 7, -8, 
}

var _lxGoto = []int32 {
	22, 29, 29, 29, 29, 30, 29, 35, 40, 45, 50, 55, 60, 29, 
29, 29, 29, 29, 29, 29, 29, 29, 6, 2, 1, 3, 4, 1, 
6, 0, 4, 2, 14, 3, 4, 4, 2, 15, 3, 4, 4, 2, 
16, 3, 4, 4, 2, 17, 3, 4, 4, 2, 18, 3, 4, 4, 
2, 19, 3, 4, 4, 2, 20, 3, 4, 
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
				return p.reduceS(
					p.sym.Peek(0).(float64),
		    )
			case 2:
				return p.reduceExpr_binary(
					p.sym.Peek(2).(float64),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(float64),
		    )
			case 3:
				return p.reduceExpr_binary(
					p.sym.Peek(2).(float64),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(float64),
		    )
			case 4:
				return p.reduceExpr_binary(
					p.sym.Peek(2).(float64),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(float64),
		    )
			case 5:
				return p.reduceExpr_binary(
					p.sym.Peek(2).(float64),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(float64),
		    )
			case 6:
				return p.reduceExpr_binary(
					p.sym.Peek(2).(float64),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(float64),
		    )
			case 7:
				return p.reduceExpr_binary(
					p.sym.Peek(2).(float64),
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(float64),
		    )
			case 8:
				return p.reduceExpr_paren(
					p.sym.Peek(2).(Token),
					p.sym.Peek(1).(float64),
					p.sym.Peek(0).(Token),
		    )
			case 9:
				return p.reduceExpr_num(
					p.sym.Peek(0).(float64),
		    )
			case 10:
				return p.reduceNum(
					p.sym.Peek(0).(Token),
		    )
			case 11:
				return p.reduceNum_minus(
					p.sym.Peek(1).(Token),
					p.sym.Peek(0).(Token),
		    )
	default:
		panic("unreachable")
	}
}
