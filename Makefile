
BINARY = server

.PHONY: build
build:
	go build -o "$(BINARY)" notify-api

.PHONY: frontend
frontend:
	go run notify-api/scripts download

.PHONY: build-production
build-production: frontend
	go build -trimpath -ldflags "-s -w" -o "$(BINARY)" notify-api

.PHONY: build-test
build-test: frontend
	go build -tags test -o "$(BINARY)" notify-api

.PHONY: unit-test
unit-test: frontend
	go test -v -tags test ./...

.PHONY: dependencies
dependencies:
	go mod vendor
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/cosmtrek/air@latest
	go install entgo.io/ent/cmd/ent@latest

.PHONY: dev
dev:
	air

.PHONY: fmt
fmt:
	swag fmt
	go fmt ./...

.PHONY: desc
desc:
	go generate ent/generate.go

.PHONY: gen
gen:
	go run ent/generate.go
	swag init
