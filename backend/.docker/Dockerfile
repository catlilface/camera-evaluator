FROM golang:1.24-alpine AS builder

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

CMD ./app