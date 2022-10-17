package helpers

import (
	"api/entities"
	"testing"
)

func TestIsSearched(t *testing.T) {

	var TestArticle = entities.Article{
		Id:      0,
		Title:   "SEAr chkey ",
		Desc:    "testDEscription",
		Context: "test content",
	}

	result := IsSearched(TestArticle, "searchkey")

	if result != true {
		t.Errorf("handler returned wrong status code: got %v want %v", result, true)
	}

}
