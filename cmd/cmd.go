package cmd

import (
	"os"

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
	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(DeleteCmd)
	rootCmd.AddCommand(OpenCmd)
}
