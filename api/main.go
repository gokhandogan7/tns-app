package main

import (
	"api/config"
	"api/entities"
	"api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

var searchedArticles []entities.Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	searchedArticles := []entities.Article{}
	enableCors(&w)
	searchKey := r.URL.Query().Get("search")
	Articles := entities.Articles
	for _, article := range Articles {
		if isSearched(article, searchKey) {

			searchedArticles = append(searchedArticles, article)

		}
	}

	json.NewEncoder(w).Encode(searchedArticles)
	if r.URL.Path != "/articles" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

}

func isSearched(article entities.Article, searchKey string) bool {

	a := strings.ReplaceAll(article.Title, " ", "")
	s := strings.ReplaceAll(searchKey, " ", "")
	return strings.Contains(strings.ToLower(a), strings.ToLower(s))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {

	enableCors(&w)
	if status == http.StatusNotFound {
		http.NotFound(w, r)
	}
}

func handleRequests() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {

	Article_CallFindAll()
	handleRequests()

}

func Article_CallFindAll() {
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		articleModel := models.ArticleModel{
			Db: db,
		}
		articles, err := articleModel.FindAll()
		entities.Articles = append(entities.Articles, articles...)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(articles)
			fmt.Println("Articles List")
			for _, article := range articles {
				fmt.Println("Title", article.Title)
				fmt.Println("Title", article.Desc)
				fmt.Println("Title", article.Content)
				fmt.Println("----------------------")
			}

		}
	}

}
