# include .envrc

run:
	@go run cmd/api/*.go
build:
	@go build -o bin/social cmd/api/*.go
run-build: build
	@./bin/social
.PHONY: migration
migration:
	@migrate create -seq -ext sql -dir ./cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))
.PHONY: migrate-up
migrate-up:
	@go run cmd/migrate/main.go up
.PHONY: migrate-down
migrate-down:
	@go run cmd/migrate/main.go down

.PHONY: seed
seed:
	@go run cmd/migrate/seed/main.go
run-frontend:
	@sh run-frontend.sh