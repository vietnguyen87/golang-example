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

# =====================================K8s=========================================
MYSQL_BIN=kubectl exec -it deployment/mysql -- mysql -uroot -p123456 -D example
MYSQL_TEST_DOCKER_CONTAINER=temporal-tables-mysql-test
MYSQL_TEST_BIN=docker exec -it $(MYSQL_TEST_DOCKER_CONTAINER) mysql -u$(MYSQL_USER) -p$(MYSQL_PASSWORD) -D$(MYSQL_DATABASE)

API_TAG=vietnt/example-api:latest

.PHONY: help
## help: shows this help message
help:
	@ echo "Usage: make [target]"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: mysql-console
## mysql-console: launches mysql local database console
mysql-console:
	@ $(MYSQL_BIN)

.PHONY: mysql-test-console
## mysql-test-console: launches mysql local test database console
mysql-test-console: export MYSQL_DATABASE=example
mysql-test-console: export MYSQL_USER=vietnt
mysql-test-console: export MYSQL_PASSWORD=123456
mysql-test-console: export MYSQL_ROOT_PASSWORD=root
mysql-test-console: export MYSQL_HOST_NAME=localhost
mysql-test-console: export MYSQL_PORT=3306
mysql-test-console:
	@ $(MYSQL_TEST_BIN)

# ==============================================================================
# Minikube

.PHONY: minikube-setup
## minikube-setup: starts minikube and build api and migration images
minikube-setup:
	@ minikube start ; \
##minikube addons enable kong ; \
	eval $$(minikube -p minikube docker-env) ; \
	docker build --no-cache -t $(API_TAG) -f k8s/api/Dockerfile .

.PHONY: minikube-dashboard
## minikube-dashboard: shows a web-based Kubernetes GUI
minikube-dashboard:
	@ minikube dashboard

# ==============================================================================

# ==============================================================================
# Deployment

.PHONY: delete-api-deployment
## delete-api-deployment: deletes api deployment. Useful for redeploying the api
delete-api-deployment:
	@ kubectl delete deployment example-api

.PHONY: deploy-api
## deploy-api: deploys the api
deploy-api:
	@ kubectl apply -f k8s/api/deployment.yaml && kubectl apply -f k8s/api/service.yaml

.PHONY: delete-db-deployment
## delete-db-deployment: deletes db deployment. Useful for redeploying the db
delete-db-deployment:
	@ kubectl delete deployment mysql

.PHONY: deploy-db
## deploy-db: deploys mysql
deploy-db:
	@ kubectl apply -f k8s/mysql/pv.yaml && kubectl apply -f k8s/mysql/pvc.yaml
	@ kubectl apply -f k8s/mysql/deployment.yaml && kubectl apply -f k8s/mysql/service.yaml

.PHONY: deploy-test-db
## deploy-test-db: deploys mysql test instance
deploy-test-db:
	@ kubectl apply -f k8s/mysql/test_db_deployment.yaml && kubectl apply -f k8s/mysql/service.yaml

.PHONY: redeploy-api
## redeploy-api: for redeploying the api
redeploy-api: delete-api-deployment build-api-img deploy-api

.PHONY: redeploy-db
## redeploy-api: for redeploying the api
redeploy-db: delete-db-deployment deploy-db

.PHONY: apply-ingress-rule
## apply-ingress-rule: applies the ingress rule
apply-ingress-rule:
	@ kubectl apply -f k8s/kong/ingress_rule.yaml

.PHONY: delete-ingress-rule
## delete-ingress-rule: deletes the ingress rule
delete-ingress-rule:
	@ kubectl delete -f k8s/kong/ingress_rule.yaml
# ==============================================================================
# Docker images

.PHONY: build-api-img
## build-api-img: builds api image
build-api-img:
	@ eval $$(minikube -p minikube docker-env) ; \
	docker build --no-cache -t $(API_TAG) -f k8s/api/Dockerfile .

.PHONY: build-migrations-img
## build-migrations-img: builds migrations image
build-migrations-img:
	@ eval $$(minikube -p minikube docker-env) ; \
	docker build --no-cache -t $(MIGRATIONS_TAG) -f k8s/mariadb/Dockerfile .

# ==============================================================================
# Database migrations

.PHONY: migrate-setup
## migrate-setup: install golang-migrate
migrate-setup:
	@ if [ -z "$$(which migrate)" ]; then echo "Installing migrate command..."; go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest; fi

.PHONY: create-migrations
## create-migration: creates up and down migration files for a given name (make create-migrations NAME=<desired_name>)
create-migration: migrate-setup
	@ if [ -z "$(NAME)" ]; then echo >&2 please set the name of the migration via the variable NAME; exit 2; fi
	@ migrate create -ext sql -dir db/migrations -seq -digits 4 $(NAME)

.PHONY: migrate-db
## migrate-db: runs database migrations
migrate-db: build-migrations-img
	@ echo "Migrating MariaDB..."
	@ until $(MARIADB_BIN) -e 'SELECT 1' >/dev/null 2>&1 && exit 0; do \
	  >&2 echo "MariaDB not ready, sleeping for 5 secs..."; \
	  sleep 5 ; \
	done
	@ echo "... MariaDB is up and running!"
	@ echo "Applying database migrations..."
	@ kubectl apply -f k8s/mariadb/migrations.yaml
	@ echo "... done."
## migrate-test-db: runs database migrations into test db
migrate-test-db:
	@ echo "Setting up test Mysql..."
	@ unset `env|grep DOCKER|cut -d\= -f1` ;\
	docker-compose up -d mysql_test migrate_test
	@ until $(MYSQL_TEST_BIN) -e 'SELECT 1' >/dev/null 2>&1 && exit 0; do \
	  >&2 echo "Mysql not ready, sleeping for 5 secs..."; \
	  sleep 5 ; \
	done
	@ echo "... Mysql is up and running!"
	
# ==============================================================================

# ==============================================================================
# Kong

.PHONY: kong
## kong: on macbooks with m1 chip, we need to keep terminal open to run Kong
kong:
	@ minikube service -n kong kong-proxy --url | head -1

# ==============================================================================
# Cleanup
.PHONY: cleanup
## cleanup: cleans everything
cleanup:
	@ kubectl delete deployment mysql
	@ kubectl delete service mysql
	@ kubectl delete deployment example-api
	@ kubectl delete service example-api
	@ kubectl delete ingress example-api-ingress

# ==============================================================================
# App's execution

.PHONY: run
## run: does all needed setup and runs the api
run: minikube-setup deploy-db deploy-api