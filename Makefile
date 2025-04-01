.PHONY: generate
generate:
	go tool templ generate

.PHONY: build
build:
	go build -o app cmd/app/main.go

.PHONY: run
run:
	go tool air

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	rm -f tmp/*

.PHONY: all
all: build test lint fmt vet clean

# Database setup

GOOSE_DRIVER=postgres
GOOSE_DBSTRING="postgres://postgres:postgres@localhost:5432/app?sslmode=disable"

.PHONY: db/start
db/start:
	docker-compose up -d

.PHONY: db/client
db/client: db/start
	docker compose exec db psql -U postgres -d app

.PHONY: db/up
db/up: db/start
	goose -dir db/migrations postgres $(GOOSE_DBSTRING) up

.PHONY: db/create
db/create:
	goose -dir db/migrations create $(name) sql

.PHONY: db/down
db/down:
	goose -dir db/migrations postgres $(GOOSE_DBSTRING) down

.PHONY: db/gen
db/gen:
	sqlc generate

.PHONY: docker/build
docker/build:
	docker build -t my-app .

.PHONY: docker/run
docker/run:
	docker run -p 8000:8000 -v ./config:/config my-app

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  generate   Generate templates"
	@echo "  build      Build the application"
	@echo "  run        Run the application"
	@echo "  test       Run tests"
	@echo "  lint       Run linter"
	@echo "  clean      Clean build artifacts"
	@echo "  all        Run all targets"
	@echo "  db/start   Start database"
	@echo "  db/client  Connect to database"
	@echo "  db/up      Migrate database"
	@echo "  db/create  Create new migration. make db/create name=<migration name>"
	@echo "  db/down    Rollback migration"
	@echo "  db/gen     Generate SQL queries"
	@echo "  help       Show this help"
