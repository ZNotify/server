package global

import (
	"golang.org/x/oauth2"

	"github.com/ZNotify/server/app/config"
	"github.com/ZNotify/server/app/db/ent/generate"
)

var App = new(Application)

type Application struct {
	DB     *generate.Client
	Config *config.Configuration
	OAuth  *oauth2.Config
}
