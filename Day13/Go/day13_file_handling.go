package main

import (
    "fmt"
    "os"
)

func main() {
    // Write file
    os.WriteFile("sample.txt", []byte("Hello Go"), 0644)

    // Read file
    data, err := os.ReadFile("sample.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(data))
}
