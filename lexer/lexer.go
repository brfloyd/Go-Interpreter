package lexer

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer { // New returns a new Lexer instance
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() { // readChar reads the next character in the input and advances the position in the input string
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}
