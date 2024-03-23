package frequency

import (
	"reflect"
	"testing"
)

func TestDefaultCalculator_CalculateFrequencies(t *testing.T) {
	contents := []byte("abbcaabbccc")
	expected := map[byte]int{'a': 3, 'b': 4, 'c': 4}

	calculator := DefaultCalculator{}

	result := calculator.CalculateFrequencies(contents)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected result. Expected: %v, Got: %v", expected, result)
	}
}
