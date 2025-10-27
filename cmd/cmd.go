package cmd

import (
	"fmt"
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
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(DeleteCmd)
	rootCmd.AddCommand(OpenCmd)
}
