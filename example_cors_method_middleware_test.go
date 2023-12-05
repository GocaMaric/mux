package mux_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

// setupRouter initializes and configures the router with the required routes and middleware.
func setupRouter() *mux.Router {
	r := mux.NewRouter()

	// Actual route handling
	r.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		// Handle the request
	}).Methods(http.MethodGet, http.MethodPut, http.MethodPatch)

	// CORS options handling
	r.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://example.com")
		w.Header().Set("Access-Control-Max-Age", "86400")
	}).Methods(http.MethodOptions)

	// CORS middleware setup
	r.Use(mux.CORSMethodMiddleware(r))

	return r
}

func ExampleCORSMethodMiddleware() {
	// Initialize and configure the router
	r := setupRouter()

	// Create a new request for testing
	req, err := http.NewRequest("OPTIONS", "/foo", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set request headers
	req.Header.Set("Access-Control-Request-Method", "POST")
	req.Header.Set("Access-Control-Request-Headers", "Authorization")
	req.Header.Set("Origin", "http://example.com")

	// Serve the request
	rw := httptest.NewRecorder()
	r.ServeHTTP(rw, req)

	// Print the results
	fmt.Println("Access-Control-Allow-Methods:", rw.Header().Get("Access-Control-Allow-Methods"))
	fmt.Println("Access-Control-Allow-Origin:", rw.Header().Get("Access-Control-Allow-Origin"))
	// Output:
	// GET,PUT,PATCH,OPTIONS
	// http://example.com
}
