SHELL := /bin/bash

.PHONY: format
format:
	go mod tidy
	goimports -local "github.com/laansdole/go-sightseeing/go-projects/subscription-service" -w .
	gofumpt -w .

.PHONY: lint
lint:
	golangci-lint run
