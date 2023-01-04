# Build the project
build:
	swag init
	go build -o bin/server notify-api

dependencies:
	go mod vendor
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/cosmtrek/air@latest
	go install entgo.io/ent/cmd/ent@latest
	go install github.com/kisielk/godepgraph@latest

# Run the project with hot reload
dev:
	air

fmt:
	swag fmt
	go fmt ./...

desc:
	go generate ent/generate.go

gen:
	go run ent/generate.go

analyze:
	godepgraph -maxlevel 16 -s -novendor -p github.com,gorm.io,modernc.com,google.golang.org,golang.org,gopkg.in,go.uber.org,go.opencensus.io,firebase.google.com,cloud.google.com notify-api | dot -Tpng -o godepgraph.png