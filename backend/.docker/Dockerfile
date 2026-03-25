FROM golang:1.24 AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

FROM debian:buster-slim

RUN addgroup --system appuser && adduser --system --group appuser --uid 1001

USER appuser

WORKDIR /app

COPY --from=builder /go/src/app/app .
COPY --from=builder /go/src/app/configs ./configs

# Переменная окружения для пути к конфигу
ENV CONFIG_PATH=./configs/docker.yaml

# Используем shell form для подстановки переменной
CMD ./app --config-file "$CONFIG_PATH"