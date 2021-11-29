package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"sort"
	"strings"
)

var t = `
	{{ if not .Books}} No Result {{ else }}
	{	{{ with .Books }} {{ range .}}
		{
		"id": {{ .ID}},
        	"book_name": "{{ .BookName }}",
        	"desc" : "{{ .Description }}",
        	"ISBN": "{{ .ISBN }}",
        	"hits_this_week": {{ .HitsThisWeek }},
        	"all_time_issues": {{ .AllTimeIssues }},
		"issue_status" : {{ .IssueStatus }},{{ if .IssueStatus }} {{ with .Issue }}
        	"current_issue": {
        	   "user": "{{ .User }}"
        	   "issue_date": "{{ .IssueDate }}",
        	   "issue_expiry": "{{ .IssueExpiry }}"
		},{{ end }} {{ end }}
        	"author": "{{ .Author }}",
        	"date_published": "{{ .DatePublished }}"
		}, {{ end }} {{ end }}
	} {{ end }}
	`

// QueryByAuthor searches using the author name ...
func QueryByAuthor(w http.ResponseWriter, r *http.Request) {
	author := mux.Vars(r)["auth"]
	tmp := Store{}

	for ix, val := range JSON.Books {
		auth := &val.Author
		if strings.Contains(strings.ToLower(*auth), strings.ToLower(author)) {
			tmp.Books = append(tmp.Books, JSON.Books[ix])
		}
	}

	renderTempate(&w, t, &tmp)
}

// QueryByISBN searches using the ISBN...
// There is no form for inputing ISBN...The will have to directly use the URL method
func QueryByISBN(w http.ResponseWriter, r *http.Request) {

	form := mux.Vars(r)["ISBN"]
	tmp := Store{}

	for ix, val := range JSON.Books {
		isbn := val.ISBN
		if isbn == form {
			tmp.Books = append(tmp.Books, JSON.Books[ix])
			break
		}
	}

	renderTempate(&w, t, &tmp)
}

// IssuedBooks searches all books that are currently issued ...
func IssuedBooks(w http.ResponseWriter, r *http.Request) {

	tmp := Store{}

	for ix, val := range JSON.Books {
		status := &val.IssueStatus
		if *status == true {
			tmp.Books = append(tmp.Books, JSON.Books[ix])
		}
	}

	renderTempate(&w, t, &tmp)
}

// ShowTopTrending shows the top trending Books rated by hit Hits_Per_Week ...
func ShowTopTrending(w http.ResponseWriter, r *http.Request) {

	tmp := Store{}
	sort.SliceStable(JSON.Books, func(i, j int) bool {
		return JSON.Books[i].HitsThisWeek > JSON.Books[j].HitsThisWeek
	})

	for i := 0; i < 5; i++ {
		tmp.Books = append(tmp.Books, JSON.Books[i])
	}

	renderTempate(&w, t, &tmp)
}

// MostIssued shows the Book with the Highest Number Of Issues
func MostIssued(w http.ResponseWriter, r *http.Request) {
	tmp := Store{}

	sort.SliceStable(JSON.Books, func(i, j int) bool {
		return JSON.Books[i].AllTimeIssues > JSON.Books[j].AllTimeIssues
	})
	tmp.Books = append(tmp.Books, JSON.Books[0])

	renderTempate(&w, t, &tmp)
}

// ShowAllBooks shows all the Books in the store
func ShowAllBooks(w http.ResponseWriter, r *http.Request) {
	renderTempate(&w, t, &JSON)
}

// ShowAvailable shows all available books
func ShowAvailable(w http.ResponseWriter, r *http.Request) {

	tmp := Store{}

	for ix, val := range JSON.Books {
		status := &val.IssueStatus
		if *status == false {
			tmp.Books = append(tmp.Books, JSON.Books[ix])
		}
	}

	renderTempate(&w, t, &tmp)
}
