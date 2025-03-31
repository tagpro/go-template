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

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: clean
clean:
	rm -f app

.PHONY: all
all: build test lint fmt vet clean

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  generate   Generate templates"
	@echo "  build      Build the application"
	@echo "  run        Run the application"
	@echo "  test       Run tests"
	@echo "  lint       Run linter"
	@echo "  fmt        Format code"
	@echo "  vet        Run vet"
	@echo "  clean      Clean build artifacts"
	@echo "  all        Run all targets"
	@echo "  help       Show this help"
