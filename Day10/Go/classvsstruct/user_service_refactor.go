// File: user_service_refactor.go
// Topic: Structs vs Java Classes — Side-by-side Refactor Example (Go Version)
//
// Scenario: Register a user, prevent duplicates, and send a welcome notification.
//
// Key Go traits:
// - Struct = data only
// - Behavior attached via functions/methods (receiver)
// - Interfaces are implicit
// - Errors are returned values (explicit control flow)

package main

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	name  string
	email string
}

func NewUser(name, email string) (User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" {
		return User{}, errors.New("name empty")
	}
	if !strings.Contains(email, "@") {
		return User{}, errors.New("email invalid")
	}
	return User{name: name, email: email}, nil
}

type Notifier interface {
	Notify(to User, message string) error
}

type EmailNotifier struct{}

func (EmailNotifier) Notify(to User, message string) error {
	fmt.Printf("Email sent to %s: %q\n", to.email, message)
	return nil
}

type UserRepo interface {
	Save(u User) error
	FindByEmail(email string) (User, bool)
}

type InMemoryUserRepo struct {
	store map[string]User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{store: make(map[string]User)}
}

func (r *InMemoryUserRepo) Save(u User) error {
	if _, exists := r.store[u.email]; exists {
		return fmt.Errorf("duplicate user: %s", u.email)
	}
	r.store[u.email] = u
	return nil
}

func (r *InMemoryUserRepo) FindByEmail(email string) (User, bool) {
	u, ok := r.store[strings.TrimSpace(email)]
	return u, ok
}

type UserService struct {
	repo     UserRepo
	notifier Notifier
}

func NewUserService(repo UserRepo, notifier Notifier) *UserService {
	return &UserService{repo: repo, notifier: notifier}
}

func (s *UserService) Register(name, email string) error {
	u, err := NewUser(name, email)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	if err := s.repo.Save(u); err != nil {
		return fmt.Errorf("save failed: %w", err)
	}
	if err := s.notifier.Notify(u, "Welcome!"); err != nil {
		return fmt.Errorf("notify failed: %w", err)
	}
	return nil
}

func (s *UserService) GetByEmail(email string) (User, error) {
	u, ok := s.repo.FindByEmail(email)
	if !ok {
		return User{}, fmt.Errorf("user not found: %s", email)
	}
	return u, nil
}

func main() {
	repo := NewInMemoryUserRepo()
	notifier := EmailNotifier{}
	svc := NewUserService(repo, notifier)

	_ = svc.Register("Deepak", "deepak@example.com")

	if err := svc.Register("Deepak", "deepak@example.com"); err != nil {
		fmt.Println("Expected error:", err)
	}

	if _, err := svc.GetByEmail("missing@example.com"); err != nil {
		fmt.Println("Expected error:", err)
	}
}
