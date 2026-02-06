package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

// --- Models ---

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Name  *string `json:"name"`  // optional
	Email *string `json:"email"` // optional
}

// --- Server ---

type Server struct {
	db *sql.DB
}

func main() {
	db := mustOpenDB()
	defer db.Close()

	if err := ensureSchema(db); err != nil {
		log.Fatal("schema error: ", err)
	}

	s := &Server{db: db}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.health)
	mux.HandleFunc("/users", s.usersCollection) // POST, GET
	mux.HandleFunc("/users/", s.usersItem)      // GET, PUT, DELETE

	addr := ":8080"
	log.Println("HTTP API running on http://localhost" + addr)
	if err := http.ListenAndServe(addr, withBasicMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}

// --- Middleware ---

func withBasicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// CORS for easy testing (optional)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)

		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

// --- Handlers ---

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "method not allowed"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"status": "ok"})
}

func (s *Server) usersCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.createUser(w, r)
	case http.MethodGet:
		s.listUsers(w, r)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "method not allowed"})
	}
}

func (s *Server) usersItem(w http.ResponseWriter, r *http.Request) {
	// Expected: /users/{id}
	id, err := parseIDFromPath(r.URL.Path)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid id"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		s.getUser(w, r, id)
	case http.MethodPut:
		s.updateUser(w, r, id)
	case http.MethodDelete:
		s.deleteUser(w, r, id)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "method not allowed"})
	}
}

// --- CRUD Implementation ---

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := decodeJSON(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid JSON body"})
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)

	if req.Name == "" || req.Email == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "name and email are required"})
		return
	}

	now := time.Now().UTC().Format(time.RFC3339)
	res, err := s.db.Exec(`INSERT INTO users(name, email, created_at) VALUES (?, ?, ?)`, req.Name, req.Email, now)
	if err != nil {
		if isUniqueViolation(err) {
			writeJSON(w, http.StatusConflict, map[string]any{"error": "email already exists"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]any{"error": "failed to create user"})
		return
	}

	id, _ := res.LastInsertId()
	user, err := findUserByID(s.db, id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]any{"error": "created but failed to read user"})
		return
	}

	writeJSON(w, http.StatusCreated, user)
}

func (s *Server) listUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.Query(`SELECT id, name, email, created_at FROM users ORDER BY id DESC`)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]any{"error": "failed to fetch users"})
		return
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var u User
		var createdAtStr string
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &createdAtStr); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]any{"error": "failed to read users"})
			return
		}
		u.CreatedAt = parseRFC3339OrZero(createdAtStr)
		users = append(users, u)
	}

	writeJSON(w, http.StatusOK, users)
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request, id int64) {
	user, err := findUserByID(s.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeJSON(w, http.StatusNotFound, map[string]any{"error": "user not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]any{"error": "failed to fetch user"})
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request, id int64) {
	var req UpdateUserRequest
	if err := decodeJSON(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid JSON body"})
		return
	}

	current, err := findUserByID(s.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeJSON(w, http.StatusNotFound, map[string]any{"error": "user not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]any{"error": "failed to fetch user"})
		return
	}

	newName := current.Name
	newEmail := current.Email

	if req.Name != nil {
		val := strings.TrimSpace(*req.Name)
		if val == "" {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": "name cannot be empty"})
			return
		}
		newName = val
	}
	if req.Email != nil {
		val := strings.TrimSpace(*req.Email)
		if val == "" {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": "email cannot be empty"})
			return
		}
		newEmail = val
	}

	_, err = s.db.Exec(`UPDATE users SET name = ?, email = ? WHERE id = ?`, newName, newEmail, id)
	if err != nil {
		if isUniqueViolation(err) {
			writeJSON(w, http.StatusConflict, map[string]any{"error": "email already exists"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]any{"error": "failed to update user"})
		return
	}

	updated, err := findUserByID(s.db, id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]any{"error": "updated but failed to read user"})
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request, id int64) {
	res, err := s.db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]any{"error": "failed to delete user"})
		return
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		writeJSON(w, http.StatusNotFound, map[string]any{"error": "user not found"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// --- DB helpers ---

func mustOpenDB() *sql.DB {
	db, err := sql.Open("sqlite", "file:day14.db")
	if err != nil {
		log.Fatal("db open error: ", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("db ping error: ", err)
	}
	return db
}

func ensureSchema(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		created_at TEXT NOT NULL
	);

	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	`
	_, err := db.Exec(schema)
	return err
}

func findUserByID(db *sql.DB, id int64) (User, error) {
	var u User
	var createdAtStr string

	row := db.QueryRow(`SELECT id, name, email, created_at FROM users WHERE id = ?`, id)
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &createdAtStr); err != nil {
		return User{}, err
	}
	u.CreatedAt = parseRFC3339OrZero(createdAtStr)
	return u, nil
}

// --- Utilities ---

func parseIDFromPath(path string) (int64, error) {
	// path: /users/{id}
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) != 2 || parts[0] != "users" {
		return 0, errors.New("bad path")
	}
	raw := strings.TrimSpace(parts[1])
	id, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid id")
	}
	return id, nil
}

func parseRFC3339OrZero(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}
	}
	return t
}

func isUniqueViolation(err error) bool {
	// Works for many SQLite error strings across drivers
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "unique") || strings.Contains(msg, "constraint")
}

func decodeJSON(r *http.Request, dst any) error {
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(dst)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
