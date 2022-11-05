package Challenge01_QuadrupleSumToTarget

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchQuadruplets(t *testing.T) {
	testCases := []struct {
		input  []float64
		target float64
		output [][4]float64
	}{
		{[]float64{4, 1, 2, -1, 1, -3}, 1, [][4]float64{{-3, -1, 1, 4}, {-3, 1, 1, 2}}},
		{[]float64{2, 0, -1, 1, -2, 2}, 2, [][4]float64{{-2, 0, 2, 2}, {-1, 0, 1, 2}}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			if result := searchQuadruplets(tc.input, tc.target); !reflect.DeepEqual(tc.output, result) {
				t.Errorf("exp %v, got %v", tc.output, result)
			}

		})
	}
}
