package main

import "fmt"

type Counter struct {
	Value int
}

// Pointer receiver
func (c *Counter) Increment() {
	c.Value++
}

func main() {
	c := Counter{Value: 1}
	c.Increment()
	fmt.Println(c.Value)
}
