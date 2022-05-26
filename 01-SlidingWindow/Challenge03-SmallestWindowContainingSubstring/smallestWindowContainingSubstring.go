package Challenge03_SmallestWindowContainingSubstring

/*
Given a string and a pattern, find the smallest substring in the given string which has all the character occurrences of the given pattern.

Example 1:
Input: String="aabdec", Pattern="abc"
Output: "abdec"
Explanation: The smallest substring having all characters of the pattern is "abdec"

Example 2:
Input: String="aabdec", Pattern="abac"
Output: "aabdec"
Explanation: The smallest substring having all characters occurrences of the pattern is "aabdec"

Example 3:
Input: String="abdbca", Pattern="abc"
Output: "bca"
Explanation: The smallest substring having all characters of the pattern is "bca".

Example 4:
Input: String="adcad", Pattern="abc"
Output: ""
Explanation: No substring in the given string has all characters of the pattern
*/

func minimumWindowSubstring(str, pattern string) string {
	windowStart, matched, subStringStart := 0, 0, 0
	minLength := len(str) + 1
	charFrequency := make(map[uint8]int, len(pattern))

	for i := 0; i < len(pattern); i++ {
		charFrequency[pattern[i]]++
	}

	// Extend the range [windowStart:windowEnd]
	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]

		if _, ok := charFrequency[rightChar]; ok {
			charFrequency[rightChar]--
			if charFrequency[rightChar] >= 0 { // Count every matching char, anytime the count is > 0
				matched++
			}
		}

		// Shrink the window if possible, finishing when a matched char is removed.
		for matched == len(pattern) {
			if minLength > windowEnd-windowStart+1 {
				minLength = windowEnd - windowStart + 1
				subStringStart = windowStart
			}

			leftChar := str[windowStart]
			windowStart++

			if _, ok := charFrequency[leftChar]; ok {
				// Decrement only when a useful occurrence of a matched char is
				// exiting the window, because there may be duplicate chars.
				if charFrequency[leftChar] == 0 {
					matched--
				}
				charFrequency[leftChar]++
			}
		}
	}

	if minLength > len(str) {
		return ""
	}

	return str[subStringStart : subStringStart+minLength]
}
