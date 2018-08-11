package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var fm = template.FuncMap{
	//"uc":     demoFunc
}

var tpl *template.Template //package.Type

func init() {
	//Binds functions to templates
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("*.html"))
}

func main() {
	router := httprouter.New()
	router.GET("/files/*path", fileServe)
	router.GET("/", index)
	router.GET("/RELOAD", reload)
	http.ListenAndServe(":80", router)
}

func index(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	templateting(&w, "index.html", struct{}{})
}

func fileServe(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	paramPath := ps.ByName("path")
	http.ServeFile(w, req, "files/"+paramPath)
}

func templateting(w *http.ResponseWriter, file string, templateValues struct{}) {
	err := tpl.ExecuteTemplate(*w, file, templateValues)
	if err != nil {
		log.Fatalln(err)
	}
}

func reload(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("*.html"))
}
