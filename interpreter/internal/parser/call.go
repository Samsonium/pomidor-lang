package parser

import "pomidor/intrp/internal/token"

type CallStatement struct {
	Name string
	Args []Expression
}

func (*CallStatement) node()        {}
func (*CallStatement) isStatement() {}

func (p *Parser) parseCallStatement() Statement {
	p.expect(token.TOK_SEE)

	var args []Expression
	if p.cur.Kind == token.TOK_STR {
		args = append(args, &StringLiteral{
			Value: p.cur.Value,
		})
		p.next()
	}

	return &CallStatement{
		Name: "see",
		Args: args,
	}
}
