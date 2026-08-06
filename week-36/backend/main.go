package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "DATABASE_URL not set"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	http.HandleFunc("/db", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Database URL: %s\n", dbURL)
	})

	addr := ":" + port
	log.Printf("Database URL: %s\n", dbURL)
	log.Printf("Server listening on http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}