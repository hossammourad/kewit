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
		item, err := db.GetItemById(id)
		if err != nil {
			fmt.Printf("Error retrieving item: %v\n", err)
			return
		}
		if item == "" {
			fmt.Printf("No item found with ID %d\n", id)
			return
		}
		var confirmation string
		fmt.Printf("Are you sure you want to delete the URL '%s' (ID: %d)? (y/N): ", item, id)
		fmt.Scanln(&confirmation)
		if confirmation != "y" && confirmation != "Y" {
			fmt.Println("Deletion cancelled.")
			return
		}
		if err := db.DeleteItemById(id); err != nil {
			fmt.Printf("Error deleting item: %v\n", err)
			return
		}
		fmt.Printf("Deleted item with URL: %s\n", item)
	},
}
