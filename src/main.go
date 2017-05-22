package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/djviolin/lanti-mvc/src/controllers"
	lib "github.com/djviolin/lanti-mvc/src/lib"
	"github.com/gorilla/mux"
)

func main() {
	// Variables to identify the build
	fmt.Println("Version: ", lib.Version)
	fmt.Println("Git commit hash: ", lib.Build)

	// Static files
	// http://stackoverflow.com/a/26563418/1442219
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/public"))))

	// Init Gorilla/mux router
	r := mux.NewRouter()

	// Routes
	//http.HandleFunc("/", controllers.Index)
	r.HandleFunc("/", controllers.Index)
	r.HandleFunc("/hello/{param}", controllers.Hello)

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
