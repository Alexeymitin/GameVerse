# syntax=docker/dockerfile:1

# 1. Используем официальный образ Go для сборки
FROM golang:1.23-alpine AS builder

WORKDIR /app

# 2. Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# 3. Копируем исходники сервиса auth и необходимые общие пакеты
COPY services/auth ./services/auth
COPY pkg ./pkg
COPY pkg/configs/ssl ./pkg/configs/ssl

WORKDIR /app/services/auth

# 4. Собираем бинарник
RUN go build -o /app/auth ./cmd/main.go

# 5. Используем минимальный образ для запуска
FROM alpine:latest

WORKDIR /app

# 6. Копируем бинарник и нужные файлы
COPY --from=builder /app/auth .
COPY --from=builder /app/pkg/configs/ssl ./pkg/configs/ssl

# 7. Копируем .env, если нужно (или монтируй через docker-compose)
# COPY ../../.env .env

# 8. Открываем порт (тот же, что в docker-compose)
EXPOSE 8000

# 9. Запускаем сервис
CMD ["./auth"]