package smallestSubarrayWithGreaterSum

import (
	"fmt"
	"testing"
)

func TestSmallestSubarraySum(t *testing.T) {
	testCases := []struct {
		input    []float64
		s        float64
		expected float64
	}{
		{[]float64{2, 1, 5, 2, 3, 2}, 7, float64(2)},
		{[]float64{2, 1, 5, 2, 8}, 7, float64(1)},
		{[]float64{3, 4, 1, 1, 6}, 8, float64(3)},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("slice: %v", tc.input), func(t *testing.T) {
			result := smallestSubarraySum(tc.s, tc.input)

			if tc.expected != result {
				t.Fatalf("Exp result %f, got %f", tc.expected, result)
			}
		})
	}
}
