````md
#  Gator — RSS Aggregator CLI (Go + PostgreSQL)

A concurrent RSS feed aggregator built in Go using PostgreSQL, SQLC, and Goose.

Gator allows users to register, follow RSS feeds, and continuously scrape and browse posts from those feeds via a CLI.

---

##  Overview

Gator is designed as a simple distributed-style system:

- Users register and authenticate locally
- Users subscribe to RSS feeds
- A background aggregator fetches feeds continuously
- Posts are stored and queried efficiently via SQLC

This project focuses on:
- database design
- concurrency
- CLI architecture
- type-safe SQL in Go

---

##  Tech Stack

- Go (CLI application)
- PostgreSQL (data storage)
- SQLC (type-safe queries)
- Goose (migrations)

---

##  Installation

### 1. Install CLI

```bash
go install github.com/MoisesASantos/BLOG-AGGREGATOR@latest
````

Ensure Go binaries are in PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

---

##  Database Setup

Create database:

```bash
createdb gator
```

Run migrations:

```bash
goose postgres "postgres://postgres:password@localhost:5432/gator?sslmode=disable" up
```

---

##  Configuration

Create:

```bash
~/.gatorconfig.json
```

```json
{
  "db_url": "postgres://postgres:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

---

##  Usage

### Run in development

```bash
go run . <command>
```

### Run installed CLI (production)

```bash
gator <command>
```

---

##  Commands

### Register user

```bash
gator register john
```

### Login

```bash
gator login john
```

### Add feed

```bash
gator addfeed "Hacker News" "https://hnrss.org/newest"
```

### Follow feed

```bash
gator follow "https://hnrss.org/newest"
```

### List followed feeds

```bash
gator following
```

### Browse posts

```bash
gator browse 10
```

(Default: 2 posts)

### Run scraper (aggregator loop)

```bash
gator agg 10s
```

---

##  Architecture

```
CLI Commands
    ↓
State Layer
    ↓
SQLC Queries
    ↓
PostgreSQL
```

### Components:

* **Goose** → schema migrations
* **SQLC** → typed SQL generation
* **Aggregator** → RSS polling loop
* **CLI layer** → command dispatch system

---

##  Scraper Behavior

* Fetches next feed using `last_fetched_at`
* Parses RSS/Atom feeds
* Ignores invalid HTML feeds
* Inserts posts into DB
* Deduplicates using `url UNIQUE`
* Runs continuously via ticker

---

##  Docker Setup (optional but recommended)

### docker-compose.yml

```yaml
version: "3.9"

services:
  db:
    image: postgres:16
    container_name: gator-db
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
      POSTGRES_DB: gator
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

volumes:
  pgdata:
```

Run:

```bash
docker compose up -d
```

---

##  Makefile

```makefile
build:
	go build -o gator .

run:
	go run .

migrate-up:
	goose postgres "$(DB_URL)" up

migrate-down:
	goose postgres "$(DB_URL)" down

install:
	go install .
```

---

##  GitHub Actions (CI)

```yaml
name: Go CI

on:
  push:
  pull_request:

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: 1.22

      - name: Build
        run: go build ./...

      - name: Test
        run: go test ./...
```

---

##  Notes

* Only RSS/Atom feeds supported
* HTML feeds are ignored safely
* Duplicate posts are automatically ignored
* Scraper runs in continuous loop using configurable interval

---

##  Build

```bash
go build -o gator .
```

Run:

```bash
./gator
```

---

##  Philosophy

This project follows a simple principle:

> “Keep the system small, predictable, and explicit.”

No ORMs. No magic. Just SQL + Go.

---

##  License

MIT

```

---

```
