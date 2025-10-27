package db

import (
	"database/sql"
	"os"
	"path/filepath"

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
        added_at TEXT NOT NULL,
        archived_at TEXT DEFAULT NULL
    )`); err != nil {
		return err
	}
	conn = c
	return nil
}
