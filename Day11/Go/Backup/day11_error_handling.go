package main

import (
    "errors"
    "fmt"
)

func readConfig(path string) (string, error) {
    if path == "" {
        return "", errors.New("path cannot be empty")
    }
    return "config-data", nil
}

func main() {
    data, err := readConfig("")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Config:", data)
}
