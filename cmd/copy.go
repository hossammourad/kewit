package cmd

import (
	"fmt"

	"github.com/hossammourad/kewit/db"
	"github.com/hossammourad/kewit/helpers"
	"github.com/spf13/cobra"
)

var CopyCmd = &cobra.Command{
	Use:   "copy [id]",
	Short: "Copy a saved URL to your clipboard by its ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := helpers.InputToInt(args[0])
		if err != nil {
			fmt.Printf("Invalid ID: %v\n", err)
			return
		}

		url, err := db.GetItemById(id)
		if err != nil {
			fmt.Printf("Error retrieving item: %v\n", err)
			return
		}

		err = helpers.CopyToClipboard(url)
		if err != nil {
			fmt.Printf("Error copying URL to clipboard: %v\n", err)
			return
		}
		fmt.Printf("'%s' has been copied to your clipboard\n", url)
	},
}
