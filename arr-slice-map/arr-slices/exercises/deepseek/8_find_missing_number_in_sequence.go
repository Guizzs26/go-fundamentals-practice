package deepseek

/*

8. Find Missing Number in Sequence:

- Task: Given an array of n-1 integers in the range 1..n, find the missing number.

*/

func FindMissingNumber(arr []int) int {
	n := len(arr) + 1

	expectedSum := n * (n + 1) / 2 // S = n.(n + 1) / 2 -> If n = 5 -> 5.(5+1)/2 = 5.(6)/2 = 15
	actualSum := 0
	for _, v := range arr {
		actualSum += v
	}
	return expectedSum - actualSum
}
