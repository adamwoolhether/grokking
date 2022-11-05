package longestSubstringWithDistinctChars

/*
Given a string, find the length of the longest substring, which has all distinct characters.

Example 1:
Input: String="aabccbb"
Output: 3
Explanation: The longest substring with distinct characters is "abc".

Example 2:
Input: String="abbbb"
Output: 2
Explanation: The longest substring with distinct characters is "ab".

Example 3:
Input: String="abccde"
Output: 3
Explanation: Longest substrings with distinct characters are "abc" & "cde".
*/

/*
Time Complexity: O(N)
Space Complexity: O(K), where K is num of distinct chars in the string.
Worst case, K <= N, if there are no duplicates. Only 26 letters means
we can say the algo runs in fixed space O(1).
*/
func nonRepeatSubstring(str string) int {
	windowStart, maxLength := 0, 0
	charIndexMap := make(map[uint8]int)

	// Try to extend the range [windowStart:windowEnd] in this loop:
	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]

		// If the map already contains 'rightChar', shrink the window from
		// the beginning, so there is only one occurrence of rightChar.
		if _, ok := charIndexMap[rightChar]; ok {
			// This will drop any 'rightChar' after its previous index. If
			// 'windowStart' is ahead of the last index of 'rightChar', we keep "windowStart'
			windowStart = max(windowStart, charIndexMap[rightChar]+1)
		}

		// Inert 'rightChar' into the map.
		charIndexMap[rightChar] = windowEnd
		// Remember the maximum length so far.
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
