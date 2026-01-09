package token

type TokenKind int

const (
	TOK_EOF TokenKind = iota

	// Keywords
	TOK_GARDEN
	TOK_PLUCK

	// Dumb callee (before stdlib impl)
	TOK_SEE

	// Identifier
	TOK_IDENT

	// Symbols
	TOK_BRACKET_OPEN
	TOK_BRACKET_CLOSE
	TOK_PAREN_OPEN
	TOK_PAREN_CLOSE

	// Literals
	TOK_STR
)
