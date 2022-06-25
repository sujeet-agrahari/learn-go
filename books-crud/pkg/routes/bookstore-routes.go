package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sujeet-agrahari/books-crud/pkg/controllers"
)

type NotFound struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(NotFound{Status: 404, Message: "Not Found"})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(res)
}

var RegisterBookStoreRoutes = func (router *mux.Router)  {
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
}