package helpers

import (
	"testing"
)

func TestInputToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		hasError bool
	}{
		{"NumericString", "123", 123, false},
		{"NegativeNumber", "-45", -45, false},
		{"Empty", "", 0, true},
		{"NonNumeric", "abc", 0, true},
		{"LeadingSpaces", "   67", 0, true},
		{"TrailingSpaces", "89   ", 0, true},
		{"MixedChars", "12abc34", 0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := InputToInt(tc.input)
			if (err != nil) != tc.hasError {
				t.Fatalf("InputToInt(%q) error = %v, want error: %v", tc.input, err != nil, tc.hasError)
			}
			if got != tc.expected {
				t.Fatalf("InputToInt(%q) = %d, want %d", tc.input, got, tc.expected)
			}
		})
	}
}
