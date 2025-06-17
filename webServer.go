package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// root, welcome page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

// endpoint 1
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

// endpoint 2
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// clean user input with Escape String as this is rendering to HTML, prevents XSS
	fmt.Fprintf(w, "Todo show: %s", html.EscapeString(vars["todoId"]))
}

func main() {
	// StrictSlash redirects /users/ to /users, for example
	router := mux.NewRouter().StrictSlash(true)
	// handle calls to multiple endpoints
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoID}", TodoShow)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
