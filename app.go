package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if host == "" || port == "" {
		log.Fatal("expecting HOST and PORT env variables to be set")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := html.EscapeString(r.URL.Path)
		fmt.Fprintf(w, "Hello, %q from %s:%s\n", path, host, port)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), mux))
}
