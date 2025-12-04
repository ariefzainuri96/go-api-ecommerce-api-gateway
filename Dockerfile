# ============================
# 1 — Build Stage
# ============================
FROM golang:1.24 AS builder

WORKDIR /app

# Cache modules first
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway ./cmd/api

# ============================
# 2 — Runtime Stage
# ============================
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/api-gateway .

# Expose the port your service uses
EXPOSE 8080

# Run the application
CMD ["./api-gateway"]