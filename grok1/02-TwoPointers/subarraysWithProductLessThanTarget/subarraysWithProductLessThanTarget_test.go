package subarraysWithProductLessThanTarget

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindSubarrays(t *testing.T) {
	testCases := []struct {
		input    []float64
		target   float64
		expected [][]float64
	}{
		{[]float64{2, 5, 3, 10}, 30, [][]float64{{2}, {5}, {2, 5}, {3}, {5, 3}, {10}}},
		{[]float64{8, 2, 6, 5}, 50, [][]float64{{8}, {2}, {8, 2}, {6}, {2, 6}, {5}, {6, 5}}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			if result := findSubarrays(tc.input, tc.target); !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("exp %v, got %v", tc.expected, result)
			}
		})
	}
}