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
// Generates the next character and advances the position in the input field
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
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()

	return tok
}

// NOTE:
// Helper function used to create tokens
// Params: tokenType (:token.tokenType), ch (:byte)
// Returns: token.Token => new token created
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
