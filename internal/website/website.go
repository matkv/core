package website

import (
	"fmt"

	"github.com/matkv/core/internal/obsidian"
)

func Scan() error {
	fmt.Println("Scanning the Obsidian vault...")

	contentFiles, err := obsidian.ScanVaultForContentFiles()
	if err != nil {
		return err
	}

	for _, file := range contentFiles {
		fmt.Println(file.Name)
	}
	return nil
}

func Sync() error {
	fmt.Println("Syncing Obsidian content to Hugo site...")
	// Placeholder for actual syncing logic
	return nil
}
