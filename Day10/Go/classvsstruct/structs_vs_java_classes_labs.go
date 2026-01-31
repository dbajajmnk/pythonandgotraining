// File: structs_vs_java_classes_labs.go
// Module: Structs vs Java Classes — Hands-on Labs (Single Go File)
// Run:
//   go run structs_vs_java_classes_labs.go
//
// What this file demonstrates (mapped from Java → Go):
// Lab 1: Class -> Struct (data model)
// Lab 2: Setters -> Pointer receivers (explicit mutation)
// Lab 3: Inheritance -> Composition (embedding)
// Lab 4: Interfaces + implements -> Implicit interfaces
// Lab 5: Service class -> Dependency inversion with interfaces
// Lab 6: Exceptions -> error return values
// Lab 7: Constructor-heavy -> factory functions / literals
// Lab 8: Testing perspective (small interfaces, fake implementations)
//
// Notes:
// - Go has package-level encapsulation (capitalization)
// - Structs are value types; use pointers when you intend to mutate shared state

package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Structs vs Java Classes — Go Hands-on Labs (Single File) ===")

	// -----------------------------
	// Lab 1: Java-style User class -> Go struct
	fmt.Println("\nLAB 1) Class -> Struct (data model)")
	u1 := NewUser("Deepak", "deepak@example.com")
	fmt.Println("User:", u1.PublicView())

	// -----------------------------
	// Lab 2: Setters -> Pointer receiver mutation (explicit)
	fmt.Println("\nLAB 2) Setter -> Pointer receiver (explicit mutation)")
	u2 := NewUser("Asha", "asha@example.com")
	fmt.Println("Before:", u2.PublicView())
	if err := u2.SetEmail("asha@company.com"); err != nil {
		fmt.Println("SetEmail error:", err)
	}
	fmt.Println("After :", u2.PublicView())

	// -----------------------------
	// Lab 3: Inheritance -> Composition (embedding)
	fmt.Println("\nLAB 3) Inheritance -> Composition (embedding)")
	admin := NewAdmin(u1, []string{"USER_READ", "USER_WRITE", "AUDIT_VIEW"})
	fmt.Println("Admin:", admin.PublicView())
	fmt.Println("Admin Permissions:", admin.Permissions)

	// -----------------------------
	// Lab 4: Java interface implements -> Go implicit interface satisfaction
	fmt.Println("\nLAB 4) Polymorphism via implicit interfaces")
	var notifier Notifier
	notifier = EmailNotifier{} // satisfies Notifier automatically
	if err := notifier.Notify(u1, "Welcome to the platform!"); err != nil {
		fmt.Println("notify error:", err)
	}

	// -----------------------------
	// Lab 5: Service class -> interface-driven service with dependency inversion
	fmt.Println("\nLAB 5) Service + dependency inversion (repo interface)")
	repo := NewInMemoryUserRepo()
	svc := NewUserService(repo, EmailNotifier{})
	if err := svc.Register("Vinay", "vinay@example.com"); err != nil {
		fmt.Println("register error:", err)
	}
	if err := svc.Register("Vinay", "vinay@example.com"); err != nil {
		fmt.Println("register error (expected):", err)
	}

	// -----------------------------
	// Lab 6: Exceptions -> explicit error returns
	fmt.Println("\nLAB 6) Error returns (no exceptions)")
	_, err := svc.GetByEmail("missing@example.com")
	fmt.Println("GetByEmail missing =>", err)

	// -----------------------------
	// Lab 7: Constructor-heavy -> factory + validation
	fmt.Println("\nLAB 7) Factory functions + validation")
	_, err = NewUserValidated("   ", "bademail")
	fmt.Println("NewUserValidated invalid =>", err)

	// -----------------------------
	// Lab 8: Testing perspective (fake repo)
	fmt.Println("\nLAB 8) Testing mindset: small interfaces + fakes")
	fakeRepo := &FakeUserRepo{
		Store: map[string]User{},
	}
	testSvc := NewUserService(fakeRepo, NoopNotifier{})
	_ = testSvc.Register("Test", "test@example.com")
	fmt.Println("Fake repo saved count:", len(fakeRepo.Store))

	fmt.Println("\n=== Done ===")
}

//
// --- LAB 1: Struct as data model ---
//

// User is the Go equivalent of a Java POJO.
// Encapsulation in Go is package-level via capitalization.
// Here we keep fields unexported (lowercase) and expose safe methods.
type User struct {
	name  string // private to package
	email string // private to package
}

// NewUser is a conventional factory (Go has no constructors).
func NewUser(name, email string) User {
	// Keep factory lightweight; deep validation can be separate (see NewUserValidated).
	return User{name: name, email: email}
}

// NewUserValidated shows constructor-heavy Java mapped to explicit factory validation.
func NewUserValidated(name, email string) (User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" {
		return User{}, errors.New("name cannot be empty")
	}
	if !strings.Contains(email, "@") {
		return User{}, errors.New("email must contain '@'")
	}
	return User{name: name, email: email}, nil
}

// PublicView is like a DTO / view model.
func (u User) PublicView() string {
	return fmt.Sprintf("{name:%q, email:%q}", u.name, u.email)
}

//
// --- LAB 2: Explicit mutation via pointer receivers ---
//

// SetEmail mutates state, so it uses a pointer receiver.
func (u *User) SetEmail(email string) error {
	email = strings.TrimSpace(email)
	if !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}
	u.email = email
	return nil
}

//
// --- LAB 3: Composition over inheritance (embedding) ---
//

type Admin struct {
	User                 // embedded: Admin *has a* User (composition)
	Permissions []string // additional state
}

func NewAdmin(user User, perms []string) Admin {
	return Admin{User: user, Permissions: perms}
}

// PublicView demonstrates that embedded fields/methods can be promoted.
func (a Admin) PublicView() string {
	return fmt.Sprintf("{role:%q, user:%s, perms:%d}", "admin", a.User.PublicView(), len(a.Permissions))
}

//
// --- LAB 4: Interfaces and implicit implementation ---
//

// Notifier is behavior-only; types satisfy it implicitly by having Notify(...).
type Notifier interface {
	Notify(to User, message string) error
}

type EmailNotifier struct{}

func (EmailNotifier) Notify(to User, message string) error {
	if strings.TrimSpace(to.email) == "" {
		return errors.New("cannot notify: user email empty")
	}
	fmt.Printf("Email sent to %s: %q\n", to.email, message)
	return nil
}

// NoopNotifier is useful in tests.
type NoopNotifier struct{}

func (NoopNotifier) Notify(to User, message string) error { return nil }

//
// --- LAB 5/6: Service + dependency inversion, errors as values ---
//

type UserRepo interface {
	Save(u User) error
	FindByEmail(email string) (User, bool)
}

type InMemoryUserRepo struct {
	store map[string]User // key: email
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{store: make(map[string]User)}
}

func (r *InMemoryUserRepo) Save(u User) error {
	email := strings.TrimSpace(u.email)
	if email == "" {
		return errors.New("cannot save user: empty email")
	}
	if _, exists := r.store[email]; exists {
		return fmt.Errorf("user already exists: %s", email)
	}
	r.store[email] = u
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
	u, err := NewUserValidated(name, email)
	if err != nil {
		return fmt.Errorf("register validation failed: %w", err)
	}

	if err := s.repo.Save(u); err != nil {
		return fmt.Errorf("register save failed: %w", err)
	}

	// Notification failure is still explicit via error return.
	if err := s.notifier.Notify(u, "Welcome!"); err != nil {
		return fmt.Errorf("register notify failed: %w", err)
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

//
// --- LAB 8: Testing with fakes ---
//

type FakeUserRepo struct {
	Store map[string]User
}

func (f *FakeUserRepo) Save(u User) error {
	if f.Store == nil {
		f.Store = map[string]User{}
	}
	f.Store[u.email] = u
	return nil
}

func (f *FakeUserRepo) FindByEmail(email string) (User, bool) {
	u, ok := f.Store[email]
	return u, ok
}
