// File: day11_builtin_functions.go
// Topic: Built-in Functions in Go (Day 11)
// Run: go run day11_builtin_functions.go

package main

import "fmt"

func main() {
	fmt.Println("=== Built-in Functions in Go ===")

	// len & cap
	s := make([]int, 0, 5)
	fmt.Println("len:", s, len(s), "cap:", cap(s))

	s = append(s, 1)
	s = append(s, 2)
	fmt.Println("after append:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, 3)
	s = append(s, 4)
	s = append(s, 5)
	s = append(s, 6)
	fmt.Println("after growth:", s, "len:", len(s), "cap:", cap(s))

	// append with ...
	a := []int{10, 20}
	b := []int{30, 40}
	a = append(a, b...)

	fmt.Println("merged slice:", a)

	// copy
	src := []int{1, 2, 3}
	//dst := make([]int, len(src))
	//fmt.Println("Dst", dst)
	//n := copy(dst, src)
	dst := src
	dst[0] = 999
	//fmt.Println("copied:", dst, "count:", n)
	fmt.Println("Dst",dst)
	fmt.Println("original:", src)

	// make
	m := make(map[string]int)
	m["x"] = 1
	fmt.Println("map:", m)

	// delete
	delete(m, "x")
	delete(m, "missing") // safe
	fmt.Println("map after delete:", m)

	fmt.Println("=== Done ===")
}
