# Pack Calculator
*Go 1.24 · Gin · GORM · SQLite / Alpine Docker*

Calculates the **least‑overshoot → least‑pack** combination to ship any customer order.  
Runs as a single container with REST API, clean HTML UI, SQLite WAL persistence and unit tests.

---

## Features
| ✔︎                            | Description                                                       |
|-------------------------------|-------------------------------------------------------------------|
| Dynamic‑programming optimiser | Minimal overshoot → minimal pack count                            |
| Clean architecture            | `domain → application → infrastructure`                           |
| SQLite persistence            | GORM + WAL file (Docker volume)                                   |
| Graceful shutdown             | Handles SIGINT / SIGTERM / SIGTSTP                                |
| Lean Docker image             | Multi‑stage Alpine, < 20 MB                                       |
| Makefile shortcuts            | `make start` · `make compose`                                     |
| Postman collection            | Included for quick API testing                                    |
| **User‑friendly UI**          | Shows result like "2 × 500 + 1 × 250 → total 1250, overshoot 249" |

---

## Quick Start

```bash
  git clone https://github.com/h6x0r/pack-calculator.git
```

<details>
<summary><b>Run locally (Go)</b></summary>

```bash
  make start
```
</details>

<details>
<summary><b>Run with Docker Compose</b></summary>

```bash
  make compose
```
</details>

---

## Make targets

| Command           | Description                                 |
|-------------------|---------------------------------------------|
| `make start`      | Run server locally via `go run`             |
| `make compose`    | Build & run stack in Docker Compose         |
| `make build`      | Build static binary to `bin/`               |
| `make test`       | Run unit tests                              |
| `make docker-run` | Build image & run `docker run -p 8081:8081` |
| `make clean`      | Remove `bin/`                               |

---

## Environment Variables

| Variable          | Default        | Purpose                       |
|-------------------|----------------|-------------------------------|
| `PACK_CALC_PORT`  | `:8081`        | HTTP bind address             |
| `PACK_CAL_DB`     | `packcalc.db`  | SQLite file path (volume)     |

Example override:

```bash
  PACK_CALC_PORT=:9090 make docker-run
```

---

## API Reference
_Base URL →_ `http://HOST:PORT/api/v1`

| Method | Path               | JSON Body                      | Description                 |
|--------|--------------------|--------------------------------|-----------------------------|
| GET    | `/packs`           | —                              | List pack sizes             |
| POST   | `/packs`           | `{ "size": 750 }`              | Add a new pack size         |
| PUT    | `/packs/{size}`    | `{ "new_size": 800 }`          | Update an existing size     |
| DELETE | `/packs/{size}`    | —                              | Delete a pack size          |
| GET    | `/calculate`       | `?items=501`                   | Calculate order             |

Example response:

```json
{
  "packs": { "500": 2, "250": 1 },
  "total": 1250,
  "overshoot": 249
}
```

---

## Postman Collection
File **`postman_collection.json`** in repo root.  
Import, set `base_url` (default `http://localhost:8081`) and try any request.

---

## Tests
```bash
  make test
```
Covers exact/inexact orders, edge‑case 500 000 (23/31/53), tie‑breakers & invalid input.

---

## Application Architecture

The project is built following Clean Architecture/DDD principles:

- **UI (`ui/`)** — static HTML user interface
- **Infrastructure (`internal/infrastructure/`)** —
  - API: HTTP handlers, routing (Gin)
  - Persistence: database access (GORM, SQLite)
- **Application (`internal/application/`)** — business logic, services, DTOs, mappers
- **Domain (`internal/domain/`)** — domain entities and repository interfaces
- **Config (`config/`)** — configuration loading
- **cmd/server/** — entry point, server startup

### Layer Interaction

```
UI → API (handlers) → Application (service) → Domain → Infrastructure (repo)
```

- **Handlers** receive requests, validate, and call services via interfaces
- **Services** implement business logic, use DTOs and mappers
- **Repositories** implement database access via interfaces

---

## Algorithm Description

- Uses dynamic programming to find the minimum number of packs with the least overshoot.
- Caching of arrays speeds up repeated calculations.
- The algorithm is robust for large orders and many pack sizes.

**Pseudocode:**
1. m = minimum pack size
2. dp[0…order+m] — minimum number of packs for sum x
3. For each size p: dp[x+p] = min(dp[x]+1, dp[x+p])
4. First reachable total ≥ order — minimal overshoot
5. Solution recovery via prev[] (minimum packs)

Cost `O(len(sizes) × (order + m))` – edge test runs in ms.

---