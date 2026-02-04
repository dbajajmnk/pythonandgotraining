package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Select Statement")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch1 <- "Message One"
	}()
	go func() {
		time.Sleep(600 * time.Millisecond)
		ch2 <- "Message Two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Recieved", msg)
		case msg := <-ch2:
			fmt.Println("Recieved", msg)
		}

	}

	select {
	case msg := <-ch1:
		fmt.Println("Recieved", msg)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout After Demo test")

	}

	test := make(chan bool)
	go worker(test)

	time.Sleep(500 * time.Millisecond)
	test <- true

	fmt.Println("All Cases Covered")

	//Select test Done
	//Select Timout test Done
	//Select test Done
}
func worker(test <-chan bool) {
	for {
		select {
		case <-test:
			fmt.Println("test Happend")
			return
		default:
			fmt.Println("Working......")
			time.Sleep(100 * time.Millisecond)

		}

	}

}
