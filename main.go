package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":4000", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from go")
}
