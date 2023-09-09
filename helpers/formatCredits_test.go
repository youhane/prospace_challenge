package helpers

import (
	"fmt"
	"testing"
)

func TestFormatCredits(t *testing.T) {
	testCases := []struct {
		credits  float32
		expected string
	}{
		{10, "10"},         // Test integer with no decimal part
		{10.5, "10.50"},    // Test float with 2 decimal places
		{10.123, "10.12"},  // Test float with more than 2 decimal places
		{10.001, "10.00"},  // Test float with fractional part close to zero
		{0.123456, "0.12"}, // Test float with fractional part
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Credits %.2f", testCase.credits), func(t *testing.T) {
			result := FormatCredits(testCase.credits)
			if result != testCase.expected {
				t.Errorf("Expected '%s' but got '%s'", testCase.expected, result)
			}
		})
	}
}
