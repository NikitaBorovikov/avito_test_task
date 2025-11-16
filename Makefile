.PHONY: up down build rebuild

up:
	docker-compose up

down:
	docker-compose down

build:
	docker-compose build

rebuild:
	docker-compose build --no-cache
