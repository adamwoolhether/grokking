package greedyTechniques

/*
Statement:There are n gas stations along a circular route, where the amount of gas at the i^{th} station is gas[i].
We have a car with an unlimited gas tank and it cost[i] of gas to travel from the i^{th}station to its next
(i+1)^{th}station. We begin the journey with an empty tank at one of the gas stations.

Given two integer arrays, gas and cost, return the starting gas station’s index if we can travel
around the circuit once in the clockwise direction. Otherwise, return −1.

If there exists a solution, it is guaranteed to be unique.

Constraints:
- gas.length == cost.length
- 1 <= gas.length <= 10^3
- 0 <= gas[i]
- cost[i] <= 10^3

How to solve:
1. Calculate total gas and cost from both arrays. If total gas < total cost, you can't complete loop.
2. Set start idx and gas available to 0
3. Traverse array, subtract current gas with corresponding element in cost array, add it to total gas.
4. If gas every becomes < 0, we can't complete the loop from that starting point, so increment starting index by one.
5. Return starting index at end of traversal.
*/

func gasStationJourney(gas []int, cost []int) int {
	var totalGas int
	var totalCost int

	steps := len(gas)
	for i := 0; i < steps; i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}

	if totalGas < totalCost {
		return -1
	}

	// If totalGas < totalCost, there must be a solution, no need to use nested loop.
	startIdx, gasAvailable := 0, 0
	for i := 0; i < steps; i++ {
		gasAvailable += gas[i] - cost[i]

		if gasAvailable < 0 {
			startIdx = i + 1
			gasAvailable = 0
		}
	}

	return startIdx
}
