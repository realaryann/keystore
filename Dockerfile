# --- Stage 1: Build ---
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# --- Stage 2: Run ---
FROM alpine:latest
WORKDIR /root/

# Copy the binary
COPY --from=builder /app/main .

ENTRYPOINT ["./main"]