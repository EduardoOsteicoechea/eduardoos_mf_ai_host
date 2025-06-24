// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for: %s from %s", r.URL.Path, r.RemoteAddr)
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	fmt.Fprintf(w, PageTop + PageBody + auiditing_love_through_the_law + PageBottom)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for: %s from %s", r.URL.Path, r.RemoteAddr)
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	fmt.Fprintf(w, todo_page_view(PageTop, PageBody, PageBottom))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received API request for: %s from %s", r.URL.Path, r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "This is an API response from eduardoos_mf_ai_host!", "path": "%s"}`+"\n", r.URL.Path)
}

func main() {
	// Register handlers for different routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/todo", todoHandler)

	// Define the port to listen on. This should match the Nginx proxy_pass port.
	port := "7007"
	addr := fmt.Sprintf(":%s", port)

	log.Printf("Eduardoos MF AI Host Go application starting on port %s...", port)

	// Start the HTTP server
	// log.Fatal will exit the application if the server fails to start, logging the error.
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed to start on port %s: %v", port, err)
	}
}