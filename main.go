package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// TODO:
// - Read a json file from stdin
// - Scan the JSON source code for tokens
// - Inintalize JSON tokens => , { } : " "
// - A key can only be a string
// - JSON data types consist of => string, number, null, boolean, object, array

type Tokens struct {
	openBracket  string
	closeBracket string
	sperator     string
	string       string
	comma        string
}

func getTokens() *Tokens {
	return &Tokens{
		"{",
		"}",
		":",
		"\"",
		",",
	}
}

func generateLexeme(fileContent []byte) []string {
	var lexeme []string
	tokens := getTokens()

	for i := range fileContent {
		jsonChar := string(fileContent[i])

		switch jsonChar {
		case tokens.openBracket:
			lexeme = append(lexeme, tokens.openBracket)
		// handle string token with characters inside
		case tokens.string:
			lexeme = append(lexeme, tokens.string)
		case tokens.sperator:
			lexeme = append(lexeme, tokens.sperator)
		case tokens.comma:
			lexeme = append(lexeme, tokens.comma)
		case tokens.closeBracket:
			lexeme = append(lexeme, tokens.closeBracket)
		}
	}

	return lexeme
}

// NOTE:
// Params: lexeme []string => lexme to be tested
// Returns: void
// Operation: dislpays the passed in lexeme for debugging, will not throw an error for invalid lexemes
func testLexeme(lexeme []string) {
	for i := range lexeme {
		cToken := lexeme[i]
		fmt.Printf("%s ", cToken)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a JSON file path")
	in, _ := reader.ReadString('\n')

	filePath := strings.Trim(in, "\n")

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file. Ensure correct path in passed: %v\n", err)
		os.Exit(400)
	}

	defer file.Close()

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("Error getting file info: %v\n", err)
	}

	log.Printf("Length of file = %d", fileInfo.Size())

	fileContentBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(fileContentBytes)
	if err != nil {
		log.Fatalf("Error reading the files contents: %v\n", err)
	}

	lexeme := generateLexeme(fileContentBytes)
	log.Printf("lexeme capacity = %d\nlexeme length = %d", cap(lexeme), len(lexeme))

	testLexeme(lexeme)
}
