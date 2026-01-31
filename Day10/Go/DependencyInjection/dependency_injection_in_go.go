// File: dependency_injection_in_go.go
// Module: Dependency Injection in Go — Single Executable File
//
// Run:
//   go run dependency_injection_in_go.go
//
// Demonstrates:
// - No DI vs DI
// - Constructor, field, and function injection
// - Interface-based decoupling
// - Testing with fake dependencies

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Dependency Injection in Go ===")

	// Without DI (anti-pattern)
	fmt.Println("\nWithout DI:")
	bad := NewBadService()
	bad.DoWork()

	// Constructor Injection
	fmt.Println("\nConstructor Injection:")
	logger := ConsoleLogger{}
	repo := NewMemoryRepo()
	service := NewService(repo, logger)
	service.DoWork()

	// Function Injection
	fmt.Println("\nFunction Injection:")
	RunTask(service, logger)

	// Testing with Fake
	fmt.Println("\nTesting with Fake Dependency:")
	fakeRepo := FakeRepo{}
	testService := NewService(fakeRepo, logger)
	testService.DoWork()

	fmt.Println("=== Done ===")
}

// ---------------- Without DI ----------------

type BadService struct {
	repo *MemoryRepo
}

func NewBadService() *BadService {
	return &BadService{repo: NewMemoryRepo()}
}

func (s *BadService) DoWork() {
	fmt.Println("BadService using repo:", s.repo.Name())
}

// ---------------- Interfaces ----------------

type Repo interface {
	Name() string
}

type Logger interface {
	Log(msg string)
}

// ---------------- Implementations ----------------

type MemoryRepo struct{}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{}
}

func (MemoryRepo) Name() string {
	return "MemoryRepo"
}

type FakeRepo struct{}

func (FakeRepo) Name() string {
	return "FakeRepo"
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(msg string) {
	fmt.Println("LOG:", msg)
}

// ---------------- Constructor Injection ----------------

type Service struct {
	repo   Repo
	logger Logger
}

func NewService(repo Repo, logger Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) DoWork() {
	s.logger.Log("Service using repo: " + s.repo.Name())
}

// ---------------- Function Injection ----------------

func RunTask(s *Service, logger Logger) {
	logger.Log("Running task with injected logger")
	s.DoWork()
}
