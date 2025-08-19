# syntax=docker/dockerfile:1

# ---------- Builder ----------
FROM golang:1.25-alpine AS builder
WORKDIR /app

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server ./cmd/api

# ---------- Final Image ----------
FROM alpine:3.20

# Add CA certificates for HTTPS requests and non-root user
RUN apk add --no-cache ca-certificates && adduser -D -u 10001 appuser

WORKDIR /home/appuser
COPY --from=builder /app/server /usr/local/bin/server

USER appuser
ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/server"]
