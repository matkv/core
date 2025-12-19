package cmd

import (
	"fmt"

	"github.com/matkv/core/internal/chooser"
	"github.com/spf13/cobra"
)

var chooseCmd = &cobra.Command{
	Use:   "choose",
	Short: "Choose a random item from the provided arguments",
	RunE: func(cmd *cobra.Command, args []string) error {
		chosenItem := chooser.Choose(args)

		if chosenItem == "" {
			fmt.Println("Usage: core choose <item1> <item2> ...")
			return nil
		}
		fmt.Printf("Result: %s\n", chosenItem)
		return nil
	}}
