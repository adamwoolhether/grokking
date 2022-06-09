package happyNumber

import "golang.org/x/exp/constraints"

/*
Any number will be called a happy number if, after repeatedly replacing it with a number equal
to the sum of the square of all of its digits, leads us to number ‘1’.
All other (not-happy) numbers will never reach ‘1’. Instead, they will be stuck in a cycle of
numbers which does not include ‘1’.

Example 1:
Input: 23
Output: true (23 is a happy number)
Explanations: Here are the steps to find out that 23 is a happy number:
2^2 + 3^2 = 4 + 9 = 13
1^2 + 3^2 = 1 + 9 = 10
1^2 + 0^2 = 1 + 0 = 1

Input: 12
Output: false (12 is not a happy number)
Explanations: Here are the steps to find out that 12 is not a happy number:

1^2 + 2^2 = 1 + 4 = 5
5^2 = 25
2^2 + 5^2 = 4 + 25 = 29
2^2 + 9^2 = 4 + 81 = 85
8^2 + 5^2 = 64 + 25 = 89
...
...
5^2 + 8^2 = 25 + 64 = 90
Last step leads back to step 5, so we can never reach 1, so 12 is not a happy number.

Solution

The process, defined above, to find out if a number is a happy number or not,
always ends in a cycle. If the number is a happy number, the process will be
stuck in a cycle on number ‘1,’ and if the number is not a happy number then the
process will be stuck in a cycle with a set of numbers. As we saw in Example-2
while determining if ‘12’ is a happy number or not, our process will get stuck in a
cycle with the following numbers: 89 -> 145 -> 42 -> 20 -> 4 -> 16 -> 37 -> 58 -> 89

We saw in the LinkedList Cycle problem that we can use the Fast & Slow
pointers method to find a cycle among a set of elements. As we have described
above, each number will definitely have a cycle. Therefore, we will use the same
fast & slow pointer strategy to find the cycle and once the cycle is found, we will
see if the cycle is stuck on number ‘1’ to find out if the number is happy or not.
*/

/*
Time Complexity: O(logN)
	All unhappy numbers eventually get stuck in cycle. This means:
	1. If N <= 1000, then '1' will be reached in at most 1001 steps.
	2. If N > 1000, supposing number has 'M' digits, and next number is 'N1',
	then the sum of sqaures of digits 'N' is at most 9^2M, or 81M(when all diging of 'N' are '9'.

	Therefore:
	1. N1 < 81M
	2. M = log(N + 1)
	3. So: N1 < 81 * log(N + 1) => N1 = O(logN)

Space Complexity: O(1)
*/

func findHappyNumber[T constraints.Integer](num T) bool {
	fast, slow := num, num

	for {
		slow = findSquareSum(slow)                // move one step
		fast = findSquareSum(findSquareSum(fast)) // move two steps
		if slow == fast {                         // found the cycle
			break
		}
	}

	return slow == 1
}

func findSquareSum[T constraints.Integer](num T) T {
	var sum T

	for num > 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}

	return sum
}
