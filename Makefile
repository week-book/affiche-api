.PHONY: test up down

up:
	docker compose up -d db migrate

test: up
	go test ./...

down:
	docker compose down -v
