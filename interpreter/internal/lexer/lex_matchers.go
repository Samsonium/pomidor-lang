package lexer

import "pomidor/intrp/internal/token"

// Match keyword or identifier
func (l *Lexer) lexIdentOrKeyword() token.Token {
	start := l.pos

	for isLetter(l.peek()) || isDigit(l.peek()) {
		l.next()
	}

	text := string(l.input[start:l.pos])
	switch text {
	case "garden":
		return token.Token{Kind: token.TOK_GARDEN}
	case "pluck":
		return token.Token{Kind: token.TOK_PLUCK}
	case "see":
		return token.Token{Kind: token.TOK_SEE}
	default:
		return token.Token{Kind: token.TOK_IDENT, Value: text}
	}
}

// Match string
func (l *Lexer) lexString() token.Token {
	l.next()
	start := l.pos

	for {
		ch := l.next()

		if ch == '"' {
			break
		}

		if ch == 0 {
			panic("Unterminated string")
		}
	}

	value := string(l.input[start : l.pos-1])
	return token.Token{Kind: token.TOK_STR, Value: value}
}

// Match symbol
func (l *Lexer) lexSymbol() token.Token {
	switch ch := l.next(); ch {
	case '{': return token.Token{Kind: token.TOK_BRACKET_OPEN}
	case '}': return token.Token{Kind: token.TOK_BRACKET_CLOSE}
	case '(': return token.Token{Kind: token.TOK_PAREN_OPEN}
	case ')': return token.Token{Kind: token.TOK_PAREN_CLOSE}
	default: panic("Unexpected symbol")
	}
}
