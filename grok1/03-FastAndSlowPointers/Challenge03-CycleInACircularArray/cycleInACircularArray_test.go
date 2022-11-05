package Challenge03_CycleInACircularArray

import (
	"fmt"
	"testing"
)

func TestCircularArrayLoopExists(t *testing.T) {
	testCases := []struct {
		input []float64
		exp   bool
	}{
		{[]float64{1, 2, -1, 2, 2}, true},
		{[]float64{2, 2, -1, 2}, true},
		{[]float64{2, 1, -1, -2}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			if got := circularArrayLoopExists(tc.input); got != tc.exp {
				t.Errorf("Exp %t, got %t", got, tc.exp)
			}
		})
	}
}