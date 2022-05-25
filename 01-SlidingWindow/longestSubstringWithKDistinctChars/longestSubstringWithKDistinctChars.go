package longestSubstringWithKDistinctChars

/*
Given a string, find the length of the longest substring
in it with no more than K distinct characters.
You can assume that K is less than or equal to the length of the given string.

Example 1:
Input: String="araaci", K=2
Output: 4
Explanation: The longest substring with no more than '2' distinct characters is "araa".

Example 2:
Input: String="araaci", K=1
Output: 2
Explanation: The longest substring with no more than '1' distinct characters is "aa".

Example 3:
Input: String="cbbebi", K=3
Output: 5
Explanation: The longest substrings with no more than '3' distinct characters are "cbbeb" & "bbebi".
*/

/*
Time Complexity: O(N+N), or O(N). Outer for loop runs for all chars,
inner loop processes each char once.
Space Complexity: O(K), storing a maximum of K+1 chars in the map.
*/
func longestSubstringWithKDistinct(str string, k int) int {
	windowStart, maxLength := 0, 0

	charFrequency := make(map[uint8]int)

	// Try to extend the range [windowStart:windowEnd] in this loop:
	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]
		charFrequency[rightChar]++

		// Shrink the sliding window until we have 'k' distinct chars in charFrequency map.
		for len(charFrequency) > k {
			leftChar := str[windowStart]
			charFrequency[leftChar]--

			if charFrequency[leftChar] == 0 {
				delete(charFrequency, leftChar)
			}

			windowStart++ // shrink the window
		}

		// remember that maximum length thus far.
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
