package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/ZNotify/server/app/bootstrap"
)

var (
	configPath string
)

func init() {
	checkCommand.Flags().StringVarP(&configPath, "config", "c", "data/config.yaml", "Load configuration from `FILE`, or use ENV to load from environment variable CONFIG.")
	rootCommand.AddCommand(checkCommand)
}

var checkCommand = &cobra.Command{
	Use:   "check",
	Short: "Check configuration file.",
	Run:   check,
}

func check(cmd *cobra.Command, args []string) {
	cfg := bootstrap.MergeConfig(bootstrap.Args{
		ConfigPath: configPath,
		Address:    "",
	})
	err, errS := cfg.Validate()
	if err != nil {
		log.Fatalf("Failed to validate configuration:\n %s", errS)
	}
	log.Printf("Configuration %s is valid.\n", configPath)
}
