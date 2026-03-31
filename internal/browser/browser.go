package browser

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/matkv/core/internal/config"
)

func OpenPredefinedURLs() error {
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

	urls := readLinksFromFile(file)
	if len(urls) == 0 {
		return fmt.Errorf("no URLs found in links file")
	}

	// TODO add check for valid URLs
	return OpenURLs(urls)
}

func OpenURLs(urls []string) error {
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

func OpenURLsFromFiles(files []string) error {
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

	return OpenURLs(allURLs)
}

func OpenURL(url string) error {
	var cmd string
	var args []string

	if url == "" {
		return fmt.Errorf("empty URL provided")
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

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

func OpenTwitchStream(username string) error {
	if username == "" {
		return fmt.Errorf("username cannot be empty")
	}

	streamURL := fmt.Sprintf(
		"https://player.twitch.tv/?channel=%s&enableExtensions=true&muted=false&parent=twitch.tv&player=popout",
		url.QueryEscape(username),
	)
	chatURL := fmt.Sprintf(
		"https://www.twitch.tv/popout/%s/chat?popout=",
		url.PathEscape(username),
	)

	htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head><title>Opening Twitch...</title></head>
<body>
<script>
window.open('%s', 'twitch_stream', 'width=1280,height=720,toolbar=no,location=no,menubar=no,status=no,scrollbars=no,resizable=yes');
window.open('%s', 'twitch_chat', 'width=360,height=720,toolbar=no,location=no,menubar=no,status=no,scrollbars=no,resizable=yes');
window.close();
</script>
</body>
</html>`, streamURL, chatURL)

	launcherPath := filepath.Join(os.TempDir(), "core-twitch-launcher.html")
	if err := os.WriteFile(launcherPath, []byte(htmlContent), 0600); err != nil {
		return fmt.Errorf("failed to write twitch launcher: %v", err)
	}

	return openLocalFile(launcherPath)
}

func openLocalFile(path string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "linux":
		cmd = "xdg-open"
		args = []string{path}
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", path}
	case "darwin":
		cmd = "open"
		args = []string{path}
	default:
		return nil
	}

	return exec.Command(cmd, args...).Start()
}
