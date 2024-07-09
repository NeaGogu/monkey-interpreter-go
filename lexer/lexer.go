package lexer

import "NeaGogu/monkey-interpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to position of ::char)
	readPosition int  // current reading position after char (used to peek later in the input)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()

	return l
}

// NOTE: this version does not support Unicode characters
// it would require "ch" to be rune instead of byte and
// l.input[l.readPosition] would not work anymore
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ascii for "NUL"
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition

	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// whitespace has no meaning in monkey
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = l.makeTwoCharToken('=', token.EQ, token.ASSIGN)
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
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '!':
		tok = l.makeTwoCharToken('=', token.NOT_EQ, token.BANG)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)

			// early return is needed because readIdentifier calls readChar repeatedly
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()

			return tok
		}

		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()

	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.readChar()
	}
}

// only supports integers
func (l *Lexer) readNumber() string {
	curPos := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[curPos:l.position]
}

func (l *Lexer) readIdentifier() string {
	curPos := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[curPos:l.position]
}

func isDigit(ch byte) bool {
	return (ch >= '0' && ch <= '9')
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func newToken(t token.TokenType, ch byte) token.Token {
	return token.Token{Type: t, Literal: string(ch)}
}

// peeks the next character. If it matches the secondChToMatch, it advances the char pointer and return a token.Token
// with the matchedToken type and the current + peeked char concatenated for the Literal
// If there is no match, newToken will be called with the current char and fallbackToken as arguments
func (l *Lexer) makeTwoCharToken(secondChToMatch byte, matchedToken token.TokenType, fallbackToken token.TokenType) token.Token {
	if l.peekChar() == secondChToMatch {
		prevChar := l.ch
		l.readChar()

		return token.Token{Type: matchedToken, Literal: string(prevChar) + string(l.ch)}
	} else {
		return newToken(fallbackToken, l.ch)
	}
}
