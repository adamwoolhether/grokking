package happyNumber

import (
	"fmt"
	"testing"
)

func TestFindHappyNumber(t *testing.T) {
	testCases := []struct {
		input int
		exp   bool
	}{
		{23, true},
		{12, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Number: %d", tc.input), func(t *testing.T) {
			if result := findHappyNumber(tc.input); tc.exp != result {
				t.Errorf("exp %t, got %t", tc.exp, result)
			}
		})
	}
}
