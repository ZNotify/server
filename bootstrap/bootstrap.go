package bootstrap

import "github.com/urfave/cli/v2"

func BootStrap(ctx *cli.Context) {
	initializeConfig(ctx)
	initializeLog()

	checkRequirements()

	initializeDatabase()
	initializeGlobalVar()
	initializePushManager()
	initializeOauth()
}
