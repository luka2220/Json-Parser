package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// NOTE:
// Runs the json file tokenizer
func ExecuteTokenizer() {
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

	fileContentBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(fileContentBytes)
	if err != nil {
		log.Fatalf("Error reading the files contents: %v\n", err)
	}

	fileStr := bytesToStr(fileContentBytes)

	fmt.Printf("File info:\nsize=%dbytes\nname=%s\ncontent(bytes)=%v\ncontent(string)=%s",
		fileInfo.Size(), fileInfo.Name(), fileContentBytes, fileStr)
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
