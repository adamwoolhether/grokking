package maxSumSubarryOfSizeK

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

type numConstraints interface {
	int | uint | int32 | int64 |
	float32 | float64
}

func max[T numConstraints](x, y T) T {
	if x > y {
		return x
	}

	return y
}

/*
Time Complexity: O(N)
Space Complexity: O(1)
*/
func findMaxSumSubArray[T numConstraints](k int, arr []T) T {
	var maxSum, windowSum T
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
func findMaxSumSubArrayBruteForce[T numConstraints](k int, arr []T) T {
	var maxSum, windowSum T

	for i := 0; i < len(arr)-k+1; i++ {
		windowSum = 0
		for j := i; j < i+k; j++ {
			windowSum += arr[j]
		}

		maxSum = max(maxSum, windowSum)
	}

	return maxSum
}
