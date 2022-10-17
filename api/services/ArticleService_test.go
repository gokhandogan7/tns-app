package services

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllArticles(t *testing.T) {

	req, err := http.NewRequest(http.MethodGet, "/articles", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllArticles)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestGetSingleArticle(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/article", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()

	q.Add("Id", "0")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetSingleArticle)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestCreateNewArticle(t *testing.T) {
	var jsonStr = []byte(`{"Id":45,"Title":"yeni title upgrade","desc":"yeni metdasdfsfahod","context":"yeni metfasfashod"}`)

	req, err := http.NewRequest(http.MethodPost, "/article", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateNewArticle)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestUpdateArticle(t *testing.T) {
	var jsonStr = []byte(`{"Id":0,"Title":"yeni title upgrade","desc":"yeni metdasdfsfahod","context":"yeni metfasfashod"}`)

	req, err := http.NewRequest(http.MethodPut, "/article", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateArticle)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteArticle(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/article", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()

	q.Add("Id", "0")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteArticle)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
