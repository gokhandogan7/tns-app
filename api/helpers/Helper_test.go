package helpers

import (
	"api/entities"
	"testing"
)

func TestIsSearched(t *testing.T) {

	expected := true

	var TestArticle = entities.Article{
		Id:      0,
		Title:   "SEArchkey ",
		Desc:    "testDEscription",
		Context: "test content",
	}

	result := IsSearched(TestArticle, "searchkey")

	if expected != result {
		t.Errorf("handler returned wrong status code: got %v want %v", result, expected)
	}

}
