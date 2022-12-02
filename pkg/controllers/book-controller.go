package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fadebowaley/Quito-Library/pkg/models"
	"github.com/fadebowaley/Quito-Library/pkg/utils"
	"github.com/gorilla/mux"
)

//var NewBook models.Book


// Controller functions, --> Get all Books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Controller functions --> Get books by its ID
func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// we use blank to remove the db result from function GetBookById in models
	book, _ := models.GetBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	createdBook := book.CreateBook()
	res, _ := json.Marshal(createdBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, bookIdConversionError := strconv.ParseInt(params["id"], 0, 0)
	if bookIdConversionError != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(bookId)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, bookIdConversionError := strconv.ParseInt(params["id"], 0, 0)
	
	if bookIdConversionError != nil {
		fmt.Println("error while parsing")
	}

	book, db := models.GetBookById(bookId)
	var payload = &models.Book{}

	utils.ParseBody(r, payload)

	if payload.Author != "" {
		book.Author = payload.Author
	}

	if payload.Name != "" {
		book.Name = payload.Name
	}

	if payload.Publication != "" {
		book.Publication = payload.Publication
	}

	db.Save(&book)

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}