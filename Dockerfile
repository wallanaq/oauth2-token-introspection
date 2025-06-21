# Build stage
FROM golang:1.24.4-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o /bin/api ./cmd/api

# Runtime image
FROM alpine:3.22

COPY --from=builder /bin/api /api

EXPOSE 8080

ENTRYPOINT ["/api"]
