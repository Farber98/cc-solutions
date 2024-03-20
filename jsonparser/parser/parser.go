package parser

// Parser defines the interface for parsing JSON tokens.
type Parser interface {
	Parse(tokens []string) bool
}

// SimpleParser implements the Parser interface for parsing JSON tokens.
type SimpleParser struct{}

// Parse checks if the given tokens represent a valid JSON object with string keys and values.
func (p *SimpleParser) Parse(tokens []string) bool {
	// Return false if tokens is empty or does not start with '{' and end with '}'
	if len(tokens) < 2 || tokens[0] != "{" || tokens[len(tokens)-1] != "}" {
		return false
	}

	// Check if tokens contain valid key-value pairs
	for i := 1; i < len(tokens)-1; i += 4 {
		// Each key-value pair should have the format: "<key>": "<value>"
		// where <key> and <value> are strings
		if i+3 >= len(tokens) || tokens[i] != "\"" || tokens[i+1] == "}" || tokens[i+2] != ":" || tokens[i+3] != "\"" {
			return false
		}
	}

	return true
}
