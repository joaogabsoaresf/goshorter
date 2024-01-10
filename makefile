.PHONY: default run build test docs clean
# Variables
APP_NAME=goshorter

# Tasks
default: run

run:
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...