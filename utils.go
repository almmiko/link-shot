package main

import (
	"encoding/json"
	"net/http"

	"github.com/teris-io/shortid"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func GetShortID() string {
	shortID, err := shortid.Generate()
	if err != nil {
		panic(err)
	}

	return shortID
}
