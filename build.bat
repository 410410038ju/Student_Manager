@echo off

echo Building Windows...
go build -o build/app.exe

echo Building Linux...
set GOOS=linux
set GOARCH=amd64
go build -o build/app-linux

echo Done!