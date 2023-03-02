//go:build schema

package config

import (
	"github.com/invopop/jsonschema"
)

func (Mode) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "string",
		Title:       "Server running mode",
		Description: "The mode the server is running in. Can be one of: test, development, production.",
		Enum: []interface{}{
			TestMode,
			DevMode,
			ProdMode,
		},
	}
}

func (Database) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "string",
		Title:       "Database type",
		Description: "The type of database to use. Can be one of: sqlite, mysql, pgsql.",
		Enum: []interface{}{
			Sqlite,
			Mysql,
			Pgsql,
		},
	}
}
