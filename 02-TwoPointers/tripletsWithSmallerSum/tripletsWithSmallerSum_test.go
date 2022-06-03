package tripletsWithSmallerSum

import (
	"fmt"
	"testing"
)

func TestTripletWithSmallerSum(t *testing.T) {
	testCases := []struct {
		input  []float64
		target float64
		output int
	}{
		{[]float64{-1, 0, 2, 3}, 3, 2},
		{[]float64{-1, 4, 2, 1, 3}, 5, 4},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			if result := tripletWithSmallerSum(tc.input, tc.target); result != tc.output {
				t.Errorf("exp %v, got %v", tc.output, result)
			}
		})
	}
}

func TestTripletWithSmallerSumTriplets(t *testing.T) {
	testCases := []struct {
		input  []float64
		target float64
		output [][3]float64
	}{
		{[]float64{-1, 0, 2, 3}, 3, [][3]float64{{-1, 0, 3}, {-1, 0, 2}}},
		{[]float64{-1, 4, 2, 1, 3}, 5, [][3]float64{{-1, 1, 4}, {-1, 1, 3}, {-1, 1, 2}, {-1, 2, 3}}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input %v", tc.input), func(t *testing.T) {
			result := tripletWithSmallerSumTriplets(tc.input, tc.target)
			for i, v := range result {
				if v != tc.output[i] {
					t.Errorf("exp %v, got %v", tc.output[i], v)
				}
			}
		})
	}
}
