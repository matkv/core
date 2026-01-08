package obsidian

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/matkv/core/internal/config"
	"github.com/matkv/core/internal/types"
	"go.yaml.in/yaml/v3"
	"golang.org/x/text/unicode/norm"
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

	vaultDir := config.C.Paths.ObsidianVault
	if err := ensureVaultPathExists(vaultDir); err != nil {
		return err
	}

	bookReviewsDir := filepath.Join(vaultDir, "Database", "Index", "Books")
	reviewFilePath := filepath.Join(bookReviewsDir, reviewFile)
	if err := ensureReviewFileExists(reviewFilePath); err != nil {
		return err
	}
	fmt.Printf("Found review file at path: %s\n", reviewFilePath)

	coversDir, err := ensureCoversDirectoryExists(bookReviewsDir)
	if err != nil {
		return err
	}

	var coverURL string
	coverURL = getCoverURLFromReviewFile(reviewFilePath)
	if coverURL == "" {
		return fmt.Errorf("no cover URL found in review file: %s", reviewFilePath)
	}

	fmt.Printf("Cover URL found: %s\n", coverURL)

	var bookSlug = slugify(reviewFile)
	fmt.Printf("Book slug generated: %s\n", bookSlug)
	downloadBookCover(coverURL, coversDir, bookSlug)

	return nil
}

func downloadBookCover(coverURL, coversDir string, bookSlug string) {
	fmt.Printf("Downloading cover from URL: %s to directory: %s\n", coverURL, coversDir)

	// Determine the file extension from the URL
	var fileExt string
	if strings.HasSuffix(coverURL, ".jpg") || strings.HasSuffix(coverURL, ".jpeg") {
		fileExt = ".jpg"
	} else if strings.HasSuffix(coverURL, ".png") {
		fileExt = ".png"
	} else {
		fileExt = ".jpg" // default to jpg
	}

	coverFilePath := filepath.Join(coversDir, bookSlug+fileExt)
	fmt.Printf("Saving cover to file: %s\n", coverFilePath)

	// Download the cover image
	err := downloadFile(coverURL, coverFilePath)
	if err != nil {
		fmt.Printf("Failed to download cover image: %v\n", err)
	} else {
		fmt.Printf("Successfully downloaded cover image to: %s\n", coverFilePath)
	}
}

func downloadFile(url string, filePath string) error {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func ensureReviewFileExists(reviewFilePath string) error {
	if _, err := os.Stat(reviewFilePath); os.IsNotExist(err) {
		return fmt.Errorf("review file does not exist: %s", reviewFilePath)
	}
	return nil
}

func ensureCoversDirectoryExists(bookReviewsDir string) (string, error) {
	coversDir := filepath.Join(bookReviewsDir, "Covers")
	if _, err := os.Stat(coversDir); os.IsNotExist(err) {
		fmt.Printf("Covers directory does not exist, creating: %s\n", coversDir)
		if err := os.Mkdir(coversDir, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create covers directory: %w", err)
		}
	}
	return coversDir, nil
}

func getCoverURLFromReviewFile(reviewFilePath string) string {
	data, err := os.ReadFile(reviewFilePath)
	if err != nil {
		return ""
	}

	frontmatter, ok := extractFrontmatter(string(data))
	if !ok {
		return ""
	}

	var fm map[string]any
	if err := yaml.Unmarshal([]byte(frontmatter), &fm); err != nil {
		return ""
	}

	if v, ok := fm["cover"]; ok {
		if s, ok := v.(string); ok {
			return strings.TrimSpace(s)
		}
	}
	return ""
}

// extractFrontmatter extracts and returns the YAML frontmatter (without the --- delimiters).
// Returns the frontmatter string and true if frontmatter was found.
func extractFrontmatter(content string) (string, bool) {
	content = strings.TrimLeft(content, "\ufeff") // handle BOM if present
	content = strings.TrimSpace(content)
	if !strings.HasPrefix(content, "---") {
		return "", false
	}
	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return "", false
	}
	fm := strings.TrimSpace(parts[1])
	if fm == "" {
		return "", false
	}
	return fm, true
}

func slugify(s string) string {
	// 1. Normalize (NFD splits accents)
	t := norm.NFD.String(s)

	// 2. Keep only ASCII letters/digits, convert spaces to hyphens
	var b strings.Builder
	b.Grow(len(t))

	prevHyphen := false

	for _, r := range t {
		switch {
		case unicode.IsLetter(r) && r <= unicode.MaxASCII:
			b.WriteRune(unicode.ToLower(r))
			prevHyphen = false

		case unicode.IsDigit(r):
			b.WriteRune(r)
			prevHyphen = false

		case r == ' ' || r == '-' || r == '_':
			if !prevHyphen {
				b.WriteRune('-')
				prevHyphen = true
			}

		default:
			// drop everything else
		}
	}

	slug := strings.Trim(b.String(), "-")
	slug = strings.TrimSuffix(slug, "md")

	return slug
}
