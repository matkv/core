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
	Use:   "pull",
	Short: "Pull updates from the GitHub repository to local repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		return dotfiles.PullFromGithub(config.C.Paths.Dotfiles.Repo)
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show git status of the dotfiles repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		return dotfiles.Status(config.C.Paths.Dotfiles.Repo)
	},
}

var pushToGithubCmd = &cobra.Command{
	Use:   "push",
	Short: "Push local dotfiles repository commits to GitHub",
	RunE: func(cmd *cobra.Command, args []string) error {
		return dotfiles.PushToGithub(config.C.Paths.Dotfiles.Repo)
	},
}

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show git diff of the dotfiles repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		return dotfiles.Diff(config.C.Paths.Dotfiles.Repo)
	},
}

var commitCmd = &cobra.Command{
	Use:   "commit <message>",
	Short: "Stage all changes and commit with the given message",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return dotfiles.Commit(config.C.Paths.Dotfiles.Repo, args[0])
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
	dotfilesCmd.AddCommand(statusCmd)
	dotfilesCmd.AddCommand(pushToGithubCmd)
	dotfilesCmd.AddCommand(diffCmd)
	dotfilesCmd.AddCommand(commitCmd)
	for appName, appConfig := range config.C.Paths.Dotfiles.Apps {
		dotfilesCmd.AddCommand(generateDotfilesSubcommand(appName, appConfig))
	}
}
