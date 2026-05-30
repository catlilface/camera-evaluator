FROM golang:1.24-alpine AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

FROM alpine:latest

RUN addgroup -g 1001 -S appuser && adduser -S appuser -u 1001

WORKDIR /app

COPY --from=builder /go/src/app/app .

RUN mkdir -p /app/.photos && chown -R appuser:appuser /app/.photos

USER appuser

CMD ["./app"]
