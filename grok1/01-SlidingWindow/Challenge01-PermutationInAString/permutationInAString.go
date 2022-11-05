package Challenge01_PermutationInAString

/*
Given a string and a pattern, find out if the string contains any permutation of the pattern.

Permutation is defined as the re-arranging of the characters of the string. For example, “abc” has the following six permutations:

abc
acb
bac
bca
cab
cba
If a string has ‘n’ distinct characters, it will have n!n! permutations.

Example 1:
Input: String="oidbcaf", Pattern="abc"
Output: true
Explanation: The string contains "bca" which is a permutation of the given pattern.

Example 2:
Input: String="odicf", Pattern="dc"
Output: false
Explanation: No permutation of the pattern is present in the given string as a substring.

Example 3:
Input: String="bcdxabcdy", Pattern="bcdyabcdx"
Output: true
Explanation: Both the string and the pattern are a permutation of each other.

Example 4:
Input: String="aaacb", Pattern="abc"
Output: true
Explanation: The string contains "acb" which is a permutation of the given pattern.
*/

/*
Time Complexity: O(N+M), where 'N' and 'M' are
the number of chars in the input and pattern strings, respectively.
Space Complexity: O(M). In worse case, the pattern has all distinct chars.
*/
func findPermutation(str, pattern string) bool {
	windowStart, matched := 0, 0
	charFrequency := make(map[uint8]int, len(pattern))

	for i := 0; i < len(pattern); i++ {
		charFrequency[pattern[i]]++
	}

	// Match all characters from 'charFrequency' with the current window,
	// trying to extend the range [windowStart, windowEnd]
	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]

		if _, ok := charFrequency[rightChar]; ok {
			// Decrement the frequency of the matched character
			charFrequency[rightChar]--
			if charFrequency[rightChar] == 0 {
				matched++
			}
		}

		// All chars in the window have been matched
		if matched == len(charFrequency) {
			return true
		}

		// Shrink the sliding window.
		if windowEnd >= len(pattern)-1 {
			leftChar := str[windowStart]
			windowStart++

			if _, ok := charFrequency[leftChar]; ok {
				if charFrequency[leftChar] == 0 {
					matched--
				}
				charFrequency[leftChar]++
			}
		}
	}

	return false
}
