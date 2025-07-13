FROM golang:1.24-alpine AS builder
# or use 1.24.4 specifically
# FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o urlshortener ./cmd/server

# Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/urlshortener .
# COPY --from=builder /app/configs ./configs # If you have configs
EXPOSE 8080

# Set ENV if you want to override from docker run
# ENV DATABASE_URL="postgres://user:password@host.docker.internal:5432/urlshortener?sslmode=disable"

CMD ["./urlshortener"]
