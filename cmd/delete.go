package cmd

import (
	"fmt"

	"github.com/hossammourad/kewit/db"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a URL from your read-later list by its ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		idStr := args[0]
		var id int
		_, err := fmt.Sscanf(idStr, "%d", &id)
		if err != nil {
			fmt.Printf("Invalid ID: %v\n", err)
			return
		}
		if err := db.DeleteItemById(id); err != nil {
			fmt.Printf("Error deleting item: %v\n", err)
			return
		}
		fmt.Printf("Deleted")
	},
}
