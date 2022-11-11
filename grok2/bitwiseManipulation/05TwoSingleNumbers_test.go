package bitwiseManipulation

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTwoSingleNumbers(t *testing.T) {
	testCases := []struct {
		input []int
		exp   []int
	}{
		{[]int{1, 2, 2, 3, 3, 4}, []int{1, 4}},
		{[]int{2, 1, 1, 8}, []int{2, 8}},
		{[]int{4, 1, 2, 1, 2, 0}, []int{4, 0}},
		{[]int{2, 5}, []int{5, 2}},
		{[]int{4, 5, 4, 1, 1, 2, 6, 6}, []int{5, 2}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			if result := twoSingleNumbers(tc.input); !cmp.Equal(result, tc.exp) {
				t.Errorf("exp: %v, got: %v", tc.exp, result)
			}
		})
	}
}
