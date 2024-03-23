package frequency

// FrequencyCalculator defines the interface for calculating character frequencies.
type FrequencyCalculator interface {
	CalculateFrequencies(contents []byte) map[byte]int
}

// DefaultCalculator implements the Calculator interface with default frequency calculation.
type DefaultCalculator struct{}

// CalculateFrequencies calculates the frequency of each character in the text.
func (c *DefaultCalculator) CalculateFrequencies(content []byte) map[byte]int {
	frequencies := make(map[byte]int)
	for _, char := range content {
		frequencies[char]++
	}
	return frequencies
}
