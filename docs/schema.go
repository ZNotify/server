package docs

import (
	_ "embed"
)

//go:embed "schema.json"
var SchemaJson string
