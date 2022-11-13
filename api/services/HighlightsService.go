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

func GetAllHighlights(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		highlightModel := mysqloperations.HighlightModel{
			Db: db,
		}
		emptyHighlightList := []entities.Highlight{}
		highlights, err := highlightModel.FindAll()
		emptyHighlightList = append(emptyHighlightList, highlights...)

		if err != nil {
			fmt.Println(err)
		}

		searchedHighlights := []entities.Highlight{}
		searchKey := r.URL.Query().Get("search")
		Highlights := emptyHighlightList

		for _, highlight := range Highlights {
			if helpers.RefactoredIsSearched(highlight.Short_Text, searchKey) {

				searchedHighlights = append(searchedHighlights, highlight)

			}
		}
		json.NewEncoder(w).Encode(searchedHighlights)
		if r.URL.Path != "/highlights" {
			helpers.ErrorHandler(w, r, http.StatusNotFound)
			return
		}
	}

}

func GetSingleHighlight(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()
	vars := mux.Vars(r)
	key := vars["id"]
	intKey, _ := strconv.ParseInt(key, 0, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		highlightModel := mysqloperations.HighlightModel{
			Db: db,
		}
		emptyHighlightList := []entities.Highlight{}
		highlight, err := highlightModel.Find(intKey)
		singleHighlight := append(emptyHighlightList, highlight)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(singleHighlight)
	}
}

func CreateNewHighlight(w http.ResponseWriter, r *http.Request) {
	db, _ := config.GetMySQLDB()
	var CreatedHighlight entities.Highlight

	err := json.NewDecoder(r.Body).Decode(&CreatedHighlight)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	highlightModel := mysqloperations.HighlightModel{
		Db: db,
	}
	highlight := entities.Highlight{
		Article_Id: CreatedHighlight.Article_Id,
		Short_Text: CreatedHighlight.Short_Text,
	}
	err2 := highlightModel.Create(&highlight)
	if err != nil {
		fmt.Println(err2)
	}

	entities.Highlights = append(entities.Highlights, highlight)

}

func UpdateHighlight(w http.ResponseWriter, r *http.Request) {
	db, _ := config.GetMySQLDB()
	vars := mux.Vars(r)
	// we will need to extract the `id` of the highlight we
	// wish to delete
	var highlightType entities.Highlight
	id := vars["id"]
	intId, _ := strconv.ParseInt(id, 0, 32)

	err := json.NewDecoder(r.Body).Decode(&highlightType)

	if err != nil {
		fmt.Println(err)
	} else {
		highlightModel := mysqloperations.HighlightModel{
			Db: db,
		}
		highlight := entities.Highlight{
			Id:         intId,
			Article_Id: highlightType.Article_Id,
			Short_Text: highlightType.Short_Text,
		}
		rows, err := highlightModel.Update(highlight)
		if err != nil {
			fmt.Println(err)
		} else {
			if rows > 0 {
				fmt.Println("done")
			}
		}

		entities.Highlights = append(entities.Highlights, highlight)

	}

}

func DeleteHighlight(w http.ResponseWriter, r *http.Request) {
	db, _ := config.GetMySQLDB()
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]
	intId, _ := strconv.ParseInt(id, 0, 32)

	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		highlightModel := mysqloperations.HighlightModel{
			Db: db,
		}
		rows, err := highlightModel.Delete(intId)
		if err != nil {
			fmt.Println(err)
		} else {
			if rows > 0 {
			}
		}
	}

}
