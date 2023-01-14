
BINARY = server

.PHONY: build
build:
	go build -o "$(BINARY)" notify-api

.PHONY: build-production
build-production:
	go build -trimpath -ldflags "-s -w -extldflags=-static" -tags osusergo,netgo,sqlite_omit_load_extension -o "$(BINARY)" notify-api

.PHONY: frontend
frontend:
	go run scripts/main.go download


.PHONY: build-test
build-test:
	go build -tags test -o "$(BINARY)" notify-api

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
