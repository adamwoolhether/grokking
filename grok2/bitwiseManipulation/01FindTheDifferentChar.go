package bitwiseManipulation

import (
	"strings"
)

/*
Statement:
Given two strings, find the index of the extra character that is present in only one of the strings.

Constraints:
- 0 <= string.length <= 1000
- The strings consist of lowercase English letters.

How to solve:
1. Create `result` var with value of 0.
2. Traverse first string till end, performing a bitwise XOR operation of `result` on
   ASCII value of every char, updating `result` as needed.
3. Repeat step 2 on second string.
4. `Result` will correspond to the ASCII val of extra char. Return index of this char from the longer string.

*/

// extraCharIndex will return the index of the extra char that is present
// in one of the (longer) strings. One of the input strings must be longer.
//
// O(n)
// Where `n` is the length of the string.
func extraCharIndex(str1, str2 string) int {
	if len(str1) == len(str2) {
		return -1
	}

	var result = 0

	for i := range str1 {
		result = result ^ int(str1[i])
	}
	for i := range str2 {
		result = result ^ int(str2[i])
	}

	if len(str1) > len(str2) {
		return strings.Index(str1, string(rune(result)))
	}

	return strings.Index(str2, string(rune(result)))
}
