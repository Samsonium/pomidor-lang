package lexer

import (
	"pomidor/intrp/internal/token"
	"testing"
)

func TestLexerBasic(t *testing.T) {
	const source = `
		garden HelloPomidor

		pluck main_tomato() {
			see "Hello, Pomidor!"
		}
		`

	var expected = []token.Token{
		{Kind: token.TOK_GARDEN, Value: ""},
		{Kind: token.TOK_IDENT, Value: "HelloPomidor"},
		{Kind: token.TOK_PLUCK, Value: ""},
		{Kind: token.TOK_IDENT, Value: "main_tomato"},
		{Kind: token.TOK_PAREN_OPEN, Value: ""},
		{Kind: token.TOK_PAREN_CLOSE, Value: ""},
		{Kind: token.TOK_BRACKET_OPEN, Value: ""},
		{Kind: token.TOK_SEE, Value: ""},
		{Kind: token.TOK_STR, Value: "Hello, Pomidor!"},
		{Kind: token.TOK_BRACKET_CLOSE, Value: ""},
		{Kind: token.TOK_EOF, Value: ""},
	}

	lex := NewLexer(source)

	for i, exp := range expected {
		tok := lex.NextToken()

		if tok.Kind != exp.Kind || tok.Value != exp.Value {
			t.Fatalf("token %d: expected %+v, got %+v", i, exp, tok)
		}
	}
}

func TestLexerComments(t *testing.T) {
	source := `
	garden A // Comment
	// Full line comment
	pluck B
	`

	var expected = []token.TokenKind{
		token.TOK_GARDEN,
		token.TOK_IDENT,
		token.TOK_PLUCK,
		token.TOK_IDENT,
	}

	lex := NewLexer(source)

	for i, k := range expected {
		tok := lex.NextToken()

		if tok.Kind != k {
			t.Fatalf("token %d: expected %v, got %v", i, k, tok.Kind)
		}
	}
}

func TestLexerString(t *testing.T) {
	source := `"Some string"`
	lex := NewLexer(source)

	tok := lex.NextToken()
	if tok.Kind != token.TOK_STR {
		t.Errorf("expected type TOK_STR (%d), got %d", token.TOK_STR, tok.Kind)
	}

	test_str := "Some string"
	if tok.Value != test_str {
		t.Errorf("expected string value '%v', got '%v'", test_str, tok.Value)
	}
}

func TestLexerEOF(t *testing.T) {
	lex := NewLexer("see")

	lex.NextToken() // skip "see"
	tok := lex.NextToken()

	if tok.Kind != token.TOK_EOF {
		t.Fatalf("Expected last token to be EOF (%d), got %d", token.TOK_EOF, tok.Kind)
	}
}
