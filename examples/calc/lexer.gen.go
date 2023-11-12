package main



var _lexerMode0 = []uint32 {
	10, 47, 52, 57, 62, 67, 72, 77, 82, 87, 36, 0, 48, 57, 
9, 0, 43, 43, 8, 0, 45, 45, 7, 0, 42, 42, 6, 0, 
47, 47, 5, 0, 37, 37, 4, 0, 94, 94, 3, 0, 40, 40, 
2, 0, 41, 41, 1, 4, 1, 0, 0, 10, 4, 1, 0, 0, 
9, 4, 1, 0, 0, 8, 4, 1, 0, 0, 7, 4, 1, 0, 
0, 6, 4, 1, 0, 0, 5, 4, 1, 0, 0, 4, 4, 1, 
0, 0, 3, 8, 0, 48, 57, 9, 1, 0, 0, 2, 
}



const (
	_lexerConsume = 0
	_lexerDiscard = 1
	_lexerAccept  = 2
	_lexerEOF     = 3
	_lexerError   = -1
)

type _LexerStateMachine struct {
	token int
	state int
	mode  []uint32
}

func (l *_LexerStateMachine) PushRune(r rune) int {
	if l.mode == nil {
		l.mode = _lexerMode0
	}

	i := int(l.mode[int(l.state)])
	count := int(l.mode[i])
	i++
	end := i + count

	for ; i < end; i += 4 {
		switch l.mode[i] {
		case 0: // Goto
			if r >= rune(l.mode[i+1]) &&
				r <= rune(l.mode[i+2]) {
				l.state = int(l.mode[i+3])
				return _lexerConsume
			}
		case 1: // Accept
			l.token = int(l.mode[i+3])
			l.state = 0
			return _lexerAccept
		default:
			panic("not-reached")
		}
	}

	if l.state == 0 && r == 0 {
		return _lexerEOF
	}

	return _lexerError
}

func (l *_LexerStateMachine) Reset() {
	l.mode = nil
	l.state = 0
}

func (l *_LexerStateMachine) Token() int {
	return l.token
}
