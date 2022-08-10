package mergeIntervals

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		name      string
		intervals []Interval
		exp       []Interval
	}{
		{"one", []Interval{{1, 4}, {2, 5}, {7, 9}}, []Interval{{1, 5}, {7, 9}}},
		{"two", []Interval{{6, 7}, {2, 4}, {5, 9}}, []Interval{{2, 4}, {5, 9}}},
		{"three", []Interval{{1, 4}, {2, 6}, {3, 5}}, []Interval{{1, 6}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := merge(tc.intervals)

			if !cmp.Equal(result, tc.exp) {
				t.Errorf("exp %v, got %v", tc.exp, result)
			}
		})
	}
}
