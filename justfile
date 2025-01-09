# List all available commands
help:
    @just --list


# Run the karma api in local
run-karma-api:
    go run cmd/karma/api/main.go