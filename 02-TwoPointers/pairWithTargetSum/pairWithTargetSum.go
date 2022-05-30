package pairWithTargetSum

import "golang.org/x/exp/constraints"

/*
Given an array of sorted numbers and a target sum,
find a pair in the array whose sum is equal to the given target.

Write a function to return the indices of the two numbers (i.e. the pair)
such that they add up to the given target.

Example 1:
Input: [1, 2, 3, 4, 6], target=6
Output: [1, 3]
Explanation: The numbers at index 1 and 3 add up to 6: 2+4=6

Example 2:
Input: [2, 5, 9, 11], target=11
Output: [0, 2]
Explanation: The numbers at index 0 and 2 add up to 11: 2+9=11

A brute force solution would have a time complexity of O(N * logN).

Time Complexity: O(N), where 'N' is the total number of elements
in the given array.
Space Complexity: O(1)
*/

type numbers interface {
	constraints.Float | constraints.Integer
}

func pairWithTargetSum[T numbers](arr []T, targetSum T) []int {
	left, right := 0, len(arr)-1

	for left < right {
		currentSum := arr[left] + arr[right]

		if currentSum == targetSum {
			return []int{left, right}
		}

		if targetSum > currentSum {
			left++
		} else {
			right--
		}
	}

	return nil
}

/*
An alternate approach is to use a hash table, and iterate through the array
one number at a time.
Ex: If at position x, we need to find 'y' such that "x + y = target".
	1. Search for 'y', which is equivalent to "target - x" in the hash table.
	If there, then we've found the answer.
	2. Otherwise, insert 'x' in the hashtable to search it for later numbers

Time Complexity: O(N)
Space Complexity: O(N). Worst case is all elements pushed to hash table.
*/

func pairWithTargetSumAlternate[T numbers](arr []T, targetSum T) []int {
	nums := make(map[T]int, len(arr))
	for i, num := range arr {
		if _, ok := nums[targetSum-num]; ok {
			return []int{nums[targetSum-num], i}
		}

		nums[num] = i
	}

	return nil
}
