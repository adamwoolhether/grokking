package subarraysWithProductLessThanTarget

import "golang.org/x/exp/constraints"

/*
Given an array with positive numbers and a positive target number,
find all of its **contiguous** subarrays whose product is less than the target number.

Example 1:
Input: [2, 5, 3, 10], target=30
Output: [2], [5], [2, 5], [3], [5, 3], [10]
Explanation: There are six contiguous subarrays whose product is less than the target.

Example 2:
Input: [8, 2, 6, 5], target=50
Output: [8], [2], [8, 2], [6], [2, 6], [5], [6, 5]
Explanation: There are seven contiguous subarrays whose product is less than the target.

Time Complexity:
	Main for-loop takes O(N), but subarrays can take up to O(N^@) in worst case.
	Overall, it takes O(N^3)
Space Complexity:
	O(n^3)
*/

type numbers interface {
	constraints.Float | constraints.Integer
}

func findSubarrays[T numbers](arr []T, target T) [][]T {
	product, left := T(1), 0
	result := [][]T{}

	for right := 0; right < len(arr); right++ {
		product *= arr[right]
		for product >= target && left < len(arr) {
			product /= arr[left]
			left++
		}

		// Because the product of all numbres from left to right is less than the target,
		// all subarrays from left to right will have a product less than the target too.
		// Therefore, start with a subarray containing only the arr[right] to avoid duplicates.
		tempList := []T{}
		for i := right; i > left-1; i-- {
			tempList = append([]T{arr[i]}, tempList...)
			result = append(result, tempList)
		}
	}

	return result
}
