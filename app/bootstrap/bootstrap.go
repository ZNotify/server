package bootstrap

type Args struct {
	Config  string
	Address string
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
