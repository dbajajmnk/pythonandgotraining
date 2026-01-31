// File: day15_building_binaries.go
// Topic: Building Binaries in Go
// Build examples:
//   go build -o app
//   GOOS=linux GOARCH=amd64 go build -o app-linux
//   go build -ldflags "-s -w -X main.version=1.0" -o app

package main

import (
	"fmt"
	"runtime"
)

var version = "dev"

func main() {
	fmt.Println("App Version:", version)
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
}
