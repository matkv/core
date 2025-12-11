package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/matkv/core/internal/config"
	"github.com/spf13/cobra"
)

var browserFileFlag bool

var browserCmd = &cobra.Command{
	Use:   "browser",
	Short: "Open multiple predefined URLs in the default web browser",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return openPredefinedURLsInBrowser()
		}

		if browserFileFlag {
			return openURLsFromFiles(args)
		}

		return openURLS(args)
	},
}

func init() {
	browserCmd.Flags().BoolVarP(
		&browserFileFlag,
		"file",
		"f",
		false,
		"treat arguments as files containing URLs")
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
		err := OpenURL(url)
		if err != nil {
			fmt.Printf("Failed to open URL %s: %v\n", url, err)
		} else {
			fmt.Printf("Opened URL: %s\n", url)
		}
	}
	return nil
}

func openURLsFromFiles(files []string) error {
	var allURLs []string

	for _, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %v", filePath, err)
		}

		urls := readLinksFromFile(file)
		file.Close()

		allURLs = append(allURLs, urls...)
	}

	if len(allURLs) == 0 {
		return fmt.Errorf("no URLs found in provided files")
	}

	return openURLS(allURLs)
}

func OpenURL(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default:
		return nil
	}

	return exec.Command(cmd, args...).Start()
}
