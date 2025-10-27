package db

import (
	"database/sql"
	"errors"
)

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
