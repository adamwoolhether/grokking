package bitwiseManipulation

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFlipAndInvertImage(t *testing.T) {
	testCases := []struct {
		input [][]int
		exp   [][]int
	}{
		{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}, [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
		{[][]int{{1, 0, 1, 0}, {1, 0, 0, 1}, {1, 1, 1, 1}, {1, 0, 0, 0}}, [][]int{{1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}, {1, 1, 1, 0}}},
		{[][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 0}}, [][]int{{1, 1, 1}, {0, 1, 0}, {1, 0, 1}}},
		{[][]int{{1, 1, 1}, {0, 0, 0}, {1, 1, 1}}, [][]int{{0, 0, 0}, {1, 1, 1}, {0, 0, 0}}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {

			if result := flipAndInvertImage(tc.input); !cmp.Equal(result, tc.exp) {
				t.Errorf("exp: %v, got: %v", tc.exp, result)
			}
		})
	}
}

func BenchmarkFlipAndInvertImage(b *testing.B) {
	input := [][]int{{1, 1, 0, 0, 1, 0, 0, 1}, {0, 0, 1, 1, 1, 0, 0, 1}, {0, 1, 0, 1, 0, 1, 1, 1}, {1, 1, 1, 0, 1, 0, 1, 0}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = flipAndInvertImage(input)
	}
}
