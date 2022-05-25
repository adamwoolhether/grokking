package fruitsIntoBaskets

/*
Given a single row of fruit trees and two baskets, pick as many fruits
as possible to be placed in the baskets.

Constraints:
- Each basked can have one type of fruit.
- Start with any tree, but can't skip it once started.
- Pick exactly one fruit from every tree until you can't and have to pick a third fruit type.

Example 1:
Input: Fruit=['A', 'B', 'C', 'A', 'C']
Output: 3
Explanation: We can put 2 'C' in one basket and one 'A' in the other from the subarray ['C', 'A', 'C']

Example 2:
Input: Fruit = ['A', 'B', 'C', 'B', 'B', 'C']
Output: 5
Explanation: We can put 3 'B' in one basket and two 'C' in the other basket.
	This can be done if we start with the second letter: ['B', 'C', 'B', 'B', 'C']
*/

/*
Time Complexity: O(N+N), or O(N), where 'N' = num chars in the input.
Outer for loop runs for all chars,
inner loop processes each char once.
Space Complexity: O(1)O(1), a maximum of three fruits stored in the map.
*/
func maxFruitCountOf2Types(fruits []byte) int {
	windowStart, maxLength := 0, 0
	fruitFrequency := make(map[uint8]int)

	// Try to extend the range [windowStart:windowEnd] in this loop:
	for windowEnd := 0; windowEnd < len(fruits); windowEnd++ {
		rightFruit := fruits[windowEnd]
		fruitFrequency[rightFruit]++

		// Shrink the sliding window until we have 'k' distinct chars in fruitFrequency map.
		for len(fruitFrequency) > 2 {
			leftFruit := fruits[windowStart]
			fruitFrequency[leftFruit]--

			if fruitFrequency[leftFruit] == 0 {
				delete(fruitFrequency, leftFruit)
			}
			windowStart++ // shrink the window
		}
		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
