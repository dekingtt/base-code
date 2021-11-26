package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func header(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%s=%v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/headers", header)
	http.ListenAndServe(":8080", nil)
}
