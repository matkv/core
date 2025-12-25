package website

import (
	"fmt"

	"github.com/matkv/core/internal/obsidian"
	"github.com/matkv/core/internal/types"
)

func Scan() error {
	fmt.Println("Scanning the Obsidian vault...")

	bookReviews, err := obsidian.LoadContentFilesOfType(types.BookReview{})
	if err != nil {
		return err
	}

	fmt.Printf("Found %d book reviews:\n", len(bookReviews))
	for _, br := range bookReviews {
		b := br.(types.BookReview)
		fmt.Printf("- %s by %s\n", b.Title, b.Author)
	}

	return nil
}

func Sync() error {
	fmt.Println("Syncing Obsidian content to Hugo site...")
	// Placeholder for actual syncing logic
	return nil
}
