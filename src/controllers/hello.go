package controllers

import (
	"log"
	"net/http"

	mw "github.com/djviolin/lanti-mvc/src/middlewares"
	"github.com/gorilla/mux"
)

// Hello : is the "/hello" route handler
func Hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	render, err := mw.ParseDirectory("./views", "index")
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	render.Execute(w, map[string]string{
		"Title": "My title",
		"Body":  "This is the body",
		"tmp":   "index",
		"param": vars["param"],
	})
}
