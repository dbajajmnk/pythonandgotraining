// File: day11_arrays.go
// Topic: Arrays in Go (Day 11 - Topic 1)
// Covers:
// - Array declaration
// - Value semantics (copy behavior)
// - Passing arrays to functions
// - Comparison with slices
//
// Run:
//   go run day11_arrays.go

package main

import "fmt"

func main() {
	fmt.Println("=== Arrays in Go ===")

	// Lab 1: Declare & Initialize
	arr1 := [3]int{1, 2, 3}
	fmt.Println("Original array:", arr1)

	// Lab 2: Assignment copy behavior
	arr2 := arr1 // COPY happens here
	arr2[0] = 100

	fmt.Println("After modification:")
	fmt.Println("arr1:", arr1) // unchanged
	fmt.Println("arr2:", arr2) // modified copy

	// Lab 3: Passing array to function (copy again)
	fmt.Println("\nPassing array to function:")
	modifyArray(arr1)
	fmt.Println("After function call, arr1:", arr1)

	// Lab 4: Compare with slice
	fmt.Println("\nArray vs Slice comparison:")
	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Println("Slice after function call:", slice)

	fmt.Println("=== Done ===")
}

// Function receives a COPY of the array
func modifyArray(a [3]int) {
	a[1] = 999
	fmt.Println("Inside modifyArray:", a)
}

// Function receives a reference-like slice
func modifySlice(s []int) {
	s[1] = 999
	fmt.Println("Inside modifySlice:", s)
}
