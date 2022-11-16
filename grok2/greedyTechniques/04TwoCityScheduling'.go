package greedyTechniques

import (
	"sort"
)

/*
Statement:
A company is planning to interview people. You're given the array costs where
costs[i] = [aCost_i, bCost_i]. The cost of flying the i^{th} person to city A
is aCost_i, and the cost of flying the i^{th}, person to city B is bCost_i.

Return the minimum cost to fly every person to a city such that exactly n people arrive in each city.

Constraints:
- costs.length = 2n
- 2 <= costs.length <= 100
- 1 <= aCost_i, bCost_i <= 1000

How to solve:
1. Init an array to store extra amount each person would cost to travel to A instead of B.
2. Along with difference, also append the cost of each person takes to travel to city A and city B.
3. Sort array in ascending order, based on difference between cost of travelling to A and B.
4. Calculate the minimum possible cost to send people evenly to the two cities.
*/
// twoCityScheduling will find the cheapest way to purchase plane tickets for people, deciding which
// city to fly them from the two given.
//
// Time Complexity: O(n log n)
// Space Complexity: O(n + m)
func twoCityScheduling(costs [][]int) int {
	amt := make([][3]int, len(costs))

	for i, v := range costs {
		amt[i] = [3]int{v[0] - v[1], v[0], v[1]}
	}

	sort.SliceStable(amt, func(i, j int) bool {
		return amt[i][0] < amt[j][0]
	})

	cost, start, end := 0, 0, len(amt)-1
	for start <= end {
		cost += amt[start][1]
		cost += amt[end][2]

		start++
		end--
	}

	return cost
}
