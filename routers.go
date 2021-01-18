package main

import (
	"net/http"
)

func routes() *http.ServeMux {
	mux := &http.ServeMux{}

	mux.HandleFunc("/get", getHandler)
	mux.HandleFunc("/load", loadHandler)

	return mux
}
