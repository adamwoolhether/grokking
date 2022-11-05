package Challenge02_ComparingStringsContainingBackspaces

/*
Problem Statement

Given two strings containing backspaces (identified by the character ‘#’), check if the two strings are equal.

Example 1:

Input: str1="xy#z", str2="xzz#"
Output: true
Explanation: After applying backspaces the strings become "xz" and "xz" respectively.
Example 2:

Input: str1="xy#z", str2="xyz#"
Output: false
Explanation: After applying backspaces the strings become "xz" and "xy" respectively.
Example 3:

Input: str1="xp#", str2="xyz##"
Output: true
Explanation: After applying backspaces the strings become "x" and "x" respectively.
In "xyz##", the first '#' removes the character 'z' and the second '#' removes the character 'y'.
Example 4:

Input: str1="xywrrmp", str2="xywrrmu#p"
Output: true
Explanation: After applying backspaces the strings become "xywrrmp" and "xywrrmp" respectively.

Solution

To compare the given strings, first, we need to apply the backspaces.
An efficient way to do this would be from the end of both the strings.
We can have separate pointers, pointing to the last element of the given strings.
We can start comparing the characters pointed out by both the pointers to see if
the strings are equal. If, at any stage, the character pointed out by any of the
pointers is a backspace (’#’), we will skip and apply the backspace until we have
a valid character available for comparison.
*/

func backspaceCompare(str1, str2 string) bool {
	idx1, idx2 := len(str1)-1, len(str2)-1

	for idx1 >= 0 || idx2 >= 0 {
		nextIdx1 := getNextValidCharIndex(str1, idx1)
		nextIdx2 := getNextValidCharIndex(str2, idx2)

		if nextIdx1 < 0 && nextIdx2 < 0 { // reached the end of both strings
			return true
		}
		if nextIdx1 < 0 || nextIdx2 < 0 { // reached end of one string
			return false
		}
		if str1[nextIdx1] != str2[nextIdx2] { // check if strings are eual
			return false
		}

		idx1 = nextIdx1 - 1
		idx2 = nextIdx2 - 1
	}

	return true
}

func getNextValidCharIndex(str string, idx int) int {
	backspaceCount := 0

	for idx >= 0 {
		if str[idx] == '#' { // backspace char found
			backspaceCount++
		} else if backspaceCount > 0 { // non-backspace char
			backspaceCount--
		} else {
			break
		}

		idx-- // skip a backspace or valid char
	}

	return idx
}
