package main

/*
Gorilla/mux:
https://github.com/gorilla/mux#examples
https://medium.com/dev-bits/understanding-the-gorilla-mux-a-sturdy-url-router-from-the-golang-7494660f4907
http://howistart.org/posts/go/1/
*/

/*
Security:
https://dev.to/joncalhoun/what-is-sql-injection-and-how-do-i-avoid-it-in-go
*/

import (
	"log"
	"net/http"
	"os"
	"strconv"

	controllers "github.com/djviolin/lanti-mvc/src/controllers"
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
	// Init Gorilla/mux router
	r := mux.NewRouter()
	// Routes
	/*r.HandleFunc("/", mw.ChainFunc(controllers.Index,
		mw.MethodFunc("GET"),
	))
	r.HandleFunc("/hello/{param}", mw.ChainFunc(controllers.Hello,
		mw.MethodFunc("GET"),
	))*/
	r.HandleFunc("/", controllers.Index).Methods("GET")
	r.HandleFunc("/hello/{param}", controllers.Hello).Methods("GET")
	// Static files
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))
	// Init routes: this comes after the routes
	http.Handle("/", mw.Chain(r, mw.Logging(logger)))

	// Server
	port := ":" + strconv.Itoa(lib.Port()) // int to string
	p := &port
	err := http.ListenAndServe(*p, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
