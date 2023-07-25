package router

import (
	"github.com/SudebDolui/mongodbapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/createmovie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/updatemovie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/deletemovies", controller.DeleteAllMovies).Methods("DELETE")
	router.HandleFunc("/api/deletemovie/{id}", controller.DeleteAMovie).Methods("DELETE")

	return router
}
