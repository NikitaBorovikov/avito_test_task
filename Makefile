.PHONY: up down build rebuild lint lint-fix

up:
	docker-compose up

down:
	docker-compose down

build:
	docker-compose build

rebuild:
	docker-compose build --no-cache
lint:
	@echo "Running linter..."
	golangci-lint run ./...
lint-fix:
	@echo "Running linter with auto-fix..."
	golangci-lint run --fix ./...