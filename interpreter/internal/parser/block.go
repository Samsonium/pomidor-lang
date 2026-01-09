package parser

import "pomidor/intrp/internal/token"

type Block struct {
	Statements []Statement
}

func (*Block) node() {}

func (p *Parser) parseBlock() *Block {
	p.expect(token.TOK_BRACKET_OPEN)

	var stmts []Statement
	for p.cur.Kind != token.TOK_BRACKET_CLOSE {
		stmt := p.parseStatement()
		if stmt == nil {
			panic("Invalid statement")
		}

		stmts = append(stmts, stmt)
	}

	p.expect(token.TOK_BRACKET_CLOSE)

	return &Block{
		Statements: stmts,
	}
}
