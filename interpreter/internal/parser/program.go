package parser

import (
	"fmt"
	"pomidor/intrp/internal/token"
)

type Program struct {
	Declarations []Declaration
}

func (p *Parser) parseProgram() *Program {
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
