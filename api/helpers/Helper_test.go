package helpers

import (
	"api/entities"
	"testing"
)

func TestIsSearched(t *testing.T) {

	var TestArticle = entities.FullArticle{
		Id:         0,
		Title:      "SEAr chkey ",
		Desc:       "testDEscription",
		Content_Id: 1,
		Author_Id:  2,
		Name:       "Name",
		Email:      "Email",
		Text:       "Text",
		Image:      "Image",
		Short_Text: "Short_Text",
		Date:       "Date",
	}

	result := IsSearched(TestArticle, "searchkey")

	if result != true {
		t.Errorf("handler returned wrong status code: got %v want %v", result, true)
	}

}
