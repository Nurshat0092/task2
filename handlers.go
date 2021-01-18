package main

import (
	"encoding/json"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		form := newForm(r.URL.Query())
		form.required("word")
		if !form.valid() {
			jsonResponse(w, form.errors, http.StatusBadRequest)
			return
		}
		word := form.Get("word")

		anns := getAnnagrams(word)

		jsonResponse(w, anns, http.StatusOK)
	}
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		words := []string{}
		err := json.NewDecoder(r.Body).Decode(&words)
		if err != nil {
			jsonResponse(w, "Invalid json structure", http.StatusBadRequest)
			return
		}
		loadAnnagrams(words)
	}
}

func jsonResponse(w http.ResponseWriter, d interface{}, c int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(c)
	err := json.NewEncoder(w).Encode(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	}
}
