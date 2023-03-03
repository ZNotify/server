BIN_DIR = bin
EXE = .exe
BINARY = server$(EXE)

GO_BUILD_FLAGS = -trimpath -ldflags '-s -w' -o "$(BIN_DIR)/$(BINARY)"

IS_LINUX = false

ifeq ($(OS),Windows_NT)
	EXE = .exe
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		IS_LINUX = true
	endif
endif

ifeq ($(IS_LINUX),true)
	GO_BUILD_FLAGS = -tags 'osusergo netgo static_build' -trimpath -ldflags '-s -w -extldflags "-fno-PIC -static"' -o "$(BIN_DIR)/$(BINARY)"
endif

.PHONY: build frontend build-production build-test unit-test dependencies dev fmt desc gen ent-gen swag-gen

build:
	go build -o "$(BIN_DIR)/$(BINARY)" github.com/ZNotify/server

app/api/web/static/index.html:
	go run github.com/ZNotify/server/scripts download

frontend: app/api/web/static/index.html

build-production: frontend
	go build $(GO_BUILD_FLAGS) github.com/ZNotify/server

build-test: frontend
	go build -tags test -o "$(BIN_DIR)/$(BINARY)" github.com/ZNotify/server

unit-test: frontend
	go test -v -tags test ./...

dependencies:
	go mod vendor
	go install github.com/swaggo/swag/cmd/swag@master
	go install github.com/cosmtrek/air@latest
	go install entgo.io/ent/cmd/ent@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

dev:
	air

fmt:
	swag fmt
	go fmt ./...

desc:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./schema

ent-gen:
	go run github.com/ZNotify/server/scripts ent

swag-gen:
	swag init

schema-gen:
	go run -tags schema github.com/ZNotify/server/scripts schema

gen: ent-gen swag-gen schema-gen

lint:
	golangci-lint run

analyze:
	go tool nm "$(BIN_DIR)/$(BINARY)"