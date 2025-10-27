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

type item struct {
	Id      int
	Url     string
	AddedAt string
}

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
	rows, err := conn.Query(`SELECT id, url, added_at FROM items`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []item
	for rows.Next() {
		var id int
		var url string
		var addedAt string
		if err := rows.Scan(&id, &url, &addedAt); err != nil {
			return nil, err
		}
		items = append(items, item{Url: url, AddedAt: addedAt, Id: id})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func DeleteItemById(id int) error {
	res, err := conn.Exec(`DELETE FROM items WHERE id = ?`, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no item found with the given id")
	}
	return nil
}

func GetItemById(id int) (string, error) {
	var url string
	err := conn.QueryRow(`SELECT url FROM items WHERE id = ?`, id).Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("no item found with the given id")
		}
		return "", err
	}
	return url, nil
}
