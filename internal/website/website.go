package website

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/matkv/core/internal/config"
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

	contentTypes := []types.ContentType{
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
