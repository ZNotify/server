package config

func IsProd() bool {
	return Config.Server.Mode == ProdMode
}
