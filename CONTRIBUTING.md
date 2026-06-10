# Contributing to DevStats Card

Thanks for your interest in contributing! This project is still under active
development, so ideas, bug reports, and pull requests are all welcome.

## Reporting Issues & Suggesting Features

- Search [existing issues](https://github.com/tico88612/devstats-card/issues) first to avoid duplicates.
- For bugs, include the URL you requested (e.g. `https://devstats.me/?username=...`), what you expected, and what you got instead.
- For feature ideas, a short description of the use case is enough — no formal template required.

## Development Setup

### Prerequisites

- [Go](https://go.dev/dl/) 1.26 or later
- (Optional) [Docker](https://www.docker.com/) if you want to test the container image

### Running locally

```bash
git clone https://github.com/<your-username>/devstats-card.git
cd devstats-card
go run .
```

The server listens on `:8080`. Try it out:

```bash
# Preview frontend page
curl http://localhost:8080/

# SVG card for a user
curl "http://localhost:8080/?username=tico88612"

# Health check
curl http://localhost:8080/health
```

### Building the Docker image

```bash
docker build -t devstats-card .
docker run -p 8080:8080 devstats-card
```

## Project Structure

| Directory   | Purpose                                           |
| ----------- | ------------------------------------------------- |
| `handlers/` | Gin route handlers (card rendering, health check) |
| `service/`  | Business logic that aggregates DevStats data      |
| `pkg/`      | Clients for external APIs (CNCF DevStats)         |
| `models/`   | Shared data structures                            |
| `svg/`      | SVG card template and rendering                   |
| `theme/`    | Card color themes                                 |
| `web/`      | Embedded frontend preview page                    |

## Before Submitting a Pull Request

1. Make sure the code builds and passes vet:

   ```bash
   go build ./...
   go vet ./...
   gofmt -l .   # should print nothing
   ```

2. If you add tests (encouraged!), make sure they pass with `go test ./...`.
3. Keep each pull request focused on a single change.

## Commit Message Convention

Use a short type prefix followed by a concise summary, matching the existing
history:

```
Feat: add HEAD method for UptimeRobot healthcheck
Fix: no PRcount case
Docs: Status Page information
Refactor: get PR count & issue count from new DevStats API
Chore: go 1.26 bump
```

Common prefixes: `Feat:`, `Fix:`, `Docs:`, `Refactor:`, `Chore:`.

## Pull Request Process

1. Fork the repository and create a branch from `main`.
2. Make your changes and verify them locally (see above).
3. Open a pull request against `main` and describe **what** changed and **why**.
4. CI builds the project (and the container image) on every pull request — make sure it stays green.

## License

By contributing, you agree that your contributions will be licensed under the
same license as the project (see [LICENSE](LICENSE)).