# Build the project
build:
	swag init
	go build -o bin/server notify-api

dependencies:
	go mod vendor
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/cosmtrek/air@latest

# Run the project with hot reload
dev:
	air

fmt:
	swag fmt
	go fmt ./...

