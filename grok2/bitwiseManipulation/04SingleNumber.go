package bitwiseManipulation

import "fmt"

/*
Statement:
Given an array of integers in which every element in the array appears twice except for one,
find the element that occurs once.

Constraints:
- 1 <= nums.length <= 3 * 10^3
- -3 x 10^3 <= nums[i] <= 3 * 10^3
- Each element in the array appears twice except for one element which only appears once.

How to solve:
1. Initialize a var `result` with value of 0.
2. Traverse input array, performing bitwise XOR of every element with `result`, updating it each time.
3. `result` will be the number that has only appeared once.
*/

// singleNumber will determine which number in a given array
// only appears once. All other numbers should appear twice.
//
// O(n)
// Where n is the total number of elements in the array.
func singleNumber(nums []int) int {
	result := 0
	for i := range nums {
		fmt.Println(result)
		result ^= nums[i]
	}

	return result
}
