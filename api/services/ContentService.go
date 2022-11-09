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

func GetAllContents(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()

	helpers.EnableCors(&w)

	if err != nil {
		fmt.Println(err)
	} else {
		contentModel := mysqloperations.ContentModel{
			Db: db,
		}
		emptyContentList := []entities.Content{}
		contents, err := contentModel.FindAll()

		emptyContentList = append(emptyContentList, contents...)

		if err != nil {
			fmt.Println(err)
		}
		searchedContents := []entities.Content{}
		searchKey := r.URL.Query().Get("search")
		fmt.Println(searchKey)
		Contents := emptyContentList

		for _, content := range Contents {
			if helpers.RefactoredIsSearched(content.Text, searchKey) {

				searchedContents = append(searchedContents, content)

			}
		}

		json.NewEncoder(w).Encode(searchedContents)
		if r.URL.Path != "/contents" {
			helpers.ErrorHandler(w, r, http.StatusNotFound)
			return
		}
	}

}

func GetSingleContent(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()
	vars := mux.Vars(r)
	key := vars["id"]
	intKey, _ := strconv.ParseInt(key, 0, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		contentModel := mysqloperations.ContentModel{
			Db: db,
		}
		emptyContentList := []entities.Content{}
		content, err := contentModel.Find(intKey)
		singleContent := append(emptyContentList, content)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(singleContent)
	}
	defer db.Close()
}

func CreateNewContent(w http.ResponseWriter, r *http.Request) {
	db, _ := config.GetMySQLDB()
	var CreatedContent entities.Content

	err := json.NewDecoder(r.Body).Decode(&CreatedContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contentModel := mysqloperations.ContentModel{
		Db: db,
	}
	content := entities.Content{
		Text:  CreatedContent.Text,
		Image: CreatedContent.Image,
	}
	err2 := contentModel.Create(&content)
	if err != nil {
		fmt.Println(err2)
	}

	entities.Contents = append(entities.Contents, content)

	defer db.Close()

}

func UpdateContent(w http.ResponseWriter, r *http.Request) {
	db, _ := config.GetMySQLDB()
	vars := mux.Vars(r)
	// we will need to extract the `id` of the content we
	// wish to delete
	var contentType entities.Content
	id := vars["id"]
	intId, _ := strconv.ParseInt(id, 0, 32)

	err := json.NewDecoder(r.Body).Decode(&contentType)

	if err != nil {
		fmt.Println(err)
	} else {
		contentModel := mysqloperations.ContentModel{
			Db: db,
		}
		content := entities.Content{
			Id:    intId,
			Text:  contentType.Text,
			Image: contentType.Image,
		}
		rows, err := contentModel.Update(content)
		if err != nil {
			fmt.Println(err)
		} else {
			if rows > 0 {
				fmt.Println("done")
			}
		}

		entities.Contents = append(entities.Contents, content)

	}
	defer db.Close()

}

func DeleteContent(w http.ResponseWriter, r *http.Request) {
	db, _ := config.GetMySQLDB()
	vars := mux.Vars(r)
	// we will need to extract the `id` of the content we
	// wish to delete
	id := vars["id"]
	intId, _ := strconv.ParseInt(id, 0, 32)

	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		contentModel := mysqloperations.ContentModel{
			Db: db,
		}
		rows, err := contentModel.Delete(intId)
		if err != nil {
			fmt.Println(err)
		} else {
			if rows > 0 {
			}
		}
	}
	defer db.Close()

}
