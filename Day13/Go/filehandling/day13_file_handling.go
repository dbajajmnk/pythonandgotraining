// File: day13_file_handling.go
// Topic: File Handling in Go (Day 13)
// Run: go run day13_file_handling.go

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== File Handling Demo ===")

	filename := "sample.txt"

	// Lab 1 & 2: Create and write file
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("Hello Go File Handling\n")

	// Lab 4: Append mode
	file2, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	file2.WriteString("Appending new line\n")

	// Lab 5: Buffered write
	writer := bufio.NewWriter(file2)
	writer.WriteString("Buffered write line\n")
	writer.Flush()

	// Lab 3 & 6: Read file
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("File content:\n", string(data))

	fmt.Println("=== Done ===")
}
