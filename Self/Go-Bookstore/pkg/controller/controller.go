package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SudebDolui/Go-Bookstore/pkg/models"
	"github.com/SudebDolui/Go-Bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBookstores(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBookstores()
	res, err := json.Marshal(newBooks)
	utils.ErrorUsingLog(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetOneBookstore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	utils.ErrorUsingLog(err)
	bookDetails, _ := models.GetOneBookstore(Id)
	res, err := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBookstore(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, err := json.Marshal(b)
	utils.ErrorUsingLog(err)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteOneBookstore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	utils.ErrorUsingLog(err)
	book := models.DeleteOneBookstore(Id)
	res, err := json.Marshal(book)
	utils.ErrorUsingLog(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// func DeleteBookstores() {

// }

func UpdateBookstore(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	utils.ErrorUsingLog(err)
	bookDetails, db := models.GetOneBookstore(Id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, err := json.Marshal(bookDetails)
	utils.ErrorUsingLog(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
