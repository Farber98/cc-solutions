package compress

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	testCases := []struct {
		name        string
		encoded     []byte
		codeTable   map[string]byte
		expected    []byte
		expectedErr error
	}{
		{
			name:        "ValidDecode",
			encoded:     []byte{165, 15},
			codeTable:   map[string]byte{"00": 99, "01": 98, "10": 97, "11": 100},
			expected:    []byte("aabbccdd"),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := &DefaultDecompressor{}
			decoded, err := d.Decode(tc.encoded, tc.codeTable)

			if !reflect.DeepEqual(decoded, tc.expected) {
				t.Errorf("Decode() failed, expected: %v, got: %v", tc.expected, decoded)
			}

			if (err != nil && tc.expectedErr == nil) || (err == nil && tc.expectedErr != nil) || (err != nil && err.Error() != tc.expectedErr.Error()) {
				t.Errorf("Decode() error, expected: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}
