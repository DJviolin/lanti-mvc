package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/djviolin/lanti-mvc/src/controllers"
	lib "github.com/djviolin/lanti-mvc/src/lib"
	"github.com/gorilla/mux"
)

func main() {
	// Variables to identify the build
	log.Print("Version: ", lib.Version)
	log.Print("Git commit hash: ", lib.Build)

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
	port := ":" + strconv.Itoa(lib.Port()) // int to string
	p := &port
	log.Printf("Listening on port %s...", *p)
	err := http.ListenAndServe(*p, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
