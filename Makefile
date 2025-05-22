.DEFAULT_GOAL := build

.PHONY: fmt vet build clean
fmt:
	@go fmt ./...
vet: fmt
	@go vet ./...
build: vet
	@go build -o bin/ecom.exe cmd/main.go
clean:
	@go clean ./...
test:
	@go test -v ./...
run: build
	@./bin/ecom.exe

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))
migrate-up:
	@go run cmd/migrate/main.go up
migrate-down:
	@go run cmd/migrate/main.go down