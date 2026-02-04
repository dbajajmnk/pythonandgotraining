// File: day12_goroutines.go
// Topic: Goroutines (Day 12)
// Run: go run day12_goroutines.go

package main
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Goroutines Demo ===")

	// Lab 1: Simple goroutine
	go sayHello("A")

	// Lab 2 & 3: WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		sayHello("B")
	}()

	go func() {
		defer wg.Done()
		sayHello("C")
	}()

	// Lab 5: Loop variable trap fixed
	for i := 0; i < 3; i++ {
		i := i // capture
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Loop goroutine:", i)
		}()
	}

	wg.Wait()

	fmt.Println("=== Done ===")
}

func sayHello(name string) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Hello from", name)
}
