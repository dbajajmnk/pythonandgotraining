# Gin CRUD (No Database) — Slice + Mutex

This is a simple **in-memory CRUD** API using:
- Gin
- A slice (`[]User`) as the storage
- `sync.RWMutex` for concurrency safety

> Data resets when the server restarts.

## Run
```bash
go mod tidy
go run .
```

Server: http://localhost:8080

## Endpoints
- GET    /health
- POST   /users
- GET    /users
- GET    /users/{id}
- PUT    /users/{id}
- DELETE /users/{id}

## curl examples (Windows-friendly)

### List
```bash
curl http://localhost:8080/users
```

### Create
```bash
curl -X POST http://localhost:8080/users ^
  -H "Content-Type: application/json" ^
  -d "{"name":"New User","email":"new@example.com"}"
```

### Get by ID
```bash
curl http://localhost:8080/users/1
```

### Update (PUT)
```bash
curl -X PUT http://localhost:8080/users/1 ^
  -H "Content-Type: application/json" ^
  -d "{"name":"Updated Name"}"
```

### Delete
```bash
curl -X DELETE http://localhost:8080/users/1
```
