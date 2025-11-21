## protoc-gen-sqlc

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=viqueen_protoc-gen-plugin&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=viqueen_protoc-gen-plugin)

Protobuf plugin to generate SQLC configuration, queries, and schema files from `.proto` definitions.

## Usage

- install it

```bash
go install github.com/labset/protoc-gen-sqlc/cmd@latest
```


## Development

### requirements

- Go 1.23.1+
- Docker (for build/lint)

### housekeeping tasks

- install dependencies

```bash
go mod download
```

- generate code

```bash
./bin/task buf generate
```

- lint the protos

```bash
./bin/task buf lint
```

- format the protos

```bash
./bin/task buf format --write
```

- lint and format the Go code

```bash
./bin/task lint --fix
```

- build the binaries

```bash
./bin/task build
```

- install it locally

```bash
go install ./cmd/...
```