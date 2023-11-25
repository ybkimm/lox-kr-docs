package main



var _lexerMode0 = []uint32 {
	29, 76, 89, 100, 116, 121, 128, 133, 146, 151, 157, 162, 166, 171, 
175, 180, 184, 189, 194, 198, 203, 208, 212, 217, 221, 226, 230, 235, 
239, 46, 15, 9, 10, 1, 13, 13, 1, 32, 32, 1, 34, 34, 
9, 44, 44, 21, 45, 45, 8, 49, 57, 3, 58, 58, 18, 91, 
91, 25, 93, 93, 23, 102, 102, 17, 110, 110, 10, 116, 116, 24, 
123, 123, 28, 125, 125, 27, 12, 3, 9, 10, 1, 13, 13, 1, 
32, 32, 1, 4, 0, 10, 3, 43, 43, 4, 45, 45, 4, 49, 
57, 5, 15, 4, 46, 46, 6, 48, 57, 3, 69, 69, 2, 101, 
101, 2, 3, 12, 4, 1, 49, 57, 5, 6, 1, 48, 57, 5, 
3, 12, 4, 1, 48, 57, 7, 12, 3, 48, 57, 7, 69, 69, 
2, 101, 101, 2, 3, 12, 4, 1, 49, 57, 3, 5, 0, 1, 
1, 5, 0, 4, 1, 117, 117, 12, 3, 0, 3, 10, 4, 1, 
108, 108, 14, 3, 0, 3, 9, 4, 1, 108, 108, 11, 3, 0, 
3, 8, 4, 1, 117, 117, 26, 4, 1, 97, 97, 19, 3, 0, 
3, 7, 4, 1, 108, 108, 20, 4, 1, 115, 115, 22, 3, 0, 
3, 6, 4, 1, 101, 101, 13, 3, 0, 3, 5, 4, 1, 114, 
114, 16, 3, 0, 3, 4, 4, 1, 101, 101, 15, 3, 0, 3, 
3, 3, 0, 3, 2, 
}



var _lexerMode1 = []uint32 {
	10, 27, 31, 27, 60, 27, 71, 82, 88, 99, 16, 5, 32, 33, 
5, 34, 34, 7, 35, 91, 5, 92, 92, 2, 93, 1114111, 5, 3, 
0, 5, 0, 28, 9, 34, 34, 3, 47, 47, 3, 92, 92, 3, 
98, 98, 3, 102, 102, 3, 110, 110, 3, 114, 114, 3, 116, 116, 
3, 117, 117, 4, 10, 3, 48, 57, 9, 65, 70, 9, 97, 102, 
9, 10, 3, 48, 57, 1, 65, 70, 1, 97, 102, 1, 5, 0, 
2, 0, 3, 11, 10, 3, 48, 57, 6, 65, 70, 6, 97, 102, 
6, 10, 3, 48, 57, 8, 65, 70, 8, 97, 102, 8, 
}




var _lexerModes = [][]uint32 {

	_lexerMode0,

	_lexerMode1,

}


const (
	_lexerConsume  = 0
	_lexerAccept   = 1
	_lexerDiscard  = 2
	_lexerTryAgain = 3
	_lexerEOF      = 4
	_lexerError    = -1
)

type _LexerStateMachine struct {
	token int
	state int
	mode  []uint32
	modeStack _Stack[[]uint32]
}

func (l *_LexerStateMachine) PushRune(r rune) int {
	if l.mode == nil {
		l.mode = _lexerMode0
	}

	mode := l.mode

	// Find the table row corresponding to state.
	i := int(mode[int(l.state)])
	count := int(mode[i])
	i++
	end := i + count

	// The format of the row is as follows:
	//
	//   gotoCount uint32
	//   [gotoCount]struct{
	//     rangeBegin uint32
	//     rangeEnd   uint32
	//     gotoState  uint32
	//   }
	//   [actionCount]struct {
	//     actionType  uint32
	//     actionParam uint32
	//   }
	//
	// Where 'actionCount' is determined by the amount of uint32 left in the row.

	gotoN := int(mode[i])
	i++

	// Use binary-search to find the next state.
	b := 0
	e := gotoN
	for b < e {
		j := b + (e-b)/2
		k := i + j*3
		switch {
		case r >= rune(mode[k]) && r <= rune(mode[k+1]):
			l.state = int(mode[k+2])
			return _lexerConsume
		case r < rune(mode[k]):
			e = j
		case r > rune(mode[k+1]):
			b = j + 1
		default:
			panic("not reached")
		}
	}

	// Move 'i' to the beginning of the actions section.
	i += gotoN * 3

	for ; i < end; i += 2 {
		switch mode[i] {
		case 1: // PushMode
			modeIndex := int(mode[i+1])
			l.modeStack.Push(mode)
			l.mode = _lexerModes[modeIndex]
		case 2: // PopMode
			l.mode = l.modeStack.Peek(0)
			l.modeStack.Pop(1)
		case 3: // Accept
			l.token = int(mode[i+1])
			l.state = 0
			return _lexerAccept
		case 4: // Discard
			l.state = 0
			return _lexerDiscard
		case 5: // Accum
			l.state = 0
			return _lexerTryAgain
		}
	}

	if l.state == 0 && r == 0 {
		return _lexerEOF
	}

	return _lexerError}

func (l *_LexerStateMachine) Reset() {
	l.mode = nil
	l.state = 0
}

func (l *_LexerStateMachine) Token() int {
	return l.token
}
