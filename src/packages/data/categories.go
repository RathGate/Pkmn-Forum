package data

import (
	"database/sql"
	"forum/packages/utils"
)

type Category struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	TopicCount int64  `json:"topic_count"`
	PostCount  int64  `json:"post_count"`
}

func GetCategories(dba utils.DB_Access) ([]Category, error) {
	db, err := sql.Open("mysql", dba.ToString())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT    c.id,    c.name AS category_name,  COUNT(DISTINCT t.id) AS topic_count, COUNT(DISTINCT p.id) AS post_count FROM    categories c LEFT JOIN    topics t ON c.id = t.category_id LEFT JOIN   posts p ON t.id = p.topic_id GROUP BY     c.id, c.name;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.TopicCount, &category.PostCount)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
