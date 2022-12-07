package config

func IsProd() bool {
	return Config.Server.Mode == ProdMode
}

func IsDev() bool {
	return Config.Server.Mode == DevMode
}

func IsTest() bool {
	return Config.Server.Mode == TestMode
}
