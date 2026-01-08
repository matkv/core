package obsidian

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/matkv/core/internal/config"
	"github.com/matkv/core/internal/types"
)

func LoadStandaloneContent(pageType types.Content) (types.Content, error) {
	vaultPath := config.C.Paths.ObsidianVault

	if err := ensureVaultPathExists(vaultPath); err != nil {
		return nil, err
	}

	fmt.Println("Loading standalone content from vault at:", pageType.PathInObsidian())

	obsidianFile, err := ScanSingleMarkdownFile(vaultPath, pageType.PathInObsidian(), pageType)
	if err != nil {
		return nil, err
	}

	contentFile := pageType.NewFromFile(obsidianFile)
	return contentFile, nil
}

func ScanSingleMarkdownFile(vaultPath string, subPath string, pageType types.Content) (types.ObsidianFile, error) {
	fullPath := filepath.Join(vaultPath, subPath)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return types.ObsidianFile{}, fmt.Errorf("file does not exist in Obsidian vault: %s", fullPath)
	}

	fmt.Printf("Scanning for single page type: %s at path: %s\n", pageType.TypeName(), fullPath)

	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		return types.ObsidianFile{}, err
	}

	if fileInfo.IsDir() {
		return types.ObsidianFile{}, fmt.Errorf("expected a file but found a directory: %s", fullPath)
	}

	return types.ObsidianFile{
		Path: fullPath,
		Name: filepath.Base(fullPath),
	}, nil
}

func LoadListContent(listType types.Content) ([]types.Content, error) {
	vaultPath := config.C.Paths.ObsidianVault

	if err := ensureVaultPathExists(vaultPath); err != nil {
		return nil, err
	}

	var obsidianFiles []types.ObsidianFile

	// Scan for index file if IndexPathInObsidian is not empty
	indexPath := listType.IndexPathInObsidian()
	if indexPath != "" {
		indexFile, err := ScanSingleMarkdownFile(vaultPath, indexPath, listType)
		if err != nil {
			return nil, fmt.Errorf("failed to scan index file for %s: %w", listType.TypeName(), err)
		}
		obsidianFiles = append(obsidianFiles, indexFile)
	}

	// Scan for all content files in the list directory
	contentFiles, err := ScanMultipleMarkdownFiles(vaultPath, listType.PathInObsidian(), listType)
	if err != nil {
		return nil, fmt.Errorf("failed to scan content files for %s: %w", listType.TypeName(), err)
	}
	obsidianFiles = append(obsidianFiles, contentFiles...)

	var contentList []types.Content
	for _, of := range obsidianFiles {
		contentList = append(contentList, listType.NewFromFile(of))
	}
	return contentList, nil
}

func ScanMultipleMarkdownFiles(vaultPath string, subPath string, listType types.Content) ([]types.ObsidianFile, error) {
	fullPath := filepath.Join(vaultPath, subPath)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("path does not exist in Obsidian vault: %s", fullPath)
	}

	fmt.Printf("Scanning for list type: %s at path: %s\n", listType.TypeName(), fullPath)

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

func ensureVaultPathExists(vaultPath string) error {
	if vaultPath == "" {
		return fmt.Errorf("Obsidian vault path is not set in the configuration")
	}

	info, err := os.Stat(vaultPath)
	if err != nil || !info.IsDir() {
		return fmt.Errorf("Obsidian vault path does not exist or is not a directory: %s", vaultPath)
	}
	return nil
}

func FixBookReviewCover(reviewFile string) error {
	fmt.Printf("Downloading book cover for review file: %s\n", reviewFile)
	// TODO download cover from cover property in the markdown file
	// store it in the appropriate location in the Obsidian vault
	// overwrite the cover property with the local path
	return nil
}
