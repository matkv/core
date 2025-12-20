package obsidian

import (
	"fmt"

	"github.com/matkv/core/internal/config"
)

func Hello() {
	println("Hello from the Obsidian package")
	fmt.Printf("Obsidian vault: %s\n", config.C.Paths.ObsidianVault)
}
