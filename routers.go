package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	mux := pat.New()
	mux.Get("/get", http.HandlerFunc(getHandler))
	mux.Post("/load", http.HandlerFunc(loadHandler))

	return mux
}
