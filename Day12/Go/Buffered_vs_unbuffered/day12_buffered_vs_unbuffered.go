// File: day12_buffered_vs_unbuffered.go
// Topic: Buffered vs Unbuffered Channels (Day 12)
// Run: go run day12_buffered_vs_unbuffered.go

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Unbuffered Channel Demo ===")
	unbuffered := make(chan string)

	go func() {
		fmt.Println("Sending on unbuffered channel")
		unbuffered <- "msg"
		fmt.Println("Send completed (unbuffered)")
	}()

	time.Sleep(300 * time.Millisecond)
	fmt.Println("Receiving:", <-unbuffered)

	fmt.Println("\n=== Buffered Channel Demo ===")
	buffered := make(chan string, 2)

	fmt.Println("Sending msg1")
	buffered <- "msg1"
	fmt.Println("Sending msg2")
	buffered <- "msg2"

	fmt.Println("Buffer len:", len(buffered), "cap:", cap(buffered))

	go func() {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Receiving:", <-buffered)
		fmt.Println("Receiving:", <-buffered)
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("=== Done ===")
}
