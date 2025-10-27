package db

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

func AddItem(url string) error {
	now := time.Now().UTC().Format(time.RFC3339)
	_, err := conn.Exec(`INSERT INTO items (url, added_at) VALUES (?, ?)`, url, now)
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "UNIQUE") {
		var archivedAt sql.NullString
		err := conn.QueryRow(`SELECT archived_at FROM items WHERE url = ?`, url).Scan(&archivedAt)
		if err != nil {
			return err
		}
		if archivedAt.Valid {
			_, err := conn.Exec(`UPDATE items SET archived_at = NULL, added_at = ? WHERE url = ?`, now, url)
			return err
		}
		return errors.New("URL already added")
	}
	return err
}
