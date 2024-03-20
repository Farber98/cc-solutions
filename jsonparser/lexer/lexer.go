package lexer

// Lexer defines the interface for tokenizing JSON input.
type Lexer interface {
	Lex(input string) []string
}

// SimpleLexer implements the Lexer interface for tokenizing JSON input.
type SimpleLexer struct{}

// Lex tokenizes the input JSON string and returns a list of tokens.
func (l *SimpleLexer) Lex(input string) []string {
	var tokens []string
	var currentToken string
	inString := false

	for _, char := range input {
		switch char {
		case '{', '}', ',':
			// If not in a string, append the current token (if any) and add the character as a token
			if !inString {
				if currentToken != "" {
					tokens = append(tokens, currentToken)
					currentToken = ""
				}
				tokens = append(tokens, string(char))
			} else {
				// If in a string, add the character to the current token
				currentToken += string(char)
			}
		case '"':
			// Toggle the inString flag and add the current token (including the opening quote) as a token
			inString = !inString
			currentToken += string(char)
			if !inString {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
		case ':':
			// If not in a string, append the current token (if any) and add the colon as a token
			if !inString {
				if currentToken != "" {
					tokens = append(tokens, currentToken)
					currentToken = ""
				}
				tokens = append(tokens, string(char))
			} else {
				// If in a string, add the character to the current token
				currentToken += string(char)
			}
		case ' ', '\t', '\n', '\r':
			// Skip whitespace characters
		default:
			// Add the character to the current token
			currentToken += string(char)
		}
	}

	// Add the last token if there's any
	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	return tokens
}
