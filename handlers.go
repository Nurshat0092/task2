package main

import (
	"encoding/json"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// word := r.URL.Query().Get("word")
		// annagrams := getAnnagrams(word)
	}
}

func loadHandler(w http.ResponseWriter, r *http.Request) {

}

func jsonResponse(w http.ResponseWriter, d interface{}, c int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(c)
	err := json.NewEncoder(w).Encode(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	}
}
