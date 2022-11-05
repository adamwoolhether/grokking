package dutchNationalFlagProblem

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDutchFlagSort(t *testing.T) {
	testCases := []struct {
		input    []float64
		expected []float64
	}{
		{[]float64{1, 0, 2, 1, 0}, []float64{0, 0, 1, 1, 2}},
		{[]float64{2, 2, 0, 1, 2, 0}, []float64{0, 0, 1, 2, 2, 2}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			if result := dutchFlagSort(tc.input); !reflect.DeepEqual(tc.expected, result) {
				t.Errorf("exp %v, gpt %v", tc.expected, result)
			}
		})
	}
}
