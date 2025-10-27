package db

func ListItems() ([]Item, error) {
	rows, err := conn.Query(`SELECT id, url, added_at FROM items WHERE archived_at IS NULL ORDER BY added_at ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var id int
		var url string
		var addedAt string
		if err := rows.Scan(&id, &url, &addedAt); err != nil {
			return nil, err
		}
		items = append(items, Item{Url: url, AddedAt: addedAt, Id: id})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
