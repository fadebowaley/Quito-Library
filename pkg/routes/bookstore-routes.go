package ro🇺tes 


import (
	"github.com/gorilla/mux"
	"github.com/fadebowaley/Quito-Library/pkg/contollers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){

	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")	
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}

var CreateNewLibaryRoutes = func(router *muxRouter){
	router.HandleFunc("/libary/", controllers.CreatLibrary).Methods("POST")
	router.HandleFunc("/library", controllers.GetLibrary).Methods("GET")
	router.HandleFunc("/libray/{libraryId}", controllers.GetLibraryById).Methods("GET")

}