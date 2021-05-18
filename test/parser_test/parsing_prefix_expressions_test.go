package parser_test

import (
	"testing"

	"github.com/juanvillacortac/palacinke/pkg/ast"
	"github.com/juanvillacortac/palacinke/pkg/lexer"
	"github.com/juanvillacortac/palacinke/pkg/parser"
)

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue interface{}
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"!true;", "!", true},
		{"!false;", "!", false},
	}
	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf(
				"program.Statements does not contain %d statements. Got: %d\n",
				1, len(program.Statements),
			)
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf(
				"program.Statements[0] is not ast.ExpressionStatement. Got: %T",
				program.Statements[0],
			)
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf(
				"stmt is not ast.PrefixExpression. Got: %T",
				stmt.Expression,
			)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. Got: %s", tt.operator, exp.Operator)
		}
		if !testLiteralExpression(t, exp.Right, tt.integerValue) {
			return
		}
	}
}
