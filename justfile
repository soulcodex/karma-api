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