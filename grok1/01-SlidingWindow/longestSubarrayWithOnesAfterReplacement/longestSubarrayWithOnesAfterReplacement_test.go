package longestSubarrayWithOnesAfterReplacement

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		input    []int
		k        int
		expected int
	}{
		{[]int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}, 2, 6},
		// {[]int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}, 3, 9},
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
