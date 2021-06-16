package easy

//
// Link: https://www.algoexpert.io/questions/Validate%20Subsequence
//
//
// Time complexity O(n)
// Space complexity O(1)

func IsValidSubsequence(array []int, sequence []int) bool {
	seqIdx, arrIdx := 0, 0
	n := len(array)
	m := len(sequence)

	for arrIdx < n && seqIdx < m {
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx++
		}
		arrIdx++
	}
	return m == seqIdx
}
