package stacks

import (
	"fmt"
	"testing"
)

func TestCalculator(t *testing.T) {
	testCases := []struct {
		input string
		exp   int
	}{
		{"12 - (6 + 2) + 5", 9},
		{"(8 + 100) + (13 - 8 - (2 + 1))", 110},
		{"40 - 25 - 5", 10},
		{"(27 + (7 + 5) - 3) + (6 + 10)", 52},
		{"34 + 45 + (23324 - 3454)", 19949},
		{"((12 + (13 + 17 - (65 - 98) + 34)) - 44)", 65},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			if result := Calculator(tc.input); result != tc.exp {
				t.Errorf("exp: %d, got %d", tc.exp, result)
			}
		})
	}
}
