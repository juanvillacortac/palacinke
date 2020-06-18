package lexer_test

import (
	"testing"

	"github.com/juandroid007/palacinke/pkg/lexer"
	"github.com/juandroid007/palacinke/pkg/token"
)

func TestNextToken(t *testing.T) {
	input := `
fn let return true false
123 =
;:,(){}[]
"foobar"
世界
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.FUNCTION, "fn"},
		{token.LET, "let"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.FALSE, "false"},
		{token.INT, "123"},
		{token.ASSIGN, "="},
		{token.SEMICOLON, ";"},
		{token.COLON, ":"},
		{token.COMMA, ","},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.STRING, "foobar"},
		{token.IDENT, "世界"},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - TokenType wrong.\n\t-> Expected: %q - Got: %q",
				i, tt.expectedType, tok.Type,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] - Literal wrong.\n\t-> Expected: %q - Got: %q",
				i, tt.expectedLiteral, tok.Literal,
			)
		}
	}
}
