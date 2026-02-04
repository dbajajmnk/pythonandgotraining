package main

import (
	"fmt"
)

func main() {
	fmt.Println("Channels lession Start")
	ch := make(chan string)
	go func() {

		ch <- "Hello From Reciever"

	}()
	msg := <-ch
	fmt.Println("Message Recieved", msg)

	numCh := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			numCh <- i
			fmt.Println(numCh)

		}
		close(numCh)

	}()
	
	for data := range numCh {

		fmt.Println("Value", data)
	}

	directional := make(chan bool)
	go producer(directional)
	<-directional

	fmt.Println("Channel Lab is Done")

}
func producer(done chan<- bool) {
	fmt.Println("Producer is active")
	done <- true

}
