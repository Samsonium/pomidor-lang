package parser

type Expression interface {
	Node
	isExpression()
}

type StringLiteral struct {
	Value string
}

func (*StringLiteral) node()         {}
func (*StringLiteral) isExpression() {}
