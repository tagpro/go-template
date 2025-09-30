# Getting Started with This Template

This guide will help you customize this template for your project.

## Initial Setup

### 1. Create Your Project

```bash
# Install gonew
go install golang.org/x/tools/cmd/gonew@latest

# Create your project
gonew github.com/tagpro/auth-api github.com/yourname/your-project
cd your-project
```

### 2. Update Module Path

The template uses `github.com/tagpro/auth-api`. Update all imports:

```bash
# On macOS/Linux
find . -type f -name "*.go" -exec sed -i '' 's|github.com/tagpro/auth-api|github.com/yourname/your-project|g' {} +

# On Linux (without macOS)
find . -type f -name "*.go" -exec sed -i 's|github.com/tagpro/auth-api|github.com/yourname/your-project|g' {} +
```

Then update `go.mod`:

```bash
go mod edit -module github.com/yourname/your-project
go mod tidy
```

### 3. Install Dependencies

```bash
go mod download
```

### 4. Install Development Tools

```bash
# Air for hot reload
go install github.com/air-verse/air@latest

# Goose for migrations
go install github.com/pressly/goose/v3/cmd/goose@latest

# Templ for templates
go install github.com/a-h/templ/cmd/templ@latest

# SQLC for SQL code generation
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

## Building Your Application

### 1. Define Your Database Schema

Edit `db/migrations/20250401042550_init.sql`:

```sql
-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS users;
```

### 2. Apply Migrations

```bash
make db/start  # Start PostgreSQL
make db/up     # Run migrations
```

### 3. Write SQLC Queries

Create `db/queries/users.sql`:

```sql
-- name: CreateUser :one
INSERT INTO users (name, email)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY created_at DESC;
```

### 4. Generate Type-Safe Code

```bash
make db/gen
go mod tidy
```

### 5. Create Handlers

Create `internal/users/handler.go`:

```go
package users

import (
    "encoding/json"
    "net/http"
    
    "github.com/yourname/your-project/internal/db"
)

type Handler struct {
    queries *db.Queries
}

func NewHandler(queries *db.Queries) *Handler {
    return &Handler{queries: queries}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
    users, err := h.queries.ListUsers(r.Context())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}
```

### 6. Register Routes

Update `internal/app/router.go`:

```go
func NewRouter(queries *db.Queries) *chi.Mux {
    r := chi.NewRouter()
    
    // ... health checks ...
    
    r.Route("/v1", func(r chi.Router) {
        userHandler := users.NewHandler(queries)
        r.Get("/users", userHandler.List)
    })
    
    return r
}
```

### 7. Update Server Setup

Update `internal/app/server.go`:

```go
import (
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/yourname/your-project/internal/db"
)

type Server struct {
    config  *config.Config
    queries *db.Queries
}

func NewServer(c *config.Config) (*Server, error) {
    // Connect to database
    pool, err := pgxpool.New(context.Background(), c.DatabaseURL)
    if err != nil {
        return nil, err
    }
    
    queries := db.New(pool)
    
    return &Server{
        config:  c,
        queries: queries,
    }, nil
}

func (s *Server) StartServer() error {
    r := NewRouter(s.queries)
    // ... rest of the code
}
```

## Development Workflow

1. **Make changes** to your code
2. **Generate templates** (if using templ): `make generate`
3. **Run with hot reload**: `make run`
4. **Test**: `make test`
5. **Build**: `make build`

## Configuration

Update `config/config.yml`:

```yaml
port: 8000
database_url: "postgres://postgres:postgres@localhost:5432/myapp?sslmode=disable"
logger:
  level: debug
  format: json
```

Add new fields to `internal/config/config.go`:

```go
type Config struct {
    Port        int
    DatabaseURL string `mapstructure:"database_url"`
    Logger      *logger.LogConfig
}
```

## Next Steps

- Add authentication/authorization
- Add input validation
- Add tests
- Add OpenAPI/Swagger documentation
- Set up CI/CD
- Configure production deployment

## Need Help?

- [Chi Documentation](https://go-chi.io/)
- [SQLC Documentation](https://docs.sqlc.dev/)
- [Goose Documentation](https://pressly.github.io/goose/)
- [Templ Documentation](https://templ.guide/)
