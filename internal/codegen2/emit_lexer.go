package codegen2

const lexerTemplate = `

{{ range m := modes() }}

var _lexerMode{{m.Name}} = []int32 {
}


{{ end }}


type _LexerStateMachine struct {
	OnToken   func(t TokenType)
	OnDiscard func()

	state int
}

func (l *_LexerStateMachine) PushRune(r rune) {
}



`
