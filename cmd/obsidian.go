package cmd

import (
	"fmt"

	"github.com/matkv/core/internal/obsidian"
	"github.com/spf13/cobra"
)

var obsidianCmd = &cobra.Command{
	Use:   "obsidian",
	Short: "Manage my personal Obsidian vault",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	}}

var obsidianFixCoverCmd = &cobra.Command{
	Use:   "cover",
	Short: "Download book cover from cover property in current book review",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("missing review file argument")
		}
		reviewFile := args[0]
		fmt.Println("Fixing book review covers...")
		return obsidian.FixBookReviewCover(reviewFile)
	}}

func init() {
	obsidianCmd.AddCommand(obsidianFixCoverCmd)
}
