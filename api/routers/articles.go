package router

import (
	"api/helpers"
	"api/services"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_ "api/docs" // docs is generated by Swag CLI, you have to import it.

	httpSwagger "github.com/swaggo/http-swagger"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		helpers.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

}

// @title Articles API
// @version 1.0
// @description This is an article service for managing articles
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email gdogan3746@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// @BasePath /
func HandleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)

	//Article Routes
	myRouter.HandleFunc("/fullarticles", services.GetAllFullArticles).Methods(http.MethodGet)                          // All Articles
	myRouter.HandleFunc("/fullarticles/{id}", services.GetUsersAllFullArticles).Methods(http.MethodGet)                // A Users's All Articles
	myRouter.HandleFunc("/fullarticles/{user_id}/{article_id}", services.GetSingleFullArticle).Methods(http.MethodGet) // A User's Single Article

	//Author Routes
	myRouter.HandleFunc("/author", services.CreateNewAuthor).Methods(http.MethodPost)
	myRouter.HandleFunc("/authors", services.GetAllAuthors).Methods(http.MethodGet)
	myRouter.HandleFunc("/author/{id}", services.GetSingleAuthor).Methods(http.MethodGet)
	myRouter.HandleFunc("/author/{id}", services.UpdateAuthor).Methods(http.MethodPut)
	myRouter.HandleFunc("/author/{id}", services.DeleteAuthor).Methods(http.MethodDelete)

	//Highlight Routes
	myRouter.HandleFunc("/highlights", services.GetAllHighlights).Methods(http.MethodGet)
	myRouter.HandleFunc("/highlight", services.CreateNewHighlight).Methods(http.MethodPost)
	myRouter.HandleFunc("/highlight/{id}", services.GetSingleHighlight).Methods(http.MethodGet)
	myRouter.HandleFunc("/highlight/{id}", services.UpdateHighlight).Methods(http.MethodPut)
	myRouter.HandleFunc("/highlight/{id}", services.DeleteHighlight).Methods(http.MethodDelete)

	//Content Routes
	myRouter.HandleFunc("/content", services.CreateNewContent).Methods(http.MethodPost)
	myRouter.HandleFunc("/contents", services.GetAllContents).Methods(http.MethodGet)
	myRouter.HandleFunc("/content/{id}", services.GetSingleContent).Methods(http.MethodGet)
	myRouter.HandleFunc("/content/{id}", services.UpdateContent).Methods(http.MethodPut)
	myRouter.HandleFunc("/content/{id}", services.DeleteContent).Methods(http.MethodDelete)

	//CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	// start server listen
	// with error handling
	handler := handlers.CORS(originsOk, headersOk, methodsOk)(myRouter)
	myRouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8081", handler))

}
