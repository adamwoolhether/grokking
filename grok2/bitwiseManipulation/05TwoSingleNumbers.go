package bitwiseManipulation

/*
Statement:
Given a non-empty array arr, in which exactly two elements appear once and all the other elements
appear twice, return the two elements that come only once.

Constraints:
2 <= arr.length <= 3*10^4
-2^31 <= arr[i] <= 2^31 -1

How to solve:
1. Init a var `bitwiseSum` with val 0
2. Traverse array, performing bitwise XOR of `bitwiseSum` on each element, updating it each time.
3. Get rightmost 1-bit of `bitwiseSum` in a var `bitMask`
4. Init a var `result = 0`. Traverse over array elements, checking if bit is set.
   If yes, get XOR with `result`, updating its value with each computation.
5. `result` and `result ^ bitwiseSum` will correspond to the two single numbers.

*/

// twoSingleNumbers will determin which two numbers in an array
// do not have a matching pair.
//
// O (n)
func twoSingleNumbers(arr []int) []int {
	bitwiseSum := 0

	for _, v := range arr {
		bitwiseSum ^= v
	}

	// Get the least significant bit.
	bitwiseMask := bitwiseSum & (-bitwiseSum)

	results := 0
	for _, v := range arr {
		if bitwiseMask&v != 0 {
			results ^= v
		}
	}

	return []int{results, bitwiseSum ^ results}
}
