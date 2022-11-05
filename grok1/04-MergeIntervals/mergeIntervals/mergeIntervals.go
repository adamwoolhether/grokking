package mergeIntervals

import "sort"

/*
Given a list of intervals, merge all the overlapping intervals to produce a list
that has only mutually exclusive intervals

Example 1:
Intervals: [[1,4], [2,5], [7,9]]
Output: [[1,5], [7,9]]
Explanation: Since the first two intervals [1,4] and [2,5] overlap, we merged them into one [1,5].

Example 2:
Intervals: [[6,7], [2,4], [5,9]]
Output: [[2,4], [5,9]]
Explanation: Since the intervals [6,7] and [5,9] overlap, we merged them into one [5,9].

Example 3:
Intervals: [[1,4], [2,6], [3,5]]
Output: [[1,6]]
Explanation: Since all the given intervals overlap, we merged them into one.
*/

// Solution:
/*
Using the example such that a.start <= b.start.
Four possible scenarios:
1. 'a' and 'b' do not overlap.
2. some part of 'b' overlaps with 'a'.
3. 'a' fully overlaps with 'b'.
4. 'b' fully overlaps 'a' both have the same start time.

Our Merge algo looks like:
1. Sort the intervals on the start time to ensure a.start <= b.start
2. If 'a' overlaps 'b'(b.start <= a.end), merge them into a new interval such that:
	c.start = a.start
	c.end = max(a.end, b.end)
3. Repeat above two steps to merge 'c' with the next interval if it overlaps with 'c'.
*/

type Interval struct {
	Start int
	End   int
}

func NewInterval(start, end int) Interval {
	return Interval{
		Start: start,
		End:   end,
	}
}

// merge will return a slice of all intervals that overlap.
// Time Complexity: O(N * log N)
func merge(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}

	// Sorts intervals based on the start time.
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i].Start <= intervals[j].Start
	})

	mergedIntervals := make([]Interval, 0, len(intervals))
	start := intervals[0].Start
	end := intervals[0].End

	for _, interval := range intervals {
		if interval.Start <= end { // intervals are overlapping, adjust the 'end'
			end = max(interval.End, end)
		} else { // non-overlapping intervals, add previous interval and reset
			mergedIntervals = append(mergedIntervals, NewInterval(start, end))
			start = interval.Start
			end = interval.End
		}
	}

	// Add the last interval.
	mergedIntervals = append(mergedIntervals, NewInterval(start, end))

	return mergedIntervals
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
