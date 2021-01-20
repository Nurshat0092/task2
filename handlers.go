package main

import (
	"encoding/json"
	"net/http"

	"github.com/Nurshat0092/task2/storage"
	"github.com/Nurshat0092/task2/utils"
)

// endPoint to get annagram
func getHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		form := utils.NewForm(r.URL.Query())
		form.Required("word")
		if !form.Valid() {
			jsonResponse(w, form.Errors, http.StatusBadRequest)
			return
		}
		word := form.Get("word")

		anns := storage.GetAnnagrams(word)

		jsonResponse(w, anns, http.StatusOK)
	}
}

// endPoint to load annagrams
func loadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		words := []string{}
		err := json.NewDecoder(r.Body).Decode(&words)
		if err != nil {
			jsonResponse(w, "Invalid json structure", http.StatusBadRequest)
			return
		}
		storage.LoadAnnagrams(words)
	}
}

// jsonResponse helper function
func jsonResponse(w http.ResponseWriter, d interface{}, c int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(c)
	err := json.NewEncoder(w).Encode(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	}
}
