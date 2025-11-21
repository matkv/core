package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version can be set via build flags: -ldflags "-X github.com/matkv/core/internal/cli.Version=1.0.0"
var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of core",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Core version: %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
