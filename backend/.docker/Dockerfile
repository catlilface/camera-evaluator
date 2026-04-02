FROM golang:1.24-alpine AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

FROM alpine:latest

RUN addgroup -g 1001 -S appuser && adduser -S appuser -u 1001

USER appuser

WORKDIR /app

COPY --from=builder /go/src/app/app .

CMD ["./app"]