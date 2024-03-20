package lexer

import (
	"strings"
	"testing"

	"github.com/Farber98/cc-solutions/jsonparser/file"
)

func TestLex(t *testing.T) {
	tests := []struct {
		name                 string
		filePath             string
		expectedTokenization string
	}{
		{
			name:                 "Valid JSON",
			filePath:             "../tests/step1/valid.json",
			expectedTokenization: "{}",
		},
		{
			name:                 "Invalid JSON",
			filePath:             "../tests/step1/invalid.json",
			expectedTokenization: "",
		},
		{
			name:                 "Valid JSON 2_1",
			filePath:             "../tests/step2/valid.json",
			expectedTokenization: "{\"key\":\"value\"}",
		},
		{
			name:                 "Valid JSON 2_2",
			filePath:             "../tests/step2/valid2.json",
			expectedTokenization: "{\"key\":\"value\",\"key2\":\"value\"}",
		},
		{
			name:                 "Invalid JSON 2_1",
			filePath:             "../tests/step2/invalid.json",
			expectedTokenization: "{\"key\":\"value\",}",
		},
		{
			name:                 "Invalid JSON 2_2",
			filePath:             "../tests/step2/invalid2.json",
			expectedTokenization: "{\"key\":\"value\",key2:\"value\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read string from test files
			f := &file.DefaultFile{Path: tt.filePath}

			// Shouldn't err when reading file content
			fileContents, err := f.ReadFileContents()
			if err != nil {
				t.Fatalf("Error reading file contents: %v", err)
			}

			// Define our lexer
			lexer := &SimpleLexer{}

			// Lex our string
			actualTokenization := lexer.Lex(string(fileContents))

			// Join the tokenization into a single string
			actualTokenizationString := strings.Join(actualTokenization, "")

			// Expect the actual tokenization to be equal to the expected tokenization
			if tt.expectedTokenization != actualTokenizationString {
				t.Errorf("Expected tokenization %q, got %q", tt.expectedTokenization, actualTokenizationString)
			}
		})
	}
}
