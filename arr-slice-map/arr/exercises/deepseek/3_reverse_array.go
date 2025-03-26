package deepseek

/*

3. Reverse an Array:

- Task: Write a function that reverses an array in-place.

*/

// [10,20,30,40,50] - Input
// [50,40,30,20,10] - Output

func ReverseArray(nums [5]int) [5]int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
