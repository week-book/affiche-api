.PHONY: test test-db down-db

test-db:
	docker compose up -d db
	sleep 2

down-db:
	docker compose down

test: test-db
	go test ./...

