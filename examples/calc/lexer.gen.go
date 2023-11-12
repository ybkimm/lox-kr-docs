package main



var _lexerMode0 = []uint32 {
	11, 52, 61, 66, 71, 76, 81, 86, 91, 96, 101, 40, 0, 48, 
57, 10, 0, 43, 43, 9, 0, 45, 45, 8, 0, 42, 42, 7, 
0, 47, 47, 6, 0, 37, 37, 5, 0, 94, 94, 4, 0, 40, 
40, 3, 0, 41, 41, 2, 0, 32, 32, 1, 8, 0, 32, 32, 
1, 4, 0, 0, 0, 4, 3, 0, 0, 10, 4, 3, 0, 0, 
9, 4, 3, 0, 0, 8, 4, 3, 0, 0, 7, 4, 3, 0, 
0, 6, 4, 3, 0, 0, 5, 4, 3, 0, 0, 4, 4, 3, 
0, 0, 3, 8, 0, 48, 57, 10, 3, 0, 0, 2, 
}



const (
	_lexerConsume  = 0
	_lexerAccept   = 1
	_lexerDiscard  = 2
	_lexerEOF      = 3
	_lexerError    = -1
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
		case 3: // Accept
			l.token = int(l.mode[i+3])
			l.state = 0
			return _lexerAccept
		case 4: // Discard
			l.state = 0
			return _lexerDiscard
		case 5: // Accum
			l.state = 0
			return _lexerConsume
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
