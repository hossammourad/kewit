package cmd

import (
	"fmt"

	"github.com/hossammourad/kewit/db"
	"github.com/hossammourad/kewit/helpers"
	"github.com/spf13/cobra"
)

var OpenCmd = &cobra.Command{
	Use:   "open [id]",
	Short: "Open a saved URL in your default browser by its ID",
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

		err = helpers.OpenBrowser(url)
		if err != nil {
			fmt.Printf("Error opening URL in your default browser: %v\n", err)
			return
		}
		fmt.Printf("'%s' is opened in your default browser\n", url)
	},
}
