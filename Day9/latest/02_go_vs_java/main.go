package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
WHAT & WHY (Go vs Java)
Go focuses on simplicity:
- No classes/inheritance; uses structs + interfaces (composition)
- Errors are explicit return values (not exceptions)
- Concurrency is built-in (goroutines)

This file covers:
1) "Struct + methods" instead of classes
2) "Interface" (duck-typing style)
3) Error-return pattern (vs try/catch)
*/

func main() {
	fmt.Println("=== GO SYNTAX vs JAVA (Engineering Mapping) ===")

	// Java: new UserService().normalizeName("  deepak  ")
	// Go: functions + structs
	svc := UserService{Prefix: "Mr."}
	name, err := svc.NormalizeName("  deepak  ")
	must(err)
	fmt.Println("Normalized:", name)

	// Interface demo (like Java interface)
	var store Repository = InMemoryRepo{}
	must(store.Save(User{ID: 1, Name: name}))
	u, err := store.FindByID(1)
	must(err)
	fmt.Println("Fetched:", u)

	// Error handling demo
	_, err = store.FindByID(999)
	fmt.Println("Expected error:", err)

	fmt.Println("\nCommon Go interview line: 'Go is OOP without classes.'")
}

type User struct {
	ID   int
	Name string
}

// "Service" as struct (like Java class fields)
type UserService struct {
	Prefix string
}

// Method receiver (like instance method in Java)
func (s UserService) NormalizeName(raw string) (string, error) {
	n := strings.TrimSpace(raw)
	if n == "" {
		return "", errors.New("name cannot be empty")
	}
	// Java: prefix + capitalized
	n = strings.ToUpper(n[:1]) + n[1:]
	return s.Prefix + " " + n, nil
}

// Java interface -> Go interface
type Repository interface {
	Save(u User) error
	FindByID(id int) (User, error)
}

type InMemoryRepo struct {
	db map[int]User
}

func (r InMemoryRepo) Save(u User) error {
	if r.db == nil {
		r.db = make(map[int]User) // IMPORTANT: make before use
	}
	if u.ID <= 0 {
		return errors.New("invalid id")
	}
	r.db[u.ID] = u
	return nil
}

func (r InMemoryRepo) FindByID(id int) (User, error) {
	if r.db == nil {
		return User{}, errors.New("repository not initialized")
	}
	u, ok := r.db[id]
	if !ok {
		return User{}, fmt.Errorf("user not found: %d", id)
	}
	return u, nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
