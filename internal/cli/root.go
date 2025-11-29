package cli

import (
	"fmt"

	"github.com/matkv/core/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "core",
	Short:             "Core CLI tools & SvelteKit web app",
	PersistentPreRunE: setupConfig(), // set up configuration before any command runs
	RunE: func(cmd *cobra.Command, args []string) error { // show help if no subcommand is provided
		fmt.Printf("Obsidian vault: %s\n", config.C.Paths.ObsidianVault) // example usage of loaded config
		return cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func setupConfig() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if _, err := config.EnsureConfigFileExists(); err != nil {
			return err
		}
		return config.Load()
	}
}
