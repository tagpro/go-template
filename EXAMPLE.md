# Example: Adding a Users Feature

This example shows how to add a complete CRUD feature to your API.

## 1. Create Migration

```bash
make db/create name=create_users_table
```

Edit the generated migration file:

```sql
-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
```

## 2. Apply Migration

```bash
make db/up
```

## 3. Create SQLC Queries

Create `db/queries/users.sql`:

```sql
-- name: CreateUser :one
INSERT INTO users (name, email)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET name = $2, email = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
```

## 4. Generate Code

```bash
make db/gen
go mod tidy
```

## 5. Create Handler

Create `internal/users/handler.go`:

```go
package users

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/google/uuid"
    "github.com/yourname/your-project/internal/db"
)

type Handler struct {
    queries *db.Queries
}

func NewHandler(queries *db.Queries) *Handler {
    return &Handler{queries: queries}
}

type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user, err := h.queries.CreateUser(r.Context(), db.CreateUserParams{
        Name:  req.Name,
        Email: req.Email,
    })
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        http.Error(w, "invalid user id", http.StatusBadRequest)
        return
    }

    user, err := h.queries.GetUserByID(r.Context(), id)
    if err != nil {
        http.Error(w, "user not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
    users, err := h.queries.ListUsers(r.Context(), db.ListUsersParams{
        Limit:  10,
        Offset: 0,
    })
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        http.Error(w, "invalid user id", http.StatusBadRequest)
        return
    }

    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user, err := h.queries.UpdateUser(r.Context(), db.UpdateUserParams{
        ID:    id,
        Name:  req.Name,
        Email: req.Email,
    })
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        http.Error(w, "invalid user id", http.StatusBadRequest)
        return
    }

    if err := h.queries.DeleteUser(r.Context(), id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
```

## 6. Register Routes

Update `internal/app/router.go`:

```go
import (
    "github.com/yourname/your-project/internal/users"
    "github.com/yourname/your-project/internal/db"
)

func NewRouter(queries *db.Queries) *chi.Mux {
    r := chi.NewRouter()

    // ... middleware and health checks ...

    r.Route("/v1", func(r chi.Router) {
        userHandler := users.NewHandler(queries)
        
        r.Route("/users", func(r chi.Router) {
            r.Get("/", userHandler.List)
            r.Post("/", userHandler.Create)
            
            r.Route("/{id}", func(r chi.Router) {
                r.Get("/", userHandler.Get)
                r.Put("/", userHandler.Update)
                r.Delete("/", userHandler.Delete)
            })
        })
    })

    return r
}
```

## 7. Update Server

Update `internal/app/server.go`:

```go
import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/yourname/your-project/internal/db"
)

type Server struct {
    config  *config.Config
    queries *db.Queries
    pool    *pgxpool.Pool
}

func NewServer(c *config.Config) (*Server, error) {
    pool, err := pgxpool.New(context.Background(), c.DatabaseURL)
    if err != nil {
        return nil, fmt.Errorf("unable to connect to database: %w", err)
    }

    queries := db.New(pool)

    return &Server{
        config:  c,
        queries: queries,
        pool:    pool,
    }, nil
}

func (s *Server) StartServer() error {
    defer s.pool.Close()
    
    r := NewRouter(s.queries)
    log.Printf("Server listening on port %d", s.config.Port)
    return http.ListenAndServe(":"+strconv.Itoa(s.config.Port), r)
}
```

## 8. Update Config

Add database URL to `internal/config/config.go`:

```go
type Config struct {
    Port        int
    DatabaseURL string `mapstructure:"database_url"`
    Logger      *logger.LogConfig
}
```

Update `config/config.yml`:

```yaml
port: 8000
database_url: "postgres://postgres:postgres@localhost:5432/app?sslmode=disable"
logger:
  level: debug
  format: json
```

## 9. Test Your API

```bash
# Start the server
make run

# Create a user
curl -X POST http://localhost:8000/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'

# List users
curl http://localhost:8000/v1/users

# Get a user
curl http://localhost:8000/v1/users/{id}

# Update a user
curl -X PUT http://localhost:8000/v1/users/{id} \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane Doe","email":"jane@example.com"}'

# Delete a user
curl -X DELETE http://localhost:8000/v1/users/{id}
```

## Dependencies Needed

Add these to your project:

```bash
go get github.com/jackc/pgx/v5/pgxpool
go get github.com/google/uuid
go mod tidy
```
