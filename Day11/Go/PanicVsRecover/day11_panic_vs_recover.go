// File: day11_panic_vs_recover.go
// Topic: panic vs recover (Day 11)
// Run: go run day11_panic_vs_recover.go

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== panic vs recover ===")

	// Lab 1 & 2: Panic without recover
	// uncomment to see crash
	// panic("fatal bug")

	// Lab 3 & 4: Recover with defer
	safeRun()

	// Lab 6: Goroutine panic with recover
	fmt.Println("\nGoroutine panic demo")
	done := make(chan bool)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in goroutine:", r)
			}
			done <- true
		}()
		panic("goroutine failure")
	}()
	<-done

	fmt.Println("=== Done ===")
}

func safeRun() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("About to panic safely...")
	panic("something went wrong")
}
