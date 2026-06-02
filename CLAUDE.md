# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What This Project Is

A full-stack courses management platform. Users can browse/enroll in courses and comment on them. Admins can create/edit/delete courses and manage user types.

## Running the Project

```bash
# Full stack (recommended) — API on :4000, frontend on :3000, MySQL on :3307
docker-compose up --build

# Backend only (with hot reload via air)
cd backend && go run main.go

# Frontend only
cd frontend && pnpm start
```

## Backend

**Stack:** Go 1.24 · Gin · GORM · MySQL 8

**Layer architecture:**
- `controllers/` — HTTP handlers (parse request, call service, return response)
- `services/` — Business logic
- `dao/` — GORM model structs + DB queries via `clients/database.go`
- `dto/` — Request/response shapes (separate from GORM models)
- `app/middlewares/` — `requireAuth` (JWT) and `requireAdmin`

**Entry points:**
- `backend/main.go` — loads `.env`, connects DB, registers routes
- `backend/app/url_mappings.go` — all route definitions
- `backend/clients/database.go` — DB connection + GORM auto-migrate

**Run tests:**
```bash
cd backend && go test ./...
```

**Environment variables** (in `backend/.env`):
- `PORT` — server port (default `4000`)
- `SECRET` — JWT signing secret
- `DB` — MySQL DSN (`user:pass@tcp(host:port)/dbname`)

## Frontend

**Stack:** React 18 · React Router v6 · Create React App · pnpm

**Package manager is pnpm** — always use `pnpm` instead of `npm`.

```bash
cd frontend && pnpm install   # install deps
cd frontend && pnpm start     # dev server
cd frontend && pnpm test      # run tests
cd frontend && pnpm build     # production build
```

Components live in `frontend/src/components/`, each paired with its own `.css` file. Routes are defined in `frontend/src/App.jsx`.

**Key dependencies:** `js-cookie` (session), `sweetalert2` (modals), `lucide-react` + `react-icons` (icons), `react-responsive` (breakpoints).

## Key Architectural Notes

- GORM auto-migration runs on startup — schema changes happen via model struct changes in `dao/`, not manual SQL.
- JWT is stored as a cookie. The `requireAuth` middleware validates it; `requireAdmin` additionally checks `userType`.
- Categories are a many-to-many relation with courses (GORM association table).
- The `dto/` types are intentionally separate from `dao/` model types — don't merge them.
- `init.sql` only seeds initial data; schema is owned by GORM.
