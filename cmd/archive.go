package cmd

import (
	"fmt"

	"github.com/hossammourad/kewit/db"
	"github.com/hossammourad/kewit/helpers"
	"github.com/spf13/cobra"
)

var ArchiveCmd = &cobra.Command{
	Use:   "archive [id]",
	Short: "Archive a URL by its ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := helpers.InputToInt(args[0])
		if err != nil {
			fmt.Printf("Invalid ID: %v\n", err)
			return
		}

		if err := db.ArchiveItemById(id); err != nil {
			fmt.Printf("Error archiving item: %v\n", err)
			return
		}
		fmt.Printf("Archived")
	},
}
