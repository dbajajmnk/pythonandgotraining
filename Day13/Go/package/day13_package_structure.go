// File: day13_package_structure.go
// Topic: Go Package Structure (Day 13)
//
// Example project layout:
//
// myapp/
// ├── go.mod
// ├── cmd/
// │   └── app/
// │       └── main.go
// ├── internal/
// │   └── service/
// │       └── service.go
// ├── pkg/
// │   └── util/
// │       └── mathutil.go
//
// This file demonstrates package usage conceptually.
// Real packages must live in separate folders.

package main

import "fmt"

func main() {
	fmt.Println("=== Go Package Structure ===")
	fmt.Println("See comments in this file for recommended layout")
}
