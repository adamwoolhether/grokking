package pairWithTargetSum

import (
	"fmt"
	"testing"
)

func TestPairWithTargetSum(t *testing.T) {
	testCases := []struct {
		input  []float64
		target float64
		exp    []int
	}{
		{[]float64{1, 2, 3, 4, 6}, 6, []int{1, 3}},
		{[]float64{2, 5, 9, 11}, 11, []int{0, 2}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input %v", tc.input), func(t *testing.T) {
			result := pairWithTargetSum(tc.input, tc.target)

			for i, res := range result {
				if tc.exp[i] != res {
					t.Errorf("index %d; exp %d, got %d", i, tc.exp[i], res)
				}
			}
		})
	}
}

func TestPairWithTargetSumAlternate(t *testing.T) {
	testCases := []struct {
		input  []float64
		target float64
		exp    []int
	}{
		{[]float64{1, 2, 3, 4, 6}, 6, []int{1, 3}},
		{[]float64{2, 5, 9, 11}, 11, []int{0, 2}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input %v", tc.input), func(t *testing.T) {
			result := pairWithTargetSumAlternate(tc.input, tc.target)

			for i, res := range result {
				if tc.exp[i] != res {
					t.Errorf("index %d; exp %d, got %d", i, tc.exp[i], res)
				}
			}
		})
	}
}

func BenchmarkPairWithTargetSum(b *testing.B) {
	arr := []float64{1, 2, 3, 4, 6}
	k := float64(6)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = pairWithTargetSum(arr, k)
	}
}

func BenchmarkPairWithTargetSumAlternate(b *testing.B) {
	arr := []float64{1, 2, 3, 4, 6}
	k := float64(6)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = pairWithTargetSumAlternate(arr, k)
	}
}
