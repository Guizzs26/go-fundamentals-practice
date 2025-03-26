package katas

import "fmt"

// This file's simple objective is to train the syntax and intrinsic
// characteristics of arrays in the Go language, without algorithms.

// Kata 01 - Declaration and Initialization
func kata01DeclarationAndInitialization() {
	// Kata 01 - Declaration and Initialization

	// 1. Basic declaration with var and short syntax
	// var arr1 [3]int    // [0, 0, 0]
	// var arr2 [2]string // ["", ""]
	// var arr3 [4]bool   // [false, false, false, false]

	// arr4 := [4]int{}    // [0, 0, 0, 0, 0]
	// arr5 := [3]string{} // ["", "", ""]
	// arr6 := [2]bool{}   // [true, false]

	// // 2. Declaration + Explicit Initialization
	// var arr7 [3]int = [3]int{10, 20, 30}     // [10, 20, 30]
	// var arr8 [2]string = [2]string{"A", "B"} // ["A", "B"]

	// var arr9 = [4]int{1, 2, 3, 4}            // [1, 2, 3, 4]
	// var arr10 = [3]string{"Go", "Rust", "C"} // ["Go", "Rust", "C"]

	// arr11 := [3]float64{1.1, 2.2, 3.3} // [1.1, 2.2, 3.3]
	// arr12 := [2]bool{true, false}      // [true, false]

	// // 3. Partial Initialization = Zero value
	// var arr13 = [5]int{1, 2}   // [1, 2, 0, 0, 0]
	// var arr14 = [3]string{"A"} // ["A", "", ""]

	// arr15 := [4]bool{true, false} // [true, false, false, false]
	// arr16 := [6]int{10, 20}       // [10, 20, 0, 0, 0, 0]

	// // 4. Size inference (...)
	// var arr17 = [...]int{10, 20, 30}     // Tamanho = 3
	// var arr18 = [...]string{"Oi", "Bem"} // Tamanho = 2

	// arr19 := [...]float64{2.2, 3.3}       // Tamanho = 2
	// arr20 := [...]bool{true, false, true} // Tamanho = 3

	// // 6. Initialization with specific indexes
	// var arr21 = [5]int{1: 10, 3: 20}      // [0, 10, 0, 20, 0]
	// var arr22 = [4]string{1: "A", 2: "C"} // ["A", "", "C", ""]

	// arr23 := [6]int{2: 100, 5: 200} // [0, 0, 100, 0, 0, 200]
	// arr24 := [3]string{1: "B"}      // ["", "B", ""]
}

// Kata 02 - Access and Modification
func kata02AccessAndModification() {
	nums := [3]int{10, 20, 30}
	i2 := nums[1]
	fmt.Println(i2)
	nums[2] = 99
	fmt.Println(nums)
	// x := nums[4] // panic: index out of range
}

func kata03ArrayAsCopy() {
	strs := [2]string{"A", "B"}
	copyStrs := strs

	copyStrs[0] = "C"
	fmt.Println(copyStrs) // 0 -> C
	fmt.Println(strs)     // But here 0 -> A, because arrays are copied and not passed by reference
}

func kata04ArrayAsValue(nums [3]int) {
	for i := range nums {
		nums[i] *= 2
	}
}

func kata05ArrayWithPointer(nums *[4]int) {
	for i := range nums {
		nums[i] += 1
	}
}

func kata06ArrayLoops() {
	nums := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(nums); i++ {
		fmt.Printf("At position %d the value is %d", i, nums[i])
	}

	for _, value := range nums {
		fmt.Println(value)
	}

	for i := range nums {
		fmt.Println(i)
	}
}

func kata07ArrayComparision() {
	nums1 := [2]int{1, 2}
	nums2 := [2]int{1, 2}
	// nums3 := [3]int{1, 2, 3}
	nums4 := [2]int{1, 3}

	fmt.Println(nums1 == nums2) // true
	// fmt.Println(nums1 == nums3) // Invalid operation: nums1 == num2 (mismatched types)
	fmt.Println(nums1 == nums4) //false
}

func kata08Challenge(nums *[3]float64) {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	nums[2] = sum / float64(len(nums))
}

func main() {
	kata01DeclarationAndInitialization()
	kata02AccessAndModification()
	kata03ArrayAsCopy()

	//
	nums := [3]int{10, 20, 30}
	kata04ArrayAsValue(nums)
	fmt.Println(nums) // It doesn't change!

	nums2 := [4]int{9, 99, 999, 9999}
	kata05ArrayWithPointer(&nums2)
	fmt.Println(nums2)
	// Now the nums2 array has been modified because we are passing the reference in its memory to the function,
	// whatever the function changes in this array will be saved in memory.

	kata06ArrayLoops()
	kata07ArrayComparision()
	//
	nums3 := [3]float64{10.0, 20.0, 30.0}
	kata08Challenge(&nums3)
	fmt.Println(nums3) // nums[2] = avg now
}

func challengeInverseArr(nums *[3]int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
