// File: day11_go_vs_java_exceptions.go
// Topic: Comparison with Java Exceptions (Go perspective)
// Run: go run day11_go_vs_java_exceptions.go
//
// NOTE:
// Java-style example is shown as comments.
// Go-style example is executable.

package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("=== Go vs Java Exceptions ===")

	// ---- Java-style thinking (conceptual) ----
	/*
	try {
	    int result = divide(10, 0);
	} catch (ArithmeticException e) {
	    System.out.println(e.getMessage());
	}
	*/

	// ---- Go-style thinking ----
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Go error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Error propagation comparison
	if err := serviceLayer(); err != nil {
		fmt.Println("Service error:", err)
	}

	// panic vs exception
	fmt.Println("\npanic vs exception demo")
	runSafely()

	fmt.Println("=== Done ===")
}

// Go-style error return
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Propagation with context
func serviceLayer() error {
	if err := repoLayer(); err != nil {
		return fmt.Errorf("serviceLayer failed: %w", err)
	}
	return nil
}

func repoLayer() error {
	return errors.New("database unavailable")
}

// panic as bug, not flow control
func runSafely() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered panic:", r)
		}
	}()
	panic("programmer bug")
}
