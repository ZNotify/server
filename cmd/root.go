package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCommand.Flags().StringP("config", "c", "data/config.yaml", "Load configuration from `FILE`, or use ENV to load from environment variable CONFIG.")
	rootCommand.Flags().StringP("address", "a", "0.0.0.0:14444", "Set listen address to `ADDRESS`.")
}

var rootCommand = &cobra.Command{
	Use:   "server",
	Short: "This is ZNotify api server.",
	Run:   Run,
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
