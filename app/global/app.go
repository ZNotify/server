package global

import (
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/ZNotify/server/app/api/router"
	"github.com/ZNotify/server/app/bootstrap"
	"github.com/ZNotify/server/app/config"
	"github.com/ZNotify/server/app/db/ent/generate"
)

var App = new(Application)

type Application struct {
	DB     *generate.Client
	Config *config.Configuration
	OAuth  *oauth2.Config
}

func (a *Application) BootUp(config string, address string) {
	bootstrap.BootStrap(bootstrap.Args{
		Config:  config,
		Address: address,
	})
	r := router.NewRouter()
	err := r.Run(a.Config.Server.Address)
	if err != nil {
		zap.S().Fatalf("Failed to start server: %+v", err)
	}
	defer a.Clean()
}

func (a *Application) Clean() {
	err := a.DB.Close()
	if err != nil {
		zap.S().Errorf("Failed to close database: %+v", err)
	}
	err = zap.L().Sync()
	if err != nil {
		zap.S().Errorf("Failed to sync logger: %+v", err)
	}
}
