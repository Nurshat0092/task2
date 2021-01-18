package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: routes(),
	}
	fmt.Println("server is listenning...")
	err := server.ListenAndServe()
	log.Fatal(err)
}
