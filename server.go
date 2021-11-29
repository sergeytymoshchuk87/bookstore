package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/styles/").Handler(http.StripPrefix("/static/styles/",
		http.FileServer(http.Dir("static/styles"))))
	r.SkipClean(true)
	r.HandleFunc("/search/author/{auth}", QueryByAuthor)
	r.HandleFunc("/search/isbn/{ISBN}", QueryByISBN)
	r.HandleFunc("/issued-books", IssuedBooks)
	r.HandleFunc("/top-trending", ShowTopTrending)
	r.HandleFunc("/most-issued", MostIssued)
	r.HandleFunc("/books", ShowAllBooks)
	r.HandleFunc("/available", ShowAvailable)

	http.ListenAndServe(":8080", r)
}
