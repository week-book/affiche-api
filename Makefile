.PHONY: test up down

up:
	docker compose up -d db migrate

test: up
	go test ./... -v

down:
	docker compose down -v
