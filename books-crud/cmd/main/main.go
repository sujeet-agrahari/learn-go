package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sujeet-agrahari/books-crud/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}