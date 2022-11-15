package greedyTechniques

/*
Statement:
In a single-player jump game, the player starts at one end of a series of squares, with the goal of reaching the last square.
At each turn, the player can take up to s steps towards the last square, where s is the value of the current square.
For example, if the value of the current square is 3, the player can take either 3 steps, or 2 steps, or 1
step in the direction of the last square. The player cannot move in the opposite direction, that is, away from the last square.

You have been tasked with writing a function to validate whether a player can win a given game or not.
You’ve been provided with the nums integer array, representing the series of squares. The player starts at the first index and,
following the rules of the game, tries to reach the last index.
If the player can reach the last index, your function returns TRUE; otherwise, it returns FALSE.

Constraints:
- 1 <= nums.length <= 10^3
- 0 <= nums[i] <= 10^3

How to solve:
1. Set last element as target
2. Traverse array from backwards
3. If current index is reachable from any preceding index based on the preceding index's value,
   update it as the new target.
4. If you reach the first index without finding a previous index that can reach the current target, return false.
5. If you can move each current target back all the way to the first index, the path is available, return true.
*/

// jumpGame will determine if a person starting at the first square is able to
// reach the last, if they can only take at most the amount of steps given by
// the square's value.
//
// Time Complexity: O(n)
// Space: O(1)
func jumpGame(nums []int) bool {
	target := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= target {
			target = i
		}
	}

	return target == 0
}
