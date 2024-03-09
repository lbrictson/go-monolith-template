.PHONY: help
help:
	@echo "The following commands are available:"
	@echo "  make test - run tests"
	@echo "  make run - start dependencies and server"
	@echo "  make hot - start server with hot reload"
	@echo "  make stop - stop dependencies"
	@echo "  make clean - delete all data"
	@echo "  make build - build server"
	@echo "  make build-docker - build docker image"
	@echo "  make generate - perform code gen"
	@echo "  make dependencies - install dependencies"

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v -cover ./pkg/...
	@rm -f local/test_dbs/*.db

.PHONY: run
run:
	@echo "Starting dependencies..."
	@docker-compose up -d
	@echo "Starting server..."
	@templ generate
	@go run cmd/server/main.go

.PHONY: stop
stop:
	@echo "Stopping dependencies..."
	@docker-compose down

.PHONY: clean
clean:
	@echo "Deleting all data..."
	@docker-compose down -v

.PHONY: build
build:
	@echo "Building server..."
	@mkdir -p bin
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server cmd/server/main.go
	@echo "Server built in bin/server"

.PHONY: generate
generate:
	@echo "Performing code gen..."
	@templ generate
	@go generate ./...

.PHONY: dependencies
dependencies:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy
	@go mod vendor
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get github.com/a-h/templ

.PHONY: hot
hot:
	@echo "Starting server with hot reload..."
	@echo "Starting dependencies..."
	@docker-compose up -d
	@echo "Starting server..."
	@air -c .air.toml

.PHONY: build-docker
build-docker:
	@echo "Building docker image..."
	@echo "Building docker image: go-monolith-template:$(shell git rev-parse HEAD)"
	@docker build -t go-monolith-template:$(shell git rev-parse HEAD) .
	@echo "Docker image built: go-monolith-template:$(shell git rev-parse HEAD)"