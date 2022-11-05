package bitwiseManipulation

import (
	"fmt"
	"testing"
)

func TestExtraCharIndex(t *testing.T) {
	testCases := []struct {
		first  string
		second string
		exp    int
	}{
		{"wxyz", "zwxgy", 3},
		{"cbda", "abc", 2},
		{"jlknm", "klmn", 0},
		{"courae", "couearg", 6},
		{"hello", "helo", 2},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if result := extraCharIndex(tc.first, tc.second); result != tc.exp {
				t.Errorf("exp %d, got %d", tc.exp, result)
			}
		})
	}
}
