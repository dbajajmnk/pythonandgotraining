// File: day11_maps.go
// Topic: Maps in Go (Day 11)
// Run: go run day11_maps.go

package main

import "fmt"

func main() {
	fmt.Println("=== Maps in Go ===")

	// Create map
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println("map:", m)

	// Lookup
	v, ok := m["a"]
	fmt.Println("lookup a:", v, "ok:", ok)

	// Missing key
	v, ok = m["c"]
	fmt.Println("lookup c:", v, "ok:", ok)

	// Delete
	delete(m, "b")
	fmt.Println("after delete:", m)

	// Pass to function
	updateMap(m)
	fmt.Println("after function:", m)

	fmt.Println("=== Done ===")
}

func updateMap(m map[string]int) {
	m["x"] = 100
}
