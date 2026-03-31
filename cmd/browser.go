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

func init() {
	browserCmd.Flags().BoolVarP(
		&browserFileFlag,
		"file",
		"f",
		false,
		"treat arguments as files containing URLs")
}
