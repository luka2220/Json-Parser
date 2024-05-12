package lexer

import (
	"os"
	"testing"

	"github.com/luka2220/json-parser/token"
)

// NOTE:
// Read the test JSON files and returns whether the files are valid JSON or not
func TestNextToken(t *testing.T) {
	json1, err := os.ReadFile("../test-data/step2/valid2.json")
	json2, err := os.ReadFile("../test-data/step3/valid.json")
	json3, err := os.ReadFile("../test-data/step4/valid.json")
	json4, err := os.ReadFile("../test-data/step4/valid2.json")

	if err != nil {
		t.Fatalf("Error opening test data file\n%v", err)
	}

	var instr1 string
	var instr2 string
	var instr3 string
	var instr4 string

	for i := range json1 {
		instr1 += string(json1[i])
	}

	for i := range json2 {
		instr2 += string(json2[i])
	}

	for i := range json3 {
		instr3 += string(json3[i])
	}

	for i := range json4 {
		instr4 += string(json4[i])
	}

	t.Logf("JSON Test 1\n%s", instr1)
	t.Logf("JSON Test 2\n%s", instr2)
	t.Logf("JSON Test 3\n%s", instr3)
	t.Logf("JSON Test 4\n%s", instr4)

	test1 := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "key"},
		{token.ASSIGN, ":"},
		{token.STRING, "value"},
		{token.COMMA, ","},
		{token.STRING, "key2"},
		{token.ASSIGN, ":"},
		{token.STRING, "value"},
		{token.RBRACE, "}"},
	}

	test2 := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "key1"},
		{token.ASSIGN, ":"},
		{token.TRUE, "true"},
		{token.COMMA, ","},
		{token.STRING, "key2"},
		{token.ASSIGN, ":"},
		{token.FALSE, "false"},
		{token.COMMA, ","},
		{token.STRING, "key3"},
		{token.ASSIGN, ":"},
		{token.NULL, "null"},
		{token.COMMA, ","},
		{token.STRING, "key4"},
		{token.ASSIGN, ":"},
		{token.STRING, "value"},
		{token.COMMA, ","},
		{token.STRING, "key5"},
		{token.ASSIGN, ":"},
		{token.NUMBER, "101"},
		{token.RBRACE, "}"},
	}

	test3 := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "key"},
		{token.ASSIGN, ":"},
		{token.STRING, "value"},
		{token.COMMA, ","},
		{token.STRING, "key-n"},
		{token.ASSIGN, ":"},
		{token.NUMBER, "101"},
		{token.COMMA, ","},
		{token.STRING, "key-o"},
		{token.ASSIGN, ":"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.STRING, "key-l"},
		{token.ASSIGN, ":"},
		{token.LSQBRACE, "["},
		{token.RSQBRACE, "]"},
		{token.RBRACE, "}"},
	}

	test4 := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "key"},
		{token.ASSIGN, ":"},
		{token.STRING, "value"},
		{token.COMMA, ","},
		{token.STRING, "key-n"},
		{token.ASSIGN, ":"},
		{token.NUMBER, "101"},
		{token.COMMA, ","},
		{token.STRING, "key-o"},
		{token.ASSIGN, ":"},
		{token.LBRACE, "{"},
		{token.STRING, "inner key"},
		{token.ASSIGN, ":"},
		{token.STRING, "inner value"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.STRING, "key-l"},
		{token.ASSIGN, ":"},
		{token.LSQBRACE, "["},
		{token.STRING, "list value"},
		{token.RSQBRACE, "]"},
		{token.RBRACE, "}"},
	}

	lexer1 := New(instr1)
	lexer2 := New(instr2)
	lexer3 := New(instr3)
	lexer4 := New(instr4)

	for i, tt := range test1 {
		tok := lexer1.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q\nexpected=%q actual=%q",
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

	for i, tt := range test3 {
		tok := lexer3.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q\nexpected=%q actucal=%q",
				i, tt.expectedType, tok.Type, tt.expectedLiteral, tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

	for i, tt := range test4 {
		tok := lexer4.NextToken()

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
