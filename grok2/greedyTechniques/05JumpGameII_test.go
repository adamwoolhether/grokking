package greedyTechniques

import (
	"fmt"
	"testing"
)

func Test_jumpGameTwo(t *testing.T) {
	tests := []struct {
		input []int
		exp   int
	}{
		{[]int{2, 3, 1, 1, 9}, 2},
		{[]int{3, 2, 1, 1, 4}, 2},
		{[]int{4, 0, 0, 0, 4}, 1},
		{[]int{1, 1}, 1},
		{[]int{1}, 0},
		{[]int{2, 0, 1, 1, 2, 3, 6}, 4},
		{[]int{1, 2, 3, 4, 5}, 3},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if result := jumpGameTwo(tt.input); result != tt.exp {
				t.Errorf("exp: %d, got %d", tt.exp, result)
			}
		})
	}
}
