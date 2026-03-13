# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run the API
go run ./cmd/api

# Build
go build ./cmd/api

# Lint (fmt + vet + build)
just lint

# Migrations
just mig-run                    # apply all pending migrations
just mig-create DESC=<name>     # create new migration files
just mig-down MANY=1            # roll back N migrations
just mig-goto VERSION=<n>       # migrate to specific version

# Kill API process on port
just kill-api
```

No test suite exists yet.

## Architecture

This is a JSON REST API built with Go's standard library + httprouter. The entry point is `cmd/api/main.go`.

**Request flow:** `routes.go` → handler in `movies.go`/`users.go`/`healthcheck.go` → `internal/data/` model → PostgreSQL

**Key types:**
- `application` struct (main.go) — holds config, slog logger, and `data.Models`; all handlers are methods on it
- `config` struct — parsed from CLI flags; DB DSN from `GREENLIGHT_DB_DSN` (set in `.zshrc`)
- `data.Models` — aggregates `MovieModel` and `UserModel`, each wrapping `*sql.DB`

**`cmd/api/` packages:**
- `main.go` — startup, DB pool init, flag parsing
- `server.go` — HTTP server with graceful shutdown
- `routes.go` — all route registrations; middleware chain: `recoverPanic → rateLimit → router`
- `middleware.go` — panic recovery, per-IP rate limiting (golang.org/x/time/rate)
- `helpers.go` — JSON read/write utilities, parameter parsing
- `errors.go` — typed JSON error responses

**`internal/data/`:**
- `models.go` — `Models` struct, shared sentinel errors (`ErrRecordNotFound`, `ErrEditConflict`)
- `movie.go` — `MovieModel` with CRUD + list/filter/sort via `Filters`
- `users.go` — `UserModel`, bcrypt password hashing via `password` type
- `filters.go` — `Filters` struct for pagination and sorting; `Metadata` for response envelope
- `runtime.go` — custom `Runtime` type that marshals/unmarshals as `"N mins"`
- `validator.go` (internal/validator) — simple `Validator` with field-level error accumulation

**Database:** PostgreSQL with `golang-migrate` for schema migrations (sequential numbered SQL files in `migrations/`). Movies table has GIN indexes for full-text search on title and genre array.

**Response envelope:** All responses use `{"movie": ...}` or `{"movies": ..., "metadata": {...}}` JSON wrappers, written via `app.writeJSON()`.
