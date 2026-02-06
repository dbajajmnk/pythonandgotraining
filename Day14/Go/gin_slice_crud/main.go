package main

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// In-memory model (slice) + mutex for concurrency safety.
// NOTE: This data resets when the server restarts.

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Name  *string `json:"name"`  // optional (PATCH-like behavior via PUT)
	Email *string `json:"email"` // optional
}

type Store struct {
	mu     sync.RWMutex
	nextID int64
	users  []User
}

func NewStore() *Store {
	return &Store{
		nextID: 1,
		users:  make([]User, 0),
	}
}

func main() {
	store := NewStore()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// CRUD routes
	r.POST("/users", func(c *gin.Context) { createUser(c, store) })
	r.GET("/users", func(c *gin.Context) { listUsers(c, store) })
	r.GET("/users/:id", func(c *gin.Context) { getUser(c, store) })
	r.PUT("/users/:id", func(c *gin.Context) { updateUser(c, store) })
	r.DELETE("/users/:id", func(c *gin.Context) { deleteUser(c, store) })

	// Seed 2 users for quick demo
	store.seedDemoUsers()

	_ = r.Run(":8080")
}

// --- Handlers ---

func createUser(c *gin.Context, store *Store) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)

	if req.Name == "" || req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name and email are required"})
		return
	}

	// "unique" email check
	if store.emailExists(req.Email, 0) {
		c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		return
	}

	now := time.Now().UTC()

	store.mu.Lock()
	id := store.nextID
	store.nextID++

	u := User{
		ID:        id,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}
	store.users = append(store.users, u)
	store.mu.Unlock()

	c.JSON(http.StatusCreated, u)
}

func listUsers(c *gin.Context, store *Store) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	// Return a copy to avoid surprises if modified later
	out := make([]User, len(store.users))
	copy(out, store.users)

	c.JSON(http.StatusOK, out)
}

func getUser(c *gin.Context, store *Store) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}

	store.mu.RLock()
	defer store.mu.RUnlock()

	idx := indexByID(store.users, id)
	if idx < 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, store.users[idx])
}

func updateUser(c *gin.Context, store *Store) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}

	store.mu.Lock()
	defer store.mu.Unlock()

	idx := indexByID(store.users, id)
	if idx < 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	current := store.users[idx]

	// Apply patch-like updates
	if req.Name != nil {
		name := strings.TrimSpace(*req.Name)
		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name cannot be empty"})
			return
		}
		current.Name = name
	}

	if req.Email != nil {
		email := strings.TrimSpace(*req.Email)
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email cannot be empty"})
			return
		}
		// Unique email check excluding current user
		if store.emailExists(email, id) {
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
			return
		}
		current.Email = email
	}

	current.UpdatedAt = time.Now().UTC()
	store.users[idx] = current

	c.JSON(http.StatusOK, current)
}

func deleteUser(c *gin.Context, store *Store) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}

	store.mu.Lock()
	defer store.mu.Unlock()

	idx := indexByID(store.users, id)
	if idx < 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// Delete from slice: preserve order
	store.users = append(store.users[:idx], store.users[idx+1:]...)
	c.Status(http.StatusNoContent)
}

// --- Helpers ---

func parseIDParam(c *gin.Context) (int64, bool) {
	raw := strings.TrimSpace(c.Param("id"))
	id, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return 0, false
	}
	return id, true
}

func indexByID(users []User, id int64) int {
	for i := range users {
		if users[i].ID == id {
			return i
		}
	}
	return -1
}

func (s *Store) emailExists(email string, excludeID int64) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, u := range s.users {
		if strings.EqualFold(u.Email, email) && u.ID != excludeID {
			return true
		}
	}
	return false
}

func (s *Store) seedDemoUsers() {
	now := time.Now().UTC()
	s.mu.Lock()
	defer s.mu.Unlock()

	s.users = append(s.users,
		User{
			ID:        s.nextID,
			Name:      "Deepak",
			Email:     "deepak@example.com",
			CreatedAt: now,
			UpdatedAt: now,
		},
	)
	s.nextID++

	s.users = append(s.users,
		User{
			ID:        s.nextID,
			Name:      "Test1",
			Email:     "test1@gmail.com",
			CreatedAt: now,
			UpdatedAt: now,
		},
	)
	s.nextID++
}
