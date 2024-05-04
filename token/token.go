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
// All valid JSON lexer tokens
const (
	ILLEGAL = "ILLEGAL" // signifies and unknown or illegal token
	EOF     = "EOF"     // signifies end of file, which lets the parser know to stop

	IDENT = "IDENT" // identifiers: "foo", "name", "age"

	// Operator
	ASSIGN = ":"

	// Delimiter
	COMMA    = ","
	LBRACE   = "{"
	RBRACE   = "}"
	LSQBRACE = "["
	RSQBRACE = "]"

	// Data Types
	STRING  = "STRING"
	NUMBER  = "NUMBER"
	OBJECT  = "OBJECT"
	ARRAY   = "ARRAY"
	BOOLEAN = "BOOLEAN"
	NULL    = "NULL"
)
