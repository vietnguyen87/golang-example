PWD = $(shell pwd)

install:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/vektra/mockery/v2@latest
	go install github.com/swaggo/swag/cmd/swag@latest

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
# golangci-lint run ./...
	docker run -v $(PWD):/app -w /app golangci/golangci-lint:v1.52.2 golangci-lint -v --tests=false --skip-dirs=vendor  --timeout=3m run ./...
.PHONY: lint-go

gen-mock:
	@cd ./internal/repository; \
	mockery --name=TaskRepository --case=snake; \
	mockery --name=Repository --case=snake; \
.PHONY: gen-mock

gen-api-docs:
	swag init --parseDependency --parseInternal true
.PHONY: gen-api-docs