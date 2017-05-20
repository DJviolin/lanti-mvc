package main

import (
	"log"
	"net/http"

	"github.com/djviolin/lanti-mvc/src/controllers"
)

func main() {
	// Static files
	// http://stackoverflow.com/a/26563418/1442219
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	// Routes
	http.HandleFunc("/", controllers.Index)

	// Server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
