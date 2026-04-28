.PHONY: help tidy run dev worker risk-calc build test fmt vet \
        infra-up infra-down infra-logs \
        migrate-up migrate-down migrate-create migrate-version \
        seed

GO            ?= go
APP_PORT      ?= 8080
PG_DSN        ?= postgres://ai_hs:ai_hs@127.0.0.1:5432/ai_hs_go?sslmode=disable
MIGRATIONS_DIR ?= ./migrations

help:
	@echo "make tidy            # go mod tidy"
	@echo "make run             # start HTTP server"
	@echo "make dev             # start HTTP server with air hot-reload"
	@echo "make worker          # start asynq worker (Phase 7)"
	@echo "make risk-calc       # enqueue a one-off risk:calculate task"
	@echo "make build           # compile all binaries"
	@echo "make test            # go test ./... -race"
	@echo "make fmt vet         # gofmt + go vet"
	@echo "make infra-up        # start postgres + redis via docker compose"
	@echo "make infra-down      # stop infra"
	@echo "make migrate-up      # apply migrations"
	@echo "make migrate-down N=1 # roll back N migrations"
	@echo "make migrate-create NAME=xxx  # create migration files"
	@echo "make seed            # seed roles/permissions/superadmin (Phase 4)"

tidy:
	$(GO) mod tidy

run:
	$(GO) run ./cmd/server -config ./configs/config.yaml

dev:
	$(GO) tool air

worker:
	$(GO) run ./cmd/worker -config ./configs/config.yaml

build:
	$(GO) build -o ./bin/server ./cmd/server
	$(GO) build -o ./bin/worker ./cmd/worker
	$(GO) build -o ./bin/migrate ./cmd/migrate
	$(GO) build -o ./bin/risk-calc ./cmd/risk-calc

risk-calc:
	$(GO) run ./cmd/risk-calc -config ./configs/config.yaml

test:
	$(GO) test ./... -race -count=1

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

infra-up:
	docker compose -f ./scripts/docker-compose.yaml up -d

infra-down:
	docker compose -f ./scripts/docker-compose.yaml down

infra-logs:
	docker compose -f ./scripts/docker-compose.yaml logs -f

migrate-up:
	$(GO) run ./cmd/migrate -dsn "$(PG_DSN)" -dir "$(MIGRATIONS_DIR)" up

migrate-down:
	$(GO) run ./cmd/migrate -dsn "$(PG_DSN)" -dir "$(MIGRATIONS_DIR)" down -n $(or $(N),1)

migrate-version:
	$(GO) run ./cmd/migrate -dsn "$(PG_DSN)" -dir "$(MIGRATIONS_DIR)" version

migrate-create:
	@if [ -z "$(NAME)" ]; then echo "Usage: make migrate-create NAME=xxx"; exit 1; fi
	$(GO) run ./cmd/migrate -dir "$(MIGRATIONS_DIR)" create -name "$(NAME)"

seed:
	$(GO) run ./scripts/seed -config ./configs/config.yaml
