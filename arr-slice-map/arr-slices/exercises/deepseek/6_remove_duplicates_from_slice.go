package deepseek

/*

6. Remove Duplicates from Slice:

- Task: Write a function that removes duplicates from an slice and returns a new slice.

*/

func RemoveDuplicate(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}
