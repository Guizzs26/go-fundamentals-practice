package main

import (
	"arr-slice-map/arr-slices/exercises/deepseek"
	"fmt"
)

func main() {
	// Deepseek
	var nums [5]int = [5]int{10, 20, 30, 40, 50}
	nums2 := []int{10, 20, 30, 10, 40, 50, 60, 10, 80, 20, 30}
	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8}

	fmt.Println(deepseek.SumArray(nums))
	fmt.Println(deepseek.FindMaxElement(nums))
	fmt.Println(deepseek.ReverseArray(nums))
	fmt.Println(deepseek.CheckIfArrContainsElement(nums, 2))
	fmt.Println(deepseek.CheckIfArrContainsElement(nums, 20))
	fmt.Println(deepseek.CountOccurrencesOfElement(nums, 10))
	fmt.Println(deepseek.RemoveDuplicate(nums2))
	fmt.Println(deepseek.MergeSortedSlices(a, b))
}
