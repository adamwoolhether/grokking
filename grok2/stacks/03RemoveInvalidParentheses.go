package stacks

/*
Statement:
Given a string with matched and unmatched parentheses, remove the minimum number of parentheses so
that the resulting string is a valid parenthesization.

Constraints:
- 1 <= string.length <= 10^5
- string consists of lowercase English letters or parentheses.

How to solve:
1. Traverse, track parenthesis and indices.
2. If match found, pop from stack.
3. Create new string omitting indices matching those still in stack.
*/

type set struct {
	char rune
	idx  int
}

// minRemoveParentheses will remove parentheses from the string that don't have
// a correlated opening/closing parentheses.
//
// Time Complexity: O(n)
// String traversal = O(n), Compile result = O(n) = O(n+n) = O(n)
// Space Complexity: O(n) in worst case, assuming string is all unmatched parentheses.
func minRemoveParentheses(s string) string {
	stack := NewStack[set]()

	for i, char := range s {
		switch char {
		case '(':
			stack.Push(set{char, i})
		case ')':
			if stack.Size() == 0 {
				continue
			}

			if stack.Peak().char == '(' {
				stack.Pop()
			} else {
				stack.Push(set{char, i})
			}
		default:
			continue
		}
	}

	for stack.Size() > 0 {
		popped := stack.Pop()
		s = s[0:popped.idx] + s[popped.idx+1:]
	}

	return s
}

// minRemoveParenthesesWithMap is slightly less efficient than the above implementation.
func minRemoveParenthesesWithMap(s string) string {
	stack := NewStack[set]()
	skipIdx := make(map[int]struct{})

	for i, char := range s {
		switch char {
		case '(':
			stack.Push(set{char, i})
			skipIdx[i] = struct{}{}
		case ')':
			if stack.Size() == 0 {
				continue
			}

			topOfStack := stack.Peak()
			if topOfStack.char == '(' {
				stack.Pop()
				delete(skipIdx, topOfStack.idx)
			} else {
				stack.Push(set{char, i})
				skipIdx[i] = struct{}{}
			}
		default:
			continue
		}
	}

	for stack.Size() > 0 {
		popped := stack.Pop()
		s = s[0:popped.idx] + s[popped.idx+1:]
	}

	return s
}
