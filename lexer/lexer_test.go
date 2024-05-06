package lexer

import (
	"os"
	"testing"

	"github.com/luka2220/json-parser/token"
)

// NOTE:
// Read the test JSON files and returns whether the files are valid JSON or not
func TestNextToken(t *testing.T) {
	input1, err := os.ReadFile("../test-data/step2/valid2.json")
	input2, err := os.ReadFile("../test-data/step3/valid.json")

	if err != nil {
		t.Fatalf("Error opening test data file\n%v", err)
	}

	var instr1 string
	var instr2 string

	for i := range input1 {
		instr1 += string(input1[i])
	}

	for j := range input2 {
		instr2 += string(input2[j])
	}

	t.Logf("Input test 1 as string\n%s", instr1)
	t.Logf("Input test 2 as string\n%s", instr2)

	test1 := []struct {
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

	test2 := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "\""},
		{token.IDENT, "key1"},
		{token.STRING, "\""},
		{token.ASSIGN, ":"},
		{token.BOOLEAN, "true"},
		{token.COMMA, ","},
		{token.STRING, "\""},
		{token.IDENT, "key2"},
		{token.STRING, "\""},
		{token.ASSIGN, ":"},
		{token.BOOLEAN, "false"},
		{token.COMMA, ","},
		{token.STRING, "\""},
		{token.IDENT, "key3"},
		{token.STRING, "\""},
		{token.ASSIGN, ":"},
		{token.NULL, "null"},
		{token.COMMA, ","},
		{token.STRING, "\""},
		{token.IDENT, "key4"},
		{token.STRING, "\""},
		{token.ASSIGN, ":"},
		{token.STRING, "\""},
		{token.IDENT, "value"},
		{token.STRING, "\""},
		{token.COMMA, ","},
		{token.STRING, "\""},
		{token.IDENT, "key5"},
		{token.STRING, "\""},
		{token.ASSIGN, ":"},
		{token.NUMBER, "101"},
		{token.RBRACE, "}"},
	}

	lexer1 := New(instr1)
	lexer2 := New(instr2)

	for i, tt := range test1 {
		tok := lexer1.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q\nexpected=%q actucal=%q",
				i, tt.expectedType, tok.Type, tt.expectedLiteral, tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

	for i, tt := range test2 {
		tok := lexer2.NextToken()

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
