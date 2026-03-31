package cmd

import (
	"github.com/matkv/core/internal/browser"
	"github.com/spf13/cobra"
)

var browserFileFlag bool

var browserCmd = &cobra.Command{
	Use:   "browser",
	Short: "Open multiple predefined URLs in the default web browser",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return browser.OpenPredefinedURLs()
		}

		if browserFileFlag {
			return browser.OpenURLsFromFiles(args)
		}

		return browser.OpenURLs(args)
	},
}

var twitchCmd = &cobra.Command{
	Use:   "twitch <username>",
	Short: "Open a Twitch stream and chat in popup windows",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return browser.OpenTwitchStream(args[0])
	},
}

func init() {
	browserCmd.Flags().BoolVarP(
		&browserFileFlag,
		"file",
		"f",
		false,
		"treat arguments as files containing URLs")
	browserCmd.AddCommand(twitchCmd)
}
