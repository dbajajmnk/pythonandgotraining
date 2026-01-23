package main

import (
    "fmt"
    "time"
)

func main() {
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

    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
}
