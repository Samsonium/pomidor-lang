package interp

import (
	"fmt"
	"io"
	"pomidor/intrp/internal/parser"
)

type Interpreter struct {
	Out io.Writer

	gardenName string
	funcs      map[string]*parser.PluckDecl
}

func New(out io.Writer) *Interpreter {
	return &Interpreter{
		Out:   out,
		funcs: make(map[string]*parser.PluckDecl),
	}
}

func (in *Interpreter) Run(prog *parser.Program) error {
	if prog == nil {
		return fmt.Errorf("nil program")
	}

	for _, d := range prog.Declarations {
		switch decl := d.(type) {
		case *parser.GardenDecl:
			in.gardenName = decl.Name
		case *parser.PluckDecl:
			in.funcs[decl.Name] = decl
		default:
			return fmt.Errorf("unknown declaration type: %T", d)
		}
	}

	mainFn, ok := in.funcs["main_pomidor"]
	if !ok {
		return fmt.Errorf("entry point pluck main_pomidor() not defined")
	}

	return in.execBlock(mainFn.Body)
}

func (in *Interpreter) execBlock(b *parser.Block) error {
	if b == nil {
		return fmt.Errorf("nil block")
	}

	for _, s := range b.Statements {
		if err := in.execStatement(s); err != nil {
			return err
		}
	}

	return nil
}

func (in *Interpreter) execStatement(s parser.Statement) error {
	switch st := s.(type) {
	case *parser.CallStatement:
		return in.execCall(st)
	default:
		return fmt.Errorf("unknown statement type: %T", s)
	}
}

func (in *Interpreter) execCall(c *parser.CallStatement) error {
	switch c.Name {
	case "see":
		for i, arg := range c.Args {
			v, err := in.evalExpr(arg)
			if err != nil {
				return err
			}

			if i > 0 {
				_, _ = io.WriteString(in.Out, " ")
			}

			_, _ = io.WriteString(in.Out, v)
		}

		_, _ = io.WriteString(in.Out, "\n")
		return nil

	default:
		return fmt.Errorf("unknown function call: %s", c.Name)
	}
}

func (in *Interpreter) evalExpr(e parser.Expression) (string, error) {
	switch ex := e.(type) {
	case *parser.StringLiteral:
		return ex.Value, nil
	default:
		return "", fmt.Errorf("unknown expression type: %T", e)
	}
}
