package middlewares

// https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
// https://www.nicolasmerouze.com/middlewares-golang-best-practices-examples/
// http://www.alexedwards.net/blog/making-and-using-middleware
// Solution: https://gowebexamples.github.io/advanced-middleware/

import (
	"log"
	"net/http"
)

// Adapter : is a function that both takes in and returns an http.Handler
type Adapter func(http.HandlerFunc) http.HandlerFunc

// Logging : adapter wraps an http.Handler with additional functionality
// It's allowing the original http.Handler `h`
// to do whatever it was already going to do in between
func Logging(l *log.Logger) Adapter {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Req before")
			log.Printf("%s %s", r.Method, r.URL.Path)
			defer log.Println("Req after")
			h.ServeHTTP(w, r)
		})
	}
}

// Adapt : takes the handler you want to adapt, and a list of our Adapter types
// h with all specified adapters.
func Adapt(h http.HandlerFunc, adapters ...Adapter) http.HandlerFunc {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
