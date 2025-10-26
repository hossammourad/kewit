package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/hossammourad/kewit/db"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kewit",
	Short: "A read-it-later app that runs on your terminal",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kewit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

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
			for i, item := range items {
				date, _ := time.Parse(time.RFC3339, item.AddedAt)
				localDate := date.Local()
				t.AppendRow(table.Row{i + 1, item.Url, localDate.Format("2006-01-02 / 15:04")})
			}
			t.Render()
		},
	}
	rootCmd.AddCommand(listCmd)
}
