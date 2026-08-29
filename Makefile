help:
	@echo "Available commands:"
	@echo "  build          Build the application"
	@echo "  run            Run the application"
	@echo "  dev            Run the application in development mode"
	@echo "  lint           Run golangci-lint on the codebase"
	@echo "  migrate-up     Apply database migrations"
	@echo "  migrate-down   Rollback database migrations"
	@echo "  docker-up      Start Docker containers"
	@echo "  docker-down    Stop Docker containers"

build:
	go build -o bin/e-commerce-go ./cmd/api

run:
	go run ./cmd/api

dev:
	go run ./cmd/api 

lint:
	golangci-lint run ./...

migrate-up:
	migrate -path db migrations -database "postgresql://postgres:REDACTED@localhost:5432/ecommerce_shop?sslmode=disable" up

migrate-down:
	migrate -path db migrations -database "postgresql://postgres:REDACTED@localhost:5432/ecommerce_shop?sslmode=disable" down

docker-up:
	docker compose -f docker/docker-compose.yml up -d

docker-down:
	docker compose -f docker/docker-compose.yml down