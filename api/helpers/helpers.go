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

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {

	if status == http.StatusNotFound {
		http.NotFound(w, r)
	}
}
