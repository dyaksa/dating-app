include .env

MIGRATE=migrate
DATABASE_URL=postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

# Nama direktori tempat file migrasi disimpan
MIGRATIONS_DIR=migrations

.PHONY : clean install build create-migrations migrate rollback create-seeder run-seed

run:
	go run ./cmd/main.go

clean:
	go mod tidy

run-this:
	echo "hello"

install:
	go mod download

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags musl -o main ./cmd/main.go

start:
	./main

test:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o cover.html

create-migrations:
ifndef name
	$(error "name is undefined. Usage: make migrate name=YOUR_MIGRATION_NAME")
endif
	goose -dir=$(MIGRATIONS_DIR) create $(name) sql

migrate:
	goose -dir=$(MIGRATIONS_DIR) postgres $(DATABASE_URL) up

rollback:
	goose -dir=$(MIGRATIONS_DIR) postgres $(DATABASE_URL) down


