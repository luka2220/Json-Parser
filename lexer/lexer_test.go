package lexer

import (
	"os"
	"testing"

	"github.com/luka2220/json-parser/token"
)

// NOTE:
// Read the test JSON files and returns whether the files are valid JSON or not
func TestNextToken(t *testing.T) {
	input, err := os.ReadFile("../test-data/step2/valid2.json")
	if err != nil {
		t.Fatalf("Error opening test data file\n%v", err)
	}

	var instr string

	for i := range input {
		instr += string(input[i])
	}

	t.Logf("Input test as string\n%s", instr)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "\""},
		{token.IDENT, "key"},
		{token.STRING, "\""},
		{token.ASSIGN, ":"},
		{token.STRING, "\""},
		{token.IDENT, "value"},
		{token.STRING, "\""},
		{token.COMMA, ","},
		{token.STRING, "\""},
		{token.IDENT, "key2"},
		{token.STRING, "\""},
		{token.ASSIGN, ":"},
		{token.STRING, "\""},
		{token.IDENT, "value"},
		{token.STRING, "\""},
		{token.RBRACE, "}"},
	}

	lexer := New(instr)

	for i, tt := range tests {
		tok := lexer.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q\nexpected=%q actucal=%q",
				i, tt.expectedType, tok.Type, tt.expectedLiteral, tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
