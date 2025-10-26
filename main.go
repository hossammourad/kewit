package main

import (
	"fmt"

	"github.com/hossammourad/kewit/cmd"
	"github.com/hossammourad/kewit/db"
)

func main() {
	if err := db.Init(); err != nil {
		fmt.Printf("DB init error: %v\n", err)
	}

	cmd.Execute()
}
