package Challenge01_PermutationInAString

import (
	"fmt"
	"testing"
)

func TestFindPermutation(t *testing.T) {
	testCases := []struct {
		input    string
		pattern  string
		expected bool
	}{
		{"oidbcaf", "abc", true},
		{"odicf", "dc", false},
		{"bcdxabcdy", "bcdyabcdx", true},
		{"aaacb", "abc", true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v, pattern: %v", tc.input, tc.pattern), func(t *testing.T) {
			result := findPermutation(tc.input, tc.pattern)

			if tc.expected != result {
				t.Fatalf("Exp result %t, got %t", tc.expected, result)
			}
		})
	}
}
