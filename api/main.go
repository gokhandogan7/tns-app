package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles = []Article{
	Article{Title: "Hello", Desc: "Articlcription", Content: "Article Content"},
	Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	Article{Title: "Hello 3", Desc: "icription", Content: "Article Content"},
	Article{Title: "Hello 4", Desc: "cle Description", Content: "Article Content"},
	Article{Title: "Hello 5", Desc: "ticleescription", Content: "Article Content"},
	Article{Title: "strasbourg", Desc: "Article Description", Content: "Article Content"},
	Article{Title: "strdurg", Desc: "Article Description", Content: "Article Content"},
	Article{Title: "sfdfg", Desc: "Article Description", Content: "Article Content"},
	Article{Title: "s", Desc: "ticlcription", Content: "Article Content"},
	Article{Title: "machine", Desc: "Article Description", Content: "Article Content"},
	Article{Title: "macht", Desc: "Article Description", Content: "Article Content"},
	Article{Title: "He", Desc: "rticle Description", Content: "Article Content"},
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

var searchedArticles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	searchedArticles := []Article{}
	enableCors(&w)
	searchKey := r.URL.Query().Get("search")
	Articles := Articles
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

func isSearched(article Article, searchKey string) bool {

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

	fmt.Println("Searched:")
	for _, u := range searchedArticles {
		fmt.Println(u)
	}

	fmt.Println(Articles)
	handleRequests()

}
