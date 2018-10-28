package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template //package.Type

var fm = template.FuncMap{}

func init() {
	//Binds functions to templates
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.html"))
}

func main() {
	router := httprouter.New()
	router.GET("/plc/*path", serveFiles)
	router.GET("/reloadTemplates", reloadTemplates)
	router.GET("/", index)
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//Anonymouse struct Datatype
	sDemo := struct {
		name  string
		lname string
	}{
		"James",
		"Bond",
	}

	err := tpl.ExecuteTemplate(w, "index.html", sDemo)
	if err != nil {
		log.Fatalln(err)
	}
}

func reloadTemplates(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.html"))
	w.Write([]byte("Reloaded templates"))
}

func serveFiles(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	paramPath := ps.ByName("path")
	if true {
		http.ServeFile(w, req, strings.Replace(paramPath, "/", "", 1))
	} else {
		w.Write([]byte("Permission Error"))
	}
}
