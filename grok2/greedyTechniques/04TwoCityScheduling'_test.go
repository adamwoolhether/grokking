package greedyTechniques

import (
	"fmt"
	"testing"
)

func Test_twoCityScheduling(t *testing.T) {
	tests := []struct {
		input [][]int
		want  int
	}{
		{[][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}, 110},
		{[][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}, 1859},
		{[][]int{{515, 563}, {451, 713}, {537, 709}, {343, 819}, {855, 779}, {457, 60}, {650, 359}, {631, 42}}, 3086},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 18},
		{[][]int{{1, 2}, {1, 2}, {1, 2}, {1, 2}}, 6},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := twoCityScheduling(tt.input); got != tt.want {
				t.Errorf("twoCityScheduling() = %v, want %v", got, tt.want)
			}
		})
	}
}
