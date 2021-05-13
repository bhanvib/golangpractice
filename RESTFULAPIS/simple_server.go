package main

import (
	"fmt"
	"io"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"project.com/DAL"
)

//Routing:
//Mapping of incoming URL with a given method

//Request: User's data
//Response : Data/Info  sent by server

func PostBooks(w http.ResponseWriter, r *http.Request) {
	// Accesing data send in Request body of Request object
	//r.Body  : JSOn
	var book DAL.Book
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&book)
	fmt.Println(err)
	fmt.Println(book)
	DAL.InsertBook(book)
}

func GetBookByIsbn(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	isbn := params["isbn"]
	book := DAL.GetBookByIsbn(isbn)
	//convert golang object to json and sending to Response object
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {

	//io.WriteString(w, "GetBooks")

	books := DAL.GetBooks() // slice of books

	//convert golang object to json and sending to Response object
	json.NewEncoder(w).Encode(books)

}
func DeleteBooks(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "DeleteBooks")

}
func PutBooks(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "PutBooks")

}

func main() {

	//http://localhost:8080/books
	//http.HandleFunc("/books", GetBooks) // ROute

	Router := mux.NewRouter()
	Router.HandleFunc("/books", GetBooks).Methods("GET")
	Router.HandleFunc("/books/{isbn}", GetBookByIsbn).Methods("GET") //localhsot:8080/books/isbn123
	
	Router.HandleFunc("/books", PostBooks).Methods("POST")
	Router.HandleFunc("/books", PutBooks).Methods("PUT")
	Router.HandleFunc("/books", DeleteBooks).Methods("DELETE")

	http.ListenAndServe(":5050", Router)

}
