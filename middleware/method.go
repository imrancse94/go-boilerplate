package middleware

import (
	"go-boilerplate/gate"
	"go-boilerplate/response"
	"net/http"
)

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) gate.Middleware {

	// Create a new gate
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			//m1 := map[string]interface{}{}
			//requests.DecodeJsonRequest(r, &m1)
			//fmt.Println("D1", m1)
			// Do middleware things
			if r.Method != m {
				response.ErrorResponse(response.ErrorResponseStruct{
					StatusCode: "E001",
					Message:    "Method not allowed",
					Error:      "",
				}, w)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}
