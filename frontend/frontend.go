package main

import (
	"CSCI566_Project/frontend/handlers"
	"fmt"
	"log"
	"net/http"
)

func main(){

	// serve vue app
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	 http.Handle("/", &handlers.IndexHandler{})
	
	fmt.Println("starting server on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}