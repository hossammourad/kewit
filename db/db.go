package db

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var conn *sql.DB

func Init() error {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".kewit.db")
	c, err := sql.Open("sqlite3", path)
	if err != nil {
		return err
	}
	if _, err = c.Exec(`CREATE TABLE IF NOT EXISTS items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        url TEXT NOT NULL UNIQUE,
        added_at TEXT NOT NULL
    )`); err != nil {
		return err
	}
	conn = c
	return nil
}

type item struct {
	Url     string
	AddedAt string
}

func AddItem(url string) error {
	_, err := conn.Exec(`INSERT INTO items (url, added_at) VALUES (?, ?)`,
		url, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return errors.New("URL already added")
		}
	}
	return nil
}

func ListItems() ([]item, error) {
	rows, err := conn.Query(`SELECT url, added_at FROM items`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []item
	for rows.Next() {
		var url string
		var addedAt string
		if err := rows.Scan(&url, &addedAt); err != nil {
			return nil, err
		}
		items = append(items, item{Url: url, AddedAt: addedAt})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
