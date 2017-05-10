package main

import (
	"log"
	"net/http"
	"fmt"
	"html"
	"/blog"
)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {	
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		blog.Start();
	})
	log.Fatal(http.ListenAndServe(":8083", nil))
}
