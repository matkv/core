package ui

import (
	"embed"
	"io/fs"
)

//go:embed all:build
var buildFS embed.FS

// GetDistFS returns the filesystem for the static web assets
func GetDistFS() (fs.FS, error) {
	// We strip the "build" prefix so the server sees the files at the root
	return fs.Sub(buildFS, "build")
}
