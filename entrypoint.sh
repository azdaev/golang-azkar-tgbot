#!/bin/sh
set -e

# Запуск миграций
echo "Running database migrations..."
goose -dir ./migrations sqlite ./repository/azkar up

# Запуск приложения
echo "Starting application..."
exec ./main
