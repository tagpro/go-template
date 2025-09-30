# Docker Guide

This project includes multiple Dockerfile configurations for different use cases.

## 📋 Available Dockerfiles

### 1. `Dockerfile` - Production Build ⭐
**Recommended for production deployments**

**Features:**
- ✅ Multi-stage build for minimal image size
- ✅ Distroless base image (minimal attack surface)
- ✅ Non-root user for security
- ✅ Statically linked binary (no dependencies)
- ✅ Optimized with `-ldflags` for smaller binary
- ✅ Build cache optimization
- ✅ Health check support

**Build:**
```bash
docker build -t myapp:latest .
```

**Run:**
```bash
docker run -p 8000:8000 -v ./config:/config myapp:latest
```

### 2. `Dockerfile.dev` - Development Build
**For local development with hot reload**

**Features:**
- ✅ Air for hot reload
- ✅ Volume mounting for live code changes
- ✅ Development tools included
- ✅ Alpine base for familiarity

**Build:**
```bash
docker build -f Dockerfile.dev -t myapp:dev .
```

**Run:**
```bash
docker run -v $(pwd):/app -p 8000:8000 myapp:dev
```

### 3. `Dockerfile.multiarch` - Multi-Platform Build
**For building AMD64 and ARM64 images**

**Features:**
- ✅ Cross-platform support (AMD64, ARM64)
- ✅ BuildKit cache mounts for faster builds
- ✅ Optional UPX compression
- ✅ Platform-aware builds

**Build:**
```bash
# Setup buildx (one-time)
docker buildx create --use

# Build for multiple platforms
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -f Dockerfile.multiarch \
  -t myapp:latest \
  --push .
```

## 🚀 Quick Start

### Using Docker Compose (Recommended)

```bash
# Start all services (app + database)
docker compose up

# Build and start
docker compose up --build

# Run in background
docker compose up -d

# View logs
docker compose logs -f app

# Stop all services
docker compose down
```

### Using Docker CLI

```bash
# Build
docker build -t myapp:latest .

# Run with environment variables
docker run -p 8000:8000 \
  -e PORT=8000 \
  -v ./config:/config \
  myapp:latest

# Run with custom config
docker run -p 8000:8000 \
  -v $(pwd)/config/config.prod.yml:/config/config.yml \
  myapp:latest
```

## 🏗️ Best Practices Implemented

### 1. **Multi-Stage Builds**
```dockerfile
FROM golang:1.24-alpine AS builder
# ... build stage ...

FROM gcr.io/distroless/static-debian12:nonroot
# ... final minimal image ...
```
- Reduces final image size by ~95%
- Builder stage discarded after compilation

### 2. **Layer Caching**
```dockerfile
# Dependencies cached separately from source code
COPY go.mod go.sum ./
RUN go mod download

# Source code copied after (changes more frequently)
COPY . .
```
- Faster rebuilds when only source code changes

### 3. **Security Hardening**
```dockerfile
# Statically linked binary
CGO_ENABLED=0

# Distroless base (no shell, no package manager)
FROM gcr.io/distroless/static-debian12:nonroot

# Non-root user
USER nonroot:nonroot
```
- Minimal attack surface
- No unnecessary binaries

### 4. **Binary Optimization**
```dockerfile
RUN CGO_ENABLED=0 go build \
    -ldflags="-s -w" \      # Strip debug info
    -trimpath \              # Remove build paths
    -o app \
    ./cmd/app
```
- Smaller binary size (~30-50% reduction)
- No debug symbols in production

### 5. **Build Arguments**
```dockerfile
ARG VERSION=dev
RUN go build -ldflags="-X main.version=${VERSION}" ...
```

Build with version:
```bash
docker build --build-arg VERSION=v1.2.3 -t myapp:v1.2.3 .
```

## 📊 Image Size Comparison

| Dockerfile | Base Image | Final Size |
|------------|------------|------------|
| Regular Go | golang:1.24 | ~1.2 GB |
| Alpine | golang:1.24-alpine | ~400 MB |
| **Distroless** | gcr.io/distroless/static | **~20 MB** ⭐ |
| Scratch | scratch | ~15 MB (no TLS/timezone) |

## 🔧 Advanced Usage

### Build with BuildKit (Faster Builds)

```bash
# Enable BuildKit
export DOCKER_BUILDKIT=1

# Or use buildx
docker buildx build -t myapp:latest .
```

### Multi-Stage Development

```dockerfile
# Target specific stage
docker build --target builder -t myapp:builder .
```

### Cache Mounts (BuildKit)

```dockerfile
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    go build -o app ./cmd/app
```

### Health Checks

```dockerfile
HEALTHCHECK --interval=30s --timeout=3s \
    CMD ["/app", "healthcheck"]
```

Test:
```bash
docker inspect --format='{{.State.Health.Status}}' container_id
```

## 🐳 Docker Compose Configuration

See `compose.yml` for the full setup with PostgreSQL.

```yaml
services:
  app:
    build: .
    ports:
      - "8000:8000"
    volumes:
      - ./config:/config
    environment:
      - DATABASE_URL=postgres://...
    depends_on:
      - db
  
  db:
    image: postgres:16-alpine
    # ... database config ...
```

## 🔐 Security Scanning

### Scan for vulnerabilities

```bash
# Using Docker Scout (built-in)
docker scout cves myapp:latest

# Using Trivy
trivy image myapp:latest

# Using Snyk
snyk container test myapp:latest
```

## 📝 Common Issues & Solutions

### 1. **Binary won't run in distroless**
- Ensure `CGO_ENABLED=0` for static linking
- Use correct architecture (GOARCH)

### 2. **Config files not found**
- Mount config directory: `-v ./config:/config`
- Or copy in Dockerfile: `COPY config /config`

### 3. **Timezone issues**
- Copy timezone data from builder
- Or use `TZ` environment variable

### 4. **CA certificates missing**
- Copy from builder: `COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/`

## 🎯 Recommended Setup by Environment

### Development
```bash
docker compose -f compose.dev.yml up
```
- Use `Dockerfile.dev`
- Volume mount source code
- Hot reload with Air

### Staging
```bash
docker build -t myapp:staging .
docker run -p 8000:8000 -v ./config/config.staging.yml:/config/config.yml myapp:staging
```
- Use production `Dockerfile`
- Staging config file

### Production
```bash
docker build -t myapp:v1.0.0 .
docker push myapp:v1.0.0
```
- Use production `Dockerfile`
- Version tags
- Deploy to Kubernetes/ECS/Cloud Run

## 📚 Additional Resources

- [Docker Best Practices](https://docs.docker.com/develop/dev-best-practices/)
- [Distroless Images](https://github.com/GoogleContainerTools/distroless)
- [Multi-stage Builds](https://docs.docker.com/build/building/multi-stage/)
- [BuildKit](https://docs.docker.com/build/buildkit/)
