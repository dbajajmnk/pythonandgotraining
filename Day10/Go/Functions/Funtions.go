// File: main.go
// Module: Functions & Multiple Return Values in Go (Single-file hands-on)
//
// Run:
//   go run main.go
//
// What this file covers (in ONE runnable program):
// - Basic functions (single return)
// - Multiple returns: (result, error), (value, ok), related values, value+metadata
// - Named returns + defer (error wrapping)
// - Variadic functions
// - Function parameters (callbacks) + returning functions (factory)
// - Practical, interview-style patterns

package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Go Functions & Multiple Return Values: Single-file Module ===")

	// 1) Single return: pure function
	fmt.Println("\n1) Single return")
	fmt.Println("add(10, 20) =>", add(10, 20))

	// 2) Multiple returns: result + error
	fmt.Println("\n2) Multiple returns (result, error)")
	q, err := safeDivide(10, 2)
	fmt.Println("safeDivide(10,2) =>", q, "err:", err)

	_, err = safeDivide(10, 0)
	fmt.Println("safeDivide(10,0) => err:", err)

	// 3) Multiple returns: value + ok (lookup pattern)
	fmt.Println("\n3) Multiple returns (value, ok) lookup")
	users := map[string]string{
		"u1": "Deepak",
		"u2": "Asha",
	}
	name, ok := lookupUser(users, "u2")
	fmt.Println("lookupUser(u2) =>", name, "ok:", ok)

	name, ok = lookupUser(users, "u9")
	fmt.Println("lookupUser(u9) =>", name, "ok:", ok)

	// 4) Multiple related results: min + max
	fmt.Println("\n4) Multiple related results (min, max)")
	minV, maxV, err := minMax([]int{8, 3, 10, 2, 7})
	fmt.Println("minMax =>", minV, maxV, "err:", err)

	// 5) Value + metadata: tokenize + count
	fmt.Println("\n5) Value + metadata (tokens, count)")
	tokens, count := tokenize("Go makes errors explicit")
	fmt.Println("tokenize =>", tokens, "count:", count)

	// 6) Named returns + defer wrapping (advanced)
	fmt.Println("\n6) Named returns + defer wrapping")
	score, err := computeScoreWithWrapping("  GOLD  ")
	fmt.Println("computeScoreWithWrapping =>", score, "err:", err)

	_, err = computeScoreWithWrapping("   ")
	fmt.Println("computeScoreWithWrapping (bad input) => err:", err)

	// 7) Variadic function
	fmt.Println("\n7) Variadic function")
	fmt.Println("sumAll(1,2,3,4) =>", sumAll(1, 2, 3, 4))
	fmt.Println("sumAll() =>", sumAll())

	// 8) Function as parameter (callback)
	fmt.Println("\n8) Function as parameter (callback validation)")
	isNonEmpty := func(s string) error {
		if strings.TrimSpace(s) == "" {
			return errors.New("empty string not allowed")
		}
		return nil
	}
	fmt.Println("validateInput('hello') =>", validateInput("hello", isNonEmpty))
	fmt.Println("validateInput('   ') =>", validateInput("   ", isNonEmpty))

	// 9) Returning function (factory)
	fmt.Println("\n9) Returning function (factory)")
	mkPrefixer := makePrefixer("ID-")
	fmt.Println("makePrefixer('ID-')('123') =>", mkPrefixer("123"))

	// 10) Show blank identifier '_' usage
	fmt.Println("\n10) Ignoring return values with '_'")
	_, ok = lookupUser(users, "u1")
	fmt.Println("ignored name, ok =>", ok)

	fmt.Println("\n=== Done ===")
}

// ------------------------------------------------------------
// 1) Single return value (pure function)
func add(a, b int) int {
	return a + b
}

// ------------------------------------------------------------
// 2) Multiple returns: (result, error)
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %.2f by zero", a)
	}
	return a / b, nil
}

// ------------------------------------------------------------
// 3) Multiple returns: (value, ok) lookup pattern
// If not found, return ("", false). This is NOT an error in many systems.
func lookupUser(users map[string]string, userID string) (string, bool) {
	v, ok := users[userID]
	return v, ok
}

// ------------------------------------------------------------
// 4) Multiple related results: (min, max) + error if input invalid
func minMax(nums []int) (int, int, error) {
	if len(nums) == 0 {
		return 0, 0, errors.New("minMax requires at least 1 number")
	}

	minV := nums[0]
	maxV := nums[0]
	for _, n := range nums[1:] {
		if n < minV {
			minV = n
		}
		if n > maxV {
			maxV = n
		}
	}
	return minV, maxV, nil
}

// ------------------------------------------------------------
// 5) Value + metadata: return data + count
func tokenize(s string) ([]string, int) {
	parts := strings.Fields(s)
	return parts, len(parts)
}

// ------------------------------------------------------------
// 6) Named returns + defer wrapping (advanced interview topic)
//
// Named returns allow `defer` to wrap/modify the return error before returning.
func computeScoreWithWrapping(rank string) (score int, err error) {
	start := time.Now()

	// Defer runs before function exits.
	// Because `err` is a named return, we can wrap it here.
	defer func() {
		// Add context if error happened
		if err != nil {
			err = fmt.Errorf("computeScoreWithWrapping failed: %w", err)
		}
		// Example of lightweight telemetry
		elapsed := time.Since(start)
		fmt.Println("  (debug) computeScoreWithWrapping took:", elapsed)
	}()

	r := strings.ToUpper(strings.TrimSpace(rank))
	if r == "" {
		return 0, errors.New("rank is empty after trimming")
	}

	switch r {
	case "GOLD":
		return 100, nil
	case "SILVER":
		return 80, nil
	case "BRONZE":
		return 60, nil
	default:
		return 0, fmt.Errorf("unknown rank: %s", r)
	}
}

// ------------------------------------------------------------
// 7) Variadic function: accepts 0..N ints
func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// ------------------------------------------------------------
// 8) Function as parameter: callback pattern
// This is the core of "strategy" / "policy injection".
func validateInput(input string, rule func(string) error) error {
	return rule(input)
}

// ------------------------------------------------------------
// 9) Returning a function: factory / middleware-ish pattern
func makePrefixer(prefix string) func(string) string {
	return func(s string) string {
		return prefix + s
	}
}
