package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// we will be using structs and slices
// its like object in java and javascript
// we are going to define the types of data inside that

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// every movies has  a director

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "4384587", Title: "Movie one", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438989", Title: "Movie Two", Director: &Director{Firstname: "Lebron", Lastname: "James"}})
	movies = append(movies, Movie{ID: "3", Isbn: "438342", Title: "Movie Three", Director: &Director{Firstname: "Trevor", Lastname: "Wallace"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("String server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
