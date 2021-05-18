package parser_test

import (
	"testing"

	"github.com/juanvillacortac/palacinke/pkg/ast"
	"github.com/juanvillacortac/palacinke/pkg/lexer"
	"github.com/juanvillacortac/palacinke/pkg/parser"
)

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input           string
		expectedBoolean bool
	}{
		{"true;", true},
		{"false;", false},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program hasn't enough statements. Got: %d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] isn't ast.ExpressionStatement. Got: %T",
				program.Statements[0])
		}

		boolean, ok := stmt.Expression.(*ast.Boolean)
		if !ok {
			t.Fatalf("exp not *ast.Boolean. Got: %T", stmt.Expression)
		}
		if boolean.Value != tt.expectedBoolean {
			t.Errorf("boolean.Value not %t. Got: %t", tt.expectedBoolean,
				boolean.Value)
		}
	}
}
