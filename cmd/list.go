package cmd

import (
	"fmt"

	"github.com/hossammourad/kewit/db"
	"github.com/hossammourad/kewit/helpers"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved URLs in your list that are not archived",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("Warning: listing shouldn't be executed with any arguments")
		}

		items, err := db.ListItems()
		if err != nil {
			fmt.Printf("Error listing items: %v\n", err)
			return
		}

		t := table.NewWriter()
		helpers.ConfigureTable(t)
		t.AppendHeader(table.Row{"ID", "URL", "Added At", "Archived At"})
		for _, item := range items {
			t.AppendRow(table.Row{
				item.Id,
				item.Url,
				helpers.LocalizeDate(item.AddedAt),
				"-",
			})
		}
		t.Render()
	},
}
