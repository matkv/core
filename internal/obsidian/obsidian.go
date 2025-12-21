package obsidian

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/matkv/core/internal/config"
)

type ObsidianFile struct {
	Path        string
	Name        string
	Description string
}

func ScanVaultForContentFiles() ([]ObsidianFile, error) {
	vaultPath := config.C.Paths.ObsidianVault
	fmt.Println("Scanning vault at:", vaultPath)

	if vaultPath == "" {
		return nil, fmt.Errorf("Obsidian vault path is not set in the configuration")
	}

	info, err := os.Stat(vaultPath)
	if err != nil || !info.IsDir() {
		return nil, fmt.Errorf("Obsidian vault path does not exist or is not a directory: %s", vaultPath)
	}

	fmt.Println("Vault directory exists")

	var contentFiles []ObsidianFile
	contentFiles, err = ScanMarkdownFiles(vaultPath)
	if err != nil {
		return nil, err
	}

	return contentFiles, nil
}

func ScanMarkdownFiles(vaultPath string) ([]ObsidianFile, error) {

	fmt.Println("Scanning for markdown files in:", vaultPath)
	var files []ObsidianFile

	err := filepath.Walk(vaultPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// temporary, contains old Obsidian setup
		if info.IsDir() && info.Name() == "!OLD VAULT SETUP" {
			return filepath.SkipDir
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			files = append(files, ObsidianFile{
				Path: path,
				Name: info.Name(),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
