# Build stage
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /bin/api ./cmd/api

# Runtime image
FROM scratch

COPY --from=builder /bin/api /api

EXPOSE 8080

ENTRYPOINT ["/api"]
