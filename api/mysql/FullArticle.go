package mysqloperations

import (
	"api/entities"
	"database/sql"
	"fmt"
)

type FullArticleModel struct {
	Db *sql.DB
}

func (fullArticleModel FullArticleModel) FindAll() ([]entities.FullArticle, error) {
	fmt.Println("calisiyor mu ful article")
	rows, err := fullArticleModel.Db.Query(`select 
	articless.id, articless.title, articless.description, articless.author_id, articless.content_id, 
	author.name, author.email, 
	content.text, content.image, 
	highlight.short_text, highlight.date 
	from articless 
	inner join author 
	inner join content 
	inner join highlight
	on 
	articless.author_id = author.id 
	and
	articless.content_id=content.id
	and 
	articless.id = highlight.article_id;`)

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
	rows, err := fullArticleModel.Db.Query(`select 
	articless.id, articless.title, articless.description, articless.author_id, articless.content_id, 
	author.name, author.email, 
	content.text, content.image, 
	highlight.short_text, highlight.date 
	from articless 
	inner join author 
	inner join content 
	inner join highlight
	on 
	articless.author_id = author.id 
	and
	articless.content_id=content.id
	and 
	articless.id = highlight.article_id
	where author_id=?;`, id)

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

	rows, err := fullArticleModel.Db.Query(`select
	articless.id, articless.title, articless.description, articless.author_id, articless.content_id,
	author.name, author.email,
	content.text, content.image,
	highlight.short_text, highlight.date
	from articless
	inner join author
	inner join content
	inner join highlight
	on
	articless.author_id = author.id
	and
	articless.content_id=content.id
	and
	articless.id = highlight.article_id
	where author_id=? and articless.id=?;`, user_id, article_id)
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
