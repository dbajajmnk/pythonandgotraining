package main

import (
    "log"
    "os"
)

func main() {
    file, _ := os.Create("app.log")
    log.SetOutput(file)

    log.Println("Application started")
}
