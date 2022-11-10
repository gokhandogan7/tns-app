package mysqloperations

import (
	"api/entities"
	"database/sql"
)

type FullArticleModel struct {
	Db *sql.DB
}

func (fullArticleModel FullArticleModel) FindAll(limit string) ([]entities.FullArticle, error) {

	rows, err := fullArticleModel.Db.Query(`SELECT 
	articles.id, articles.title, articles.description, articles.author_id, articles.content_id, 
	author.name, author.email, 
	content.text, content.image, 
	highlight.short_text, highlight.date 
	from articles 
	inner join author 
	inner join content 
	inner join highlight
	on 
	articles.author_id = author.id 
	and
	articles.content_id=content.id
	and 
	articles.id = highlight.article_id LIMIT ?;`, limit)

	if err != nil {
		return nil, err
	}
	fullArticles := []entities.FullArticle{}
	for rows.Next() {
		var id int64
		var title string
		var desc string
		var content_id int64
		var author_id int64
		var name string
		var email string
		var text string
		var image string
		var short_text string
		var date string
		err := rows.Scan(&id, &title, &desc, &content_id, &author_id, &name, &email, &text, &image, &short_text, &date)
		if err != nil {
			return nil, err
		}
		fullArticle := entities.FullArticle{Id: id, Title: title, Desc: desc, Content_Id: content_id, Author_Id: author_id, Name: name, Email: email, Text: text, Image: image, Short_Text: short_text, Date: date}

		fullArticles = append(fullArticles, fullArticle)

	}
	return fullArticles, nil

}

func (fullArticleModel FullArticleModel) FindUsersAll(id int64) ([]entities.FullArticle, error) {
	rows, err := fullArticleModel.Db.Query(`SELECT
	articles.id, articles.title, articles.description, articles.author_id, articles.content_id,
	author.name, author.email,
	content.text, content.image,
	highlight.short_text, highlight.date
	FROM articles
	INNER JOIN author
	INNER JOIN content
	INNER JOIN highlight
	ON
	articles.author_id = author.id
	AND
	articles.content_id=content.id
	AND
	articles.id = highlight.article_id
	WHERE author_id=?;`, id)

	if err != nil {
		return nil, err
	}
	fullArticles := []entities.FullArticle{}
	for rows.Next() {
		var id int64
		var title string
		var desc string
		var content_id int64
		var author_id int64
		var name string
		var email string
		var text string
		var image string
		var short_text string
		var date string
		err := rows.Scan(&id, &title, &desc, &content_id, &author_id, &name, &email, &text, &image, &short_text, &date)
		if err != nil {
			return nil, err
		}
		fullArticle := entities.FullArticle{Id: id, Title: title, Desc: desc, Content_Id: content_id, Author_Id: author_id, Name: name, Email: email, Text: text, Image: image, Short_Text: short_text, Date: date}

		fullArticles = append(fullArticles, fullArticle)

	}
	return fullArticles, nil

}

func (fullArticleModel FullArticleModel) Find(user_id int64, article_id int64) (entities.FullArticle, error) {

	rows, err := fullArticleModel.Db.Query(`SELECT
	articles.id, articles.title, articles.description, articles.author_id, articles.content_id,
	author.name, author.email,
	content.text, content.image,
	highlight.short_text, highlight.date
	FROM articles
	INNER JOIN author
	INNER JOIN content
	INNER JOIN highlight
	ON
	articles.author_id = author.id
	AND
	articles.content_id=content.id
	AND
	articles.id = highlight.article_id
	WHERE author_id=? AND articles.id=?;`, user_id, article_id)
	if err != nil {
		return entities.FullArticle{}, err
	} else {
		fullArticle := entities.FullArticle{}
		for rows.Next() {
			var id int64
			var title string
			var desc string
			var content_id int64
			var author_id int64
			var name string
			var email string
			var text string
			var image string
			var short_text string
			var date string
			err := rows.Scan(&id, &title, &desc, &content_id, &author_id, &name, &email, &text, &image, &short_text, &date)
			if err != nil {
				return entities.FullArticle{}, err
			}
			fullArticle = entities.FullArticle{Id: id, Title: title, Desc: desc, Content_Id: content_id, Author_Id: author_id, Name: name, Email: email, Text: text, Image: image, Short_Text: short_text, Date: date}
		}
		return fullArticle, nil
	}

}
