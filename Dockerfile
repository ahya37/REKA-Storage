# ===============================
# Build stage
# ===============================
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git (dibutuhkan untuk go mod)
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -o app ./cmd/server/main.go

# ===============================
# Runtime stage
# ===============================
FROM alpine:latest

WORKDIR /app

# CA certs for HTTPS (MinIO / Mongo)
RUN apk add --no-cache ca-certificates

# Copy binary
COPY --from=builder /app/app .

# Copy env file (optional)
# COPY .env .env

EXPOSE 4001

CMD ["./app"]
