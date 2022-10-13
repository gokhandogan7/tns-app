package models

import (
	"api/entities"
	"database/sql"
)

type ArticleModel struct {
	Db *sql.DB
}

func (articleModel ArticleModel) FindAll() ([]entities.Article, error) {

	rows, err := articleModel.Db.Query("SELECT * FROM `articles`")

	if err != nil {
		return nil, err
	} else {
		articles := []entities.Article{}
		for rows.Next() {
			var id int
			var title string
			var desc string
			var content string
			err2 := rows.Scan(&id, &title, &desc, &content)
			if err2 != nil {
				return nil, err2
			} else {
				article := entities.Article{title, desc, content}

				articles = append(articles, article)
			}
		}
		return articles, nil
	}

}
