package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/luka2220/json-parser/lexer"
	"github.com/luka2220/json-parser/token"
)

// NOTE:
// Runs the json file tokenizer
func ExecuteTokenizer(filePath string) {
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

	fileContentBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(fileContentBytes)
	if err != nil {
		log.Fatalf("Error reading the files contents: %v\n", err)
	}

	fileStr := bytesToStr(fileContentBytes)
	lexer := lexer.New(fileStr)

	for {
		tok := lexer.NextToken()
		out := fmt.Sprintf("> Token: %v > Literal: %v", tok.Type, tok.Literal)

		if tok.Type == token.EOF {
			fmt.Println(out)
			break
		}
		out = fmt.Sprintf("> Token: %v > Literal: %v", tok.Type, tok.Literal)
		fmt.Println(out)
	}
}

// NOTE:
// Helper function for converting a file of bytes to string content
func bytesToStr(f []byte) string {
	var s string

	for i := range f {
		s += string(f[i])
	}

	return s
}
