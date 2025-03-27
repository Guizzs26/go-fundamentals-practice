package deepseek

/*

7. Merge Two Sorted Arrays:

- Task: Write a function that merges two sorted arrays into a single sorted array.

*/

func MergeSortedSlices(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[i] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[i])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}
