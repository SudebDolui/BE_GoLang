package main

import (
	"log"
	"net/http"

	"github.com/SudebDolui/Go-Bookstore/pkg/router"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	router.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))

}
