package cmd

import (
	"github.com/matkv/core/internal/config"
	"github.com/matkv/core/internal/dotfiles"
	"github.com/spf13/cobra"
)

var dotfilesCmd = &cobra.Command{
	Use:   "dotfiles",
	Short: "Sync dotfiles to/from a git repository on Windows",
}

// generate parent command for each application, e.g. "core dotfiles neovim"
func generateDotfilesSubcommand(appName string, appConfig config.Application) *cobra.Command {
	return &cobra.Command{
		Use:   appName,
		Short: "Manage " + appName + " dotfiles",
	}
}

var pullFromGithubCmd = &cobra.Command{
	Use:   "pull-github",
	Short: "Pull updates from the GitHub repository to local repo)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return cmd.Help()
		}

		// TODO adapt config - store the path to the dotfiles repo
		// and then only the relative path of each app IN the dotfiles repo
		// example just (.config/nvim) for the neovim config

		return nil
	}}

var pushCmd = &cobra.Command{
	Use:   "push <app>",
	Short: "Push local config changes to the dotfiles repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return cmd.Help()
		}

		// get Application based on the app name argument
		appName := args[0]
		application, exists := config.C.Paths.Dotfiles[appName]
		if !exists { // TODO show error + list of valid app names
			return cmd.Help()
		}

		return dotfiles.Push(application)
	},
}

var pullCmd = &cobra.Command{
	Use:   "pull <app>",
	Short: "Pull config changes from the dotfiles repository to local",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return cmd.Help()
		}
		// get Application based on the app name argument
		appName := args[0]
		application, exists := config.C.Paths.Dotfiles[appName]
		if !exists {
			return cmd.Help()
		}
		return dotfiles.Pull(application)
	},
}

func setupDotfilesSubCommands() {
	for appName, appConfig := range config.C.Paths.Dotfiles {
		appCmd := generateDotfilesSubcommand(appName, appConfig)
		appCmd.AddCommand(pushCmd)
		appCmd.AddCommand(pullCmd)
		dotfilesCmd.AddCommand(appCmd)
	}
}
