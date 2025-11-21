## protoc-gen-sqlc

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=viqueen_protoc-gen-plugin&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=viqueen_protoc-gen-plugin)

Template repository for creating protoc plugins in Go.

## Prerequisites

- Go 1.23.1+
- Docker (for build/lint)

## Development

```bash
# Lint
./bin/task lint

# Build (creates dist/ with cross-platform binaries)
./bin/task build

# Install locally
go install ./cmd
```

## Usage

```bash
go install github.com/<username>/go-protoc-gen-plugin/cmd@latest
```
