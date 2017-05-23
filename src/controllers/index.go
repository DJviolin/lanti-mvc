package controllers

import (
	"log"
	"net/http"

	lib "github.com/djviolin/lanti-mvc/src/lib"
	mw "github.com/djviolin/lanti-mvc/src/middlewares"
)

// Student : constructor for template
/*type Student struct {
	//exported field since it begins
	//with a capital letter
	Name string
}*/

// Index : is the index handler
func Index(w http.ResponseWriter, r *http.Request) {
	render, err := mw.ParseDirectory("./views", "index")
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	render.Execute(w, map[string]string{
		"Title": "My title",
		"Body":  "This is the body",
		"tmp":   "index",
		"user":  lib.User,
	})
}
