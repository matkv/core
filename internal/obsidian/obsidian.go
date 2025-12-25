package obsidian

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/matkv/core/internal/config"
	"github.com/matkv/core/internal/types"
)

func LoadContentFilesOfType[T types.ContentType](contentType T) ([]types.ContentType, error) {
	vaultPath := config.C.Paths.ObsidianVault

	if vaultPath == "" {
		return nil, fmt.Errorf("Obsidian vault path is not set in the configuration")
	}

	info, err := os.Stat(vaultPath)
	if err != nil || !info.IsDir() {
		return nil, fmt.Errorf("Obsidian vault path does not exist or is not a directory: %s", vaultPath)
	}

	fmt.Println("Vault directory exists")
	fmt.Println("Loading content files from vault at:", contentType.ObsidianRootPath())

	var obsidianFiles []types.ObsidianFile
	obsidianFiles, err = ScanMarkdownFilesInPath(vaultPath, contentType.ObsidianRootPath(), contentType)
	if err != nil {
		return nil, err
	}

	var contentFiles []types.ContentType
	for _, file := range obsidianFiles {
		contentFile := contentType.CreateNew(file)
		contentFiles = append(contentFiles, contentFile)
	}

	return contentFiles, nil
}

func ScanMarkdownFilesInPath(vaultPath string, subPath string, contentType types.ContentType) ([]types.ObsidianFile, error) {
	fullPath := filepath.Join(vaultPath, subPath)

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
