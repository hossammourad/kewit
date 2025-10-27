package helpers

import (
	"fmt"
	"time"
)

func LocalizeDate(dateStr string) string {
	addedAtDate, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return dateStr
	}
	localizedAddedAtDate := addedAtDate.Local()
	return localizedAddedAtDate.Format("2006-01-02 / 15:04")
}
