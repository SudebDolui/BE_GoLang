package router

import (
	"github.com/SudebDolui/Go-Bookstore/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/createbookstore", controller.CreateBookstore).Methods("POST")
	router.HandleFunc("/api/getbookstores", controller.GetBookstores).Methods("GET")
	router.HandleFunc("/api/getonebookstore/{bookid}", controller.GetOneBookstore).Methods("GET")
	router.HandleFunc("/api/updatebookstore/{bookid}", controller.UpdateBookstore).Methods("PUT")
	router.HandleFunc("/api/deleteonebookstore/{bookid}", controller.DeleteOneBookstore).Methods("DELETE")
	// router.HandleFunc("/api/updatebookstores", controller.DeleteBookstores).Methods("DELETE")

}
