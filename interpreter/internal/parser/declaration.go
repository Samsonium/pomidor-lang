package parser

import "pomidor/intrp/internal/token"

type Declaration interface {
	Node
	isDeclaration()
}

type GardenDecl struct {
	Name string
}

func (*GardenDecl) node()          {}
func (*GardenDecl) isDeclaration() {}

type PluckDecl struct {
	Name string
	Body *Block
}

func (*PluckDecl) node()          {}
func (*PluckDecl) isDeclaration() {}

func (p *Parser) parseGardenDecl() Declaration {
	p.expect(token.TOK_GARDEN)
	name := p.expect(token.TOK_IDENT)

	return &GardenDecl{
		Name: name.Value,
	}
}

func (p *Parser) parsePluckDecl() Declaration {
	p.expect(token.TOK_PLUCK)
	name := p.expect(token.TOK_IDENT)
	p.expect(token.TOK_PAREN_OPEN)
	p.expect(token.TOK_PAREN_CLOSE)

	body := p.parseBlock()

	return &PluckDecl{
		Name: name.Value,
		Body: body,
	}
}

func (p *Parser) parseDeclaration() Declaration {
	switch p.cur.Kind {
	case token.TOK_GARDEN:
		return p.parseGardenDecl()
	case token.TOK_PLUCK:
		return p.parsePluckDecl()
	default:
		return nil
	}
}
