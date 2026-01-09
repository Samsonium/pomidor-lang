package main

import (
	"fmt"
	"os"

	"pomidor/intrp/internal/interp"
	"pomidor/intrp/internal/lexer"
	"pomidor/intrp/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: pomidor <file.pmd>")
		os.Exit(2)
	}

	path := os.Args[1]
	src, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		os.Exit(1)
	}

	l := lexer.NewLexer(string(src))
	p := parser.NewParser(l)

	prog := p.ParseProgram()
	in := interp.New(os.Stdout)

	if err := in.Run(prog); err != nil {
		fmt.Fprintln(os.Stderr, "runtime error:", err)
		os.Exit(1)
	}
}
