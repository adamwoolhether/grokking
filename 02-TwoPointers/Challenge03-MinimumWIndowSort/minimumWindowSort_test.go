package Challenge03_MinimumWIndowSort

import (
	"fmt"
	"testing"
)

func TestShortestWindow(t *testing.T) {
	testCases := []struct {
		input    []float64
		expected int
	}{
		{[]float64{1, 2, 5, 3, 7, 10, 9, 12}, 5},
		{[]float64{1, 3, 2, 0, -1, 7, 10}, 5},
		{[]float64{1, 2, 3}, 0},
		{[]float64{3, 2, 1}, 3},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			if result := shortestWindow(tc.input); result != tc.expected {
				t.Errorf("exp %v, got %v", tc.expected, result)
			}
		})
	}
}
