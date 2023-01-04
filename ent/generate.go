package main

//go:generate go run -mod=mod entgo.io/ent/cmd/ent describe ./schema

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./ent/schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureUpsert,
		},
		Target:  "./ent/generate",
		Package: "notify-api/ent/generate",
		Schema:  "notify-api/ent/schema",
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	} else {
		log.Println("ent codegen completed successfully")
	}
}
