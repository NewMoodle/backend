include .env.local

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# MIGRATIONS
# ==================================================================================== #

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all database up migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database postgres://${POSTGRE_USERNAME}:${POSTGRE_PASSWORD}@${POSTGRE_HOST}:${POSTGRE_PORT}/${POSTGRE_DATABASE}?sslmode=${POSTGRE_SSLMODE} up

## db/migrations/up: apply all database down migrations
.PHONY: db/migrations/down
db/migrations/down:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database postgres://${POSTGRE_USERNAME}:${POSTGRE_PASSWORD}@${POSTGRE_HOST}:${POSTGRE_PORT}/${POSTGRE_DATABASE}?sslmode=${POSTGRE_SSLMODE} down

## testdb/migrations/up: apply all test database up migrations
.PHONY: testdb/migrations/up
testdb/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database postgres://tester:tester@$localhost:5432/testing?sslmode=disable up

## testdb/migrations/up: apply all test database down migrations
.PHONY: testdb/migrations/down
testdb/migrations/down:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database postgres://tester:tester@$localhost:5432/testing?sslmode=disable down

# ==================================================================================== #
# TESTS
# ==================================================================================== #

.PHONY: test
test:
	go test -v --cover ./...