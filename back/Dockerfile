# Build stage
FROM golang:1.21-alpine AS builder

# Enable Go modules caching
ENV GOCACHE=/root/.cache/go-build
ENV GOMODCACHE=/go/pkg/mod

# Set working directory
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev git bash

# Copy go mod files for caching
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy the source code
COPY . .

# Build with caching
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# Final stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose port
EXPOSE 3000

# Command to run
CMD ["./main"]
