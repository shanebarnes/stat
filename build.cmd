@echo off

go env || exit /b
go vet -v ./... || exit /b
go build -o bin\stat-windows.exe -v ./... || exit /b
