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

var websiteBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the Hugo website",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello from the subcommand")
		return website.Hello()
	}}

func init() {
	websiteCmd.AddCommand(websiteBuildCmd)
}
