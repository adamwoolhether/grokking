package longestSubstringWithDistinctChars

import (
	"fmt"
	"testing"
)

func TestNonRepeatSubstring(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"aabccbb", 3},
		{"abbbb", 2},
		{"abccde", 3},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("string: %v", tc.input), func(t *testing.T) {
			result := nonRepeatSubstring(tc.input)

			if tc.expected != result {
				t.Fatalf("Exp result %d, got %d", tc.expected, result)
			}
		})
	}
}
