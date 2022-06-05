package Challenge01_QuadrupleSumToTarget

import (
	"sort"

	"golang.org/x/exp/constraints"
)

/*
Problem Statement

Given an array of unsorted numbers and a target number, find all unique quadruplets in it,
whose sum is equal to the target number.

Example 1:
Input: [4, 1, 2, -1, 1, -3], target=1
Output: [-3, -1, 1, 4], [-3, 1, 1, 2]
Explanation: Both the quadruplets add up to the target.

Example 2:
Input: [2, 0, -1, 1, -2, 2], target=2
Output: [-2, 0, 2, 2], [-1, 0, 1, 2]
Explanation: Both the quadruplets add up to the target.

Time Complexity:
	O(M + N), where 'M' and 'N' are the lengths
	of the two input strings respectively.
Space Complexity:
	O(1)
*/

type numbers interface {
	constraints.Float | constraints.Integer
}

func searchQuadruplets[T numbers](arr []T, target T) [][4]T {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	quadruplets := [][4]T{}

	for i := 0; i < len(arr)-3; i++ {
		// skip the same element to avoid deplicate quadruplets
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		for j := i + 1; j < len(arr)-2; j++ {
			// skip the same element to avoid deplicate quadruplets
			if j > 1 && arr[j] == arr[j-1] {
				continue
			}

			quadruplets = searchPairs(arr, target, i, j, quadruplets)
		}
	}

	return quadruplets
}

func searchPairs[T numbers](arr []T, target T, first, second int, quadruplets [][4]T) [][4]T {
	left, right := second+1, len(arr)-1

	for left < right {
		currentSum := arr[first] + arr[second] + arr[left] + arr[right]

		if currentSum == target {
			quadruplets = append(quadruplets, [4]T{arr[first], arr[second], arr[left], arr[right]})
			left++
			right--

			// Skip the same elements on left or right to avoid duplicates
			for left < right && arr[left] == arr[left-1] {
				left++
			}
			for left < right && arr[right] == arr[right+1] {
				right--
			}
		} else if currentSum < target { // we need a pair with larger currentSum
			left++
		} else { // we need a pair with smaller currentSum
			right--
		}
	}

	return quadruplets
}
