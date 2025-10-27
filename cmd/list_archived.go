package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/hossammourad/kewit/db"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var ListArchivedCmd = &cobra.Command{
	Use:   "list-archived",
	Short: "List archived URLs in your list",
	Run: func(cmd *cobra.Command, args []string) {
		items, err := db.ListArchivedItems()
		if err != nil {
			fmt.Printf("Error listing archived items: %v\n", err)
			return
		}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleRounded)
		t.Style().Options.SeparateRows = true
		t.AppendHeader(table.Row{"ID", "URL", "Added At", "Archived At"})
		for _, item := range items {
			addedAtDate, _ := time.Parse(time.RFC3339, item.AddedAt)
			localizedAddedAtDate := addedAtDate.Local()
			archivedAtDate, _ := time.Parse(time.RFC3339, item.ArchivedAt)
			localizedArchivedAtDate := archivedAtDate.Local()
			t.AppendRow(table.Row{
				item.Id,
				item.Url,
				localizedAddedAtDate.Format("2006-01-02 / 15:04"),
				localizedArchivedAtDate.Format("2006-01-02 / 15:04"),
			})
		}
		t.Render()
	},
}
