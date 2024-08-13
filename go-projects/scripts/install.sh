#!/bin/bash

# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1

golangci-lint --version
go install golang.org/x/tools/cmd/goimports@latest
go install mvdan.cc/gofumpt@v0.6.0
go mod tidy