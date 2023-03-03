package bootstrap

type Args struct {
	ConfigPath string
	Address    string
}

func BootStrap(args Args) {
	initializeConfig(args)
	initializeLog()

	checkRequirements()

	initializeDatabase()
	initializeGlobalVar()
	initializePushManager()
	initializeOauth()
}
