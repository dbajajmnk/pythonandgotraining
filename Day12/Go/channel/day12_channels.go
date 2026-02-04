// File: day12_channels.go
// Topic: Channels (Day 12)
// Run: go run day12_channels.go

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Channels Demo ===")

	// Lab 1 & 2: Simple channel + synchronization
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine"
	}()

	msg := <-ch
	fmt.Println("Received:", msg)

	// Lab 3: Range and close
	numCh := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			numCh <- i
		}
		close(numCh)
	}()

	for v := range numCh {
		fmt.Println("Range received:", v)
	}

	// Lab 4: Directional channels
	done := make(chan bool)
	go producer(done)
	<-done

	fmt.Println("=== Done ===")
}

func producer(done chan<- bool) {
	fmt.Println("Producer done")
	done <- true
}
