.PHONY: up down test run

up:
	docker-compose up -d

down:
	docker-compose down

test:
	go test ./...

run:
	go run cmd/main.go
