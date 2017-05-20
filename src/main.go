package main

import (
	"log"
	"net/http"

	"github.com/djviolin/lanti-mvc/src/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Index)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
