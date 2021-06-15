package easy

//
// Link: https://www.algoexpert.io/questions/Two%20Number%20Sum
//
//
// Time complexity O(n)
// Space complexity O(n)

func TwoNumberSum(array []int, target int) []int {
	lookupSet := make(map[int]bool)

	for i, value := range array {
		if i == 0 {
			lookupSet[value] = true
			continue
		}

		expectedValue := target - value

		if lookupSet[expectedValue] {
			return []int{expectedValue, value}
		} else {
			lookupSet[value] = true
		}
	}

	return []int{}
}
