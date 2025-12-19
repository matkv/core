package cmd

import (
	"fmt"

	"github.com/matkv/core/internal/picker"
	"github.com/spf13/cobra"
)

var pickCmd = &cobra.Command{
	Use:   "pick",
	Short: "Pick a random item from the provided arguments",
	RunE: func(cmd *cobra.Command, args []string) error {
		pickedItem := picker.Pick(args)

		if pickedItem == "" {
			fmt.Println("Usage: core pick <item1> <item2> ...")
			return nil
		}
		fmt.Printf("Result: %s\n", pickedItem)
		return nil
	}}
