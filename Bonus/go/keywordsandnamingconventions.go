// go_master_keywords_naming_cheatsheet.go
//
// This file demonstrates:
// 1. Go keywords
// 2. Go naming conventions
// 3. WHY each rule exists (commented)
//


package main // package: ownership + namespace

import ( // import: explicit dependencies
	"fmt"
	"time"
)

// const: immutable compile-time value
const maxRetries = 3

// type: new named type
type UserID int

// struct: real-world entity (no class, no inheritance)
type User struct {
	ID     UserID // Exported field
	Name   string // Exported field
	active bool   // Private field
}

// interface: behavior contract (implicit implementation)
type Greeter interface {
	Greet() string
}

// func: method with short receiver name
func (u User) Greet() string {
	return "Hello " + u.Name
}

// error naming convention
var ErrNotFound = fmt.Errorf("resource not found")

func main() { // func main: entry point

	// var: mutable variable
	var retryCount int

	// if / else
	if retryCount < maxRetries {
		retryCount++
	} else {
		fmt.Println("Max retries reached")
	}

	// map: key-value store
	users := map[int]string{
		1: "Alice",
		2: "Bob",
	}

	// for / range / continue
	for id, name := range users {
		if id == 2 {
			continue
		}
		fmt.Println(id, name)
	}

	// chan: concurrency communication
	msgCh := make(chan string, 1)

	// go: goroutine
	go func() {
		msgCh <- "background work done"
	}()

	// select: channel coordination
	select {
	case msg := <-msgCh:
		fmt.Println(msg)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	default:
		fmt.Println("no message yet")
	}

	// defer: guaranteed cleanup
	defer fmt.Println("program exiting")

	// switch / case / default
	status := "ok"
	switch status {
	case "ok":
		fmt.Println("status ok")
	default:
		fmt.Println("unknown status")
	}

	// panic / recover (demonstration only)
	safeDivide(10, 0)

	return // return: explicit function exit
}

// panic + recover example
func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	}
	fmt.Println(a / b)
}
