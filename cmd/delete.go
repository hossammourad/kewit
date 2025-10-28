package cmd

import (
	"fmt"

	"github.com/hossammourad/kewit/db"
	"github.com/hossammourad/kewit/helpers"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a URL from your read-later list by its ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Convert the input argument to an integer ID
		id, err := helpers.InputToInt(args[0])
		if err != nil {
			fmt.Printf("Invalid ID: %v\n", err)
			return
		}

		// Retrieve the item before deletion for confirmation
		item, err := db.GetItemById(id)
		if err != nil {
			fmt.Printf("Error retrieving item: %v\n", err)
			return
		}

		// Check if the item exists
		if item == "" {
			fmt.Printf("No item found with ID %d\n", id)
			return
		}

		// Confirm deletion with the user
		var confirmation string
		fmt.Printf("Are you sure you want to delete the URL '%s' (ID: %d)? (y/N): ", item, id)
		fmt.Scanln(&confirmation)
		if confirmation != "y" && confirmation != "Y" {
			fmt.Println("Deletion cancelled.")
			return
		}

		// Proceed to delete the item
		if err := db.DeleteItemById(id); err != nil {
			fmt.Printf("Error deleting item: %v\n", err)
			return
		}
		fmt.Printf("Deleted item with URL: %s\n", item)
	},
}
