DOCKER_MOCKERY := "docker run --user $(id -u):$(id -g) --rm -v $(pwd):/src -w /src  vektra/mockery:v2.51.0"

# List all available commands
help:
    @just --list

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
    cp .env.example-local .env
    go run cmd/karma/api/main.go

# interfaceName: The interface name which your aim to generate a mock struct
# outputPkg: The package used by the mock generator to be used as package name for mocks
# Generate mocks given an interface name and the output package name
generate-mock $interfaceName $outputPkg:
    #!/usr/bin/env bash
    export DIR=$(find . -type d -name .git -prune -o -type d -print | fzf --header="Select input directory to search interfaces") && \
    export MOCK_FILE_NAME=$(echo $interfaceName | sed -r 's/([a-z])([A-Z])/\1_\L\2/g' | tr '[:upper:]' '[:lower:]') && \
    {{ DOCKER_MOCKERY }} --dir $DIR \
        --output $DIR/mocks \
        --name {{ interfaceName }} \
        --outpkg {{ outputPkg }} \
        --filename "$MOCK_FILE_NAME.go" \
        --structname {{ interfaceName }}Mock

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