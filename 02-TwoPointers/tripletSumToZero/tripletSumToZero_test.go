package tripletSumToZero

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchTriplets(t *testing.T) {
	testCases := []struct {
		input  []float64
		output [][3]float64
	}{
		{[]float64{-3, 0, 1, 2, -1, 1, -2}, [][3]float64{{-3, 1, 2}, {-2, 0, 2}, {-2, 1, 1}, {-1, 0, 1}}},
		{[]float64{-5, 2, -1, -2, 3}, [][3]float64{{-5, 2, 3}, {-2, -1, 3}}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			if result := searchTriplets(tc.input); !reflect.DeepEqual(tc.output, result) {
				t.Errorf("exp %v, got %v", tc.output, result)
			}

		})
	}
}
