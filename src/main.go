package main

import (
	"log"
	"net/http"

	"github.com/djviolin/lanti-mvc/src/controllers"
	"github.com/gorilla/mux"
)

func main() {
	// Static files
	// http://stackoverflow.com/a/26563418/1442219
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/public"))))

	// Init Gorilla/Mux router
	r := mux.NewRouter()

	// Routes
	//http.HandleFunc("/", controllers.Index)
	r.HandleFunc("/", controllers.Index) // Root page

	// Static files
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))

	// This comes after the routes
	http.Handle("/", r)

	// Server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
