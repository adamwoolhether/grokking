package smallestSubarrayWithGreaterSum

import (
	"math"

	"golang.org/x/exp/constraints"
)

/*
Given an array of positive numbers and a positive number ‘S,’
find the length of the smallest contiguous subarray whose sum
is greater than or equal to ‘S’. Return 0 if no such subarray exists.

Example 1:
Input: [2, 1, 5, 2, 3, 2], S=7
Output: 2
Explanation: The smallest subarray with a sum greater than or equal to '7' is [5, 2].

Example 2:
Input: [2, 1, 5, 2, 8], S=7
Output: 1
Explanation: The smallest subarray with a sum greater than or equal to '7' is [8].

Example 3:
Input: [3, 4, 1, 1, 6], S=8
Output: 3
Explanation: Smallest subarrays with a sum greater than or equal to '8' are [3, 4, 1]
or [1, 1, 6].
*/

type numbers interface {
	constraints.Float | constraints.Integer
}

/*
Time Complexity: O(N)
	Outer for loop runs for all elements, inner loop processes each
	element only once, resulting in O(N+N), which is asymptotically equivalent to O(N)
*/
func smallestSubarraySum[T numbers](s T, arr []T) T {
	windowSum := T(0)
	windowStart := 0
	minLength := T(math.Inf(1))

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd] // add the next element

		// Shrink the window as small as possible until 'windowSum' is smaller than 's'.
		for windowSum >= s {
			minLength = min(minLength, T(windowEnd-windowStart+1))
			windowSum -= arr[windowStart]
			windowStart += 1
		}
	}

	if minLength == T(math.Inf(1)) {
		return 0
	}

	return minLength
}

func min[T numbers](x, y T) T {
	if x < y {
		return x
	}

	return y
}
