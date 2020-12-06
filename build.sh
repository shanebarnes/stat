#!/bin/bash

set -eu

go env
go vet -v ./...
go build -v ./...
