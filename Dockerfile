# Goose build stage - cached separately
FROM golang:1.23-alpine AS goose-builder

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go install github.com/pressly/goose/v3/cmd/goose@latest

# Application build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies for CGO (needed for sqlite3)
RUN apk add --no-cache gcc musl-dev

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies with cache mount
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && go mod verify

# Copy application source code
COPY . .

# Build static binary with CGO for sqlite3
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

# Copy goose binary from goose-builder
COPY --from=goose-builder /go/bin/goose /usr/local/bin/goose

# Copy the bot binary from builder
COPY --from=builder /app/main .

# Copy migrations directory (needed for goose)
COPY --from=builder /app/migrations ./migrations

CMD ["./main"]
