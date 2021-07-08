package easy

import "sort"

//
// Link: https://www.algoexpert.io/questions/Non-Constructible%20Change
//
// Time complexity O(nlogn)
// Space complexity O(1)

// input: coins=[5, 7, 1, 1, 2, 3, 22]
// output: 20

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)
	change := 0
	for _, coin := range coins {
		if coin > change+1 {
			return change + 1
		}
		change += coin
	}
	return change + 1
}
