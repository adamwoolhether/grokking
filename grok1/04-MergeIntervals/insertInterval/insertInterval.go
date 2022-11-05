package insertInterval

/*
Given a list of non-overlapping intervals sorted by their start time,
insert a given interval at the correct position and merge all necessary
intervals to produce a list that has only mutually exclusive intervals.

Example 1:
Input: Intervals=[[1,3], [5,7], [8,12]], New Interval=[4,6]
Output: [[1,3], [4,7], [8,12]]
Explanation: After insertion, since [4,6] overlaps with [5,7], we merged them into one [4,7].

Example 2:
Input: Intervals=[[1,3], [5,7], [8,12]], New Interval=[4,10]
Output: [[1,3], [4,12]]
Explanation: After insertion, since [4,10] overlaps with [5,7] & [8,12], we merged them into [4,12].

Example 3:
Input: Intervals=[[2,3],[5,7]], New Interval=[1,4]
Output: [[1,4], [5,7]]
Explanation: After insertion, since [1,4] overlaps with [2,3], we merged them into one [1,4].
*/

// Solution
/*
Because the input list is already sorted, we can come up with a faster solution
than simply running merge().

We need to find the correct index with the new interval can be placed, skipping all intervals
that end before the start of the new interval. (skip all with this condition: intervals[i].end < newInterval.Start

After finding the correct place, we can use a similar approach as merge().

1. Skip all intervals ending before the start of the new interval.
2. Call the last interval 'b' that doesn't satisfy the above condition.
	if b.start <= a.end; merge them into a new interval 'c'.
3. Repeat above steps to merge 'c' with the next overlapping level.
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

func insert(intervals []Interval, newInterval Interval) []Interval {
	merged := make([]Interval, 0, len(intervals)+1)
	i := 0

	// Skip and add to output all intervals that come before 'newInterval'.
	for i < len(intervals) && intervals[i].End <= newInterval.Start {
		merged = append(merged, intervals[i])
		i++
	}

	// Merge intervals that overlap with the newInterval.
	for i < len(intervals) && intervals[i].Start <= newInterval.End {
		newInterval.Start = min(intervals[i].Start, newInterval.Start)
		newInterval.End = max(intervals[i].End, newInterval.End)
		i++
	}

	// Insert the new interval.
	merged = append(merged, newInterval)

	// All remaining intervals to the merged slice.
	for i < len(intervals) {
		merged = append(merged, intervals[i])
		i++
	}

	return merged
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
