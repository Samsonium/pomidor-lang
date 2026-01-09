package parser

import (
	"fmt"
	"pomidor/intrp/internal/lexer"
	"pomidor/intrp/internal/token"
)

type Parser struct {
	lexer *lexer.Lexer
	cur   token.Token
	peek  token.Token
}

func (p *Parser) next() {
	p.cur = p.peek
	p.peek = p.lexer.NextToken()
}

// Instantiate parser
func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}

	p.next()
	p.next()

	return p
}

// Check for expecting token in lookahead buffer
func (p *Parser) expect(k token.TokenKind) token.Token {
	if p.cur.Kind != k {
		panic(fmt.Sprintf("expected %v, got %v", k, p.cur.Kind))
	}

	t := p.cur
	p.next()

	return t
}

