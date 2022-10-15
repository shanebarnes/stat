#!/bin/bash

set -eu

go env
GOARCH=$(go env GOARCH)
GOOS=$(go env GOOS)
gofmt -d .
go vet -v ./...
go test -v ./... -cover
go build -v -o "bin/stat-${GOOS}-${GOARCH}" cmd/app/stat.go
