package greedyTechniques

import (
	"fmt"
	"testing"
)

func Test_gasStationJourney(t *testing.T) {
	tests := []struct {
		gas  []int
		cost []int
		exp  int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
		{[]int{1, 1, 1, 1, 10}, []int{2, 2, 1, 3, 1}, 4},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}, -1},
		{[]int{1}, []int{1}, 0},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := gasStationJourney(tc.gas, tc.cost); got != tc.exp {
				t.Errorf("gasStationJourney() = %v, exp %v", got, tc.exp)
			}
		})
	}
}
