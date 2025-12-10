package cli

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/matkv/core/internal/config"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var browserCmd = &cobra.Command{
	Use:   "browser",
	Short: "Open multiple predefined URLs in the default web browser",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return openPredefinedURLsInBrowser()
		}
		return openURLInBrowser(args)
	},
}

func openURLInBrowser(urls []string) error {
	return openURLS(urls)
}

func openPredefinedURLsInBrowser() error {
	configPath, err := config.ConfigPath()
	if err != nil {
		return err
	}

	configDir := filepath.Dir(configPath)
	linksFile := filepath.Join(configDir, "links.txt")

	if _, err := os.Stat(linksFile); os.IsNotExist(err) {
		return fmt.Errorf("links file does not exist: %s", linksFile)
	}

	file, err := os.Open(linksFile)
	if err != nil {
		return fmt.Errorf("failed to open links file: %v", err)
	}
	defer file.Close()

	var urls []string
	urls = readLinksFromFile(file)
	if len(urls) == 0 {
		return fmt.Errorf("no URLs found in links file")
	}

	// TODO add check for valid URLs

	openURLS(urls)
	return nil
}

func readLinksFromFile(file *os.File) []string {
	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			urls = append(urls, line)
		}
	}

	return urls
}

func openURLS(urls []string) error {
	for _, url := range urls {
		err := browser.OpenURL(url)
		if err != nil {
			fmt.Printf("Failed to open URL %s: %v\n", url, err)
		} else {
			fmt.Printf("Opened URL: %s\n", url)
		}
	}
	return nil
}
