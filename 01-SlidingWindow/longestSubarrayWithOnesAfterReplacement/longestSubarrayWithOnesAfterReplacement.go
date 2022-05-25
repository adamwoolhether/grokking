package longestSubarrayWithOnesAfterReplacement

/*
Given an array containing 0s and 1s, if you are allowed to replace no more
than ‘k’ 0s with 1s, find the length of the longest contiguous subarray having all 1s.

Example 1:
Input: Array=[0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1], k=2
Output: 6
Explanation: Replace the '0' at index 5 and 8 to have the longest contiguous subarray of 1s having length 6.

Example 2:
Input: Array=[0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1], k=3
Output: 9
Explanation: Replace the '0' at index 6, 9, and 10 to have the longest contiguous subarray of 1s having length 9.
*/

/*
Time Complexity: O(N), where 'N' is the count of numbers in input aray.
*/
func lengthOfLongestSubstring(arr []int, k int) int {
	windowStart, maxLength, maxOnesCount := 0, 0, 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		if arr[windowEnd] == 1 {
			maxOnesCount++
		}

		// Current window size is from windowStart to windowEnd. We have a max of 1s
		// repeating a max of 'maxOnesCount' times. This means we can have a window with 'maxOnesCount'
		// 1s and the remaining 0s which should be replaced with 1s. If the remaining 0s are more than
		// 'k', then we need to shrink the window, because we can't replace more than 'k' 0s.
		if windowEnd-windowStart+1-maxOnesCount > k {
			if arr[windowStart] == 1 {
				maxOnesCount--
			}
			windowStart++
		}
		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
