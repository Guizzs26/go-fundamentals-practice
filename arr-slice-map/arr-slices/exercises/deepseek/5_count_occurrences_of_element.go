package deepseek

/*

5. Count Occurrences of Element:

- Task: Write a function that counts how many times a given integer appears in an array.

*/

func CountOccurrencesOfElement(nums [5]int, target int) int {
	count := 0
	for _, num := range nums {
		if num == target {
			count++
		}
	}
	return count
}
