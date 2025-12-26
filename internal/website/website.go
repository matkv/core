package website

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/matkv/core/internal/obsidian"
	"github.com/matkv/core/internal/types"
)

func Scan() error {
	fmt.Println("Scanning the Obsidian vault...")

	err := scanStandaloneContent()
	if err != nil {
		return err
	}

	err = scanListContent()
	if err != nil {
		return err
	}
	return nil
}

func scanStandaloneContent() error {
	fmt.Println("Scanning standalone content types...")

	for _, ct := range types.StandalonePages {

		file, err := obsidian.LoadStandaloneContent(ct)
		if err != nil {
			return err
		}

		fmt.Printf("Loaded standalone content: %s\n", file.PathInWebsite())
		switch v := file.(type) {
		case types.NowPage:
			fmt.Printf("Now Page Content: %s\n", v.Content)
		default:
			fmt.Printf("Unknown standalone content type: %v\n", v)
		}
	}
	return nil
}

func scanListContent() error {
	fmt.Println("Scanning list content types...")

	for _, ct := range types.ListPages {

		files, err := obsidian.LoadListContent(ct)
		if err != nil {
			return err
		}

		fmt.Printf("Loaded %d items of type %s\n", len(files), ct.TypeName())
		for _, f := range files {
			switch v := f.(type) {
			case types.BookReview:
				fmt.Printf("- Book Review: %s by %s\n", v.Title, v.Author)
			case types.MovieReview:
				fmt.Printf("- Movie Review: %s\n", v.Title)
			case types.Project:
				fmt.Printf("- Project: %s\n", v.Title)
			default:
				fmt.Printf("- Unknown list content type: %v\n", v)
			}
		}
	}
	return nil
}

// TODO rewrite
func Sync() error {
	return nil
	/*
		 	fmt.Println("Scanning the Obsidian vault...")

			// TODO remove duplicate check in LoadContentFilesOfType
			// check if obsidian vault path exists
			vaultPath := config.C.Paths.ObsidianVault

			if vaultPath == "" {
				return fmt.Errorf("Obsidian vault path is not set in the configuration")
			}

			info, err := os.Stat(vaultPath)
			if err != nil || !info.IsDir() {
				return fmt.Errorf("Obsidian vault path does not exist or is not a directory: %s", vaultPath)
			}

			websitePath := config.C.Paths.Website

			if websitePath == "" {
				return fmt.Errorf("Website path is not set in the configuration")
			}

			var contentPath string = filepath.Join(websitePath, "content")

			info, err = os.Stat(contentPath)
			if err != nil || !info.IsDir() {
				return fmt.Errorf("Content path does not exist or is not a directory: %s", websitePath)
			}

			fmt.Printf("Content path: %s\n", contentPath)

			contentTypes := []types.Content{
				types.BookReview{},
				// types.MovieReview{},
			}

			if err := clearContentDirAndRecreate(contentPath); err != nil {
				return fmt.Errorf("failed to clear content directory: %w", err)
			}

			setupFolderStructure(contentPath)
			generateStandaloneFiles()

			for _, ct := range contentTypes {
				files, err := obsidian.LoadContentFilesOfType(ct)
				if err != nil {
					return fmt.Errorf("failed to load files for %s: %w", ct.TypeName(), err)
				}

				fmt.Printf("Syncing %d %s(s)...\n", len(files), ct.TypeName())
				for _, f := range files {
					fmt.Printf("Syncing file: %v\n", f)
				}
			}

			fmt.Println("Do you want to sync the content?")

			return nil
	*/
}

func clearContentDirAndRecreate(contentPath string) error {
	fmt.Println("Clearing existing content in the website directory...")

	if err := os.RemoveAll(contentPath); err != nil {
		return err
	}

	return os.MkdirAll(contentPath, 0o755)
}

func setupFolderStructure(contentPath string) {
	dirs := []string{
		filepath.Join(contentPath, "garden"),
		filepath.Join(contentPath, "library", "books"),
		filepath.Join(contentPath, "log"),
		filepath.Join(contentPath, "projects"),
	}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			fmt.Printf("Failed to create directory %s: %v\n", dir, err)
		}
	}
}

func generateStandaloneFiles() {
	fmt.Println("Generating standalone files...")

}
