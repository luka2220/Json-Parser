package lexer

import (
	"github.com/luka2220/json-parser/token"
)

// NOTE:
// Data structure for creating lexer tokens for json file
// input (:string) => string to generate tokens from
// position (:int) => current position in the input (points to the current char)
// readPosition (:int) => current reading position in the input (after the current char)
// ch (:byte) => current character under exmination in the input
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// NOTE:
// Initialize a new Lexer data structure
// Params: input (:string) => json file text as string
// Returns: *Lexer
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

// NOTE:
// Generates the next character in the input and advances the position in the input field
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// NOTE:
// Creates a token based on the current character under examination, and advances
// the position in the input.
// Returns: token.Token => token type based on the current character
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case ':':
		tok = newToken(token.ASSIGN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LSQBRACE, l.ch)
	case ']':
		tok = newToken(token.RSQBRACE, l.ch)
	case '"':
		tok = newToken(token.STRING, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.IDENT
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readDigit()
			tok.Type = token.IDENT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

// NOTE:
// Helper function to skip whitespaces, tabs, newline, and returns
// Addvances to the next char in the input while any are found
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// NOTE:
// Reads the identifier string and advances the lexers position until it encounters a non-letter byte
// Returns: string
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// NOTE:
// Reads the identifier digit and advances the lexers position until it encounters a non-digit byte
// Returns: string
func (l *Lexer) readDigit() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// NOTE:
// Helper function to determine if the current character is a letter or number which are valid in strings
// Params: ch (:byte)
// Returns: bool
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || isDigit(ch)
}

// NOTE:
// Helper function to determine if the current character is a number
// Patams: ch (:byte)
// Returns: bool
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// NOTE:
// Helper function used to create tokens
// Params: tokenType (:token.tokenType), ch (:byte)
// Returns: token.Token => new token created
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
