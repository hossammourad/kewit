package cmd

import (
	"fmt"

	"github.com/hossammourad/kewit/db"
	"github.com/hossammourad/kewit/helpers"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add [url]",
	Short: "Add a URL to your read-later list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		if !helpers.IsValidURL(url) {
			fmt.Printf("Invalid URL: %s\n", url)
			return
		}
		if err := db.AddItem(url); err != nil {
			fmt.Printf("Error adding URL: %v\n", err)
			return
		}
		fmt.Printf("Added URL: %s\n", url)
	},
}
