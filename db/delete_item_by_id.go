package db

import (
	"errors"
)

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
