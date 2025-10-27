package db

func ListArchivedItems() ([]Item, error) {
	rows, err := conn.Query(`SELECT id, url, added_at, archived_at FROM items WHERE archived_at IS NOT NULL ORDER BY archived_at ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var id int
		var url string
		var addedAt string
		var archivedAt string
		if err := rows.Scan(&id, &url, &addedAt, &archivedAt); err != nil {
			return nil, err
		}
		items = append(items, Item{Url: url, AddedAt: addedAt, ArchivedAt: archivedAt, Id: id})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
