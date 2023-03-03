package global

import (
	"github.com/ZNotify/server/app/config"
)

func IsProd() bool {
	return App.Config.Server.Mode == config.ProdMode
}

func IsDev() bool {
	return App.Config.Server.Mode == config.DevMode
}

func IsTest() bool {
	return App.Config.Server.Mode == config.TestMode
}
