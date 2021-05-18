package parser_test

import (
	"testing"

	"github.com/juanvillacortac/palacinke/pkg/ast"
	"github.com/juanvillacortac/palacinke/pkg/lexer"
	"github.com/juanvillacortac/palacinke/pkg/parser"
)

func TestIntegerLiteral(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. Got: %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. Got: %T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. Got: %T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. Got: %d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. Got: %s", "5",
			literal.TokenLiteral())
	}
}
