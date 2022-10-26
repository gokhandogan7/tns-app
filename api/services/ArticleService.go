package services

import (
	"api/config"
	"api/entities"
	"api/helpers"
	mysqloperations "api/mysql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Article represents the model for an article
type Article struct {
	Id      int64  `json:"Id" example: "1"`
	Title   string `json:"Title" example: "Title 1"`
	Desc    string `json:"desc" example: "Description 1"`
	Context string `json:"context" example: "Content 1"`
}

// GetAllArticles godoc
// @Summary Get all articles
// @Description Get details of all articles
// @Tags articles
// @Accept  json
// @Produce  json
// @Success 200 {array} Article
// @Router /articles [get]
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()

	helpers.EnableCors(&w)

	if err != nil {
		fmt.Println(err)
	} else {
		articleModel := mysqloperations.ArticleModel{
			Db: db,
		}
		emptyArticleList := []entities.Article{}
		articles, err := articleModel.FindAll()

		emptyArticleList = append(emptyArticleList, articles...)

		if err != nil {
			fmt.Println(err)
		}

		searchedArticles := []entities.Article{}
		searchKey := r.URL.Query().Get("search")
		Articles := emptyArticleList
		for _, article := range Articles {
			if helpers.IsSearched(article, searchKey) {

				searchedArticles = append(searchedArticles, article)

			}
		}
		json.NewEncoder(w).Encode(searchedArticles)
		if r.URL.Path != "/articles" {
			helpers.ErrorHandler(w, r, http.StatusNotFound)
			return
		}
	}

}

// GetSingleArticle godoc
// @Summary Get Single Article
// @Description Get single article with the input paylod
// @Tags articles
// @Accept  json
// @Produce  json
// @Param article body Article true "Update article"
// @Success 200 {object} Article
// @Router /article/{id} [get]
func GetSingleArticle(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()
	vars := mux.Vars(r)
	key := vars["id"]
	intKey, _ := strconv.ParseInt(key, 0, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		articleModel := mysqloperations.ArticleModel{
			Db: db,
		}
		emptyArticleList := []entities.Article{}
		article, err := articleModel.Find(intKey)
		singleArticle := append(emptyArticleList, article)
		fmt.Println(singleArticle)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(singleArticle)
	}
	defer db.Close()
}

// CreateArticle godoc
// @Summary Create a new article
// @Description Create a new article with the input paylod
// @Tags articles
// @Accept  json
// @Produce  json
// @Param article body Article true "Create article"
// @Success 200 {object} Article
// @Router /article [post]
func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	db, _ := config.GetMySQLDB()
	var article2 entities.Article

	err := json.NewDecoder(r.Body).Decode(&article2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	articleModel := mysqloperations.ArticleModel{
		Db: db,
	}
	fmt.Println("", article2)
	article := entities.Article{
		Id:      article2.Id,
		Title:   article2.Title,
		Desc:    article2.Desc,
		Context: article2.Context,
	}
	err2 := articleModel.Create(&article)
	if err != nil {
		fmt.Println(err2)
	}

	entities.Articles = append(entities.Articles, article)

	defer db.Close()

}

// UpdateArticle godoc
// @Summary Update an article
// @Description Update an article with the input paylod
// @Tags articles
// @Accept  json
// @Produce  json
// @Param article body Article true "Update article"
// @Success 200 {object} Article
// @Router /article/{id} [put]
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	db, _ := config.GetMySQLDB()
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	var articleType entities.Article
	id := vars["id"]
	intId, _ := strconv.ParseInt(id, 0, 32)

	err := json.NewDecoder(r.Body).Decode(&articleType)

	if err != nil {
		fmt.Println(err)
	} else {
		articleModel := mysqloperations.ArticleModel{
			Db: db,
		}
		article := entities.Article{
			Id:      intId,
			Title:   articleType.Title,
			Desc:    articleType.Desc,
			Context: articleType.Context,
		}
		rows, err := articleModel.Update(article)
		if err != nil {
			fmt.Println(err)
		} else {
			if rows > 0 {
				fmt.Println("done")
			}
		}

		entities.Articles = append(entities.Articles, article)

	}
	defer db.Close()

}

// DeleteArticle godoc
// @Summary Delete an article
// @Description Delete an article with the input paylod
// @Tags articles
// @Accept  json
// @Produce  json
// @Param article body Article true "Delete article"
// @Success 200 {object} Article
// @Router /article/{id} [delete]
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	db, _ := config.GetMySQLDB()
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]
	intId, _ := strconv.ParseInt(id, 0, 32)

	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		articleModel := mysqloperations.ArticleModel{
			Db: db,
		}
		rows, err := articleModel.Delete(intId)
		if err != nil {
			fmt.Println(err)
		} else {
			if rows > 0 {
			}
		}
	}
	defer db.Close()

}
