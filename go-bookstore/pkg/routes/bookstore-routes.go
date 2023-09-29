package routes 
// golang has the absolute path and not relative path 
// "github.com/

import (
	"github.com/goriila/mux"
	"github.com/bloodgroupcplusplus/fcc_go_akhil/pkg/controllers"

)


var RegisterBookStore = func (router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book{bookId}",controllers.GetBookById).Methods("GET")

	router.HandleFunc("/book{bookId}",controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{bookdId}",controllers.UpdateBook).Methods("PUT")
}
	

