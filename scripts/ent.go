package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func ent() {
	err := entc.Generate("./db/ent/schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureUpsert,
		},
		Target:  "./db/ent/generate",
		Package: "notify-api/db/ent/generate",
		Schema:  "notify-api/db/ent/schema",
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	} else {
		log.Println("ent codegen completed successfully")
	}
}
