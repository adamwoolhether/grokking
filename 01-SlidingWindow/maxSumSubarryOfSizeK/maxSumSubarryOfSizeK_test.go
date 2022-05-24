package maxSumSubarryOfSizeK

import (
	"fmt"
	"testing"
)

func TestFindMaxSumSubArray(t *testing.T) {
	testCases := []struct {
		input    []float64
		k        int
		expected float64
	}{
		{[]float64{2, 1, 5, 1, 3, 2}, 3, float64(9)},
		{[]float64{2, 3, 4, 1, 5}, 2, float64(7)},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("slice: %v", tc.input), func(t *testing.T) {
			result := findMaxSumSubArray(tc.k, tc.input)

			if tc.expected != result {
				t.Fatalf("Exp result %f, got %f", tc.expected, result)
			}
		})
	}
}

func TestFindMaxSumSubArrayBruteForce(t *testing.T) {
	testCases := []struct {
		input    []float64
		k        int
		expected float64
	}{
		{[]float64{2, 1, 5, 1, 3, 2}, 3, float64(9)},
		{[]float64{2, 3, 4, 1, 5}, 2, float64(7)},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("slice: %v", tc.input), func(t *testing.T) {
			result := findMaxSumSubArrayBruteForce(tc.k, tc.input)

			if tc.expected != result {
				t.Fatalf("Exp result %f, got %f", tc.expected, result)
			}
		})
	}
}

func BenchmarkFindMaxSumSubarrays(b *testing.B) {
	b.ResetTimer()

	arr := []float64{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 5

	for i := 0; i < b.N; i++ {
		_ = findMaxSumSubArray(k, arr)
	}
}

func BenchmarkFindMaxSumSubarraysBruteForce(b *testing.B) {
	b.ResetTimer()

	arr := []float64{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 5

	for i := 0; i < b.N; i++ {
		_ = findMaxSumSubArrayBruteForce(k, arr)
	}
}
