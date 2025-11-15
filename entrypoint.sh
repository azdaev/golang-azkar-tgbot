#!/bin/sh
set -e

# Запуск миграций
echo "Running database migrations..."
DB_PATH=${DB_PATH:-./repository/azkar}
goose -dir ./migrations sqlite "$DB_PATH" up

# Запуск приложения
echo "Starting application..."
exec ./main
