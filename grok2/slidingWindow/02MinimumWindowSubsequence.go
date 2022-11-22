package slidingWindow

import "math"

/*
Statement:
Given strings str1 and str2, find the minimum (contiguous) substring subStr of str1,
such that every character of str2 appears in subStr in the same order as it is present in str2.
If there is no window in str1 that covers all characters in str2, return an empty string.
If there are multiple minimum-length windows, return the one with the leftmost starting index.

Constraints:
- 1 <= str1.length <= 2*10^3
- 1<= str2.length <= 100
- str1 and str2 consist of uppercase and lowercase English letters.

Steps:
1. Init two pointers `index1` and `index2` at index 0, one for each string.
2. If char at index for each string is the same, increment both pointers.
3. Create two new pointers, `start` and `end` when `index2` reaches the end of str2
   `start` = `index1` and `end` = `start` + 1
4. Decrement `start` until `index2` is less than 0
5. Decrement `index2` only if char at `start` in str1 is equal to `index2` at str2
6. Calculate substring length by subtracting values of end and start vars
7. If length is less than current minimum length, update length and `minSubsequence` string.
8. Repeat until `index1` reaches end of str1.
*/

func minWindow(str1 string, str2 string) string {
	index1, index2 := 0, 0
	length := math.MaxInt64

	var minSubsequence string

	for index1 < len(str1) {
		if str1[index1] == str2[index2] {
			index2++

			if index2 == len(str2) {
				start, end := index1, index1+1

				index2--

				for index2 >= 0 {
					if str1[start] == str2[index2] {
						index2--
					}

					start--
				}
				start++

				if end-start < length {
					length = end - start

					minSubsequence = str1[start:end]
				}
				index1 = start
				index2 = 0
			}
		}

		index1++
	}

	return minSubsequence
}
