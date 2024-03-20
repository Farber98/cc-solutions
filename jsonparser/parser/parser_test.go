package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name            string
		tokenization    []string
		expectedParsing bool
	}{
		{
			name:            "STEP1: Valid JSON",
			tokenization:    []string{"{", "}"},
			expectedParsing: true,
		},
		{
			name:            "STEP1: Invalid JSON",
			tokenization:    []string{},
			expectedParsing: false,
		},
		{
			name:            "STEP2: Valid JSON 2_1",
			tokenization:    []string{"{\"key\":\"value\"}"},
			expectedParsing: true,
		},
		{
			name:            "STEP2: Valid JSON 2_2",
			tokenization:    []string{"{\"key\":\"value\",\"key2\":\"value\"}"},
			expectedParsing: true,
		},
		{
			name:            "STEP2: Invalid JSON 2_1",
			tokenization:    []string{"{\"key\":\"value\",}"},
			expectedParsing: false,
		},
		{
			name:            "STEP2: Invalid JSON 2_2",
			tokenization:    []string{"{\"key\":\"value\",key2:\"value\"}"},
			expectedParsing: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Define our parser
			p := &SimpleParser{}

			// Parse our tokenization
			actualParsing := p.Parse(tt.tokenization)

			// Expect the actual parsing to be equal to the expected parsing
			if actualParsing != tt.expectedParsing {
				t.Errorf("Expected parsing %t, got %t", tt.expectedParsing, actualParsing)
			}
		})
	}
}
