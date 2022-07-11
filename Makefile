PWD = $(shell pwd)

install:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/vektra/mockery/v2@latest

goose.create:
	@read -p "Migration filename: " filename; \
	cd ./database/migrations; \
	goose create $$filename sql;

goose.up:
	@read -p "Database DSN: " dsn; \
	cd ./database/migrations; \
	goose mysql "$$dsn" up

goose.down:
	@read -p "Database DSN: " dsn; \
	cd ./database/migrations; \
	goose mysql "$$dsn" down

goose.status:
	@read -p "Database DSN: " dsn; \
	cd ./database/migrations; \
	goose mysql "$$dsn" status

lint-go:
# golangci-lint run ./backend/..
	docker run --rm -v $(PWD):/app -w /app registry.gitlab.com/gitlab-org/gitlab-build-images:golangci-lint-alpine golangci-lint --timeout=3m run ./...
.PHONY: lint-go

gen-mock:
	@cd ./internal/repository; \
	mockery --name=TaskRepository --case=snake; \
	mockery --name=Repository --case=snake; \
.PHONY: gen-mock

gen-api-docs:
	swag init
.PHONY: gen-api-docs