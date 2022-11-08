package helpers

import (
	"api/entities"
	"net/http"
	"strings"
)

// articles search helper
func IsSearched(article entities.FullArticle, searchKey string) bool {

	a := strings.ReplaceAll(article.Title, " ", "")
	s := strings.ReplaceAll(searchKey, " ", "")

	return strings.Contains(strings.ToLower(a), strings.ToLower(s))
}

// author search helper
func IsSearchedAuthor(author entities.Author, searchKey string) bool {

	a := strings.ReplaceAll(author.Name, " ", "")
	s := strings.ReplaceAll(searchKey, " ", "")

	return strings.Contains(strings.ToLower(a), strings.ToLower(s))
}

// content serach helper
func IsSearchedContent(author entities.Content, searchKey string) bool {

	a := strings.ReplaceAll(author.Text, " ", "")
	s := strings.ReplaceAll(searchKey, " ", "")

	return strings.Contains(strings.ToLower(a), strings.ToLower(s))
}

func IsSearchedHighlight(author entities.Highlight, searchKey string) bool {

	a := strings.ReplaceAll(author.Short_Text, " ", "")
	s := strings.ReplaceAll(searchKey, " ", "")

	return strings.Contains(strings.ToLower(a), strings.ToLower(s))
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {

	EnableCors(&w)
	if status == http.StatusNotFound {
		http.NotFound(w, r)
	}
}
