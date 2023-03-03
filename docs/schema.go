package docs

import (
	_ "embed"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

//go:embed "schema.json"
var schemaJson string

var Schema = jsonschema.MustCompileString("schema.json", schemaJson)
