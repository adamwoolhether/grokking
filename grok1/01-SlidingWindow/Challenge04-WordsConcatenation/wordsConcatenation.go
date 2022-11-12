package Challenge04_WordsConcatenation

/*
Given a string and a list of words, find all the starting indices of substrings in the given string that are a concatenation of all the given words exactly once without any overlapping of words. It is given that all words are of the same length.

Example 1:
Input: String="catfoxcat", Words=["cat", "fox"]
Output: [0, 3]
Explanation: The two substring containing both the words are "catfox" & "foxcat".

Example 2:
Input: String="catcatfoxfox", Words=["cat", "fox"]
Output: [3]
Explanation: The only substring containing both the words is "catfox".
*/

/*
Time Complexity: O(N*M*Len), where 'N' and 'M' are
the number of chars in the input and pattern strings, respectively.
Len is the length of a word.
Space Complexity: O(M+N). In worse case, O(N) is needed for the
resulting list.
*/
func wordConcatenation(str string, words []string) []int {
	if len(words) == 0 || len(words[0]) == 0 {
		return nil
	}

	resultIdx := []int(nil)
	wordsCnt := len(words)
	wordLen := len(words[0])
	subStrLen := wordsCnt * wordLen
	wordFrequency := make(map[string]int, len(words))

	for _, w := range words {
		wordFrequency[w]++
	}

	// This solution isn't actually a proper sliding window. TODO: Implement new version.
	for i := 0; i < len(str)-subStrLen+1; i++ {
		wordsSeen := make(map[string]int, len(words))

		for j := 0; j < wordsCnt; j++ {
			nextWordIdx := i + j*wordLen

			// Get the next word from the string
			word := str[nextWordIdx : nextWordIdx+wordLen]
			if _, ok := wordFrequency[word]; !ok {
				break
			}

			wordsSeen[word]++

			// If word has higher frequency than required, no need to process further
			if wordsSeen[word] > wordFrequency[word] {
				break
			}

			if j == wordsCnt-1 { // store index if all words are found.
				resultIdx = append(resultIdx, i)
			}

		}
	}

	return resultIdx
}
