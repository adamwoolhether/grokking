package maxSumSubarryOfSizeK

import "golang.org/x/exp/constraints"

/*
Given an array of positive numbers and a positive number ‘k’,
find the maximum sum of any contiguous subarray of size ‘k’.

Example 1:
Input: [2, 1, 5, 1, 3, 2], k=3
Output: 9
Explanation: Subarray with maximum sum is [5, 1, 3].

Example 2:
Input: [2, 3, 4, 1, 5], k=2
Output: 7
Explanation: Subarray with maximum sum is [3, 4].
*/

type numbers interface {
	constraints.Float | constraints.Integer
}

/*
Time Complexity: O(N)
Space Complexity: O(1)
*/
func findMaxSumSubArray[T numbers](k int, arr []T) T {
	maxSum, windowSum := T(0), T(0)
	windowStart := 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd] // add the next element

		// Slide the window if we haven't hit window size of 'k'.
		if windowEnd >= k-1 {
			maxSum = max(maxSum, windowSum)
			windowSum -= arr[windowStart] // subtract the element leaving the window
			windowStart++                 // slide the window ahead
		}
	}

	return maxSum
}

/*
Time: O(N*K)
*/
func findMaxSumSubArrayBruteForce[T numbers](k int, arr []T) T {
	maxSum, windowSum := T(0), T(0)

	for i := 0; i < len(arr)-k+1; i++ {
		windowSum = 0
		for j := i; j < i+k; j++ {
			windowSum += arr[j]
		}

		maxSum = max(maxSum, windowSum)
	}

	return maxSum
}

func max[T numbers](x, y T) T {
	if x > y {
		return x
	}

	return y
}
