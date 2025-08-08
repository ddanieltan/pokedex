package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", index)
	router.HandleFunc("/about", index)
	router.HandleFunc("/skills", index)

	server := http.Server{
		Addr:    ":3000",
		Handler: loggingMiddleware(router),
	}

	fmt.Println("Server running...")
	server.ListenAndServe()
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, "Error parsing "+tmpl, http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}
