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

// Check for expecting token in lookahead buffer
func (p *Parser) expect(k token.TokenKind) token.Token {
	if p.cur.Kind != k {
		panic(fmt.Sprintf("expected %v, got %v", k, p.cur.Kind))
	}

	t := p.cur
	p.next()

	return t
}

// Instantiate parser
func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}

	p.next()
	p.next()

	return p
}

// Parse source code and return AST
func (p *Parser) ParseProgram() *Program {
	prog := &Program{}

	for p.cur.Kind != token.TOK_EOF {
		decl := p.parseDeclaration()
		if decl == nil {
			panic(fmt.Sprintf("unexpected token: %d", p.cur.Kind))
		}

		prog.Declarations = append(prog.Declarations, decl)
	}

	return prog
}
