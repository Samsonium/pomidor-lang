package parser

import "pomidor/intrp/internal/token"

type Statement interface {
	Node
	isStatement()
}

func (p *Parser) parseStatement() Statement {
	switch p.cur.Kind {
	case token.TOK_SEE:
		return p.parseCallStatement()
	default:
		return nil
	}
}
