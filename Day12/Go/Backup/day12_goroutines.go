package main

import (
    "fmt"
    "time"
)

func worker(id int) {
    fmt.Println("Worker", id, "started")
    time.Sleep(500 * time.Millisecond)
    fmt.Println("Worker", id, "finished")
}

func main() {
    for i := 1; i <= 3; i++ {
        go worker(i) // goroutine
    }

    time.Sleep(1 * time.Second) // wait for goroutines
}
