// USED: mongodb+srv://Demo:<password>@cluster0.6422n.mongodb.net/myFirstDatabase?retryWrites=true&w=majority
// USED: go get -u github.com/gorilla/mux
// USED: go get go.mongodb.org/mongo-driver/mongo

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SudebDolui/mongodbapi/router"
)

func main() {
	fmt.Println("Welcome to MONGODB SPI")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
}
