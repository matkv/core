package website

import (
	"fmt"

	"github.com/matkv/core/internal/obsidian"
)

func Hello() error {
	fmt.Println("Hello from the website package")

	obsidian.Hello()
	return nil
}
