#!/usr/bin/env bash
# Stops the process if something fails
set -xe

# get all of the dependencies needed
go get "github.com/gin-gonic/gin v1.5.0"
go get "github.com/go-sql-driver/mysql v1.4.1"
go get "github.com/joho/godotenv v1.3.0"

# create the application binary that eb uses
GOOS=linux GOARCH=amd64 go build -o bin/application -ldflags="-s -w"
