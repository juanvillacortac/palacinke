package parser_test

import (
	"testing"

	"github.com/juandroid007/palacinke/pkg/ast"
	"github.com/juandroid007/palacinke/pkg/lexer"
	"github.com/juandroid007/palacinke/pkg/parser"
)

func TestReturnStatements(t *testing.T) {
	input := `return 10;
return 5;
return 838383;`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Got: %d",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. Got: %T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf(`returnStmt.TokenLiteral not "return". Got: %q`,
				returnStmt.TokenLiteral())
		}
	}
}
