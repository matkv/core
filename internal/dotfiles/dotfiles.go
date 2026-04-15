package dotfiles

import (
	"os"
	"os/exec"

	"github.com/matkv/core/internal/config"
)

func Push(appConfig config.Application) error {
	return nil
}

func Pull(appConfig config.Application) error {
	return nil
}

func PullFromGithub(repoPath string) error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
