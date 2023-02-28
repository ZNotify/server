package global

import (
	"golang.org/x/oauth2"

	"notify-api/config"
	"notify-api/db/ent/generate"
)

var App = new(Application)

type Application struct {
	DB     *generate.Client
	Config *config.Configuration
	OAuth  *oauth2.Config
}
