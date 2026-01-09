package lexer

import "pomidor/intrp/internal/token"

type Lexer struct {
	input []rune
	pos   int
}

// Get current rune
func (l *Lexer) peek() rune {
	if l.pos >= len(l.input) {
		return 0 // EOF
	}

	return l.input[l.pos]
}

// Go to the next rune
func (l *Lexer) next() rune {
	ch := l.peek()
	l.pos++
	return ch
}

// Get the next rune without incrementing lexer position
func (l *Lexer) peekNext() rune {
	l.pos++
	ch := l.peek()
	l.pos--
	return ch
}

// Instantiate lexer
func NewLexer(source string) *Lexer {
	if len([]rune(source)) < 1 {
		panic("Empty source")
	}

	return &Lexer{
		input: []rune(source),
		pos:   0,
	}
}

// Common token match loop
func (l *Lexer) NextToken() token.Token {
	for {
		ch := l.peek()

		if ch == 0 {
			return token.Token{Kind: token.TOK_EOF}
		}

		if isWhitespace(ch) {
			l.skipWhitespace()
			continue
		}

		if ch == '/' && l.peekNext() == '/' {
			l.skipComment()
			continue
		}

		if isLetter(ch) {
			return l.lexIdentOrKeyword()
		}

		if ch == '"' {
			return l.lexString()
		}

		return l.lexSymbol()
	}
}
