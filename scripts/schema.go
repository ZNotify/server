//go:build schema

package main

import (
	"encoding/json"
	"log"
	"os"
	"reflect"

	"github.com/invopop/jsonschema"

	"github.com/ZNotify/server/app/config"
)

func schema() {
	log.Println("Generating JSON schema")
	r := jsonschema.Reflector{
		RequiredFromJSONSchemaTags: true,
	}
	err := r.AddGoComments("github.com/ZNotify/server/app/config", "./")
	if err != nil {
		panic(err)
	}
	err = r.AddGoComments("github.com/ZNotify/server/app/config/sender", "./")
	if err != nil {
		panic(err)
	}
	s := r.Reflect(&config.Configuration{})

	// hacks

	// hack to add minProperties to the sender
	// get sender config name
	senderConfigName := reflect.TypeOf(config.SenderConfiguration{}).Name()
	// get sender config schema
	senderConfigSchema := s.Definitions[senderConfigName]
	senderConfigSchema.MinProperties = 1

	// hack to add title
	s.Title = "ZNotify server configuration"
	s.Description = "The configuration schema of ZNotify server."

	content, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("./docs/schema.json", content, 0644)
	if err != nil {
		panic(err)
	}
	log.Println("JSON schema generated successfully")
}
