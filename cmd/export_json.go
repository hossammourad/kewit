package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hossammourad/kewit/db"
	"github.com/spf13/cobra"
)

var ExportJsonCmd = &cobra.Command{
	Use:   "export-json",
	Short: "Export your list as JSON",
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch all items from the database
		items, err := db.ListItems()
		if err != nil {
			fmt.Println("Error fetching items:", err)
			return
		}

		// Convert items to JSON
		jsonData, err := json.MarshalIndent(items, "", "  ")
		if err != nil {
			fmt.Println("Error converting items to JSON:", err)
			return
		}

		// Get user's home directory and prepare file path
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return
		}
		fileName := "kewit_list.json"
		filePath := filepath.Join(homeDir, fileName)

		// Check if file already exists
		_, err = os.Stat(filePath)
		if err == nil {
			fmt.Printf("File %q already exists at your home directory, please remove or rename it before exporting\n", fileName)
			return
		}

		// Write JSON data to file
		err = os.WriteFile(filePath, jsonData, 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		fmt.Printf("Items successfully exported to %q and saved in your home directory\n", fileName)
	},
}
