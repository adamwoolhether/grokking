package Challenge03_SmallestWindowContainingSubstring

import (
	"fmt"
	"testing"
)

func TestMinimumWindowSubstring(t *testing.T) {
	testCases := []struct {
		input    string
		pattern  string
		expected string
	}{
		{"aabdec", "abc", "abdec"},
		{"aabdec", "abac", "aabdec"},
		{"abdbca", "abc", "bca"},
		{"adcad", "abc", ""},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v, pattern: %v", tc.input, tc.pattern), func(t *testing.T) {
			result := minimumWindowSubstring(tc.input, tc.pattern)

			if tc.expected != result {
				t.Fatalf("Exp result %s, got %s", tc.expected, result)
			}
		})
	}
}
