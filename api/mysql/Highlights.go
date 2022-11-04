package mysqloperations

import (
	"api/entities"
	"database/sql"
	"fmt"
)

type HighlightModel struct {
	Db *sql.DB
}

func (highlightModel HighlightModel) Delete(id int64) (int64, error) {

	result, err := highlightModel.Db.Exec("DELETE from highlight where id=?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (highlightModel HighlightModel) Update(highlight entities.Highlight) (int64, error) {

	result, err := highlightModel.Db.Exec("UPDATE highlight set article_id=?, short_text=? where id=?", highlight.Article_Id, highlight.Short_Text, highlight.Id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (highlightModel HighlightModel) Create(highlight *entities.Highlight) error {

	result, err := highlightModel.Db.Exec("INSERT INTO highlight(article_id, short_text) VALUES (?,?)", highlight.Article_Id, highlight.Short_Text)
	if err != nil {
		return err
	} else {
		highlight.Id, _ = result.LastInsertId()
		return nil
	}

}

func (highlightModel HighlightModel) Find(id int64) (entities.Highlight, error) {

	rows, err := highlightModel.Db.Query("SELECT * FROM `highlight` where id=?", id)
	fmt.Println(entities.Highlight{}.Id)
	if err != nil {
		return entities.Highlight{}, err
	} else {
		highlight := entities.Highlight{}
		for rows.Next() {
			var id int64
			var article_id int64
			var short_text string
			var date string
			err := rows.Scan(&id, &article_id, &short_text, &date)
			if err != nil {
				return entities.Highlight{}, err
			}
			highlight = entities.Highlight{Id: id, Article_Id: article_id, Short_Text: short_text, Date: date}
		}
		return highlight, nil
	}

}

func (highlightModel HighlightModel) FindAll() ([]entities.Highlight, error) {

	rows, err := highlightModel.Db.Query("SELECT * FROM highlight")

	if err != nil {
		return nil, err
	}
	highlights := []entities.Highlight{}
	for rows.Next() {
		var id int64
		var article_id int64
		var short_text string
		var date string
		err := rows.Scan(&id, &article_id, &short_text, &date)
		if err != nil {
			return nil, err
		}
		highlight := entities.Highlight{Id: id, Article_Id: article_id, Short_Text: short_text, Date: date}

		highlights = append(highlights, highlight)

	}
	return highlights, nil

}
