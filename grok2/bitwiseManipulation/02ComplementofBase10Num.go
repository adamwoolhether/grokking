package bitwiseManipulation

import "math"

/*
Statement:
For any n positive number in base 10, return the complement of its
binary representation as an integer in base 10.

Constraints:
- 0 <= n <= 10^9

How to solve:
1. Calculate the number of bits require to store the positive integer.
   a. Calculate the binary logarithm
   b. Round down to next whole integer.
   c. Add 1
2. Create an all bits set against the number of bits of input value.
3. Flip all occurrences of 1s to 0s, and 0s to 1s by computing XOR
4. Convert the binary back to base 10 and return the complement.
*/

// power calculates the power of a given integer.
// O(1)
func power(digit int, power int) int {
	res := 1

	for i := 0; i < power; i++ {
		res *= digit
	}

	return res
}

func findBitwiseComplement(num int) int {
	if num == 0 {
		return 1
	}

	// 1.
	bitCount := math.Floor(math.Log2(float64(num))) + 1

	// 2.
	allBitsSet := power(2, int(bitCount)) - 1

	// 3 & 4
	return num ^ allBitsSet
}
