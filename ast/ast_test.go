package ast

import (
	"github.com/dpakach/pongo/token"
	"testing"
)

func TestString(t *testing.T) {
	programString := "let myVar = anotherVar;"
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != programString {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
