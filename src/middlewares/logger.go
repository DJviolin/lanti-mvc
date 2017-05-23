package middlewares

// https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
// https://www.nicolasmerouze.com/middlewares-golang-best-practices-examples/
// http://www.alexedwards.net/blog/making-and-using-middleware

import (
	"log"
	"net/http"
)

// Adapter : is a function that both takes in and returns an http.Handler
type Adapter func(http.Handler) http.Handler

// Notify : adapt an http.Handler to write out the “before” and “after” strings,
// allowing the original http.Handler `h`
// to do whatever it was already going to do in between
func Notify(logger *log.Logger) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println("before")
			defer logger.Println("after")
			h.ServeHTTP(w, r)
		})
	}
}
