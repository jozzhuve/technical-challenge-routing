.PHONY: run test quality build

run:
	go run ./cmd/api

test:
	go test ./... -cover

quality:
	gofmt -w .
	go vet ./...
	go test ./... -race -cover

build:
	CGO_ENABLED=0 go build -o bin/routing ./cmd/api
