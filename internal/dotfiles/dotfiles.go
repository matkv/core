package dotfiles

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/matkv/core/internal/config"
)

func Push(appConfig config.Application) error {
	repoPath := config.C.Paths.Dotfiles.Repo
	for _, source := range appConfig.Sources {
		src := filepath.Join(appConfig.Target, filepath.Base(source))
		dst := filepath.Join(repoPath, source)

		if err := copyFile(src, dst); err != nil {
			return fmt.Errorf("push %s: %w", source, err)
		}
		fmt.Printf("pushed %s → %s\n", src, dst)
	}
	return nil
}

func Pull(appConfig config.Application) error {
	repoPath := config.C.Paths.Dotfiles.Repo
	for _, source := range appConfig.Sources {
		src := filepath.Join(repoPath, source)
		dst := filepath.Join(appConfig.Target, filepath.Base(source))

		if err := copyFile(src, dst); err != nil {
			return fmt.Errorf("pull %s: %w", source, err)
		}
		fmt.Printf("pulled %s → %s\n", src, dst)
	}
	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

func PullFromGithub(repoPath string) error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Status(repoPath string) error {
	cmd := exec.Command("git", "status")
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func PushToGithub(repoPath string) error {
	cmd := exec.Command("git", "push")
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Diff(repoPath string) error {
	cmd := exec.Command("git", "diff")
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Commit(repoPath, message string) error {
	add := exec.Command("git", "add", "-A")
	add.Dir = repoPath
	add.Stdout = os.Stdout
	add.Stderr = os.Stderr
	if err := add.Run(); err != nil {
		return err
	}

	commit := exec.Command("git", "commit", "-m", message)
	commit.Dir = repoPath
	commit.Stdout = os.Stdout
	commit.Stderr = os.Stderr
	return commit.Run()
}
