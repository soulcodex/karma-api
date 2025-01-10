# List all available commands
help:
    @just --list

# Run the karma api in local
run-karma-api:
    go run cmd/karma/api/main.go

# Start necessary infra to run the integration tests
start:
    #!/usr/bin/env bash
    cp .env.example .env
    docker compose up -d

# Shutdown development infra
stop flags="--volumes --remove-orphans":
    #!/usr/bin/env bash
    docker compose down {{flags}}

# Run karma api binary
run-api:
    #!/usr/bin/env bash
    go run cmd/karma/api/main.go

# Run tests
test:
    #!/usr/bin/env bash
    go test -race ./... -timeout 300s -count 1

# Create new database migration
new-migration:
    [ ! -f .env ] || export $(grep -v '^#' .env | xargs) && \
    sql-migrate new -config ./migrations/dbconfig.yaml

# Fetch database migrations individual status
migrate-status:
    [ ! -f .env ] || export $(grep -v '^#' .env | xargs) && \
    sql-migrate status -config ./migrations/dbconfig.yaml

# Run migration up
migrate-up:
    [ ! -f .env ] || export $(grep -v '^#' .env | xargs) && \
    sql-migrate up -config ./migrations/dbconfig.yaml

# Run migration down
migrate-down:
    [ ! -f .env ] || export $(grep -v '^#' .env | xargs) && \
    sql-migrate down -config ./migrations/dbconfig.yaml