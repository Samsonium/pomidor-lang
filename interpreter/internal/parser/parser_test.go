package parser

import (
	"pomidor/intrp/internal/lexer"
	"testing"
)

func TestPraserGolden(t *testing.T) {
	source := `
	garden Tomatoes

	pluck main_pomidor() { see "Hello, Pomidor!" }
	`

	l := lexer.NewLexer(source)
	p := NewParser(l)

	prog := p.parseProgram()
	if len(prog.Declarations) != 2 {
		t.Fatalf("expected 2 declarations in program, got %d", len(prog.Declarations))
	}
}
