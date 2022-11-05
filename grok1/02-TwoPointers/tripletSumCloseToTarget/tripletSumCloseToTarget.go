package tripletSumCloseToTarget

import (
	"math"
	"sort"

	"golang.org/x/exp/constraints"
)

/*
Given an array of unsorted numbers and a target number,
find a triplet in the array whose sum is as close to the target number as possible,
return the sum of the triplet.
If there are more than one such triplet, return the sum of the triplet with the smallest sum.

Similar to the Triplet Sum Close to Target, but we save the difference between the
triplet and target number at each step.

Example 1:
Input: [-2, 0, 1, 2], target=2
Output: 1
Explanation: The triplet [-2, 1, 2] has the closest sum to the target.

Example 2:
Input: [-3, -1, 1, 2], target=1
Output: 0
Explanation: The triplet [-3, 1, 2] has the closest sum to the target.

Example 3:
Input: [1, 0, 1, 1], target=100
Output: 3
Explanation: The triplet [1, 1, 1] has the closest sum to the target.

Time Complexity: Sorting takes O(N * logN)
	Overall is O(N * logN + N^2), which is asymptotically equivalent to O(N^2)
Space Complexity: O(N)
*/

type numbers interface {
	constraints.Float | constraints.Integer
}

func tripletSumCloseToTarget[T numbers](arr []T, targetSum T) T {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	smallestDiff := T(math.Inf(1))

	for i := 0; i < len(arr)-2; i++ {
		left, right := i+1, len(arr)-1

		for left < right {
			targetDiff := targetSum - arr[i] - arr[left] - arr[right]

			if targetDiff == 0 { // Triplet with exact sum found.
				return targetSum // Return sum of all numbers.
			}

			// Handle the smalled sum when we have more than one solution:
			if (math.Abs(float64(targetDiff)) < math.Abs(float64(smallestDiff))) ||
				(math.Abs(float64(targetDiff)) == math.Abs(float64(smallestDiff)) && targetDiff > smallestDiff) {
				smallestDiff = targetDiff
			}

			if targetDiff > 0 {
				left++
			} else {
				right--
			}
		}
	}

	return targetSum - smallestDiff
}
