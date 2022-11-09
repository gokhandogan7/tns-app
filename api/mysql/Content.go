package mysqloperations

import (
	"api/entities"
	"database/sql"
	"fmt"
)

type ContentModel struct {
	Db *sql.DB
}

func (contentModel ContentModel) Delete(id int64) (int64, error) {

	result, err := contentModel.Db.Exec("DELETE from content WHERE id=?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (contentModel ContentModel) Update(content entities.Content) (int64, error) {

	result, err := contentModel.Db.Exec("UPDATE content set text=?, image=? WHERE id=?", content.Text, content.Image, content.Id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (contentModel ContentModel) Create(content *entities.Content) error {

	result, err := contentModel.Db.Exec("INSERT INTO content(text, image) VALUES (?,?)", content.Text, content.Image)
	if err != nil {
		return err
	} else {
		content.Id, _ = result.LastInsertId()
		return nil
	}

}

func (contentModel ContentModel) Find(id int64) (entities.Content, error) {

	rows, err := contentModel.Db.Query("SELECT * FROM `content` WHERE id=?", id)
	fmt.Println(entities.Content{}.Id)
	if err != nil {
		return entities.Content{}, err
	} else {
		content := entities.Content{}
		for rows.Next() {
			var id int64
			var text string
			var image string
			err := rows.Scan(&id, &text, &image)
			if err != nil {
				return entities.Content{}, err
			}
			content = entities.Content{Id: id, Text: text, Image: image}
		}
		return content, nil
	}

}

func (contentModel ContentModel) FindAll() ([]entities.Content, error) {

	rows, err := contentModel.Db.Query("SELECT * FROM content")

	if err != nil {
		return nil, err
	}
	contents := []entities.Content{}
	for rows.Next() {
		var id int64
		var text string
		var image string
		err := rows.Scan(&id, &text, &image)
		if err != nil {
			return nil, err
		}
		content := entities.Content{Id: id, Text: text, Image: image}

		contents = append(contents, content)

	}
	return contents, nil

}
