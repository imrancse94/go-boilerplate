package middleware

import (
	"go-boilerplate/gate"
	"log"
	"net/http"
	"time"
)

func Logging() gate.Middleware {

	// Create a new gate
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			/*m1 := map[string]interface{}{}
			requests.DecodeJsonRequest(r, &m1)
			fmt.Println("m1", m1)
			w.Write([]byte("<h1>Logging</h1>"))*/
			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}
