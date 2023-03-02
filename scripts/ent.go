package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func ent() {
	log.Println("running ent codegen")
	err := entc.Generate("./app/db/ent/schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureUpsert,
		},
		Target:  "./app/db/ent/generate",
		Package: "notify-api/app/db/ent/generate",
		Schema:  "notify-api/app/db/ent/schema",
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	} else {
		log.Println("ent codegen completed successfully")
	}
}
