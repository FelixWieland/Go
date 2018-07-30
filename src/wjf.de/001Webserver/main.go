package main

import (
	"net/http"

	"wjf.de/001Webserver/lib"
)

func serverOutput(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = lib.Demo()
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", serverOutput)
	http.ListenAndServe(":8080", nil)
}
