package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var permissions = []string{
	"/demo.jpg",
}

func main() {
	router := httprouter.New()
	router.GET("/pvt/*path", pvtFileServe)
	http.ListenAndServe(":80", router)
}

func pvtFileServe(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	paramPath := ps.ByName("path")
	if searchInSlice(permissions, paramPath) != -1 {
		http.ServeFile(w, req, "pvt"+paramPath)
	} else {
		w.Write([]byte("Permission Error"))
	}
}

func searchInSlice(slice []string, toSearch string) int {
	for index := range slice {
		if slice[index] == toSearch {
			return index
		}
	}
	return -1
}
