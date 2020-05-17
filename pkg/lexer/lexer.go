package lexer

import "github.com/juandroid007/palacinke/pkg/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = l.newTwoCharToken(token.EQ)
		} else {
			tok = l.newToken(token.ASSIGN)
		}
	case '!':
		if l.peekChar() == '=' {
			tok = l.newTwoCharToken(token.NOT_EQ)
		} else {
			tok = l.newToken(token.BANG)
		}
	case '+':
		tok = l.newToken(token.PLUS)
	case '-':
		tok = l.newToken(token.MINUS)
	case '/':
		tok = l.newToken(token.SLASH)
	case '*':
		tok = l.newToken(token.ASTERISK)
	case '^':
		tok = l.newToken(token.POW)
	case '<':
		if l.peekChar() == '=' {
			tok = l.newTwoCharToken(token.LEQT)
		} else {
			tok = l.newToken(token.LT)
		}
	case '>':
		if l.peekChar() == '=' {
			tok = l.newTwoCharToken(token.GEQT)
		} else {
			tok = l.newToken(token.GT)
		}
	case ';':
		tok = l.newToken(token.SEMICOLON)
	case ',':
		tok = l.newToken(token.COMMA)
	case '(':
		tok = l.newToken(token.LPAREN)
	case ')':
		tok = l.newToken(token.RPAREN)
	case '{':
		tok = l.newToken(token.LBRACE)
	case '}':
		tok = l.newToken(token.RBRACE)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = l.newToken(token.ILLEGAL)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) newToken(tokenType token.TokenType) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(l.ch),
	}
}

func (l *Lexer) newTwoCharToken(tokenType token.TokenType) token.Token {
	ch := l.ch
	l.readChar()
	return token.Token{
		Type:    tokenType,
		Literal: string(ch) + string(l.ch),
	}
}
