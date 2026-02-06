# Day 14 - net/http + SQLite (Pure Go) CRUD

This is a **working CRUD API** using:
- Standard library `net/http` (no Gin)
- `database/sql`
- SQLite driver: `modernc.org/sqlite` (**pure Go**, no CGO, no gcc)

## Run
```bash
go mod tidy
go run .
```

Server: http://localhost:8080
- Health: GET /health
- DB file created: day14.db

## Endpoints
- POST   /users
- GET    /users
- GET    /users/{id}
- PUT    /users/{id}
- DELETE /users/{id}

## curl examples (Windows CMD / PowerShell friendly)

### Create
```bash
curl -X POST http://localhost:8080/users ^
  -H "Content-Type: application/json" ^
  -d "{"name":"Deepak","email":"deepak@example.com"}"
```

### List
```bash
curl http://localhost:8080/users
```

### Get by ID
```bash
curl http://localhost:8080/users/1
```

### Update
```bash
curl -X PUT http://localhost:8080/users/1 ^
  -H "Content-Type: application/json" ^
  -d "{"name":"Deepak Bajaj"}"
```

### Delete
```bash
curl -X DELETE http://localhost:8080/users/1
```
