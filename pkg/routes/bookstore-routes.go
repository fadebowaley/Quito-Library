package routes




import (
	"github.com/gorilla/mux"
	"github.com/fadebowaley/Quito-Library/pkg/controllers"
)

var RegisterBookStoreRouters = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
}

