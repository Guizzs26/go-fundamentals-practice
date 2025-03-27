package deepseek

/*

4. Check if Array Contains Element:

- Task: Write a function that checks if an array contains a given integer.

*/

func CheckIfArrContainsElement(nums [5]int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}
