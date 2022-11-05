package tripletSumCloseToTarget

import (
	"fmt"
	"testing"
)

func TestTripletSumCloseToTarget(t *testing.T) {
	testCases := []struct {
		input  []float64
		target float64
		output float64
	}{
		{[]float64{-2, 0, 1, 2}, 2, 1},
		{[]float64{-3, -1, 1, 2}, 1, 0},
		{[]float64{1, 0, 1, 1}, 100, 3},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			if result := tripletSumCloseToTarget(tc.input, tc.target); result != tc.output {
				t.Errorf("exp %v, got %v", tc.output, result)
			}
		})
	}
}
