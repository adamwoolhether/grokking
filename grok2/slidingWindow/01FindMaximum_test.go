package slidingWindow

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_findMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		input      []int
		windowSize int
		exp        []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, []int{3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3}, 4, []int{3, 3, 3, 3, 3, 3, 3}},
		{[]int{10, 6, 9, -3, 23, -1, 34, 56, 67, -1, -4, -8, -2, 9, 10, 34, 67}, 2, []int{10, 9, 9, 23, 23, 34, 56, 67, 67, -1, -4, -2, 9, 10, 34, 67}},
		{[]int{4, 5, 6, 1, 2, 3}, 1, []int{4, 5, 6, 1, 2, 3}},
		{[]int{9, 5, 3, 1, 6, 3}, 2, []int{9, 5, 3, 6, 6}},
		{[]int{1, 2}, 2, []int{2}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := findMaxSlidingWindow(tt.input, tt.windowSize); !cmp.Equal(got, tt.exp) {
				t.Errorf("findMaxSlidingWindow() = %v, exp %v", got, tt.exp)
			}
		})
	}
}
