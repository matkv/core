package cli

import (
	"github.com/matkv/core/internal/random"
	"github.com/spf13/cobra"
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Print a random number (example command)",
	Run: func(cmd *cobra.Command, args []string) {
		value := random.Int(1000)
		cmd.Printf("%d\n", value)
	},
}
