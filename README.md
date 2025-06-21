# oauth2-token-introspection

[![Go Report Card](https://goreportcard.com/badge/github.com/wallanaq/oauth2-token-introspection)](https://goreportcard.com/report/github.com/wallanaq/oauth2-token-introspection)
[![Build](https://img.shields.io/github/actions/workflow/status/wallanaq/oauth2-token-introspection/ci.yml?branch=main)](https://github.com/wallanaq/oauth2-token-introspection/actions)
[![License](https://img.shields.io/github/license/wallanaq/oauth2-token-introspection)](LICENSE)

> A lightweight OAuth2 Token Introspection API implementation following [RFC 7662](https://datatracker.ietf.org/doc/html/rfc7662), built in Go.


## 🚀 Features

- RFC-compliant `/introspect` endpoint
- Graceful shutdown and production-ready HTTP server
- Docker-ready
- Configurable timeouts
- Simple and extensible architecture


## 📦 Requirements

- Go 1.22+
- (Optional) [Docker](https://www.docker.com/)
- (Optional) [Make](https://www.gnu.org/software/make/)


## 🔧 Build and Run

### Using Make

```bash
make build       # build binary to ./bin
make run         # run the application locally
```

### Without Make

```bash
go build -o bin/introspect ./cmd/api
./bin/introspect
```

## 🐳 Docker

### Build Docker image

```bash
make build-image
```

### Run Docker container

```bash
make docker-run
```

### Clean

```bash
make docker-clean
```

## 🔐 Token Introspection (RFC 7662)

### Request
```h
POST /introspect HTTP/1.1
Content-Type: application/x-www-form-urlencoded
Authorization: Basic base64(client_id:client_secret)

token=abc123
```

### Response
```json
{
  "active": true,
  "scope": "read write",
  "client_id": "my-client",
  "username": "user@example.com",
  "exp": 1685582590,
  "iat": 1685581990,
  "sub": "user123",
  "aud": "resource-server",
  "iss": "https://sso.example.com",
  "token_type": "access_token"
}
```

## 🧹 Clean build artifacts

```bash
make clean
```
