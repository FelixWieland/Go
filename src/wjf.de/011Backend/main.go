package main

import (
	"flag"
	"log"
	"net/http"
)

func backend(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir(*flag.String("d", ".", "/")))
}

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "/")

	flag.Parse()

	//Check for Backend
	http.HandleFunc("/backend", backend)

	//Server
	http.Handle("/", http.FileServer(http.Dir(*directory)))
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
