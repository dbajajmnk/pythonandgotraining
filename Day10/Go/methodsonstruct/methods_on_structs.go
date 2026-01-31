// File: methods_on_structs.go
// Module: Methods on Structs (Go) — Single Executable File
//
// Run:
//   go run methods_on_structs.go
//
// Covers:
// - Value receivers vs pointer receivers
// - Encapsulation with methods
// - Interface implementation via methods
// - Embedding and method promotion

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Methods on Structs (Go) ===")

	// Lab 1: Value receiver (read-only)
	u := User{name: "Deepak", age: 30}
	fmt.Println("Name:", u.Name())
	fmt.Println("Age:", u.Age())

	// Lab 2: Pointer receiver (mutation)
	u.Birthday()
	fmt.Println("After Birthday, Age:", u.Age())

	// Lab 3: Interface satisfaction
	var p Printer
	p = u // User satisfies Printer
	p.Print()

	// Lab 4: Embedding and method promotion
	admin := Admin{
		User:  u,
		Role: "ADMIN",
	}
	admin.Print() // promoted method
	fmt.Println("Role:", admin.Role)

	fmt.Println("=== Done ===")
}

// ------------------ Data Model ------------------

type User struct {
	name string
	age  int
}

// ------------------ Methods ------------------

// Value receiver: read-only
func (u User) Name() string {
	return u.name
}

func (u User) Age() int {
	return u.age
}

// Pointer receiver: mutation
func (u *User) Birthday() {
	u.age++
}

// Method used for interface satisfaction
func (u User) Print() {
	fmt.Printf("User{name:%q, age:%d}\n", u.name, u.age)
}

// ------------------ Interface ------------------

type Printer interface {
	Print()
}

// ------------------ Embedding ------------------

type Admin struct {
	User
	Role string
}
