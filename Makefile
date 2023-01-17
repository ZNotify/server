
BINARY = server

.PHONY: build frontend build-production build-test unit-test dependencies dev fmt desc gen

build:
	go build -o "$(BINARY)" notify-api

frontend:
	go run notify-api/scripts download

build-production: frontend
	go build -trimpath -ldflags "-s -w" -o "$(BINARY)" notify-api

build-test: frontend
	go build -tags test -o "$(BINARY)" notify-api

unit-test: frontend
	go test -v -tags test ./...

dependencies:
	go mod vendor
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/cosmtrek/air@latest
	go install entgo.io/ent/cmd/ent@latest

dev:
	air

fmt:
	swag fmt
	go fmt ./...

desc:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./schema

gen:
	go run notify-api/scripts ent
	swag init
