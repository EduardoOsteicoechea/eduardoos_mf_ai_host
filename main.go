//go:generate go run ./_mf_sccmp_todo_app/generate.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/EduardoOsteicoechea/_mf_sccmp_todo_app"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for: %s from %s", r.URL.Path, r.RemoteAddr)
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var component_css_tag string = _mf_sccmp_todo_app.PrintCssTags()
	var component_javascript_tag string = _mf_sccmp_todo_app.PrintJavaScriptTags()
	var component_markup string = _mf_sccmp_todo_app.PrintComponentMarkup()

	cssTags := []string{component_css_tag}
	javascriptTags := []string{component_javascript_tag}
	componentMarkup := []string{component_markup}

	htmlOutput := PrintPage(cssTags, javascriptTags, componentMarkup)

	fmt.Fprint(w, htmlOutput)
}

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Received request for: %s from %s", r.URL.Path, r.RemoteAddr)
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	fmt.Fprint(w, todo_page_view(PageTop, PageBody, PageBottom))
// }

// func todoHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Received request for: %s from %s", r.URL.Path, r.RemoteAddr)
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	fmt.Fprint(w, todo_page_view(PageTop, PageBody, PageBottom))
// }

func apiHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received API request for: %s from %s", r.URL.Path, r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "This is an API response from eduardoos_mf_ai_host!", "path": "%s"}`+"\n", r.URL.Path)
}


// func configureStaticFiles(mux *http.ServeMux) {
// 	staticDir := "./static"
// 	fs := http.FileServer(http.Dir(staticDir))
// 	mux.Handle("/static/", http.StripPrefix("/static/", fs))
// 	fmt.Printf("Serving static files from '%s' under the '/static/' URL path.\n", staticDir)
// }
// func main() {
// 	////////////////////
// 	////////////////////
// 	// FOR WINDOWS DEVELOPMENT ENVIRONMENT ONLY
// 	// COMPILE WITH:
// 	// go generate ./...; $env:GOOS = "windows";  $env:GOARCH = "amd64";  go build -ldflags="-s -w" -o eduardoosminfrw.exe; .\eduardoosminfrw.exe;
// 	////////////////////
// 	////////////////////

// 	mux := http.NewServeMux()

// 	// Configure static files
// 	configureStaticFiles(mux)

// 	mux.HandleFunc("/", homeHandler)
// 	mux.HandleFunc("/api", apiHandler)

// 	// Define the port to listen on. This should match the Nginx proxy_pass port.
// 	port := "7007"
// 	addr := fmt.Sprintf(":%s", port)

// 	log.Printf("Eduardoos MF AI Host Go application starting on port %s...", port)

// 	// Start the HTTP server, passing YOUR custom mux (NOT nil)
// 	if err := http.ListenAndServe(addr, mux); err != nil { // FIX: Pass 'mux' here
// 		log.Fatalf("Server failed to start on port %s: %v", port, err)
// 	}
// }

func main() {
	//////////////////////
	//////////////////////
	// FOR AWS
	//////////////////////
	//////////////////////

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
