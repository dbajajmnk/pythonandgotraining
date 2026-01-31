package main

import "fmt"

func safeDivide(a, b int) int {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    if b == 0 {
        panic("division by zero")
    }
    return a / b
}

func main() {
    result := safeDivide(10, 0)
    fmt.Println("Result:", result)
}
