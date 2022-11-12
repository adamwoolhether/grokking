package stacks

import "unicode"

/*
Statement:
Given a string containing an arithmetic expression, implement a basic calculator that evaluates the expression string.
The expression string can contain integer numeric values and should be able to handle the “+” and “-” operators,
as well as “()” parentheses.

Constraints:
- 1 <= s.length <= 3 × 10^3
- consists of digits, “+”, “-”, “(”, and “)”.
- s represents a valid expression.
- “+” is not used as a unary operation
- “-” could be used as a unary operation
- There will be no two consecutive operators in the input.
- Every number and running calculation will fit in a signed 32-bit integer.

How to solve:
1. Init three vars:
   `number` will memory int form of current number in string = 0
   `sign` to memory multiplication value to change sign = 1
   `result` to memory the evaluated result = 0
2. When encountering a digit, update `number` by multiplying existing value by 10
   and adding integer value of digit char to it:
   number = number * 10 + digit
3. If encountering a '(', push value to result var and the sign on to stack. Reset `sign` and `result` to 1 and 0.
4. If encountering a '+' or '-', change `sign` var to 1 or -1 respectively. Evaluate the expression
   on left by multiplying the existing value of result var by sign value of the var and adding
   the number to this: result = number + (result * sign value).
   Also reset value of `number` to 0,
5. If encounter a ')', char. Update result var to evaluate the expression within parenthesis:
   result = number + ( result * sign value) + digit
   Then reset `number` to 0.
*/

type calculator struct {
	sign       int
	popSign    int
	workingNum int
	result     int
	memory     *Stack[int]
}

func newCalc() *calculator {
	return &calculator{
		sign:       1,
		popSign:    0,
		workingNum: 0,
		result:     0,
		memory:     NewStack[int](),
	}
}

func Calculator(expression string) int {
	calc := newCalc()

	for _, char := range expression {
		if unicode.IsDigit(char) {
			calc.workingNum = calc.workingNum*10 + int(char-'0')
			continue
		}

		switch char {
		case '+', '-':
			calc.addSub()
			calc.setSign(char)
		case '(':
			calc.store()
		case ')':
			calc.retrieve()
		}
	}

	return calc.sum()
}

func (c *calculator) addSub() {
	c.result += c.workingNum * c.sign
	c.clearWorkingNum()
}

func (c *calculator) store() {
	c.memory.Push(c.result)
	c.memory.Push(c.sign)
	c.setSign()
	c.clearResult()
}

func (c *calculator) retrieve() {
	c.result += c.sign * c.workingNum

	popSign := c.memory.Pop()
	c.result *= popSign

	secondVal := c.memory.Pop()
	c.result += secondVal
	c.clearWorkingNum()
}

func (c *calculator) sum() int {
	return c.result + c.workingNum*c.sign
}

// /////////////////////////////////////////////////////////////////

func (c *calculator) setSign(input ...rune) {
	if len(input) == 0 {
		c.sign = 1
		return
	}

	if input[0] == '+' {
		c.sign = 1
		return
	}

	c.sign = -1
}

func (c *calculator) clearWorkingNum() {
	c.workingNum = 0
}

func (c *calculator) clearResult() {
	c.result = 0
}
