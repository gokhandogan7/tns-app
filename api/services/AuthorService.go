package services

import (
	"api/config"
	"api/entities"
	"api/helpers"
	mysqloperations "api/mysql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		authorModel := mysqloperations.AuthorModel{
			Db: db,
		}
		emptyAuthorList := []entities.Author{}
		authors, err := authorModel.FindAll()

		emptyAuthorList = append(emptyAuthorList, authors...)

		if err != nil {
			fmt.Println(err)
		}
		searchedAuthors := []entities.Author{}
		searchKey := r.URL.Query().Get("search")
		Authors := emptyAuthorList

		for _, author := range Authors {
			if helpers.RefactoredIsSearched(author.Name, searchKey) {

				searchedAuthors = append(searchedAuthors, author)

			}
		}
		json.NewEncoder(w).Encode(searchedAuthors)
		if r.URL.Path != "/authors" {
			helpers.ErrorHandler(w, r, http.StatusNotFound)
			return
		}
	}

}

func GetSingleAuthor(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()
	vars := mux.Vars(r)
	key := vars["id"]
	intKey, _ := strconv.ParseInt(key, 0, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		authorModel := mysqloperations.AuthorModel{
			Db: db,
		}
		emptyAuthorList := []entities.Author{}
		author, err := authorModel.Find(intKey)
		singleAuthor := append(emptyAuthorList, author)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(singleAuthor)
	}
}

func CreateNewAuthor(w http.ResponseWriter, r *http.Request) {
	db, _ := config.GetMySQLDB()
	var CreatedAuthor entities.Author

	err := json.NewDecoder(r.Body).Decode(&CreatedAuthor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authorModel := mysqloperations.AuthorModel{
		Db: db,
	}
	author := entities.Author{
		Name:  CreatedAuthor.Name,
		Email: CreatedAuthor.Email,
	}
	err2 := authorModel.Create(&author)
	if err != nil {
		fmt.Println(err2)
	}

	entities.Authors = append(entities.Authors, author)

}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	db, _ := config.GetMySQLDB()
	vars := mux.Vars(r)
	// we will need to extract the `id` of the author we
	// wish to delete
	var authorType entities.Author
	id := vars["id"]
	intId, _ := strconv.ParseInt(id, 0, 32)

	err := json.NewDecoder(r.Body).Decode(&authorType)

	if err != nil {
		fmt.Println(err)
	} else {
		authorModel := mysqloperations.AuthorModel{
			Db: db,
		}
		author := entities.Author{
			Id:    intId,
			Name:  authorType.Name,
			Email: authorType.Email,
		}
		rows, err := authorModel.Update(author)
		if err != nil {
			fmt.Println(err)
		} else {
			if rows > 0 {
				fmt.Println("done")
			}
		}

		entities.Authors = append(entities.Authors, author)

	}

}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	db, _ := config.GetMySQLDB()
	vars := mux.Vars(r)
	// we will need to extract the `id` of the author we
	// wish to delete
	id := vars["id"]
	intId, _ := strconv.ParseInt(id, 0, 32)

	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		authorModel := mysqloperations.AuthorModel{
			Db: db,
		}
		rows, err := authorModel.Delete(intId)
		if err != nil {
			fmt.Println(err)
		} else {
			if rows > 0 {
			}
		}
	}

}
