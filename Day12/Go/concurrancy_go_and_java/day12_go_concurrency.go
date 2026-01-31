// File: day12_go_concurrency.go
// Topic: Go Concurrency vs Java Threads (Go side)
// Run: go run day12_go_concurrency.go

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Go Concurrency Demo ===")

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(200 * time.Millisecond)
			fmt.Println("Worker", id, "done")
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed")
}
