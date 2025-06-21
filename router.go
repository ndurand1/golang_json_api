package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// create nwe router instance, returns a pointer to mux.Router
func NewRouter() *mux.Router {
	// StrictSlash redirects /users/ to /users, for example
	router := mux.NewRouter().StrictSlash(true)
	// _ b/c we don't need index of element in slice
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
