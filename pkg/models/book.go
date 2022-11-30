package models

import (
	"github.com/fadebowaley/Quito-Library/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Initialize Database Config Function
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// function to create Books
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// function to get all Books
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// function to get books by its ID
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}

// function to delete a book by its ID
func DeleteBook(ID int64) *Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return &book
}