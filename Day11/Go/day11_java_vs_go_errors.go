package main

import "fmt"

func javaStyle() {
    // In Java, exception might be thrown here
    fmt.Println("Java-style: try/catch")
}

func goStyle() error {
    return fmt.Errorf("something went wrong")
}

func main() {
    err := goStyle()
    if err != nil {
        fmt.Println("Handled error explicitly:", err)
    }
}
