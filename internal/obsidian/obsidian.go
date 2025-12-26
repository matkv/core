package obsidian

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/matkv/core/internal/config"
	"github.com/matkv/core/internal/types"
)

func LoadContentFilesOfType[T types.Content](contentType T) ([]types.Content, error) {
	vaultPath := config.C.Paths.ObsidianVault

	if vaultPath == "" {
		return nil, fmt.Errorf("Obsidian vault path is not set in the configuration")
	}

	info, err := os.Stat(vaultPath)
	if err != nil || !info.IsDir() {
		return nil, fmt.Errorf("Obsidian vault path does not exist or is not a directory: %s", vaultPath)
	}

	fmt.Println("Vault directory exists")
	fmt.Println("Loading content files from vault at:", contentType.PathInObsidian())

	var obsidianFiles []types.ObsidianFile
	obsidianFiles, err = ScanMarkdownFilesInPath(vaultPath, contentType.PathInObsidian(), contentType)
	if err != nil {
		return nil, err
	}

	var contentFiles []types.Content
	for _, file := range obsidianFiles {
		contentFile := contentType.NewFromFile(file)
		contentFiles = append(contentFiles, contentFile)
	}

	return contentFiles, nil
}

func ScanMarkdownFilesInPath(vaultPath string, subPath string, contentType types.Content) ([]types.ObsidianFile, error) {
	fullPath := filepath.Join(vaultPath, subPath)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("path does not exist in Obsidian vault: %s", fullPath)
	}

	fmt.Printf("Scanning for type: %s at path: %s\n", contentType.TypeName(), fullPath)

	var files []types.ObsidianFile

	err := filepath.Walk(fullPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// temporary, contains old Obsidian setup
		if info.IsDir() && info.Name() == "!OLD VAULT SETUP" {
			return filepath.SkipDir
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			files = append(files, types.ObsidianFile{
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
