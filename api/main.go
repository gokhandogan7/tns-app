package main

import (
	"api/config"
	"api/entities"
	"api/helpers"
	"api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

/* func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
} */

/* func isSearched(article entities.Article, searchKey string) bool {

	a := strings.ReplaceAll(article.Title, " ", "")
	s := strings.ReplaceAll(searchKey, " ", "")
	return strings.Contains(strings.ToLower(a), strings.ToLower(s))
} */

func returnAllArticles(w http.ResponseWriter, r *http.Request) {

	searchedArticles := []entities.Article{}
	helpers.EnableCors(&w)
	searchKey := r.URL.Query().Get("search")
	Articles := entities.Articles
	for _, article := range Articles {
		if helpers.IsSearched(article, searchKey) {

			searchedArticles = append(searchedArticles, article)

		}
	}
	json.NewEncoder(w).Encode(searchedArticles)
	if r.URL.Path != "/articles" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {

	helpers.EnableCors(&w)
	if status == http.StatusNotFound {
		http.NotFound(w, r)
	}
}

func handleRequests() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {

	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
		defer db.Close()
	}
	fmt.Println("success")

	handleRequests()

}

func Article_Delete() {
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(">>>>>>0", err)
	} else {
		articleModel := models.ArticleModel{
			Db: db,
		}
		rows, err := articleModel.Delete(8)
		if err != nil {
			fmt.Println(">>>>>1", err)
		} else {
			if rows > 0 {
				fmt.Println("delete done")
			}
		}
	}
	defer db.Close()
}

func Article_Update() {
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(">>>>>>0", err)
	} else {
		articleModel := models.ArticleModel{
			Db: db,
		}
		article := entities.Article{
			Id:      2,
			Title:   "Title 10222222",
			Desc:    "Description hzfnew",
			Context: "new content d",
		}
		rows, err := articleModel.Update(article)
		if err != nil {
			fmt.Println(">>>>>1", err)
		} else {
			if rows > 0 {
				fmt.Println("done")
			}
		}

		entities.Articles = append(entities.Articles, article)

	}
	defer db.Close()
}

func Article_Create() {
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(">>>>>>0", err)
	} else {
		articleModel := models.ArticleModel{
			Db: db,
		}
		article := entities.Article{
			Id:      15,
			Title:   "Title 10",
			Desc:    "Description 10 new",
			Context: "new content with create method",
		}
		err := articleModel.Create(&article)
		if err != nil {
			fmt.Println(">>>>>1", err)
		} else {
			fmt.Println("lastest id", article.Id)
		}

		entities.Articles = append(entities.Articles, article)

	}
	defer db.Close()
}

func Article_CallFind() {
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		articleModel := models.ArticleModel{
			Db: db,
		}
		article, err := articleModel.Find(3)
		entities.Articles = append(entities.Articles, article)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(article)
			for _, article := range entities.Articles {
				fmt.Println("Title", article.Title)
				fmt.Println("Title", article.Desc)
				fmt.Println("Title", article.Context)
			}

		}
	}
	defer db.Close()
}

func Article_CallFindAll() {
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		articleModel := models.ArticleModel{
			Db: db,
		}
		articles, err := articleModel.FindAll()
		entities.Articles = append(entities.Articles, articles...)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(articles)
			fmt.Println("Articles List")
			for _, article := range articles {
				fmt.Println("Title", article.Title)
				fmt.Println("Title", article.Desc)
				fmt.Println("Title", article.Context)
				fmt.Println("----------------------")
			}

		}
	}

}
