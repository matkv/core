package cli

import (
	"fmt"

	"github.com/matkv/core/internal/config"
	"github.com/spf13/cobra"
)

type cliCommand struct {
	cmd            *cobra.Command
	allowedDevices []config.Device
}

var (
	commands           []cliCommand
	commandsRegistered bool
)

var rootCmd = &cobra.Command{
	Use:   "core",
	Short: "Core CLI tools & SvelteKit web app",
	RunE: func(cmd *cobra.Command, args []string) error { // show help if no subcommand is provided
		// example usage of loaded config
		fmt.Printf("Obsidian vault: %s\n", config.C.Paths.ObsidianVault)
		fmt.Printf("Device type: %s\n", config.C.Device)
		return cmd.Help()
	},
}

func setupConfig() error {
	if _, err := config.EnsureConfigFileExists(); err != nil {
		return err
	}
	if err := config.Load(); err != nil {
		return err
	}
	if !config.C.Device.IsValidDevice() {
		return fmt.Errorf("invalid device %q in config", config.C.Device)
	}
	return nil
}

func setupCommands() {
	if commandsRegistered {
		return
	}
	commandsRegistered = true
	currentDevice := config.C.Device
	for _, c := range commands {
		if isDeviceAllowed(currentDevice, c.allowedDevices) {
			rootCmd.AddCommand(c.cmd)
		}
	}
}

func init() {
	addCommand(versionCmd, config.Desktop, config.Laptop, config.WSL)
	addCommand(randomCmd, config.Desktop, config.Laptop)
	addCommand(serveCmd, config.Desktop, config.Laptop, config.WSL)
}

func addCommand(command *cobra.Command, devices ...config.Device) {
	commands = append(commands, cliCommand{cmd: command, allowedDevices: devices})
}

func isDeviceAllowed(currentDevice config.Device, allowedDevices []config.Device) bool {
	for _, d := range allowedDevices {
		if currentDevice == d {
			return true
		}
	}
	return false
}

func Execute() {
	if err := setupConfig(); err != nil {
		panic(err)
	}

	setupCommands()

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
