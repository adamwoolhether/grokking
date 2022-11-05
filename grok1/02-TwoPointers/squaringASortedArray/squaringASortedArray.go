package squaringASortedArray

import "golang.org/x/exp/constraints"

/*
Given a sorted array, create a new array containing squares
of all the numbers of the input array in the sorted order.

Due to presence of negative numbers, we find the first non-
negative number, and use two pointers to iterate the array.
The number returning a bigger square at any step is added to
the output array.
Alternatively, we could use two pointers starting at both ends.
Whichever pointer gives us the bigger square, add it to the
result and mobe to the next/previous number.

Example 1:
Input: [-2, -1, 0, 2, 3]
Output: [0, 1, 4, 4, 9]

Example 2:
Input: [-3, -1, 0, 1, 2]
Output: [0, 1, 1, 4, 9]

Time Complexity: O(N)
Space Complexity: O(N)
*/

type numbers interface {
	constraints.Float | constraints.Integer
}

func makeSquares[T numbers](arr []T) []T {
	arrLen := len(arr)
	squares := make([]T, arrLen)

	highestSquareIdx := arrLen - 1
	left, right := 0, arrLen-1

	for left <= right {
		leftSquare := arr[left] * arr[left]
		rightSquare := arr[right] * arr[right]

		if leftSquare > rightSquare {
			squares[highestSquareIdx] = leftSquare
			left++
		} else {
			squares[highestSquareIdx] = rightSquare
			right--
		}

		highestSquareIdx--
	}

	return squares
}
