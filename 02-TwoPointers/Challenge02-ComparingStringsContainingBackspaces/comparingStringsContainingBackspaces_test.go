package Challenge02_ComparingStringsContainingBackspaces

import (
	"fmt"
	"testing"
)

func TestBackspaceCompare(t *testing.T) {
	testCases := []struct {
		str1 string
		str2 string
		exp  bool
	}{
		{"xy#z", "xzz#", true},
		{"xy#z", "xyz#", false},
		{"xp#", "xyz##", true},
		{"xywrrmp", "xywrrmu#p", true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("string1: %v , string2: %v", tc.str1, tc.str2), func(t *testing.T) {
			if result := backspaceCompare(tc.str1, tc.str2); tc.exp != result {
				t.Errorf("exp %t, got %t", tc.exp, result)
			}
		})
	}
}
