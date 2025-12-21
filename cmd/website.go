package cmd

import (
	"fmt"

	"github.com/matkv/core/internal/website"
	"github.com/spf13/cobra"
)

var websiteCmd = &cobra.Command{
	Use:   "website",
	Short: "Manage my personal hugo website",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	}}

var websiteScanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the Obsidian vault & print stats",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello from the scan subcommand")
		return website.Scan()
	}}

var websiteSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync and copy Obsidian content to the hugo site",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello from the sync subcommand")
		return website.Sync()
	}}

func init() {
	websiteCmd.AddCommand(websiteScanCmd)
	websiteCmd.AddCommand(websiteSyncCmd)
}
