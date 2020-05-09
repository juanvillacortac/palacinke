package parser_test

import (
	"testing"

	"github.com/juandroid007/palacinke/pkg/ast"
	"github.com/juandroid007/palacinke/pkg/lexer"
	"github.com/juandroid007/palacinke/pkg/parser"
)

func TestFunctionLiteralParsing(t *testing.T) {
	input := `fn(x, y) { x + y; }`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Body doesn't contain %d statements. Got: %d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] isn't ast.ExpressionStatement. Got: %T", program.Statements[0])
	}

	fn, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt.Expression isn't ast.FunctionLiteral. Got: %T\n", stmt.Expression)
	}

	if len(fn.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. Want 2, got: %d", len(fn.Parameters))
	}

	testLiteralExpression(t, fn.Parameters[0], "x")
	testLiteralExpression(t, fn.Parameters[1], "y")

	body, ok := fn.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("function body stmt isn't ast.ExpressionStatement. Got: %T", fn.Body.Statements[0])
	}

	testInfixExpression(t, body.Expression, "x", "+", "y")
}
