build:
	@go build -o bin/jwt-auth cmd/api/main.go

run: build
	@./bin/jwt-auth

test:
	@go test -v ./...

watch:
	@air

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

