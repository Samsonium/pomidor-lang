package lexer

import "unicode"

// Match whitespace
func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// Match digit
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// Match letter
func isLetter(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}
