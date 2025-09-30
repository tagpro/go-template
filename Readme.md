# Go API Template

A production-ready Go API template with database migrations, SQLC code generation, and modern tooling.

## Features

- **HTTP Server**: [Chi Router](https://github.com/go-chi/chi) - Lightweight, idiomatic HTTP router
- **Database Migrations**: [Goose](https://github.com/pressly/goose) - Database migration tool
- **Type-Safe SQL**: [SQLC](https://github.com/sqlc-dev/sqlc) - Generate type-safe Go code from SQL
- **Templating**: [Templ](https://github.com/a-h/templ) - Type-safe HTML templating
- **Configuration**: [Viper](https://github.com/spf13/viper) - Configuration management
- **Hot Reload**: [Air](https://github.com/air-verse/air) - Live reload for development
- **Docker**: Ready-to-use Docker setup with PostgreSQL

## Quick Start

### Prerequisites

- Go 1.24.1 or higher
- Docker & Docker Compose
- Make

### Installation

1. **Create your new project using gonew:**

   ```bash
   # Install gonew if you haven't already
   go install golang.org/x/tools/cmd/gonew@latest
   
   # Create a new project from this template
   gonew github.com/tagpro/auth-api example.com/your-project-name
   cd your-project-name
   ```

2. **Install dependencies:**

   ```bash
   go mod download
   ```

3. **Install development tools:**

   ```bash
   # Install Air for hot reload
   go install github.com/air-verse/air@latest
   
   # Install Goose for migrations
   go install github.com/pressly/goose/v3/cmd/goose@latest
   
   # Install Templ for template generation
   go install github.com/a-h/templ/cmd/templ@latest
   ```

## Project Structure

```
.
├── cmd/
│   └── app/
│       └── main.go           # Application entry point
├── internal/
│   ├── app/
│   │   ├── server.go         # HTTP server setup
│   │   └── router.go         # Route definitions
│   └── config/
│       └── config.go         # Configuration loader
├── pkg/
│   └── logger/
│       └── logger.go         # Logging utilities
├── db/
│   ├── migrations/           # SQL migration files (Goose)
│   │   └── 20250401042550_init.sql
│   └── queries/              # SQL query files (SQLC)
│       └── example.sql
├── templates/                # Templ templates
│   └── example/
│       └── hello.templ
├── config/
│   ├── config.yml            # Development config
│   └── config.prod.yml       # Production config
├── Makefile                  # Build commands
├── sqlc.yml                  # SQLC configuration
├── .air.toml                 # Air (hot reload) configuration
├── compose.yml               # Docker Compose setup
└── Dockerfile                # Container image
```

## Development

### Database Setup

1. **Start the database:**

   ```bash
   make db/start
   ```

2. **Create your schema in `db/migrations/20250401042550_init.sql`:**

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

3. **Run migrations:**

   ```bash
   make db/up
   ```

4. **Create SQLC queries in `db/queries/`:**

   ```sql
   -- name: CreateUser :one
   INSERT INTO users (name, email)
   VALUES ($1, $2)
   RETURNING *;

   -- name: GetUserByID :one
   SELECT * FROM users WHERE id = $1;
   ```

5. **Generate type-safe Go code:**

   ```bash
   make db/gen
   go mod tidy
   ```

### Running the Application

**Development mode (with hot reload):**

```bash
make run
```

**Build and run:**

```bash
make build
./app
```

The server will start on `http://localhost:8000` (configurable in `config/config.yml`).

### Creating New Migrations

```bash
make db/create name=add_users_table
```

This creates a new migration file in `db/migrations/`.

## Available Make Commands

```bash
make help              # Show all available commands
make generate          # Generate templ templates
make build             # Build the application
make run               # Run with hot reload (Air)
make test              # Run tests
make db/start          # Start PostgreSQL with Docker
make db/client         # Connect to database CLI
make db/up             # Run migrations
make db/down           # Rollback last migration
make db/create         # Create new migration
make db/gen            # Generate SQLC code
make docker/build      # Build Docker image
make docker/run        # Run Docker container
```

## Configuration

Configuration files are located in the `config/` directory:

- `config.yml` - Development configuration
- `config.prod.yml` - Production configuration

Example `config.yml`:

```yaml
port: 8000
logger:
  level: debug
  format: json
```

## Adding Your Business Logic

1. **Create a new handler in `internal/`:**

   ```go
   // internal/users/handler.go
   package users
   
   import (
       "net/http"
       "github.com/your-project/internal/db"
   )
   
   type Handler struct {
       queries *db.Queries
   }
   
   func NewHandler(queries *db.Queries) *Handler {
       return &Handler{queries: queries}
   }
   
   func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
       // Your logic here
   }
   ```

2. **Register routes in `internal/app/router.go`:**

   ```go
   r.Route("/v1", func(r chi.Router) {
       userHandler := users.NewHandler(queries)
       r.Get("/users", userHandler.List)
   })
   ```

## Docker

**Build image:**

```bash
docker build -t myapp:latest .
```

**Run container:**

```bash
docker run -p 8000:8000 -v ./config:/config myapp:latest
```

**Using Docker Compose:**

```bash
# Production
docker compose up

# Development (with hot reload)
docker compose -f compose.dev.yml up
```

For detailed Docker documentation, see [DOCKER.md](DOCKER.md).

## Testing

```bash
make test
```

## Production Deployment

1. Update `config/config.prod.yml` with production settings
2. Build the application: `make build`
3. Or use Docker: `make docker/build && make docker/run`

## Libraries Used

- [Chi](https://github.com/go-chi/chi) - HTTP router
- [SQLC](https://github.com/sqlc-dev/sqlc) - Type-safe SQL
- [Goose](https://github.com/pressly/goose) - Database migrations
- [Templ](https://github.com/a-h/templ) - HTML templating
- [Viper](https://github.com/spf13/viper) - Configuration
- [Air](https://github.com/air-verse/air) - Hot reload

## License

MIT

