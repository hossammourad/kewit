package helpers

import "testing"

func TestIsValidURL(t *testing.T) {
	tests := []struct {
		name string
		input   string
		expected bool
	}{
		{"Empty", "", false},
		{"String", "test", false},
		{"Integer", "1", false},
		{"Space", " ", false},
		{"Valid HTTP", "http://example.com", true},
		{"Valid HTTPS", "https://example.com", true},
		{"Valid with path", "https://example.com/path", true},
		{"Valid with query", "https://example.com/path?query=1", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValidURL(tc.input)
			if got != tc.expected {
				t.Fatalf("IsValidURL(%q) = %v, want %v", tc.input, got, tc.expected)
			}
		})
	}
}
