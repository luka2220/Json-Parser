package token

// NOTE:
// Token data structure is defined below
// Type (:TokenType: string) => holds the token type to distinguish between strings, numbers, right brackets, semi-colons etc.
// Literal (:string) => holds the literal string value of the token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// NOTE:
// Below are all valid lexer tokens for JSON
const (
	ILLEGAL = "ILLEGAL" // signifies and unknown or illegal token
	EOF     = "EOF"     // signifies end of file, which lets the parser know to stop

	// Operator
	ASSIGN = ":"

	// Delimiter
	COMMA    = ","
	LBRACE   = "{"
	RBRACE   = "}"
	LSQBRACE = "["
	RSQBRACE = "]"

	// Types
	STRING = "STRING"
	NUMBER = "NUMBER"
	OBJECT = "OBJECT"
	ARRAY  = "ARRAY"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	NULL   = "NULL"
)