package stacks

import "strings"

/*
Statement:
Given a string consisting of lowercase English letters, repeatedly remove adjacent duplicate letters,
one pair at a time. Both members of a pair of adjacent duplicate letters need to be removed.

Constraints:
- 1 <= string.length <= 10^5
- string consists of lowercase English letters

How to solve:
1. Init stack
2. Iterate over string. If current char and char at stack's top match, increment it count.
   Otherwise, push it to stack and set its count to one.
3. If char at top of stack == 2, a duplicate is found, pop it from stack.
4. After traversal, return chars still stored in stack.
*/

// removeDuplicates will remove all subsequently duplicated chars from a string.
//
// Time Complexity: O(n)
// Space Complexity: O(n) in worst case.
func removeDuplicates(toCleanUp string) string {
	stack := NewStack[uint8]()

	for i := 0; i < len(toCleanUp); i++ {
		char := toCleanUp[i]

		if char == stack.Peak() {
			stack.Pop()
			continue
		}

		stack.Push(char)
	}

	var sb strings.Builder
	for _, c := range stack.data {
		sb.WriteByte(c)
	}

	return sb.String()
}

/*// removeDuplicates will remove all subsequently duplicated chars from a string.
// Uses an extra var to reduce stack lookups. Gives arbitrary, but noticeable speedup.
func removeDuplicatesWithPrev(toCleanUp string) string {
	stack := NewStack[uint8]()
	var prev uint8

	for i := 0; i < len(toCleanUp); i++ {
		char := toCleanUp[i]

		if char == prev {
			stack.Pop()
			prev = 0
			continue
		}

		if char == stack.Peak() {
			stack.Pop()
			continue
		}

		stack.Push(char)
		prev = char
	}

	var sb strings.Builder
	for _, c := range stack.data {
		sb.WriteByte(c)
	}

	return sb.String()
}*/
