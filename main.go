package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

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

	if !isJSON(fileInfo.Name()) {
		log.Printf("Enter a json file")
		os.Exit(400)
	}

	fileContentBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(fileContentBytes)
	if err != nil {
		log.Fatalf("Error reading the files contents: %v\n", err)
	}

	fileStr := bytesToStr(fileContentBytes)

	fmt.Printf("File info:\nsize=%dbytes\nname=%s\ncontent(bytes)=%v\ncontent(string)=%s", fileInfo.Size(), fileInfo.Name(), fileContentBytes, fileStr)
}

// NOTE:
// Convert file from bytes into a string
// Params: f (:[]byte)
// Returns: string
func bytesToStr(f []byte) string {
	var s string

	for i := range f {
		s += string(f[i])
	}

	return s
}

// NOTE:
// Check if a file type is json or not
// Params: fname (:string)
// Returns: bool
func isJSON(fname string) bool {
	return path.Ext(fname) == ".json"
}
