package longestSubstringWithSameLettersAfterReplacement

/*
Given a string with lowercase letters only, if you are allowed to replace no more than ‘k’
letters with any letter, find the length of the longest substring having the same letters after replacement.

Example 1:
Input: String="aabccbb", k=2
Output: 5
Explanation: Replace the two 'c' with 'b' to have a longest repeating substring "bbbbb".

Example 2:
Input: String="abbcb", k=1
Output: 4
Explanation: Replace the 'c' with 'b' to have a longest repeating substring "bbbb".

Example 3:
Input: String="abccde", k=1
Output: 3
Explanation: Replace the 'b' or 'd' with 'c' to have the longest repeating substring "ccc".
*/

/*
/*
Time Complexity: O(N), where 'N' is number of letters in string
Space Complexity: O(26), for each letter of alphabet, or O(1).
*/

func lengthOfLongestSubstring(str string, k int) int {
	windowStart, maxLength, maxRepeatLetterCount := 0, 0, 0
	frequencyMap := make(map[uint8]int)

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]
		frequencyMap[rightChar]++

		maxRepeatLetterCount := max(maxRepeatLetterCount, frequencyMap[rightChar])

		if windowEnd-windowStart+1-maxRepeatLetterCount > k {
			leftChar := str[windowStart]
			frequencyMap[leftChar]++
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
