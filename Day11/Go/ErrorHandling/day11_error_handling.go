// File: day11_error_handling.go
// Topic: Error Handling in Go (Day 11)
// Run: go run day11_error_handling.go

package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("=== Error Handling in Go ===")

	// Lab 1 & 2: Error return + early return
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("safeDivide error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Lab 3: Error propagation
	if err := processRequest(""); err != nil {
		fmt.Println("processRequest error:", err)
	}

	// Lab 4: Custom error
	err = validateAge(-1)
	fmt.Println("validateAge:", err)

	// Lab 5 & 6: panic + recover
	fmt.Println("\nPanic / Recover demo")
	runSafely()

	fmt.Println("=== Done ===")
}

// Standard error return
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Error propagation with context
func processRequest(input string) error {
	if err := validateInput(input); err != nil {
		return fmt.Errorf("processRequest failed: %w", err)
	}
	return nil
}

func validateInput(input string) error {
	if input == "" {
		return errors.New("input cannot be empty")
	}
	return nil
}

// Custom error type
type AgeError struct {
	Age int
}

func (e AgeError) Error() string {
	return fmt.Sprintf("invalid age: %d", e.Age)
}

func validateAge(age int) error {
	if age < 0 {
		return AgeError{Age: age}
	}
	return nil
}

// panic + recover boundary
func runSafely() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("About to panic...")
	panic("unexpected system failure")
}
