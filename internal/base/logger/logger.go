package logger

import (
	"bytes"
	"fmt"
	"io"
)

type Logger struct {
	w      io.Writer
	indent int
}

func New(w io.Writer) *Logger {
	return &Logger{w: w}
}

func (l *Logger) WithIndent() *Logger {
	return &Logger{
		w:      l.w,
		indent: l.indent + 1,
	}
}

func (l *Logger) Logf(msg string, args ...any) {
	indent := make([]byte, l.indent*2+1)
	indent[0] = '\n'
	for i := 1; i < len(indent); i++ {
		indent[i] = ' '
	}

	buf := bytes.Buffer{}
	buf.Write(indent[1:])
	fmt.Fprintf(&buf, msg, args...)
	msgBytes := bytes.ReplaceAll(buf.Bytes(), []byte{'\n'}, indent)
	msgBytes = append(msgBytes, '\n')
	l.w.Write(msgBytes)
}
