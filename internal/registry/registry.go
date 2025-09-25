package registry

import (
	"github.com/matkv/utils/config"
	"github.com/spf13/cobra"
)

var AllCommands = []*cobra.Command{}

// Annotation keys
const (
	AnnWindowsOnly = "IsWindowsOnly"
	AnnLinuxOnly   = "IsLinuxOnly"
	AnnArchived    = "IsArchived"
)

// RegisterCommand adds a command to the AllCommands slice
func RegisterCommand(cmd *cobra.Command) {
	AllCommands = append(AllCommands, cmd)
}

// GetAllCommands returns all registered commands
func GetAllCommands() []*cobra.Command {
	return AllCommands
}

// ShouldRegister determines if a command should be registered based on annotations and current config type
func ShouldRegister(cmd *cobra.Command) bool {
	if cmd == nil {
		return false
	}

	if v, ok := cmd.Annotations[AnnArchived]; ok && v == "true" {
		return false
	}

	if v, ok := cmd.Annotations[AnnWindowsOnly]; ok && v == "true" {
		return config.IsWindows()
	}

	if v, ok := cmd.Annotations[AnnLinuxOnly]; ok && v == "true" {
		return config.IsLinux()
	}

	return true
}
