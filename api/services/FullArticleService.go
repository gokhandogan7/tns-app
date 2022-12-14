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

// GetAllArticles godoc
// @Summary Get all articles
// @Description Get details of all articles
// @Tags articles
// @Accept  json
// @Produce  json
// @Success 200 {array} Article
// @Router /articles [get]
func GetAllFullArticles(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		fullArticleModel := mysqloperations.FullArticleModel{
			Db: db,
		}
		emptyFullArticleList := []entities.FullArticle{}
		limit := r.URL.Query().Get("limit")
		fullArticles, err := fullArticleModel.FindAll(limit)

		emptyFullArticleList = append(emptyFullArticleList, fullArticles...)

		if err != nil {
			fmt.Println(err)
		}

		searchedFullArticles := []entities.FullArticle{}
		searchKey := r.URL.Query().Get("search")
		Articles := emptyFullArticleList

		for _, article := range Articles {
			if helpers.RefactoredIsSearched(article.Title, searchKey) {

				searchedFullArticles = append(searchedFullArticles, article)

			}
		}
		json.NewEncoder(w).Encode(searchedFullArticles)
		if r.URL.Path != "/fullarticles" {
			helpers.ErrorHandler(w, r, http.StatusNotFound)
			return
		}
	}

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
func GetUsersAllFullArticles(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()
	vars := mux.Vars(r)
	key := vars["id"]
	intKey, _ := strconv.ParseInt(key, 0, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		fullArticleModel := mysqloperations.FullArticleModel{
			Db: db,
		}
		emptyFullArticleList := []entities.FullArticle{}
		fullArticles, err := fullArticleModel.FindUsersAll(intKey)
		emptyFullArticleList = append(emptyFullArticleList, fullArticles...)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(emptyFullArticleList)
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
func GetSingleFullArticle(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()
	vars := mux.Vars(r)
	user_id := vars["user_id"]
	article_id := vars["article_id"]
	intUser_id, _ := strconv.ParseInt(user_id, 0, 32)
	intArticle_id, _ := strconv.ParseInt(article_id, 0, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		fullArticleModel := mysqloperations.FullArticleModel{
			Db: db,
		}
		emptyFullArticleList := []entities.FullArticle{}
		singleArticle, err := fullArticleModel.Find(intUser_id, intArticle_id)
		singleFullArticleList := append(emptyFullArticleList, singleArticle)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(singleFullArticleList)
	}
}
