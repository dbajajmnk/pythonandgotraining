// File: day13_go_modules.go
// Topic: Go Modules (Day 13)
//
// Steps:
//   go mod init github.com/yourname/myapp
//   go get github.com/google/uuid@latest
//   go mod tidy
//
// Example go.mod:
//
// module github.com/yourname/myapp
//
// go 1.21
//
// require github.com/google/uuid v1.6.0
//
// --------------------------------

package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("=== Go Modules Demo ===")
	id := uuid.New()
	fmt.Println("Generated UUID:", id.String())
}
