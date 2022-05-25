package longestSubstringWithSameLettersAfterReplacement

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		input    string
		k        int
		expected int
	}{
		{"aabccbb", 2, 5},
		{"abbcb", 1, 4},
		{"abccde", 1, 3},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("string: %v", tc.input), func(t *testing.T) {
			result := lengthOfLongestSubstring(tc.input, tc.k)

			if tc.expected != result {
				t.Fatalf("Exp result %d, got %d", tc.expected, result)
			}
		})
	}
}
