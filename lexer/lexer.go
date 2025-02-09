package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
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

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = l.handleAssign()
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = l.handleBang()
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = l.handleLessThen()
	case '>':
		tok = l.handleGreaterThen()
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) peek() byte {

	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func isLetter(ch byte) bool {
	// Check if the byte value falls within the range of alphanumeric characters
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {

	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]

}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) isEOF() bool {
	return l.readPosition+1 > len(l.input)
}

func (l *Lexer) handleAssign() token.Token {
	if l.peek() == '=' {
		ch := l.ch
		l.readChar()
		literal := string(ch) + string(l.ch)
		return token.Token{Type: token.EQ, Literal: literal}
	} else {
		return newToken(token.ASSIGN, l.ch)
	}
}

func (l *Lexer) handleBang() token.Token {
	if l.peek() == '=' {
		ch := l.ch
		l.readChar()
		literal := string(ch) + string(l.ch)
		return token.Token{Type: token.NOT_EQ, Literal: literal}
	} else {
		return newToken(token.BANG, l.ch)
	}
}

func (l *Lexer) handleLessThen() token.Token {
	if l.peek() == '=' {
		ch := l.ch
		l.readChar()
		literal := string(ch) + string(l.ch)
		return token.Token{Type: token.LT_OR_EQ, Literal: literal}
	} else {
		return newToken(token.LT, l.ch)
	}
}

func (l *Lexer) handleGreaterThen() token.Token {
	if l.peek() == '=' {
		ch := l.ch
		l.readChar()
		literal := string(ch) + string(l.ch)
		return token.Token{Type: token.GT_OR_EQ, Literal: literal}
	} else {
		return newToken(token.GT, l.ch)
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
