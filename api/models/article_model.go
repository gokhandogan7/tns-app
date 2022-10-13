package models

import (
	"api/entities"
	"database/sql"
)

type ArticleModel struct {
	Db *sql.DB
}

func (articleModel ArticleModel) Delete(id int64) (int64, error) {

	result, err := articleModel.Db.Exec("DELETE from articles where id=?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (articleModel ArticleModel) Update(article entities.Article) (int64, error) {

	result, err := articleModel.Db.Exec("UPDATE articles set title=?, description=?, context=? where id=?", article.Title, article.Desc, article.Context, article.Id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (articleModel ArticleModel) Create(article *entities.Article) error {

	result, err := articleModel.Db.Exec("INSERT INTO articles(id, title, description, context) VALUES (?,?,?,?)", article.Id, article.Title, article.Desc, article.Context)
	if err != nil {
		return err
	} else {
		article.Id, _ = result.LastInsertId()
		return nil
	}

}

func (articleModel ArticleModel) Find(id int64) (entities.Article, error) {

	rows, err := articleModel.Db.Query("SELECT * FROM `articles` where id=?", id)

	if err != nil {
		return entities.Article{}, err
	} else {
		article := entities.Article{}
		for rows.Next() {
			var id int64
			var title string
			var desc string
			var context string
			err2 := rows.Scan(&id, &title, &desc, &context)
			if err2 != nil {
				return entities.Article{}, err2
			} else {
				article = entities.Article{Id: id, Title: title, Desc: desc, Context: context}

			}
		}
		return article, nil
	}

}

func (articleModel ArticleModel) FindAll() ([]entities.Article, error) {

	rows, err := articleModel.Db.Query("SELECT * FROM articles")

	if err != nil {
		return nil, err
	}
	articles := []entities.Article{}
	for rows.Next() {
		var id int64
		var title string
		var desc string
		var context string
		err := rows.Scan(&id, &title, &desc, &context)
		if err != nil {
			return nil, err
		}
		article := entities.Article{Id: id, Title: title, Desc: desc, Context: context}

		articles = append(articles, article)

	}
	return articles, nil

}
