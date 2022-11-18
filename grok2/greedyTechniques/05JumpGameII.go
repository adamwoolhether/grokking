package greedyTechniques

/*
Statement:
In a single-player jump game, the player starts at one end of a series of squares, with the goal of reaching the last square.

At each turn, the player can take up to s steps towards the last square, where is the value of the current square.
For example, if the value of the current square is 3 the player can take either 3 steps, or 2 steps, or 1
step in the direction of the last square. The player cannot move in the opposite direction, that is,
away from the last square.
You’ve been provided with the nums integer array, representing the series of squares.
You’re initially positioned at the first index of the array. Find the minimum number
of jumps needed to reach the last index of the array.
You may assume that you can always reach the last index.

Constraints:
- 1<= nums.length <=10^3
- 0 <= nums[i] <= 10^3

How to solve:
1. Init three vars:
	`farthestJump` holds the farthest reachable index,
	`currentJump` marks the end of our current jump,
	`jump` stores the number of jumps.
2. Traverse array, reaching second to last index. On each ith iteration, update farthestJump += `nums[i]`
3. If i == `currentJump`, jump is completed. increment `jumps` by 1, set `currentJump` = `farthestJump`
4. Otherwise, don't update either `jumps` or `currentJump`, since jump isn't completed.
5. At end, `jumps` will contain the minimum number of jump required.
*/

func jumpGameTwo(nums []int) int {
	farthestJump := 0
	currentJump := 0
	jumps := 0

	for i := 0; i < len(nums)-1; i++ {
		farthestJump += nums[i]
		if i == currentJump {
			jumps++
			currentJump = farthestJump
		}
	}

	return jumps
}
