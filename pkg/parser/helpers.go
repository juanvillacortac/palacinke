package parser

import (
	"fmt"

	"github.com/juanvillacortac/palacinke/pkg/token"
)

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) appendError(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	line, col := p.currentToken.Pos.Line, p.currentToken.Pos.Col
	p.errors = append(p.errors, fmt.Sprintf("[%d:%d] %s", line, col, msg))
}

func (p *Parser) peekError(t token.TokenType) {
	p.appendError("Expected next token to be %s, got %s instead", t, p.peekToken.Type)
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	p.appendError("No prefix parse function for %s found", t)
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) currentPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}
	return LOWEST
}
