package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// JSON is struct representation of the store.json file ...
var JSON Store

// CurrentIssue represents the Current Issue info ...
type CurrentIssue struct {
	User        string `json:"user"`
	IssueDate   string `json:"issue_date"`
	IssueExpiry string `json:"issue_expiry"`
}

// Details is the struct representation of the Book details ...
type Details struct {
	ID            int          `json:"id"`
	BookName      string       `json:"book_name"`
	Description   string       `json:"desc"`
	ISBN          string       `json:"ISBN"`
	HitsThisWeek  int          `json:"hits_this_week"`
	AllTimeIssues int          `json:"all_time_issues"`
	IssueStatus   bool         `json:"issue_status"`
	Issue         CurrentIssue `json:"current_issue"`
	Author        string       `json:"author"`
	DatePublished string       `json:"date_published"`
}

// Store holds the store.json array of Books ...
type Store struct {
	Books []Details `json:"books"`
}

// init initializes the data and JSON variables
func init() {

	// data  is the byte slice of the store.json file
	var data []byte

	data, err := ioutil.ReadFile("store.json")
	JSON = Store{}
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	json.Unmarshal(data, &JSON)
}
