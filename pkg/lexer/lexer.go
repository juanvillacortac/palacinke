package lexer

import (
	"encoding/hex"
	"strings"
	"text/scanner"

	"github.com/juanvillacortac/palacinke/pkg/token"
)

type Lexer struct {
	s    scanner.Scanner
	curr rune
}

func New(input string) *Lexer {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	l := &Lexer{s: s}
	l.readRune()
	return l
}

func (l *Lexer) readRune() {
	l.curr = l.s.Scan()
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.curr {
	case '=':
		tok = l.eitherToken('=', token.EQ, token.ASSIGN)
	case '!':
		tok = l.eitherToken('=', token.NOT_EQ, token.BANG)
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
	case '%':
		tok = l.newToken(token.MOD)
	case '<':
		tok = l.eitherToken('=', token.LEQT, token.LT)
	case '>':
		tok = l.eitherToken('=', token.GEQT, token.GT)
	case ';':
		tok = l.newToken(token.SEMICOLON)
	case ',':
		tok = l.newToken(token.COMMA)
	case ':':
		tok = l.newToken(token.COLON)
	case '(':
		tok = l.newToken(token.LPAREN)
	case ')':
		tok = l.newToken(token.RPAREN)
	case '{':
		tok = l.newToken(token.LBRACE)
	case '}':
		tok = l.newToken(token.RBRACE)
	case '[':
		tok = l.newToken(token.LBRACKET)
	case ']':
		tok = l.newToken(token.RBRACKET)
	case scanner.Ident:
		p := l.s.Pos()
		lit := l.s.TokenText()
		col := p.Column
		if la := l.s.Peek(); la == '?' || la == '!' {
			l.readRune()
			lit += l.s.TokenText()
			col += 1
		}
		tok = token.Token{
			Type:    token.LookupIdent(lit),
			Literal: lit,
			Pos: token.TokenPos{
				Line: p.Line,
				Col:  p.Column,
			},
		}
	case scanner.Int:
		p := l.s.Pos()
		lit := l.s.TokenText()
		tok = token.Token{
			Type:    token.INT,
			Literal: lit,
			Pos: token.TokenPos{
				Line: p.Line,
				Col:  p.Column,
			},
		}
	case scanner.String:
		p := l.s.Pos()
		txt := l.s.TokenText()
		str := txt[1 : len(txt)-1]
		lit := &strings.Builder{}
		for i := 0; i < len(str); i++ {
			ch := str[i]
			if ch == '\\' {
				switch str[i+1] {
				case '"':
					lit.WriteByte('"')
				case 'n':
					lit.WriteByte('\n')
				case 'r':
					lit.WriteByte('\r')
				case 't':
					lit.WriteByte('\t')
				case '\\':
					lit.WriteByte('\\')
				case 'x':
					// Skip over the the '\\', 'x' and the next two bytes (hex)
					i += 3
					src := string([]byte{str[i-1], ch})
					dst, err := hex.DecodeString(src)
					if err != nil {
						return l.newToken(token.ILLEGAL)
					}
					lit.Write(dst)
					continue
				}

				// Skip over the '\\' and the matched single escape char
				i++
				continue
			} else {
				if ch == '"' || ch == 0 {
					break
				}
			}
			lit.WriteByte(ch)
		}
		tok = token.Token{
			Type:    token.STRING,
			Literal: lit.String(),
			Pos: token.TokenPos{
				Line: p.Line,
				Col:  p.Column - 2,
			},
		}
	case scanner.EOF:
		p := l.s.Pos()
		tok = token.Token{
			Type:    token.EOF,
			Literal: "",
			Pos: token.TokenPos{
				Line: p.Line,
				Col:  p.Column,
			},
		}
	default:
		tok = l.newToken(token.ILLEGAL)
	}

	l.readRune()
	return tok
}

func (l *Lexer) newToken(tokenType token.TokenType) token.Token {
	p := l.s.Pos()
	lit := l.s.TokenText()
	return token.Token{
		Type:    tokenType,
		Literal: lit,
		Pos: token.TokenPos{
			Line: p.Line,
			Col:  p.Column,
		},
	}
}

func (l *Lexer) eitherToken(
	lookAhead rune,
	optionToken, alternativeToken token.TokenType,
) token.Token {
	p := l.s.Pos()
	lit := l.s.TokenText()
	col := p.Column
	if l.s.Peek() == lookAhead {
		l.readRune()
		lit += l.s.TokenText()
		col += 1
		return token.Token{
			Type:    optionToken,
			Literal: lit,
			Pos: token.TokenPos{
				Line: p.Line,
				Col:  p.Column,
			},
		}
	} else {
		return token.Token{
			Type:    alternativeToken,
			Literal: lit,
			Pos: token.TokenPos{
				Line: p.Line,
				Col:  p.Column,
			},
		}
	}
}
