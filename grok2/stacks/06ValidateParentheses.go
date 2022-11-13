package stacks

/*
Statement:
Given a string that may consist of opening and closing parentheses, your task is to check if the string contains valid parenthesization or not.
The conditions to validate are:
Every opening parenthesis should be closed by the same kind of parenthesis. So, {)and [(]) strings are invalid.
Every opening parenthesis must be closed in the correct order. So, )( and ()(() are invalid.



Constraints:
- 1 <= n <= 100
- 1 <= logs.length <=500
- 0 <= functionID < n
- No two start/end events happen at the same time.
- Each function has a "end" log entry for each "start"

How to solve:

*/

func validParentheses(expression string) bool {
	stack := Stack[rune]{}

	brace := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	// Loop through each char.
	for _, char := range expression {
		switch char {
		case '(', '[', '{':
			stack.Push(char)
		case ')', ']', '}':
			poppedOpenBrace := stack.Pop()

			// If the stack was empty.
			if poppedOpenBrace == 0 {
				return false
			}

			// If the popped brace doesn't correspond to the closing brace.
			if check := brace[char]; check != poppedOpenBrace {
				return false
			}
		}
	}

	// There are remaining chars in the stack, it lacks a closing brace.
	if stack.Size() != 0 {
		return false
	}

	return true
}
