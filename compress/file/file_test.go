package file

import (
	"bytes"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestReadFileContents_NonExistentFile(t *testing.T) {
	f := &DefaultFile{}
	path := "non_existent_file.txt"
	_, err := f.ReadFileContents(path)
	if err == nil {
		t.Error("Expected an error when opening a non-existent file, but got nil")
	}
}

func TestReadFileContents_Success(t *testing.T) {
	data := []byte("ja jaja jajaja")
	f := &DefaultFile{}

	filePath, cleanup, err := f.CreateTempFileWithData(data)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	fileContents, err := f.ReadFileContents(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	if !bytes.Equal(fileContents, data) {
		t.Errorf("Expected %q, got %q", string(data), string(fileContents))
	}
}

func TestWriteCodeTableOnHeader(t *testing.T) {
	testCases := []struct {
		name          string
		codesTable    map[byte]string
		outputPath    string
		expectedError string
	}{
		{
			name:       "Success",
			codesTable: map[byte]string{'a': "00", 'b': "01", 'c': "10", 'd': "11"},
			outputPath: "test_header.txt",
		},
		{
			name:          "FileCreationError",
			codesTable:    map[byte]string{'a': "00", 'b': "01", 'c': "10", 'd': "11"},
			outputPath:    "", // Invalid path to force an error
			expectedError: "open : no such file or directory",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f := &DefaultFile{}
			err := f.WriteCodeTableOnHeader(tc.codesTable, tc.outputPath)

			if tc.expectedError != "" {
				if err == nil {
					t.Error("Expected an error but got nil")
				} else if !strings.Contains(err.Error(), tc.expectedError) {
					t.Errorf("Expected error message to contain %q, got %q", tc.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}

				// Clean up the temporary file
				os.Remove(tc.outputPath)
			}
		})
	}
}

func TestWriteEncodedText(t *testing.T) {
	// Prepare test data
	encodedText := []byte{0b10000000, 0b11000000, 0b11100000} // Example encoded text

	f := &DefaultFile{}

	// Create a temporary file with a mock header
	filePath, cleanup, err := f.CreateTempFileWithData([]byte(HeaderStart + "\n" + HeaderEnd + "\n"))
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer cleanup() // Clean up the temporary file

	// Use the DefaultFile implementation to write the encoded text to the file
	err = f.WriteTextAfterHeader(encodedText, filePath)
	if err != nil {
		t.Fatalf("Failed to write encoded text: %v", err)
	}

	// Read the contents of the file to verify the encoded text was written correctly
	fileContents, err := f.ReadFileContents(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// Verify that the encoded text was written correctly
	expectedContents := "HS\nHE\n\x80\xc0\xe0"
	if string(fileContents) != expectedContents {
		t.Errorf("Expected file contents %q, got %q", expectedContents, string(fileContents))
	}
}

func TestReadReverseLookupCodesTableFromHeader(t *testing.T) {
	testCases := []struct {
		name          string
		contents      string
		expected      map[string]byte
		expectedError string
	}{
		{
			name:     "Valid_header_with_codes",
			contents: "HS\n4\n00,a\n01,b\n10,c\n11,d\nHE\n",
			expected: map[string]byte{"00": 'a', "01": 'b', "10": 'c', "11": 'd'},
		},
		{
			name:          "Empty_header",
			contents:      "HS\n0\nHE\n",
			expected:      map[string]byte{},
			expectedError: "",
		},
		{
			name:          "Missing_header_end_marker",
			contents:      "HS\n4\n00,a\n01,b\n10,c\n11,d",
			expectedError: "invalid header end marker",
		},
		{
			name:          "Missing_header_start_marker",
			contents:      "4\n00,a\n01,b\n10,c\n11,d\nHE\n",
			expectedError: "invalid header start marker",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a temporary test file with contents
			f := &DefaultFile{}
			filePath, cleanup, err := f.CreateTempFileWithData([]byte(tc.contents))
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer cleanup()

			// Read codes table from the header
			codesTable, err := f.ReadReverseLookupCodesTableFromHeader(filePath)
			// Check for expected error
			if tc.expectedError != "" {
				if err == nil {
					t.Error("Expected an error but got nil")
				} else if !strings.Contains(err.Error(), tc.expectedError) {
					t.Errorf("Expected error message to contain %q, got %q", tc.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if !reflect.DeepEqual(codesTable, tc.expected) {
					t.Errorf("Expected codesTable %v, got %v", tc.expected, codesTable)
				}
			}
		})
	}
}

func TestCreateNewFile(t *testing.T) {
	f := &DefaultFile{}

	// Create a temporary test file
	fileName := "testfile.txt"
	createdFile, err := f.CreateNewFile(fileName)
	if err != nil {
		t.Fatalf("Failed to create new file: %v", err)
	}
	defer func() {
		// Close the file
		err := createdFile.Close()
		if err != nil {
			t.Fatalf("Failed to close file: %v", err)
		}

		// Remove the file
		err = os.Remove(fileName)
		if err != nil {
			t.Fatalf("Failed to remove file: %v", err)
		}
	}()

	// Check if the file was created
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Error("Expected file to be created, but it was not")
	}
}

func TestReadTextAfterHeader(t *testing.T) {
	testCases := []struct {
		name          string
		contents      string
		expected      []byte
		expectedError string
	}{
		{
			name:     "Valid_encoded_content",
			contents: "HS\n3\na,1\nb,2\nc,3\nHE\n010101\n",
			expected: []byte("010101"),
		},
		{
			name:          "Empty_file",
			contents:      "",
			expected:      []byte{},
			expectedError: "header end marker not found",
		},
		{
			name:          "Invalid_header_format",
			contents:      "INVALID_HEADER\n010101\n",
			expected:      []byte{},
			expectedError: "header end marker not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a temporary test file with contents
			f := &DefaultFile{}
			filePath, cleanup, err := f.CreateTempFileWithData([]byte(tc.contents))
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer cleanup()

			// Read encoded content from the file
			encodedContent, err := f.ReadTextAfterHeader(filePath)
			// Check for expected error
			if tc.expectedError != "" {
				if err == nil {
					t.Error("Expected an error but got nil")
				} else if !strings.Contains(err.Error(), tc.expectedError) {
					t.Errorf("Expected error message to contain %q, got %q", tc.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if !reflect.DeepEqual(encodedContent, tc.expected) {
					t.Errorf("Expected encoded content %v, got %v", tc.expected, encodedContent)
				}
			}
		})
	}
}
