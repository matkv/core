package cmd

import (
	"errors"
	"fmt"
	"strings"

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

var obsidianCoverCmd = &cobra.Command{
	Use:   "cover",
	Short: "Download review covers and point the review to the local file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	}}

var obsidianCoverBooksCmd = &cobra.Command{
	Use:   "books <review file>",
	Short: "Download the cover of a book review",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return fixSingleCover(obsidian.BookReviewsDir, strings.Join(args, " "))
	}}

var coverMoviesAll bool

var obsidianCoverMoviesCmd = &cobra.Command{
	Use:   "movies [review file]",
	Short: "Download the cover of a movie/TV show review, or all remote covers with --all",
	RunE: func(cmd *cobra.Command, args []string) error {
		if coverMoviesAll == (len(args) > 0) {
			return fmt.Errorf("pass either a review file or --all")
		}
		if !coverMoviesAll {
			return fixSingleCover(obsidian.MovieReviewsDir, strings.Join(args, " "))
		}

		summary, err := obsidian.FixAllReviewCovers(obsidian.MovieReviewsDir)
		if err != nil {
			return err
		}
		fmt.Printf("\nDownloaded: %d, skipped: %d, failed: %d\n",
			len(summary.Downloaded), len(summary.Skipped), len(summary.Failed))
		for _, f := range summary.Failed {
			fmt.Printf("  FAILED %s: %v\n", f.File, f.Err)
		}
		return nil
	}}

// fixSingleCover treats an already-local cover as "nothing to do" instead of an error.
func fixSingleCover(reviewsDir, reviewFile string) error {
	err := obsidian.FixReviewCover(reviewsDir, reviewFile)
	if errors.Is(err, obsidian.ErrCoverAlreadyLocal) {
		fmt.Println("Cover is already a local file, nothing to do.")
		return nil
	}
	return err
}

func init() {
	obsidianCoverMoviesCmd.Flags().BoolVar(&coverMoviesAll, "all", false, "download all remote covers in the folder")
	obsidianCoverCmd.AddCommand(obsidianCoverBooksCmd, obsidianCoverMoviesCmd)
	obsidianCmd.AddCommand(obsidianCoverCmd)
}
