package greedyTechniques

import (
	"fmt"
	"testing"
)

func Test_rescueBoats(t *testing.T) {
	tests := []struct {
		people []int
		limit  int
		exp    int
	}{
		{[]int{3, 1, 4, 2, 4}, 4, 4},
		{[]int{1, 1, 1, 1, 2}, 3, 3},
		{[]int{1, 2}, 3, 1},
		{[]int{5, 5, 5, 5}, 5, 4},
		{[]int{3, 2, 5, 5}, 5, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.people, tt.limit), func(t *testing.T) {
			if got := rescueBoats(tt.people, tt.limit); got != tt.exp {
				t.Errorf("rescueBoats() = %v, exp %v", got, tt.exp)
			}
		})
	}
}
