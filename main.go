package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	// ListenAndServe(addr string, handler http.handler returns error)
	log.Fatal(http.ListenAndServe(":8080", router))
}
