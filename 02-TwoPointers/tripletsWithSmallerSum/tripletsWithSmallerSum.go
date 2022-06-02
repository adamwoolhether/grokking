package tripletsWithSmallerSum

import (
	"sort"

	"golang.org/x/exp/constraints"
)

/*
Given an array arr of unsorted numbers and a target sum, count all triplets in it
such that arr[i] + arr[j] + arr[k] < target where i, j, and k are three different indices.
Write a function to return the count of such triplets.

 Let’s say during our iteration we are at number ‘X’, so we need to find ‘Y’ and ‘Z’ such that
X + Y + Z < target. At this stage, our problem translates into finding a pair whose
sum is less than “target - X” (as from the above equation Y + Z == target - X)

Example 1:
Input: [-1, 0, 2, 3], target=3
Output: 2
Explanation: There are two triplets whose sum is less than the target: [-1, 0, 3], [-1, 0, 2]

Example 2:
Input: [-1, 4, 2, 1, 3], target=5
Output: 4
Explanation: There are four triplets whose sum is less than the target:
   [-1, 1, 4], [-1, 1, 3], [-1, 1, 2], [-1, 2, 3]


Time Complexity:
	Sorting -  O(N * logN)
	searchPair func - O(N^2)
		Main while loop runs in O(N), nested for loop may take O(N)
	tripletWithSmallerSum overall takes O(N * logN + N^3), asymptotically
	equivalent to O(N^3).
*/

type numbers interface {
	constraints.Float | constraints.Integer
}

func tripletWithSmallerSum[T numbers](arr []T, target T) int {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	count := 0

	for i := 0; i < len(arr)-2; i++ {
		left, right := i+1, len(arr)-1

		for left < right {
			currSum := arr[i] + arr[left] + arr[right]

			if currSum < target {
				// We found the triplet, because arr[right] >= arr[left], so replace arr[right]
				// with any number between left and right to get a sum less that the target sum.
				count += right - left
				left++
			} else {
				right--
			}
		}
	}

	return count
}
