package parser

import (
	"errors"
	gotoken "go/token"

	"github.com/dcaiafa/loxlex/simplelexer"
)

type state int

const (
	stateStart state = iota
	stateExtend
	stateNL
)

type lexer struct {
	lexer       *simplelexer.Lexer
	file        *gotoken.File
	queuedToken Token
}

func newLexer(cfg simplelexer.Config) *lexer {
	return &lexer{
		lexer: simplelexer.New(cfg),
		file:  cfg.File,
	}
}

func (l *lexer) ReadToken() (Token, int) {
	if l.queuedToken.Type != EOF {
		tok := l.queuedToken
		l.queuedToken = Token{}
		return tok, tok.Type
	}

	state := stateStart
	for {
		tok, _ := l.lexer.ReadToken()

		switch state {
		case stateStart:
			if tok.Type == EXTEND {
				state = stateExtend
			} else if tok.Type == NL {
				state = stateNL
			} else {
				return tok, tok.Type
			}

		case stateExtend:
			if tok.Type != NL {
				tok = Token{
					Type: ERROR,
					Err:  errors.New("token \\ must be followed by a new-line"),
					Pos:  tok.Pos,
				}
				return tok, ERROR
			}
			state = stateStart

		case stateNL:
			if tok.Type == OR {
				return tok, tok.Type
			} else if tok.Type == NL {
				continue
			} else {
				l.queuedToken = tok
				tok = Token{
					Type: NL,
					Pos:  tok.Pos,
				}
				return tok, NL
			}

		default:
			panic("unreachable")
		}
	}
}
