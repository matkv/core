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

var pullFromGithubCmd = &cobra.Command{
	Use:   "pull-github",
	Short: "Pull updates from the GitHub repository to local repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		return dotfiles.PullFromGithub(config.C.Paths.Dotfiles.Repo)
	},
}

// generate parent command for each application, e.g. "core dotfiles neovim"
func generateDotfilesSubcommand(appName string, appConfig config.Application) *cobra.Command {
	appCmd := &cobra.Command{
		Use:   appName,
		Short: "Manage " + appName + " dotfiles",
	}

	appCmd.AddCommand(&cobra.Command{
		Use:   "push",
		Short: "Push local config changes to the dotfiles repository",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dotfiles.Push(appConfig)
		},
	})

	appCmd.AddCommand(&cobra.Command{
		Use:   "pull",
		Short: "Pull config changes from the dotfiles repository to local",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dotfiles.Pull(appConfig)
		},
	})

	return appCmd
}

func setupDotfilesSubCommands() {
	dotfilesCmd.AddCommand(pullFromGithubCmd)
	for appName, appConfig := range config.C.Paths.Dotfiles.Apps {
		dotfilesCmd.AddCommand(generateDotfilesSubcommand(appName, appConfig))
	}
}
