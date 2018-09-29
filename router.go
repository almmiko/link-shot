package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Endpoint struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Endpoints []Endpoint

var endpoints = Endpoints{
	Endpoint{
		Name:        "new",
		Method:      "POST",
		Pattern:     "/new",
		HandlerFunc: CreateShortUrl,
	},
	Endpoint{
		Name:        "root",
		Method:      "GET",
		Pattern:     "/{code}",
		HandlerFunc: RedirectToUrl,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, endpoint := range endpoints {
		var handler http.Handler
		handler = endpoint.HandlerFunc
		handler = Logger(handler, endpoint.Name)

		router.
			Methods(endpoint.Method).
			Path(endpoint.Pattern).
			Name(endpoint.Name).
			Handler(handler)
	}

	return router
}
