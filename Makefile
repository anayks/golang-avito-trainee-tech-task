.PHONY: test
test:
	go test -v -race ./...

.PHONY: build
build:
	./wait-for-postgres.sh db ./main

.PHONY: wait
wait:
	./wait-for-postgres.sh db

.PHONY: migrate_up
migrate_up:
	migrate -database "postgres://postgres:HEYO@db/postgres?sslmode=disable" -path migrations up

.PHONY: run
run: wait migrate_up build  

.DEFAULT_GOAL := build