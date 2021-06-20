package easy

//
// Link: https://www.algoexpert.io/questions/Two%20Number%20Sum
//
//
// Time complexity O(n)
// Space complexity O(1)

// input: [1, 2, 3, 5, 6, 8, 9]
// output: [1, 4, 9, 15, 36, 64, 81]

func SortedSquaredArray(array []int) []int {
	n := len(array)
	sortedSquared := make([]int, n)
	smIdx := 0
	lgIdx := n - 1

	for i := n - 1; i >= 0; i-- {
		smValue := array[smIdx]
		lgValue := array[lgIdx]

		if Abs(smValue) > Abs(lgValue) {
			sortedSquared[i] = smValue * smValue
			smIdx++
		} else {
			sortedSquared[i] = lgValue * lgValue
			lgIdx--
		}
	}

	return sortedSquared
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
