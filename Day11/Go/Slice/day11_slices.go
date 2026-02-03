// File: day11_slices.go
// Topic: Slices in Go (Day 11)
// Run: go run day11_slices.go

package main

import "fmt"

func main() {
	fmt.Println("=== Slices in Go ===")

	// Create slice
	s := []int{1, 2, 3}
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	// append growth
	s2 := s[:1]
	fmt.Println("s2 before append:", s2, "cap:", cap(s2))

	s2 = append(s2, 100)
	fmt.Println("s2 after append:", s2)
	fmt.Println("s after append:", s)

	// copy to avoid sharing
	dst := make([]int, len(s))
	copy(dst, s)
	dst[0] = 999
	fmt.Println("dst:", dst)
	fmt.Println("s unchanged:", s)

	fmt.Println("=== Done ===")
}
