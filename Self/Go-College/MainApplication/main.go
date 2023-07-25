package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/SudebDolui/CollegeApi/Packages/Config"
	router "github.com/SudebDolui/CollegeApi/Packages/Router"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to College API Made Using MongoDB")
	myRouter := mux.NewRouter()
	router.RegisterRouter(myRouter)
	fmt.Println("Server is getting Started")
	config.Connect()
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
