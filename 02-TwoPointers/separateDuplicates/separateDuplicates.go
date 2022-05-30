package separateDuplicates

import "golang.org/x/exp/constraints"

/*
Given an array of sorted numbers, separate all duplicates from it in-place.
You should not use any extra space;
move all duplicates at the end of the array and after moving
return the length of the subarray that has no duplicate in it.

Example 1:
Input: [2, 3, 3, 3, 6, 9, 9]
Output: 4
Explanation: The first four elements after removing the duplicates will be [2, 3, 6, 9].

Example 2:
Input: [2, 2, 2, 11]
Output: 2
Explanation: The first two elements after removing the duplicates will be [2, 11].

Time Complexity: O(N), where 'N' is the total number of elements in the array
Space Complexity: O(1)
*/

type numbers interface {
	constraints.Float | constraints.Integer
}

func removeDuplicates[T numbers](arr []T) int {
	nextNonDuplicate := 1

	for _, v := range arr {
		if arr[nextNonDuplicate-1] != v {
			arr[nextNonDuplicate] = v
			nextNonDuplicate++
		}
	}

	return nextNonDuplicate
}

/*
Problem 1: Given an unsorted array of numbers and a target ‘key’,
remove all instances of ‘key’ in-place and return the new length of the array.

Example 1:
Input: [3, 2, 3, 6, 3, 10, 9, 3], Key=3
Output: 4
Explanation: The first four elements after removing every 'Key' will be [2, 6, 10, 9].

Example 2:
Input: [2, 11, 2, 2, 1], Key=2
Output: 2
Explanation: The first two elements after removing every 'Key' will be [11, 1].

Time Complexity: O(N), where 'N' is the total number of elements in the array.
Space Complexity: O(1)
*/

func removeElement[T numbers](arr []T, key T) int {
	nextElement := 0

	for _, v := range arr {
		if v != key {
			arr[nextElement] = v
			nextElement++
		}
	}

	return nextElement
}
