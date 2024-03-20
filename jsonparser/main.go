package main

import (
	"log"
	"os"

	"github.com/Farber98/cc-solutions/jsonparser/file"
	"github.com/Farber98/cc-solutions/jsonparser/lexer"
	"github.com/Farber98/cc-solutions/jsonparser/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: jsonparser <file_path>")
	}

	filePath := os.Args[1]

	// Read file contents
	f := &file.DefaultFile{Path: filePath}
	fileContents, err := f.ReadFileContents()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Tokenize the input
	lexer := &lexer.SimpleLexer{}
	tokens := lexer.Lex(string(fileContents))

	// Parse the tokens
	parser := &parser.SimpleParser{}
	isValid := parser.Parse(tokens)

	// Output the result
	if isValid {
		log.Println("Valid JSON")
		os.Exit(0)
	} else {
		log.Println("Invalid JSON")
		os.Exit(1)
	}
}
