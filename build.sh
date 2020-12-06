#!/bin/bash

set -eu

go env
GOOS=$(go env GOOS)
go vet -v ./...
go test -v ./... -cover
go build -v -o "bin/stat-$GOOS"
