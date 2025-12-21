package website

import (
	"fmt"

	"github.com/matkv/core/internal/obsidian"
)

func Scan() error {
	fmt.Println("Scanning the Obsidian vault...")
	// Placeholder for actual scanning logic
	obsidian.Hello()
	return nil
}

func Sync() error {
	fmt.Println("Syncing Obsidian content to Hugo site...")
	// Placeholder for actual syncing logic
	obsidian.Hello()
	return nil
}
