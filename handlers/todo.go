package handlers

import (
	"encoding/json"
	"fmt"
	"golang_json_api/models"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

// root, welcome page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

// endpoint 1: sending back JSON
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	// struct in models/todo.go (Todos is ordered collection of Todo struct)
	todos := models.Todos{
		models.Todo{Name: "Write presentation"},
		models.Todo{Name: "Ship laptop"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// endpoint 2
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// clean user input with Escape String as this is rendering to HTML, prevents XSS
	fmt.Fprintf(w, "Todo show: %s", html.EscapeString(vars["todoId"]))
}
