package main

import (
  _i0 "errors"
  _i1 "github.com/dcaiafa/lox/internal/base/baselexer"
)

var _LHS = []int32 {
	0, 1, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 
}

var _TermCounts = []int32 {
	1, 1, 3, 3, 3, 3, 3, 3, 3, 1, 1, 2, 	
}

var _Actions = []int32 {
	22, 29, 22, 46, 49, 52, 67, 22, 22, 22, 22, 22, 22, 84, 
99, 116, 133, 150, 167, 184, 201, 218, 6, 2, 1, 9, 2, 4, 
4, 16, 3, -10, 10, -10, 6, -10, 0, -10, 5, -10, 8, -10, 
7, -10, 4, -10, 2, 0, 2147483647, 2, 2, 14, 14, 3, 7, 6, 
8, 0, -1, 5, 9, 8, 10, 7, 11, 4, 12, 16, 3, -9, 
10, -9, 6, -9, 0, -9, 5, -9, 8, -9, 7, -9, 4, -9, 
14, 3, 7, 10, 15, 6, 8, 5, 9, 8, 10, 7, 11, 4, 
12, 16, 3, -11, 10, -11, 6, -11, 0, -11, 5, -11, 8, -11, 
7, -11, 4, -11, 16, 3, -8, 10, -8, 6, -8, 0, -8, 5, 
-8, 8, -8, 7, -8, 4, -8, 16, 3, -2, 10, -2, 6, 8, 
0, -2, 5, 9, 8, 10, 7, 11, 4, -2, 16, 3, -3, 10, 
-3, 6, 8, 0, -3, 5, 9, 8, 10, 7, 11, 4, -3, 16, 
3, -4, 10, -4, 6, -4, 0, -4, 5, -4, 8, 10, 7, -4, 
4, -4, 16, 3, -5, 10, -5, 6, -5, 0, -5, 5, -5, 8, 
10, 7, -5, 4, -5, 16, 3, -6, 10, -6, 6, -6, 0, -6, 
5, -6, 8, 10, 7, -6, 4, -6, 16, 3, -7, 10, -7, 6, 
-7, 0, -7, 5, -7, 8, -7, 7, -7, 4, -7, 
}

var _Goto = []int32 {
	22, 29, 30, 29, 29, 29, 29, 35, 40, 45, 50, 55, 60, 29, 
29, 29, 29, 29, 29, 29, 29, 29, 6, 1, 3, 2, 5, 3, 
6, 0, 4, 2, 13, 3, 6, 4, 2, 16, 3, 6, 4, 2, 
19, 3, 6, 4, 2, 18, 3, 6, 4, 2, 21, 3, 6, 4, 
2, 20, 3, 6, 4, 2, 17, 3, 6, 
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

	_lookahead     Token
	_lookaheadType int
	_errorToken    Token
}

func (p *calcParser) parse(lex _Lexer) bool {
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
			p._ReadToken()
		} else { // reduce
			prod := -action
			termCount := _TermCounts[int(prod)]
			rule := _LHS[int(prod)]
			res := p._Act(prod)
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

func (p *calcParser) errorToken() Token {
	return p._errorToken
}

func (p *calcParser) _ReadToken() {
	p._lookahead, p._lookaheadType = p._lex.ReadToken()
}

func (p *calcParser) _Recover() (int32, bool) {
	p._errorToken = p._lookahead

	for {
		for p._lookaheadType == ERROR {
			p._ReadToken()
		}

		saveState := p._state
		saveSym := p._sym

		for len(p._state) > 1 {
			topState := p._state.Peek(0)
			action, ok := _Find(_Actions, topState, int32(ERROR))
			if ok {
				action2, ok := _Find(
					_Actions, action, int32(p._lookaheadType))
				if ok {
					p._state.Push(action)
					p._sym.Push(_errorPlaceholder)
					return action2, true
				}
			}
			p._state.Pop(1)
			p._sym.Pop(1)
		}

		if p._lookaheadType == EOF {
			p.onError()
			return 0, false
		}

		p._ReadToken()
		p._state = saveState
		p._sym = saveSym
	}
}

func (p *calcParser) _Act(prod int32) any {
	switch prod {
			case 1:
				return p.on_S__foo(
				  _cast[float64](p._sym.Peek(0)),
		    )
			case 2:
				return p.on_expr__binary(
				  _cast[float64](p._sym.Peek(2)),
				  _cast[_i1.Token](p._sym.Peek(1)),
				  _cast[float64](p._sym.Peek(0)),
		    )
			case 3:
				return p.on_expr__binary(
				  _cast[float64](p._sym.Peek(2)),
				  _cast[_i1.Token](p._sym.Peek(1)),
				  _cast[float64](p._sym.Peek(0)),
		    )
			case 4:
				return p.on_expr__binary(
				  _cast[float64](p._sym.Peek(2)),
				  _cast[_i1.Token](p._sym.Peek(1)),
				  _cast[float64](p._sym.Peek(0)),
		    )
			case 5:
				return p.on_expr__binary(
				  _cast[float64](p._sym.Peek(2)),
				  _cast[_i1.Token](p._sym.Peek(1)),
				  _cast[float64](p._sym.Peek(0)),
		    )
			case 6:
				return p.on_expr__binary(
				  _cast[float64](p._sym.Peek(2)),
				  _cast[_i1.Token](p._sym.Peek(1)),
				  _cast[float64](p._sym.Peek(0)),
		    )
			case 7:
				return p.on_expr__binary(
				  _cast[float64](p._sym.Peek(2)),
				  _cast[_i1.Token](p._sym.Peek(1)),
				  _cast[float64](p._sym.Peek(0)),
		    )
			case 8:
				return p.on_expr__paren(
				  _cast[_i1.Token](p._sym.Peek(2)),
				  _cast[float64](p._sym.Peek(1)),
				  _cast[_i1.Token](p._sym.Peek(0)),
		    )
			case 9:
				return p.on_expr__num(
				  _cast[float64](p._sym.Peek(0)),
		    )
			case 10:
				return p.on_num(
				  _cast[_i1.Token](p._sym.Peek(0)),
		    )
			case 11:
				return p.on_num__minus(
				  _cast[_i1.Token](p._sym.Peek(1)),
				  _cast[_i1.Token](p._sym.Peek(0)),
		    )
	default:
		panic("unreachable")
	}
}
