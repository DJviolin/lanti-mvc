package main

/*
Gorilla/mux:
https://github.com/gorilla/mux#examples
https://medium.com/dev-bits/understanding-the-gorilla-mux-a-sturdy-url-router-from-the-golang-7494660f4907
http://howistart.org/posts/go/1/
*/

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/djviolin/lanti-mvc/src/controllers"
	lib "github.com/djviolin/lanti-mvc/src/lib"
	mw "github.com/djviolin/lanti-mvc/src/middlewares"
	"github.com/gorilla/mux"
)

func main() {
	// Variables to identify the build
	log.Print("Version: ", lib.Version)
	log.Print("Git commit hash: ", lib.Build)

	// Init logger
	logger := log.New(os.Stdout, "server: ", log.Lshortfile)

	// Static files
	// http://stackoverflow.com/a/26563418/1442219
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/public"))))

	// Init Gorilla/mux router
	r := mux.NewRouter()
	// Routes
	//r.HandleFunc("/", controllers.Index)
	http.Handle("/", mw.Notify(logger)(controllers.Index)) // route with logger middleware
	r.HandleFunc("/hello/{param}", controllers.Hello)
	// Static files
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))
	// This comes after the routes
	http.Handle("/", r)

	// Server
	port := ":" + strconv.Itoa(lib.Port()) // int to string
	p := &port
	err := http.ListenAndServe(*p, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
