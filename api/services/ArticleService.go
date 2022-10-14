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

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()

	helpers.EnableCors(&w)

	if err != nil {
		fmt.Println(err)
	} else {
		articleModel := mysqloperations.ArticleModel{
			Db: db,
		}

		articles, err := articleModel.FindAll()
		entities.Articles = append(entities.Articles, articles...)

		if err != nil {
			fmt.Println(err)
		}
	}

	searchedArticles := []entities.Article{}
	searchKey := r.URL.Query().Get("search")
	Articles := entities.Articles
	for _, article := range Articles {
		if helpers.IsSearched(article, searchKey) {

			searchedArticles = append(searchedArticles, article)

		}
	}
	json.NewEncoder(w).Encode(searchedArticles)
	if r.URL.Path != "/articles" {
		helpers.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

}

func GetSingleArticle(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()
	vars := mux.Vars(r)
	key := vars["id"]
	intKey, _ := strconv.ParseInt(key, 0, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		articleModel := mysqloperations.ArticleModel{
			Db: db,
		}
		emptyArticleList := []entities.Article{}
		article, err := articleModel.Find(intKey)
		singleArticle := append(emptyArticleList, article)
		fmt.Println(singleArticle)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(singleArticle)
	}
	defer db.Close()
}

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	db, _ := config.GetMySQLDB()
	var article2 entities.Article

	err := json.NewDecoder(r.Body).Decode(&article2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	articleModel := mysqloperations.ArticleModel{
		Db: db,
	}
	fmt.Println("", article2)
	article := entities.Article{
		Id:      article2.Id,
		Title:   article2.Title,
		Desc:    article2.Desc,
		Context: article2.Context,
	}
	err2 := articleModel.Create(&article)
	if err != nil {
		fmt.Println(err2)
	}

	entities.Articles = append(entities.Articles, article)

	defer db.Close()

}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	db, _ := config.GetMySQLDB()
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	var articleType entities.Article
	id := vars["id"]
	intId, _ := strconv.ParseInt(id, 0, 32)

	err := json.NewDecoder(r.Body).Decode(&articleType)

	if err != nil {
		fmt.Println(">>>>>>0", err)
	} else {
		articleModel := mysqloperations.ArticleModel{
			Db: db,
		}
		article := entities.Article{
			Id:      intId,
			Title:   articleType.Title,
			Desc:    articleType.Desc,
			Context: articleType.Context,
		}
		rows, err := articleModel.Update(article)
		if err != nil {
			fmt.Println(err)
		} else {
			if rows > 0 {
				fmt.Println("done")
			}
		}

		entities.Articles = append(entities.Articles, article)

	}
	defer db.Close()

}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
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
		articleModel := mysqloperations.ArticleModel{
			Db: db,
		}
		rows, err := articleModel.Delete(intId)
		if err != nil {
			fmt.Println(err)
		} else {
			if rows > 0 {
				fmt.Println("delete done")
			}
		}
	}
	defer db.Close()

}

func RemoveArticle(w http.ResponseWriter, r *http.Request) {
	var article2 entities.Article
	db, _ := config.GetMySQLDB()

	err := json.NewDecoder(r.Body).Decode(&article2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	articleModel := mysqloperations.ArticleModel{
		Db: db,
	}
	rows, err := articleModel.Delete(article2.Id)
	if err != nil {
		fmt.Println(">>>>>1", err)
	} else {
		if rows > 0 {
			fmt.Println("delete done")
		}
	}

	defer db.Close()
}
