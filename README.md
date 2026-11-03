<div align="center">

# HTTP Server

**A minimal, production-minded HTTP API built entirely with the Go standard library.**

[![Go Version](https://img.shields.io/badge/Go-1.27.3-00ADD8?style=flat&logo=go&logoColor=white)](https://go.dev/)
[![Zero Dependencies](https://img.shields.io/badge/dependencies-0-2ecc71?style=flat)](#features)
[![Standard Library](https://img.shields.io/badge/built%20with-net%2Fhttp-00ADD8?style=flat)](https://pkg.go.dev/net/http)
[![Status](https://img.shields.io/badge/status-early%20development-orange?style=flat)](#roadmap)

</div>

## Overview

A small task API that demonstrates how to structure an HTTP server in Go without any third-party frameworks: a dedicated server wrapper that owns configuration and routing, plus one file per resource handler.

> [!NOTE]
> This project is in an early stage. Handlers return stub data and there is no persistence layer yet — the focus is on a clean, idiomatic foundation.

## Features

- **Standard library only** — `net/http`, `encoding/json`, zero third-party dependencies.
- **Encapsulated server** — a `Server` type wraps `http.Server` and `http.ServeMux` behind a small interface (`New` / `Start` / `Handle`).
- **Hardened timeouts** — read, read-header, write and idle timeouts are all set to 10s from day one.
- **Method-based routing** — dispatching via `r.Method` with a `405 Method Not Allowed` fallback.
- **Layered layout** — `main` (composition) → `server` (infrastructure) → handlers (business logic).

## Project structure

```
http_server/
├── main.go    # Composition root: wires the server and registers handlers
├── server.go  # Server wrapper: http.Server config + route registration
├── tasks.go   # /api/v1/tasks handler (GET, POST)
├── go.mod     # Module definition (no go.sum: no external deps)
└── app        # Compiled binary (build artifact)
```

## Requirements

- [Go](https://go.dev/dl/) **1.27.3** or newer.

## Getting started

Run the server (listens on `:8080`):

```bash
go run .
```

Or build a binary:

```bash
go build -o app .
./app
```

Sanity checks:

```bash
go vet ./...
go build ./...
```

## API

Base URL: `http://localhost:8080`

| Method | Endpoint              | Description                    | Response                                        |
| ------ | --------------------- | ------------------------------ | ----------------------------------------------- |
| `GET`  | `/api/v1/tasks`       | Retrieve tasks                 | `200 OK` — plain text                           |
| `POST` | `/api/v1/tasks`       | Create a task (stub)           | `200 OK` — `application/json`                   |
| `*`    | any other method      | Unsupported method             | `405 Method Not Allowed`                        |

### Examples

```bash
curl http://localhost:8080/api/v1/tasks
# Getting task

curl -X POST http://localhost:8080/api/v1/tasks
# {"id":"01-task","title":"Task number 01"}
```

## Configuration

There is no external configuration. The listen address is passed in `main.go`:

```go
srv := New(":8080")
```

Server timeouts live in `server.go` and are currently fixed at `10 * time.Second`.

## Architecture

```
main.go
  └── Server (server.go) ── http.Server + ServeMux, 10s timeouts
        └── /api/v1/tasks ── TaskHandler (tasks.go)
                               ├── GET  → text stub
                               └── POST → JSON stub
```

Each layer depends only on the one below it, so handlers can be tested in isolation with `net/http/httptest` once tests are added.

## Roadmap

Planned improvements, roughly in priority order:

- [ ] Graceful shutdown on `SIGINT` / `SIGTERM` (`http.Server.Shutdown`).
- [ ] Handler tests with `net/http/httptest`.
- [ ] Structured logging with `log/slog` and a middleware chain (logging, panic recovery, CORS).
- [ ] Consistent JSON response convention for success and error cases.
- [ ] Path parameters (`GET /api/v1/tasks/{id}`) via `r.PathValue`.
- [ ] Input validation with `json.Decoder` and a real persistence layer.

### Known limitations

- `405` responses do not yet include an `Allow` header.
- Response formats are inconsistent (`GET` returns text, `POST` returns JSON).
- Error handling in the `POST` path attempts to change the status code after the response has been written.
- All responses are hardcoded stubs.
