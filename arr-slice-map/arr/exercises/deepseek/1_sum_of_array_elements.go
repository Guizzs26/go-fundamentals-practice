package deepseek

/*

1. Sum of Array Elements:

- Task: Write a function that returns the sum of all elements in an array of integers.

*/

func SumArray(nums [5]int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
