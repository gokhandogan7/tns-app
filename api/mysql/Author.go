package mysqloperations

import (
	"api/entities"
	"database/sql"
	"fmt"
)

type AuthorModel struct {
	Db *sql.DB
}

func (authorModel AuthorModel) Delete(id int64) (int64, error) {

	result, err := authorModel.Db.Exec("DELETE from author where id=?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (authorModel AuthorModel) Update(author entities.Author) (int64, error) {

	result, err := authorModel.Db.Exec("UPDATE author set name=?, email=?, where id=?", author.Name, author.Email, author.Id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (authorModel AuthorModel) Create(author *entities.Author) error {

	result, err := authorModel.Db.Exec("INSERT INTO author(name, email) VALUES (?,?)", author.Name, author.Email)
	if err != nil {
		return err
	} else {
		author.Id, _ = result.LastInsertId()
		return nil
	}

}

func (authorModel AuthorModel) Find(id int64) (entities.Author, error) {

	rows, err := authorModel.Db.Query("SELECT * FROM `author` where id=?", id)
	fmt.Println(entities.Author{}.Id)
	if err != nil {
		return entities.Author{}, err
	} else {
		author := entities.Author{}
		for rows.Next() {
			var id int64
			var name string
			var email string
			err := rows.Scan(&id, &name, &email)
			if err != nil {
				return entities.Author{}, err
			}
			author = entities.Author{Id: id, Name: name, Email: email}
		}
		return author, nil
	}

}

func (authorModel AuthorModel) FindAll() ([]entities.Author, error) {

	rows, err := authorModel.Db.Query("SELECT * FROM author")

	if err != nil {
		return nil, err
	}
	authors := []entities.Author{}
	for rows.Next() {
		var id int64
		var name string
		var email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			return nil, err
		}
		author := entities.Author{Id: id, Name: name, Email: email}

		authors = append(authors, author)

	}
	return authors, nil

}
