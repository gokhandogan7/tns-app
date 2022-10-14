package router

import (
	"api/controllers"
	"api/helpers"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		helpers.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

}

func HandleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", controllers.GetAllArticles).Methods(http.MethodGet)
	myRouter.HandleFunc("/article/{id}", controllers.GetSingleArticle).Methods(http.MethodGet)
	myRouter.HandleFunc("/article", controllers.CreateNewArticle).Methods(http.MethodPost)
	myRouter.HandleFunc("/article/{id}", controllers.DeleteArticle).Methods(http.MethodDelete)
	myRouter.HandleFunc("/article/{id}", controllers.UpdateArticle).Methods(http.MethodPut)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	// start server listen
	// with error handling
	handler := handlers.CORS(originsOk, headersOk, methodsOk)(myRouter)
	log.Fatal(http.ListenAndServe(":8081", handler))

}
