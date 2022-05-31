package tripletSumToZero

import (
	"sort"

	"golang.org/x/exp/constraints"
)

/*
Given an array of unsorted numbers,
find all unique triplets in it that add up to zero.

Example 1:
Input: [-3, 0, 1, 2, -1, 1, -2]
Output: [-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]
Explanation: There are four unique triplets whose sum is equal to zero.

Example 2:
Input: [-5, 2, -1, -2, 3]
Output: [[-5, 2, 3], [-2, -1, 3]]
Explanation: There are two unique triplets whose sum is equal to zero.

Sorting -  O(N * logN)
	searchPair func - O(N)
	Because searchPair is called for every number in the input array,
	searchTriplets takes O(N * logN + N^2)
		This is asymptotically equivalent to O(N^2)
*/
type numbers interface {
	constraints.Float | constraints.Integer
}

func searchTriplets[T numbers](arr []T) [][3]T {
	// Sort the array, so all duplicates are next to each other.
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	triplets := [][3]T{}

	for i := 0; i < len(arr); i++ {
		// Skip identical elements to avoid duplicate triplets.
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		triplets = searchPair(arr, -arr[i], i+1, triplets)
	}

	return triplets
}

func searchPair[T numbers](arr []T, targetSum T, left int, triplets [][3]T) [][3]T {
	right := len(arr) - 1
	for left < right {
		currentSum := arr[left] + arr[right]
		if currentSum == targetSum {
			triplets = append(triplets, [3]T{-targetSum, arr[left], arr[right]})
			left++
			right--

			for left < right && arr[left] == arr[left-1] {
				left++
			}

			for left < right && arr[right] == arr[right+1] {
				right++
			}
		} else if targetSum > currentSum {
			left++
		} else {
			right--
		}
	}

	return triplets
}
