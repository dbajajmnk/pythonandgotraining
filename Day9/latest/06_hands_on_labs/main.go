package main

import (
	"errors"
	"fmt"
)

/*
HANDS-ON LABS (Deep Practice)
This file combines multiple labs:
Lab A) Variables & zero values
Lab B) Struct + slice (data modeling)
Lab C) Map-based index (fast lookup)
Lab D) Pointer-based update (mutability control)
Lab E) Error-return pattern (production style)
*/

func main() {
	fmt.Println("=== HANDS-ON LABS (GO) ===")

	// Lab A
	var count int // 0
	fmt.Println("Lab A - zero value count:", count)

	// Lab B
	employees := []Employee{
		{ID: 1, Name: "Deepak", Dept: "IT"},
		{ID: 2, Name: "Asha", Dept: "HR"},
	}
	employees = append(employees, Employee{ID: 3, Name: "Ravi", Dept: "Finance"})
	fmt.Println("\nLab B - employees slice:", employees)

	// Lab C
	index := make(map[int]Employee)
	for _, e := range employees {
		index[e.ID] = e
	}
	fmt.Println("\nLab C - map index:", index)

	// Lab D (pointer update)
	err := UpdateDept(&employees[0], "Platform")
	must(err)
	fmt.Println("\nLab D - updated employee:", employees[0])

	// Lab E (error return, no exceptions)
	e, err := FindByID(index, 2)
	must(err)
	fmt.Println("\nLab E - found:", e)

	_, err = FindByID(index, 999)
	fmt.Println("\nLab E - expected error:", err)

	// Extra practice prompts
	fmt.Println("\n--- Extra Practice ---")
	fmt.Println("1) Add validation: Dept must be non-empty")
	fmt.Println("2) Add function: Filter employees by Dept")
	fmt.Println("3) Add function: Promote employee (change Dept + return copy)")
}

type Employee struct {
	ID   int
	Name string
	Dept string
}

func UpdateDept(e *Employee, dept string) error {
	if e == nil {
		return errors.New("nil employee")
	}
	if dept == "" {
		return errors.New("dept cannot be empty")
	}
	e.Dept = dept
	return nil
}

func FindByID(index map[int]Employee, id int) (Employee, error) {
	if index == nil {
		return Employee{}, errors.New("index map is nil")
	}
	e, ok := index[id]
	if !ok {
		return Employee{}, fmt.Errorf("employee not found: %d", id)
	}
	return e, nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
