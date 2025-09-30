# This dockerfile runs the go app using go 1.24 and runs cmd/app on a distroless image
FROM golang:1.24-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app ./cmd/app

FROM gcr.io/distroless/static-debian12

LABEL maintainer=tagpro
USER nonroot:nonroot
COPY --from=builder --chown=nonroot:nonroot /app /app

ENTRYPOINT ["/app"]
