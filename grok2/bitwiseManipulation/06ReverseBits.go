package bitwiseManipulation

/*
Statement:
Given an integer n, we need to calculate the 32-bit unsigned integer it would be if we reversed its bits.
When we say “reverse” we don’t mean flipping the 0s to 1s and vice versa,
but simply reversing the order in which they appear – from left-to-right to right-to-left.
We need to return the integer the reversed bits result in.

Constraints:
0 <= n <= n^31

How to solve:
1. Use a step counter at 31.
2. Loop while `number` is greater than 0.
3. Perform bitwise AND operation: 1 AND `number`.
4. Left-shift result of AND operation as many times as the value of step counter.
5. Add result of left-shift operation to an accumulator.
6. Divide `number` by 2 and decrement step counter.
7. Return the accumulator.


Not sure why tests say this is wrong...
*/

func reverseBits(n int) uint32 {
	var result uint32
	num := uint32(n)

	for i := 0; i < 32; i++ {
		result = result<<1 + num&1
		n >>= 1
	}

	return result
}
