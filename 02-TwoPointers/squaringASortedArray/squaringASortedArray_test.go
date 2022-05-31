package squaringASortedArray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMakeSquares(t *testing.T) {
	testCases := []struct {
		input  []float64
		output []float64
	}{
		{[]float64{-2, -1, 0, 2, 3}, []float64{0, 1, 4, 4, 9}},
		{[]float64{-3, -1, 0, 1, 2}, []float64{0, 1, 1, 4, 9}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			result := makeSquares(tc.input)

			if !reflect.DeepEqual(tc.output, result) {
				t.Errorf("exp %v, got %v", tc.output, result)
			}
		})
	}
}