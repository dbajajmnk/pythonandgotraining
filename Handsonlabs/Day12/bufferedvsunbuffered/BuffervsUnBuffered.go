package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("============Buffered Vs Un Buffered=============")
	unbuffered := make(chan string)
	go func() {
		fmt.Println("Unbuffered Send Message")
		unbuffered <- "Go is Great"
		fmt.Println("Unbuffered Send Message Done")
	}()
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Recieving", <-unbuffered)

	//Create a channel with Buffer
	buffered := make(chan string, 4)

	fmt.Println("Sending message1")
	buffered <- "Message1"
	fmt.Println("Sending message2")
	buffered <- "Message2"
	fmt.Println("Sending message3")
	buffered <- "Message3"
	fmt.Println("Sending message4")
	buffered <- "Message4"

	fmt.Println("Buffered Channel", len(buffered), cap(buffered))

	go func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Print("Recive Message1", <-buffered)
		fmt.Print("Recive Message2", <-buffered)
		fmt.Print("Recive Message3", <-buffered)
		fmt.Print("Recive Message4", <-buffered)
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Thanks Sunil", "We are Done")

	//Send message as per capacity to channel Done

	// with Go routine recieve the message from channel putting some delay Done

	//Put More Delay Done

	//complete with done message Done

}
