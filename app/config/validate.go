package config

import (
	"fmt"

	"github.com/invopop/yaml"

	"github.com/ZNotify/server/docs"
)

func (c *Configuration) Validate() (error, string) {
	var raw any
	err := yaml.Unmarshal(c.Raw, &raw)
	if err != nil {
		return err, err.Error()
	}

	err = docs.Schema.Validate(raw)
	errString := fmt.Sprintf("%#v", err)

	return err, errString
}
