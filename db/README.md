# Database

This directory contains database migrations and queries.

## Structure

- `migrations/` - Goose migration files
- `queries/` - SQLC query files

## Workflow

### 1. Create a Migration

```bash
make db/create name=create_users_table
```

This creates a new migration file in `migrations/` with timestamps.

### 2. Write Your Schema

Edit the generated migration file:

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

### 3. Apply Migrations

```bash
make db/up
```

### 4. Write SQLC Queries

Create a file in `queries/` (e.g., `users.sql`):

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

### 5. Generate Type-Safe Go Code

```bash
make db/gen
go mod tidy
```

This generates Go code in `internal/db/`.

### 6. Use in Your Code

```go
import "github.com/tagpro/auth-api/internal/db"

// In your handler
queries := db.New(dbConn)
user, err := queries.GetUserByID(ctx, userID)
```

## Migration Commands

- `make db/start` - Start PostgreSQL database
- `make db/up` - Run migrations
- `make db/down` - Rollback last migration
- `make db/create name=<name>` - Create new migration
- `make db/client` - Connect to database CLI

## SQLC Query Annotations

- `:one` - Returns a single row
- `:many` - Returns multiple rows
- `:exec` - Executes without returning data
- `:execrows` - Returns number of affected rows

See [SQLC documentation](https://docs.sqlc.dev/) for more details.
