package separateDuplicates

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		input []float64
		exp   int
	}{
		{[]float64{2, 3, 3, 3, 6, 9, 9}, 4},
		{[]float64{2, 2, 2, 11}, 2},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input %v", tc.input), func(t *testing.T) {
			result := removeDuplicates(tc.input)

			if tc.exp != result {
				t.Errorf("exp %d, got %d", tc.exp, result)
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		input []float64
		key   float64
		exp   int
	}{
		{[]float64{3, 2, 3, 6, 3, 10, 9, 3}, 3, 4},
		{[]float64{2, 11, 2, 2, 1}, 2, 2},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input %v", tc.input), func(t *testing.T) {
			result := removeElement(tc.input, tc.key)

			if tc.exp != result {
				t.Errorf("exp %d, got %d", tc.exp, result)
			}
		})
	}
}
