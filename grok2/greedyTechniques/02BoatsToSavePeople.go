package greedyTechniques

import (
	"sort"
)

/*
Statement:
You’re given an array, people, where people[i] is the weight of the i^ith person,
and an infinite number of boats, where each boat can carry a maximum weight, limit.
Each boat carries, at most, two people at the same time.
This is provided that the sum of the weight of those people is under or equal to the weight limit.
You need to return the minimum number of boats to carry every person in the array.

Constraints:
- 1 <= people.length <= 5x10^30
- 1 <= people[i] <= limit <= 3 x 10^3

How to solve:
1. Sort the array. Init three vars to mark the heaviest and lightest person in list and calculate num of needed boats.
2. Lightest var will point fo first element, heaviest to last (because we already sorted it)
3. If weight of heaviest + lightest person < limit, then
   increment lightest var, decrement heaviest var by one, and increment boats by one.
   If not, only decrement the heaviest pointer and increase boat count.
4. Return the boat count.
*/

// rescueBoats will return the amount of boats needed to rescue people
// given the weight limit for each boat, with a cap of two people per boat.
//
// Time Complexity: O(n log n)
// Space: O(n)
func rescueBoats(people []int, limit int) int {
	sort.Ints(people)

	lightest := 0
	heaviest := len(people) - 1
	boats := 0

	for lightest <= heaviest {
		if people[lightest]+people[heaviest] <= limit {
			lightest++
			heaviest--
			boats++
		} else {
			heaviest--
			boats++
		}
	}

	return boats
}
