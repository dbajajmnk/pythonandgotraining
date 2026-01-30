package main

import (
	"fmt"
)

/*
WHAT & WHY
Data types are the "contract" for memory + behavior.
Go is statically typed (compiler verifies types), but supports type inference.

This file covers:
- Primitive: int, float64, bool, string
- Composite: array, slice, map, struct
- Reference-like: pointer, slice, map
- Value vs reference behavior (important for interviews)
*/

func main() {
	fmt.Println("=== GO DATA TYPES & VARIABLES ===")

	// 1) Variables: var, :=, zero values
	var a int          // zero value = 0
	var b = 10         // inferred as int
	c := 20            // short declaration (function scope)
	var ok bool        // false
	var msg string     // ""
	var rate float64   // 0.0

	fmt.Println("Zero values:", a, ok, msg, rate)
	fmt.Println("Inferred:", b, c)

	// 2) Primitive
	age := 33
	salary := 95000.75
	active := true
	name := "Deepak"
	fmt.Println("Primitive:", age, salary, active, name)

	// 3) Array (fixed size, value type)
	var arr [3]int = [3]int{1, 2, 3}
	arr2 := arr // copy
	arr2[0] = 99
	fmt.Println("Array arr :", arr)
	fmt.Println("Array arr2:", arr2, " (arr unchanged because arrays are value-copied)")

	// 4) Slice (dynamic, reference-like)
	s := []int{1, 2, 3}
	s2 := s // points to same backing array (header copy)
	s2[0] = 77
	fmt.Println("Slice s :", s, " (changed!)")
	fmt.Println("Slice s2:", s2)

	// 5) Map (reference-like)
	m := map[string]int{"go": 90, "java": 85}
	m2 := m
	m2["go"] = 100
	fmt.Println("Map m :", m, " (changed!)")
	fmt.Println("Map m2:", m2)

	// 6) Struct (custom type)
	e := Employee{ID: 1, Name: "Deepak", Dept: "IT"}
	fmt.Println("Struct:", e)

	// 7) Pointers (explicit reference)
	x := 10
	increment(&x)
	fmt.Println("Pointer updated x:", x)

	// Practice mini-labs
	fmt.Println("\n--- Practice ---")
	fmt.Println("1) Create a slice of Employees and append 3 employees.")
	fmt.Println("2) Create a map[ int ]Employee as an index by ID.")
	fmt.Println("3) Write a function that updates Dept using pointer receiver.")
}

type Employee struct {
	ID   int
	Name string
	Dept string
}

func increment(n *int) {
	*n = *n + 1
}
