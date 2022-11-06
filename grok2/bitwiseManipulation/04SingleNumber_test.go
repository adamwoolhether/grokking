package bitwiseManipulation

import "testing"

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		input []int
		exp   int
	}{
		{[]int{1, 2, 2, 3, 3, 1, 4}, 4},
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{25}, 25},
		{[]int{4, 5, 5, 2, 1, 1, 4}, 2},
	}

	for _, tc := range testCases {
		if result := singleNumber(tc.input); result != tc.exp {
			t.Errorf("exp: %d, got %d", tc.exp, result)
		}
	}
}
