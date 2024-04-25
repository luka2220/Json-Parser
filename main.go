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
}

func getTokens() *Tokens {
	return &Tokens{
		"{",
		"}",
		":",
		"\"",
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

	lexeme := make([]string, fileInfo.Size())

	for i := range fileContentBytes {
		char := string(fileContentBytes[i])
		tokens := getTokens()

		switch char {
		case tokens.openBracket:
			lexeme = append(lexeme, tokens.openBracket)
		case tokens.string:
			lexeme = append(lexeme, tokens.string)
		case tokens.sperator:
			lexeme = append(lexeme, tokens.sperator)
		case tokens.closeBracket:
			lexeme = append(lexeme, tokens.closeBracket)
		}
	}

	log.Printf("Test lexeme = %v", lexeme)
	log.Printf("lexeme capacity = %d\nlexeme length = %d", cap(lexeme), len(lexeme))
}
