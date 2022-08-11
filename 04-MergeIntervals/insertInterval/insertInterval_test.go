package insertInterval

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		name        string
		intervals   []Interval
		newInterval Interval
		exp         []Interval
	}{
		{"one", []Interval{{1, 3}, {5, 7}, {8, 12}}, Interval{4, 6}, []Interval{{1, 3}, {4, 7}, {8, 12}}},
		{"two", []Interval{{1, 3}, {5, 7}, {8, 12}}, Interval{4, 10}, []Interval{{1, 3}, {4, 12}}},
		{"three", []Interval{{2, 3}, {5, 7}}, Interval{1, 4}, []Interval{{1, 4}, {5, 7}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := insert(tc.intervals, tc.newInterval)

			if !cmp.Equal(result, tc.exp) {
				t.Errorf("exp %v, got %v", tc.exp, result)
			}
		})
	}
}
