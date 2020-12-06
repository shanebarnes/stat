#!/bin/bash

set -eu

go env
GOOS=$(go env GOOS)
go vet -v ./...
go build -o "bin/stat-$GOOS" -v ./...
