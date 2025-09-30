# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.25-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /build

# Copy go mod files and download dependencies
# This layer is cached unless go.mod or go.sum changes
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=0: Build a statically-linked binary
# -ldflags="-s -w": Strip debug information to reduce binary size
# -trimpath: Remove file system paths from the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w -X main.version=${VERSION:-dev}" \
    -trimpath \
    -o app \
    ./cmd/app

# Final stage - using distroless for minimal attack surface
FROM gcr.io/distroless/static-debian12:nonroot

# Copy timezone data and CA certificates from builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary
COPY --from=builder /build/app /app

# Copy config files (optional - can be mounted as volume)
COPY --from=builder /build/config /config

# Use non-root user (distroless nonroot user has UID 65532)
USER nonroot:nonroot

# Expose port (default 8000, change as needed)
EXPOSE 8000

# Health check (adjust path and interval as needed)
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/app", "healthcheck"]

# Set the binary as entrypoint
ENTRYPOINT ["/app"]
