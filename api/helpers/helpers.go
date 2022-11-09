package helpers

import (
	"net/http"
	"strings"
)

func RefactoredIsSearched(entry string, searchKey string) bool {

	a := strings.ReplaceAll(entry, " ", "")
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
