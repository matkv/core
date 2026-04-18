package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/matkv/core/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Open the config directory in Explorer",
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath, err := config.ConfigPath()
		if err != nil {
			return err
		}
		configDir := filepath.Dir(configPath)
		fmt.Println(configDir)
		return exec.Command("explorer.exe", configDir).Start()
	},
}
