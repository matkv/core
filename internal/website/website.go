package website

import (
	"fmt"

	"github.com/matkv/core/internal/obsidian"
	"github.com/matkv/core/internal/types"
)

func Scan() error {
	fmt.Println("Scanning the Obsidian vault...")

	contentTypes := []types.ContentType{
		types.BookReview{},
		// types.MovieReview{},
	}

	for _, ct := range contentTypes {
		files, err := obsidian.LoadContentFilesOfType(ct)
		if err != nil {
			return err
		}

		fmt.Printf("Found %d %s(s):\n", len(files), ct.TypeName())
		for _, f := range files {
			switch v := f.(type) {
			case types.BookReview:
				fmt.Printf("- %s by %s\n", v.Title, v.Author)
			case types.MovieReview:
				fmt.Printf("- %s\n", v.TypeName())
			// Add more cases for additional types
			default:
				fmt.Printf("- %v\n", v)
			}
		}
	}

	return nil
}

func Sync() error {
	fmt.Println("Syncing Obsidian content to Hugo site...")
	// Placeholder for actual syncing logic
	return nil
}
