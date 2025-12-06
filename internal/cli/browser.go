package cli

import "github.com/spf13/cobra"

var browserCmd = &cobra.Command{
	Use:   "browser",
	Short: "Open multiple predefined URLs in the default web browser",
	RunE: func(cmd *cobra.Command, args []string) error {
		if args == nil || len(args) == 0 {
			return openPredefinedURLsInBrowser()
		}
		return openURLInBrowser(args)
	},
}

func openPredefinedURLsInBrowser() error {}

func openURLInBrowser(urls []string) error {
	// check if URLs are valid (basic check)

	var url = args[0]

}
