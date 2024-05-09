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
	NULL    = "NULL"
	BOOLEAN = "BOOLEAN"
)

// NOTE:
// Map of string (key), TokenType (value) to represent keyword data types
var keywords = map[string]TokenType{
	"true":  BOOLEAN,
	"false": BOOLEAN,
	"null":  NULL,
}

// NOTE:
// Checks the keywords table to see if an identifier is a keyword or not
// Params: ident (:string)
// Returns: TokenType
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
