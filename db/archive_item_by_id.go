package db

import (
	"errors"
	"time"
)

func ArchiveItemById(id int) error {
	res, err := conn.Exec(`UPDATE items SET archived_at = ? WHERE id = ?`,
		time.Now().UTC().Format(time.RFC3339), id)
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
