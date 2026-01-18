# Stage 1: Goose builder (для миграций)
FROM golang:1.25-alpine3.22 AS goose-builder

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go install github.com/pressly/goose/v3/cmd/goose@latest

# Stage 2: Application builder
FROM golang:1.25-alpine3.22 AS builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Скачиваем зависимости с кешем
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && go mod verify

# Копируем исходники
COPY . .

# Сборка приложения (CGO не нужен)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -ldflags="-s -w" -o main .

# Stage 3: Runtime
FROM alpine:3.19

WORKDIR /app

# Устанавливаем сертификаты для HTTPS
RUN apk --no-cache add ca-certificates tzdata

# Копируем goose
COPY --from=goose-builder /go/bin/goose /usr/local/bin/goose

# Копируем приложение
COPY --from=builder /app/main .

# Копируем миграции
COPY --from=builder /app/migrations ./migrations

# Копируем entrypoint скрипт
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

CMD ["./entrypoint.sh"]
