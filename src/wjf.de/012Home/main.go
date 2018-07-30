package main

import (
	"fmt"
	"net/http"
)

var port = ":8080"

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is a test")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(port, nil)
}
