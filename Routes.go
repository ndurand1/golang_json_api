package main

import (
	"golang_json_api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		handlers.TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"todos/{todoId}",
		handlers.TodoShow,
	},
}

// create nwe router instance, returns a pointer to mux.Router
func NewRouter() *mux.Router {
	// StrictSlash redirects /users/ to /users, for example
	router := mux.NewRouter().StrictSlash(true)
	// _ b/c we don't need index of element in slice
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
