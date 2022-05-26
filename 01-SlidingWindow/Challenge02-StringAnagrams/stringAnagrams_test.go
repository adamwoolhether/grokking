package Challenge02_StringAnagrams

import (
	"fmt"
	"testing"
)

func TestFindStringAnagrams(t *testing.T) {
	testCases := []struct {
		input    string
		pattern  string
		expected []int
	}{
		{"ppqp", "pq", []int{1, 2}},
		{"abbcabc", "abc", []int{2, 3, 4}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v, pattern: %v", tc.input, tc.pattern), func(t *testing.T) {
			result := findStringAnagrams(tc.input, tc.pattern)

			for i, v := range tc.expected {
				if v != result[i] {
					t.Errorf("Index %d; exp %d, got %d", i, v, result[i])
				}
			}
		})
	}
}
