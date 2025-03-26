package main

import (
	"arr-slice-map/arr/exercises/deepseek"
	"fmt"
)

func main() {
	// Deepseek
	var nums [5]int = [5]int{10, 20, 30, 40, 50}

	fmt.Println(deepseek.SumArray(nums))
	fmt.Println(deepseek.FindMaxElement(nums))
	fmt.Println(deepseek.ReverseArray(nums))
	fmt.Println(deepseek.CheckIfArrContainsElement(nums, 2))
	fmt.Println(deepseek.CheckIfArrContainsElement(nums, 20))
	fmt.Println(deepseek.CountOccurrencesOfElement(nums, 10))
}
