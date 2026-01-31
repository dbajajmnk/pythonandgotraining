// File: day12_select.go
// Topic: select statement (Day 12)
// Run: go run day12_select.go

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== select Demo ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(600 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	// Lab 1: select on multiple channels
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Received:", msg)
		case msg := <-ch2:
			fmt.Println("Received:", msg)
		}
	}

	// Lab 4: Timeout pattern
	select {
	case msg := <-ch1:
		fmt.Println("Received:", msg)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout reached")
	}

	// Lab 5: done channel
	done := make(chan bool)
	go worker(done)

	time.Sleep(500 * time.Millisecond)
	done <- true

	time.Sleep(200 * time.Millisecond)
	fmt.Println("=== Done ===")
}

func worker(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
