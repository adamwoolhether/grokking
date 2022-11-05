package Challenge02_StringAnagrams

/*
Given a string and a pattern, find all anagrams of the pattern in the given string.

Every anagram is a permutation of a string. As we know, when we are not allowed to repeat characters while finding permutations of a string, we get N!N! permutations (or anagrams) of a string having NN characters. For example, here are the six anagrams of the string “abc”:

abc
acb
bac
bca
cab
cba
Write a function to return a list of starting indices of the anagrams of the pattern in the given string.

Example 1:
Input: String="ppqp", Pattern="pq"
Output: [1, 2]
Explanation: The two anagrams of the pattern in the given string are "pq" and "qp".

Example 2:
Input: String="abbcabc", Pattern="abc"
Output: [2, 3, 4]
Explanation: The three anagrams of the pattern in the given string are "bca", "cab", and "abc".
*/

/*
Time Complexity: O(N+M), where 'N' and 'M' are
the number of chars in the input and pattern strings, respectively.
Space Complexity: O(M). In worse case, the pattern has all distinct chars.
*/
func findStringAnagrams(str, pattern string) []int {

	windowStart, matched := 0, 0
	charFrequency := make(map[uint8]int, len(pattern))
	resultIdx := make([]int, 0, len(str))

	for i := 0; i < len(pattern); i++ {
		charFrequency[pattern[i]]++
	}

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]

		if _, ok := charFrequency[rightChar]; ok {
			charFrequency[rightChar]--
			if charFrequency[rightChar] == 0 {
				matched++
			}
		}

		// All chars matched, record the index.
		if matched == len(charFrequency) {
			resultIdx = append(resultIdx, windowStart)
		}

		// Shrink the sliding window.
		if windowEnd >= len(pattern)-1 {
			leftChar := str[windowStart]
			windowStart++

			if _, ok := charFrequency[leftChar]; ok {
				if charFrequency[leftChar] == 0 {
					matched-- // decrement counter before butting char back
				}
				charFrequency[leftChar]++ // put the char back
			}
		}
	}

	return resultIdx
}
