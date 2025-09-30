# Template Conversion Summary

This repository has been converted from a specific authentication API into a **Go API template** that can be used to bootstrap new projects.

## What Was Removed

### Business Logic
- ✂️ `internal/auth/` - Authentication handlers and logic
- ✂️ `internal/hello/` - Example hello world handlers
- ✂️ `templates/helloworld/` - Business-specific templ templates

### Database Business Logic
- ✂️ User authentication tables from initial migration
- ✂️ Sessions and auth_tokens tables
- ✂️ All SQLC queries for users, sessions, and auth tokens

### Dependencies
- ✂️ `github.com/golang-jwt/jwt/v5` - JWT library (not needed in template)

## What Was Kept

### Infrastructure ✅
- ✅ Database migration setup (Goose)
- ✅ SQLC configuration and setup
- ✅ Docker & Docker Compose configuration
- ✅ Configuration management (Viper)
- ✅ Logger setup
- ✅ HTTP server setup (Chi router)
- ✅ Makefile with all commands
- ✅ Air configuration for hot reload
- ✅ Templ template setup
- ✅ All development tools

### Project Structure ✅
```
.
├── cmd/app/              # Application entry point
├── internal/
│   ├── app/              # Server and routing setup
│   ├── config/           # Configuration loader
│   └── db/               # SQLC generated code (gitignored)
├── pkg/logger/           # Logging utilities
├── db/
│   ├── migrations/       # Database migrations
│   └── queries/          # SQLC query files
├── templates/            # Templ templates
└── config/               # Configuration files
```

## What Was Added

### Documentation 📚
- ✅ **README.md** - Comprehensive project documentation
- ✅ **GETTING_STARTED.md** - Step-by-step guide for using the template
- ✅ **EXAMPLE.md** - Complete example of adding a Users feature
- ✅ **db/README.md** - Database workflow documentation
- ✅ **templates/README.md** - Templ usage guide

### Examples 📝
- ✅ `db/migrations/20250401042550_init.sql` - Empty migration template with examples
- ✅ `db/queries/example.sql` - Commented SQLC query examples
- ✅ `templates/example/hello.templ` - Simple templ example

### Infrastructure Improvements 🔧
- ✅ Updated `.gitignore` to ignore generated files
- ✅ Added `internal/db/.gitignore` to keep directory but ignore generated code
- ✅ Updated router with middleware and better structure
- ✅ Improved health check endpoints with JSON responses
- ✅ Updated SQLC config to use directory for queries

## How to Use This Template

### Option 1: Using gonew (Recommended)

```bash
go install golang.org/x/tools/cmd/gonew@latest
gonew github.com/tagpro/auth-api github.com/yourname/your-project
```

### Option 2: Clone and Modify

```bash
git clone <this-repo> your-project
cd your-project
# Update module name in go.mod
# Find and replace github.com/tagpro/auth-api with your module name
```

## Next Steps

1. Read **GETTING_STARTED.md** for setup instructions
2. Follow **EXAMPLE.md** to add your first feature
3. Customize the migration in `db/migrations/20250401042550_init.sql`
4. Add your SQLC queries in `db/queries/`
5. Run `make db/gen` to generate type-safe database code
6. Create your handlers in `internal/`
7. Register routes in `internal/app/router.go`

## Key Features

✨ **Type-Safe SQL** - SQLC generates Go code from SQL
✨ **Database Migrations** - Goose for versioned schema changes
✨ **Hot Reload** - Air for development
✨ **HTML Templates** - Templ for type-safe HTML
✨ **Configuration** - Viper for flexible config management
✨ **Docker Ready** - Docker & Compose files included
✨ **Production Ready** - Structured logging, health checks, middleware

## Available Commands

```bash
make help              # Show all commands
make build             # Build application
make run               # Run with hot reload
make test              # Run tests
make db/start          # Start PostgreSQL
make db/up             # Run migrations
make db/create         # Create new migration
make db/gen            # Generate SQLC code
make generate          # Generate templ templates
```

## Questions?

- Check **GETTING_STARTED.md** for detailed setup
- See **EXAMPLE.md** for a complete feature example
- Read the **README.md** for project overview
