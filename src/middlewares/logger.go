package middlewares

// https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
// https://gowebexamples.github.io/advanced-middleware/
// https://www.nicolasmerouze.com/middlewares-golang-best-practices-examples/
// http://www.alexedwards.net/blog/making-and-using-middleware
// Solution: https://gowebexamples.github.io/advanced-middleware/
// Status code:
// http://ndersson.me/post/capturing_status_code_in_net_http/
// https://gist.github.com/Boerworz/b683e46ae0761056a636

import (
	"log"
	"net/http"
	"time"
)

// Middleware : is a function that both takes in and returns an http.Handler
//type Middleware func(http.HandlerFunc) http.HandlerFunc
type Middleware func(http.Handler) http.Handler

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// NewLoggingResponseWriter :
func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	// WriteHeader(int) is not called if our response implicitly returns 200 OK, so
	// we default to that status code.
	return &loggingResponseWriter{w, http.StatusOK}
}

// Logging : adapter wraps an http.Handler with additional functionality
// It's allowing the original http.Handler `h`
// to do whatever it was already going to do in between
/*func Logging(l *log.Logger) Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Printf("--> %s %s | %s", r.Method, r.URL.Path, time.Since(start))
				log.Printf("<-- %d %s", http.StatusOK, http.StatusText(http.StatusOK))
			}()
			h.ServeHTTP(w, r)
		})
	}
}*/
func Logging(l *log.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			lrw := NewLoggingResponseWriter(w)
			defer func() {
				log.Printf("--> %s %s | %s", r.Method, r.URL.Path, time.Since(start))
				//log.Printf("<-- %d %s", http.StatusOK, http.StatusText(http.StatusOK))
				//lrw := NewLoggingResponseWriter(w)
				//h.ServeHTTP(lrw, r)
				statusCode := lrw.statusCode
				log.Printf("<-- %d %s", statusCode, http.StatusText(statusCode))
			}()
			//h.ServeHTTP(w, r)
			h.ServeHTTP(lrw, r)
		})
	}
}

// Method : ensures that url can only be requested with a specific method,
// else returns a 400 Bad Request
// curl -I -X POST 127.0.0.1:8080
/*func Method(m string) Middleware {
	// Create a new Middleware
	return func(h http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			// Call the next middleware/handler in chain
			h(w, r)
		}
	}
}*/

// Method : ensures that url can only be requested with a specific method,
// else returns a 400 Bad Request
// curl -I -X POST 127.0.0.1:8080
func Method(m string) Middleware {
	// Create a new Middleware
	return func(h http.Handler) http.Handler {
		// Define the http.HandlerFunc
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			// Call the next middleware/handler in chain
			h.ServeHTTP(w, r)
		})
	}
}

// Chain : takes the handler you want to adapt, and a list of our Adapter types
// h with all specified adapters.
/*func Chain(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}*/
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
