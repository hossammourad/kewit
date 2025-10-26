package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/hossammourad/kewit/db"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kewit",
	Short: "A read-it-later app that runs on your terminal",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	addCmd := &cobra.Command{
		Use:   "add [url]",
		Short: "Add a URL to your read-later list",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]
			if err := db.AddItem(url); err != nil {
				fmt.Printf("Error adding URL: %v\n", err)
				return
			}
			fmt.Printf("Added URL: %s\n", url)
		},
	}
	rootCmd.AddCommand(addCmd)

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List saved URLs in your list",
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
			t.SetOutputMirror(os.Stdout)
			t.SetStyle(table.StyleRounded)
			t.Style().Options.SeparateRows = true
			t.AppendHeader(table.Row{"#", "URL", "Added At"})
			for _, item := range items {
				date, _ := time.Parse(time.RFC3339, item.AddedAt)
				localDate := date.Local()
				t.AppendRow(table.Row{item.Id, item.Url, localDate.Format("2006-01-02 / 15:04")})
			}
			t.Render()
		},
	}
	rootCmd.AddCommand(listCmd)

	deleteCmd := &cobra.Command{
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
	rootCmd.AddCommand(deleteCmd)
}
