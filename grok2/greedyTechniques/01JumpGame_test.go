package greedyTechniques

import (
	"fmt"
	"testing"
)

func Test_jumpGame(t *testing.T) {
	tests := []struct {
		input []int
		exp   bool
	}{
		{[]int{2, 3, 1, 1, 9}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{4, 0, 0, 0, 4}, true},
		{[]int{0}, true},
		{[]int{1}, true},
		{[]int{3, 2, 2, 0, 1, 4}, true},
		{[]int{2, 3, 1, 1, 9}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{1, 2, 3, 4, 5, 6, 7}, true},
		{[]int{3, 3, 3, 3, 3}, true},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			if got := jumpGame(tc.input); got != tc.exp {
				t.Errorf("jumpGame() = %v, want %v", got, tc.exp)
			}
		})
	}
}
