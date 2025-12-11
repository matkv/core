package cmd

import (
	"fmt"

	"github.com/matkv/core/internal/config"
	"github.com/spf13/cobra"
)

// Version can be set via build flags: -ldflags "-X github.com/matkv/core/cmd.Version=1.0.0"
var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of core",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello from the version command!")
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Core version: %s\n", Version)

		// example usage of loaded config
		fmt.Printf("Obsidian vault FROM VERSION COMMAND: %s\n", config.C.Paths.ObsidianVault)
		fmt.Printf("Device type: %s\n", config.C.Device)
	},
}
