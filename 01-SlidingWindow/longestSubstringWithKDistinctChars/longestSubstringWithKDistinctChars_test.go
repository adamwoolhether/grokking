package longestSubstringWithKDistinctChars

import (
	"fmt"
	"testing"
)

func TestFindMaxSumSubArray(t *testing.T) {
	testCases := []struct {
		input    string
		k        int
		expected int
	}{
		{"araaci", 2, 4},
		{"araaci", 1, 2},
		{"cbbebi", 3, 5},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("string: %v", tc.input), func(t *testing.T) {
			result := longestSubstringWithKDistinct(tc.input, tc.k)

			if tc.expected != result {
				t.Fatalf("Exp result %d, got %d", tc.expected, result)
			}
		})
	}
}
