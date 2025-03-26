package deepseek

/*

2. Find Maximum Element:

- Task: Write a function that returns the maximum value in an array of integers.

*/

func FindMaxElement(nums [5]int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
