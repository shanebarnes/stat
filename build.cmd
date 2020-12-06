@echo off

go env || exit /b
go vet -v ./... || exit /b
go test -v ./... -cover || exit /b
go build -v -o bin\stat-windows.exe || exit /b
