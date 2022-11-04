package services

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllFullArticles(t *testing.T) {

	req, err := http.NewRequest(http.MethodGet, "/fullarticles", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllFullArticles)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestGetSingleFullArticle(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/fullarticles", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("Id", "0")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetSingleFullArticle)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//expected := `[{"Id":0,"Title":"","desc":"","context":""}]`
	/* var expected = entities.Article{
		Id:      0,
		Title:   "",
		Desc:    "",
		Context: "",
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	} */

}
